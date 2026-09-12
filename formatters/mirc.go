package formatters

import (
	"fmt"
	"io"
	"iter"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/alecthomas/chroma/v3"
)

// mIRC extended color codes
var MIRC100 = Register("mirc100", chroma.FormatterFunc(mIRCExtendedColorFormatter))

// mIRC16m is a true-colour formatter.
var MIRC16m = Register("mirc16m", chroma.FormatterFunc(mIRCHexColorFormatter))

// Print the text with the given formatting, resetting the formatting at the end
// of each line and resuming it on the next line.
//
// This way, a pager (like https://github.com/walles/moar for example) can show
// any line in the output by itself, and it will get the right formatting.
func mIRCwriteToken(w io.Writer, formatting string, text string) error {
	if formatting == "" {
		_, err := fmt.Fprint(w, text)
		return err
	}

	newlineIndices := crOrCrLf.FindAllStringIndex(text, -1)

	afterLastNewline := 0
	for _, indices := range newlineIndices {
		newlineStart, afterNewline := indices[0], indices[1]
		_, err := fmt.Fprint(w, formatting, text[afterLastNewline:newlineStart], "\x0f", text[newlineStart:afterNewline])
		if err != nil {
			return err
		}
		afterLastNewline = afterNewline
	}

	if afterLastNewline < len(text) {
		// Print whatever is left after the last newline
		_, err := fmt.Fprint(w, formatting, text[afterLastNewline:], "\x0f")
		return err
	}
	return nil
}

// mIRCColor represents a pre-parsed IRC color for fast distance calculations.
type mIRCColor struct {
	Code    string
	R, G, B int
}

// Precalculate the palette on first access and cache the returned slice.
var getMIRCPalette = sync.OnceValue(func() []mIRCColor {
	// 00-98 raw hex codes
	rawColors := []string{
		// 00-15
		"FFFFFF", "000000", "00007F", "009300", "FF0000", "7F0000", "9C009C", "FC7F00",
		"FFFF00", "00FC00", "009393", "00FFFF", "0000FC", "FF00FF", "7F7F7F", "D2D2D2",
		// 16-31
		"470000", "472100", "474700", "324700", "004700", "00472C", "004747", "002747",
		"000047", "2E0047", "470047", "47002A", "740000", "743A00", "747400", "517400",
		// 32-47
		"007400", "007449", "007474", "004074", "000074", "4B0074", "740074", "740045",
		"B50000", "B56300", "B5B500", "7DB500", "00B500", "00B571", "00B5B5", "0063B5",
		// 48-63
		"0000B5", "7500B5", "B500B5", "B5006B", "FF0000", "FF8C00", "FFFF00", "B2FF00",
		"00FF00", "00FFA0", "00FFFF", "008CFF", "0000FF", "A500FF", "FF00FF", "FF0098",
		// 64-79
		"FF5959", "FFB459", "FFFF71", "CFFF60", "6FFF6F", "65FFC9", "6DFFFF", "59B4FF",
		"5959FF", "C459FF", "FF66FF", "FF59BC", "FF9C9C", "FFD39C", "FFFF9C", "E2FF9C",
		// 80-95
		"9CFF9C", "9CFFDB", "9CFFFF", "9CD3FF", "9C9CFF", "DC9CFF", "FF9CFF", "FF94D3",
		"000000", "131313", "282828", "363636", "4D4D4D", "656565", "818181", "9F9F9F",
		// 96-98
		"BCBCBC", "E2E2E2", "FFFFFF",
	}

	// Use a map to track and exclude duplicates.
	// Because we iterate from 0 to 98, standard codes (0-15) are saved first.
	seenHex := make(map[string]bool)
	var palette []mIRCColor

	for i, hex := range rawColors {
		if seenHex[hex] {
			continue
		}
		seenHex[hex] = true

		r, _ := strconv.ParseInt(hex[0:2], 16, 32)
		g, _ := strconv.ParseInt(hex[2:4], 16, 32)
		b, _ := strconv.ParseInt(hex[4:6], 16, 32)

		palette = append(palette, mIRCColor{
			Code: fmt.Sprintf("%02d", i),
			R:    int(r), G: int(g), B: int(b),
		})
	}

	return palette
})

