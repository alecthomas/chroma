// Command automerge approves CI and enables auto-merge for pull requests that
// only change lexer rules, style entries or lexer test data.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing/fstest"
	"time"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/repr"
)

const (
	bot           = "github-actions[bot]"
	goTestsMarker = "<!-- auto-merge:go-tests -->"
)

var (
	lexerPath    = regexp.MustCompile(`^lexers/embedded/[^/]+\.xml$`)
	stylePath    = regexp.MustCompile(`^styles/[^/]+\.xml$`)
	testdataPath = regexp.MustCompile(`^lexers/testdata/.+\.(actual|expected)$`)
)

type pullRequest struct {
	repo      string
	number    string
	head      string
	mergeBase string
	content   func(ctx context.Context, path, ref string) (string, error)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("automerge: ")
	ctx := context.Background()
	changedFiles, err := strconv.Atoi(mustEnv("CHANGED_FILES"))
	if err != nil {
		log.Fatalf("CHANGED_FILES: %s", err)
	}
	pr := &pullRequest{repo: mustEnv("REPO"), number: mustEnv("PR"), head: mustEnv("HEAD_SHA")}
	pr.content = pr.fetchContent

	reasons, err := pr.ineligibleReasons(ctx, mustEnv("BASE_SHA"), changedFiles)
	if err != nil {
		pr.disableAutoMerge(ctx)
		log.Fatal(err)
	}
	if len(reasons) > 0 {
		fmt.Printf("Not eligible: %s\n", strings.Join(reasons, "; "))
		pr.disableAutoMerge(ctx)
		return
	}

	fmt.Println("Eligible: approving CI and enabling auto-merge")
	if err := pr.approveCIRuns(ctx); err != nil {
		log.Fatal(err)
	}
	if _, err := gh(ctx, "pr", "merge", pr.number, "--repo", pr.repo, "--auto", "--squash", "--match-head-commit", pr.head); err != nil {
		log.Fatal(err)
	}
}

func mustEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s is not set", name)
	}
	return value
}

func gh(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "gh", args...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("gh %s: %w", strings.Join(args, " "), err)
	}
	return string(out), nil
}

func (pr *pullRequest) fetchContent(ctx context.Context, path, ref string) (string, error) {
	segments := strings.Split(path, "/")
	for i, segment := range segments {
		segments[i] = url.PathEscape(segment)
	}
	endpoint := fmt.Sprintf("repos/%s/contents/%s?ref=%s", pr.repo, strings.Join(segments, "/"), url.QueryEscape(ref))
	return gh(ctx, "api", "-H", "Accept: application/vnd.github.raw", endpoint)
}

func (pr *pullRequest) ineligibleReasons(ctx context.Context, baseSHA string, changedFiles int) ([]string, error) {
	mergeBase, err := gh(ctx, "api", fmt.Sprintf("repos/%s/compare/%s...%s", pr.repo, baseSHA, pr.head), "--jq", ".merge_base_commit.sha")
	if err != nil {
		return nil, err
	}
	pr.mergeBase = strings.TrimSpace(mergeBase)

	files, err := pr.files(ctx)
	if err != nil {
		return nil, err
	}
	// The files API silently truncates at 3000 entries.
	if len(files) != changedFiles {
		return []string{fmt.Sprintf("only %d of %d changed files could be listed", len(files), changedFiles)}, nil
	}
	var goTests, reasons []string
	for _, file := range files {
		if file.Status == "added" && strings.HasSuffix(file.Filename, "_test.go") {
			goTests = append(goTests, file.Filename)
			continue
		}
		reason, err := pr.fileReason(ctx, file.Status, file.Filename)
		if err != nil {
			return nil, err
		}
		if reason != "" {
			reasons = append(reasons, reason)
		}
	}
	if len(goTests) == 0 || len(reasons) > 0 {
		return reasons, nil
	}
	if err := pr.commentOnGoTests(ctx, goTests); err != nil {
		return nil, err
	}
	for _, path := range goTests {
		reasons = append(reasons, path+" adds a Go test")
	}
	return reasons, nil
}

type changedFile struct {
	Status   string `json:"status"`
	Filename string `json:"filename"`
}

func (pr *pullRequest) files(ctx context.Context) ([]changedFile, error) {
	out, err := gh(ctx, "api", "--paginate", fmt.Sprintf("repos/%s/pulls/%s/files", pr.repo, pr.number), "--jq", ".[] | {status, filename}")
	if err != nil {
		return nil, err
	}
	var files []changedFile
	dec := json.NewDecoder(strings.NewReader(out))
	for {
		var file changedFile
		err := dec.Decode(&file)
		if errors.Is(err, io.EOF) {
			return files, nil
		}
		if err != nil {
			return nil, fmt.Errorf("decoding files: %w", err)
		}
		files = append(files, file)
	}
}

