library("magrittr")

# Find color passing --------------------

get_letter <- function(rgb, letter) {
  if (rgb@coords[1, letter] <= 0.03928) {
    return(rgb@coords[1, letter]/12.92)
  } else {
    return(((rgb@coords[1, letter] +0.055)/1.055) ^ 2.4)
  }
}

calculate_luminance <- function(hex) {
  # https://www.w3.org/TR/WCAG20-TECHS/G17.html
  
  rgb <- colorspace::hex2RGB(hex)
  
  R <- get_letter(rgb, "R")
  G <- get_letter(rgb, "G")
  B <- get_letter(rgb, "B")
  
  0.2126 * R + 0.7152 * G + 0.0722 * B
}

calculate_contrast <- function(color, background) {
  # https://www.w3.org/TR/WCAG20-TECHS/G17.html
  l1 <- calculate_luminance(color)
  l2 <- calculate_luminance(background)
  
  
  (max(l1, l2) + 0.05) / (min(l1, l2) + 0.05)
  

}

replace_color <- function(original_color, background, verbose = FALSE) {

  color <- original_color

  while (calculate_contrast(color, background) < 4.5) {
    if(verbose) {
      message(
        glue::glue(
          "still searching a replacement for {color} on {background}"
          )
        )
    }
    
    color <- colorspace::darken(color, space = "HCL")
  }
  
  if (verbose) {
    message(
      glue::glue(
        "The color {original_color} should be replaced with {color} to ensure a contrast of {calculate_contrast(color, background)}!"))
  }
    
  return(color)
}

# Transform style --------------------

extract_colors <- function(style_file) {
  style <- readLines(style_file)
  # original colors
  colors <- unique(
    tolower(
      stringr::str_extract(style, "\\#[:alnum:]*")
    )
  ) 
  colors <- colors[!is.na(colors)]
  purrr::map_chr(colors, lengthen_color)
}

lengthen_color <- function(color) {
  if (stringr::str_length(color) == 7) {
    return(color)
  } else {
    color <- glue::glue_collapse(c("#",
                      rep(substring(color, 2, 2), 2),
                      rep(substring(color, 3, 3), 2),
                      rep(substring(color, 4, 4), 2)))
    
    return(color)
  }
  
}

amend_style <- function(style_file, wcag = "AAA") {
  style <- readLines(style_file)
  
  
  
  bg_line <- style[grepl("chroma\\.Background", style)]
  bg <- stringr::str_extract(bg_line, "\\#[:alnum:]*")
  
  treat_line <- function(line, bg = bg) {
    # don't amend background
    if (!grepl("chroma\\.Background", line)) {
      # only amend lines with a color
      if (grepl('[\" ]\\#[:alnum:]*', line)) {
        # extract the color
        line_color <- trimws(
          stringr::str_remove(stringr::str_extract(
                                line, '[\" ]\\#[:alnum:]*'),
                              '"'
                              )
          )
        
        # if available extract the background
        if (grepl("bg\\:\\#[:alnum:]*", line)) {
          bg <- stringr::str_remove(
            stringr::str_extract(line, "bg\\:\\#[:alnum:]*"),
            "bg\\:"
            )
        }

        new_color <- replace_color(lengthen_color(color), 
                                   background = bg)
        line <- stringr::str_replace(line, color, new_color)
        
      }
    }
    
    return(line)
  }
  
  newlines <- purrr::map_chr(style, treat_line)
  
  name <- style %>%
    glue::glue_collapse() %>%
    stringr::str_extract("var .* \\=") %>%
    stringr::str_remove("var") %>%
    stringr::str_remove("\\=") %>%
    stringr::str_squish()
  
  newlines <- stringr::str_replace(name, paste0(name, "-hc"))
  newlines <- stringr::str_replace(tolower(name), 
                                   paste0(tolower(name), "-hc"))
  
  new_path <- file.path(styles,
                        paste0(tolower(name), "-hc.go"))
  writeLines(newlines, new_path)

  pals::pal.bands(original = extract_colors(style_file), 
                  new = extract_colors(new_path),
                  show.names = TRUE)
}



# Transform styleS --------------------

style_file <- file.path("styles", "monokai.go")
amend_style(style_file)