// FindClosestIRCColor uses the precalculated palette to quickly find the nearest match.
func findClosestIRCColor(hexColor string) string {
	palette := getMIRCPalette()

	hexColor = strings.ToUpper(strings.TrimPrefix(hexColor, "#"))
	if len(hexColor) != 6 {
		return "01"
	}

	r64, _ := strconv.ParseInt(hexColor[0:2], 16, 32)
	g64, _ := strconv.ParseInt(hexColor[2:4], 16, 32)
	b64, _ := strconv.ParseInt(hexColor[4:6], 16, 32)
	r1, g1, b1 := int(r64), int(g64), int(b64)

	minDist := math.MaxInt32
	bestCode := "01"

	for _, c := range palette {
		// Calculate the mean red level to adjust weights dynamically
		rMean := (r1 + c.R) / 2

		rDiff := r1 - c.R
		gDiff := g1 - c.G
		bDiff := b1 - c.B

		// "Redmean" perceptual distance approximation.
		// This heavily weights Green (which controls perceived luminosity) and
		// dynamically scales Red/Blue weighting based on how bright the red channel is.
		dist := (((512 + rMean) * rDiff * rDiff) >> 8) + (4 * gDiff * gDiff) + (((767 - rMean) * bDiff * bDiff) >> 8)

		if dist == 0 {
			return c.Code
		}

		if dist < minDist {
			minDist = dist
			bestCode = c.Code
		}
	}

	return bestCode
}

func mIRCExtendedColorFormatter(w io.Writer, style *chroma.Style, it iter.Seq[chroma.Token]) error {
	style = clearBackground(style)
	for token := range it {
		entry := style.Get(token.Type)
		if entry.IsZero() {
			if _, err := fmt.Fprint(w, token.Value); err != nil {
				return err
			}
			continue
		}

		formatting := ""
		if entry.Bold == chroma.Yes {
			formatting += "\x02"
		}
		if entry.Underline == chroma.Yes {
			formatting += "\x1f"
		}
		if entry.Italic == chroma.Yes {
			formatting += "\x1d"
		}
		if entry.Colour.IsSet() {
			hexColorCode := fmt.Sprintf("%02X%02X%02X", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
			formatting += fmt.Sprintf("\x03%s", findClosestIRCColor(hexColorCode))
			if entry.Background.IsSet() {
				bgHexColorCode := fmt.Sprintf("%02X%02X%02X", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
				formatting += fmt.Sprintf(",%s", findClosestIRCColor(bgHexColorCode))
			}
		}

		if err := mIRCwriteToken(w, formatting, token.Value); err != nil {
			return err
		}
	}
	return nil
}

func mIRCHexColorFormatter(w io.Writer, style *chroma.Style, it iter.Seq[chroma.Token]) error {
	style = clearBackground(style)
	for token := range it {
		entry := style.Get(token.Type)
		if entry.IsZero() {
			if _, err := fmt.Fprint(w, token.Value); err != nil {
				return err
			}
			continue
		}

		formatting := ""
		if entry.Bold == chroma.Yes {
			formatting += "\x02"
		}
		if entry.Underline == chroma.Yes {
			formatting += "\x1f"
		}
		if entry.Italic == chroma.Yes {
			formatting += "\x1d"
		}
		if entry.Colour.IsSet() {
			formatting += fmt.Sprintf("\x04%02X%02X%02X", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
			if entry.Background.IsSet() {
				formatting += fmt.Sprintf(",%02X%02X%02X", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
			}
		}

		if err := mIRCwriteToken(w, formatting, token.Value); err != nil {
			return err
		}
	}
	return nil
}