func (pr *pullRequest) fileReason(ctx context.Context, status, path string) (string, error) {
	switch {
	case testdataPath.MatchString(path):
		if status != "added" && status != "modified" {
			return fmt.Sprintf("%s is %s", path, status), nil
		}
		return "", nil
	case lexerPath.MatchString(path):
		return pr.metadataReason(ctx, status, path, "lexer", "changes lexer config")
	case stylePath.MatchString(path):
		return pr.metadataReason(ctx, status, path, "style", "changes style attributes")
	default:
		return path + " is not a lexer, style or testdata file", nil
	}
}

func (pr *pullRequest) metadataReason(ctx context.Context, status, path, kind, change string) (string, error) {
	if status != "modified" {
		return fmt.Sprintf("%s is %s", path, status), nil
	}
	var loaded [2]string
	for i, ref := range []string{pr.mergeBase, pr.head} {
		content, err := pr.content(ctx, path, ref)
		if err != nil {
			return "", err
		}
		loaded[i], err = metadata(kind, content)
		if err != nil {
			return "", fmt.Errorf("%s at %s: %w", path, ref, err)
		}
	}
	if loaded[0] != loaded[1] {
		return fmt.Sprintf("%s %s", path, change), nil
	}
	return "", nil
}

// metadata uses Chroma's own parsers so the check cannot diverge from what Chroma loads.
func metadata(kind, content string) (string, error) {
	switch kind {
	case "lexer":
		lexer, err := chroma.NewXMLLexer(fstest.MapFS{"lexer.xml": {Data: []byte(content)}}, "lexer.xml")
		if err != nil {
			return "", fmt.Errorf("lexer: %w", err)
		}
		return repr.String(*lexer.Config()), nil
	case "style":
		style, err := chroma.NewXMLStyle(strings.NewReader(content))
		if err != nil {
			return "", fmt.Errorf("style: %w", err)
		}
		return repr.String([]string{style.Name, style.Counterpart}), nil
	default:
		return "", fmt.Errorf("unknown kind %q", kind)
	}
}

func (pr *pullRequest) commentOnGoTests(ctx context.Context, paths []string) error {
	jq := fmt.Sprintf(".[] | select(.user.login == %q) | .body", bot)
	comments, err := gh(ctx, "api", "--paginate", fmt.Sprintf("repos/%s/issues/%s/comments", pr.repo, pr.number), "--jq", jq)
	if err != nil {
		return err
	}
	if strings.Contains(comments, goTestsMarker) {
		return nil
	}
	body := &strings.Builder{}
	body.WriteString(goTestsMarker + "\n" +
		"Lexer and style changes should be tested with `lexers/testdata` files, not Go tests. " +
		"Add a `.actual` input, generate its `.expected` output with " +
		"`RECORD=true go test ./lexers`, and remove:\n\n")
	for _, path := range paths {
		fmt.Fprintf(body, "- `%s`\n", path)
	}
	_, err = gh(ctx, "pr", "comment", pr.number, "--repo", pr.repo, "--body", body.String())
	return err
}

func (pr *pullRequest) disableAutoMerge(ctx context.Context) {
	enabledBy, err := gh(ctx, "api", fmt.Sprintf("repos/%s/pulls/%s", pr.repo, pr.number), "--jq", ".auto_merge.enabled_by.login // empty")
	if err != nil {
		log.Printf("checking auto-merge: %s", err)
		return
	}
	if strings.TrimSpace(enabledBy) != bot {
		return
	}
	if _, err := gh(ctx, "pr", "merge", pr.number, "--repo", pr.repo, "--disable-auto"); err != nil {
		log.Printf("disabling auto-merge: %s", err)
	}
}

func (pr *pullRequest) approveCIRuns(ctx context.Context) error {
	endpoint := fmt.Sprintf("repos/%s/actions/runs?event=pull_request&head_sha=%s", pr.repo, pr.head)
	var runs string
	// CI runs for this push are created alongside this one, so may not be listed yet.
	for range 12 {
		var err error
		runs, err = gh(ctx, "api", endpoint, "--jq", ".workflow_runs[] | [.id, .status] | @tsv")
		if err != nil {
			return err
		}
		if runs != "" {
			break
		}
		time.Sleep(5 * time.Second)
	}
	for line := range strings.Lines(runs) {
		id, status, _ := strings.Cut(strings.TrimSpace(line), "\t")
		if status != "action_required" {
			continue
		}
		fmt.Printf("Approving CI run %s\n", id)
		if _, err := gh(ctx, "api", "-X", "POST", fmt.Sprintf("repos/%s/actions/runs/%s/approve", pr.repo, id)); err != nil {
			return err
		}
	}
	return nil
}
