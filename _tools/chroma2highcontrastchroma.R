library("magrittr")

# Find color passing --------------------

get_letter <- function(rgb, letter) {
  if (rgb@coords[1, letter] <= 0.03928) {
    return(rgb@coords[1, letter]/12.92)
  } else {
    return(((rgb@coords[1, letter] +0.055)/1.055) ^ 2.4)
  }
}

.calculate_luminance <- function(hex) {
  # https://www.w3.org/TR/WCAG20-TECHS/G17.html
  
  rgb <- colorspace::hex2RGB(hex)
  
  R <- get_letter(rgb, "R")
  G <- get_letter(rgb, "G")
  B <- get_letter(rgb, "B")
  
  0.2126 * R + 0.7152 * G + 0.0722 * B
}

calculate_luminance <- memoise::memoise(.calculate_luminance)

calculate_contrast <- function(color, background) {
  # https://www.w3.org/TR/WCAG20-TECHS/G17.html
  l1 <- calculate_luminance(color)
  l2 <- calculate_luminance(background)
  
  
  (max(l1, l2) + 0.05) / (min(l1, l2) + 0.05)
  

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

replace_color <- function(original_color, background, wcag, verbose = FALSE) {

  if(wcag == "AAA") {
    threshold <- 7
  } else{
    threshold <- 4.5
  }
  
  original_color <- lengthen_color(original_color)
  background <- lengthen_color(background)
  
  color <- original_color

  while (calculate_contrast(color, background) < threshold) {
    if(verbose) {
      message(
        glue::glue(
          "still searching a replacement for {color} on {background}"
          )
        )
    }
    
    if (calculate_luminance(original_color) > calculate_luminance(background)) {
      if (calculate_luminance(color) == 1) {
        message(glue::glue("We cannot do better than {color} for {original_color} on {background}"))
        return(color)
      }
      color <- colorspace::lighten(color, space = "HCL")
    } else {
      color <- colorspace::darken(color, space = "HCL")
      if (calculate_luminance(color) == 0) {
        message(glue::glue("We cannot do better than {color} for {original_color} on {background}"))
        return(color)
      }
    }
    
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
  colors <- (
    tolower(
      stringr::str_extract(style, "\\#[:alnum:]*")
    )
  ) 
  colors <- colors[!is.na(colors)]
  purrr::map_chr(colors, lengthen_color)
}

treat_line <- function(line, bg = bg, wcag) {

  # special case of e.g. solarized-light
  line <- gsub("bg\\: ", "bg:", line)
  
  # don't amend background
  if (!grepl("chroma\\.Background", line)) {
    # only amend lines with a color
    if (grepl('[\" ]\\#[:alnum:]*', line)) {
      # extract the color
      line_color <- trimws(
        stringr::str_remove(
          stringr::str_extract(
            line, '[\" ]\\#[:alnum:]*'),
          '"'
        )
      )
      
      # if available extract the background
      if (grepl("bg\\:\\#[:alnum:]*", line)) {
        background <- stringr::str_remove(
          stringr::str_extract(line, "bg\\:\\#[:alnum:]*"),
          "bg\\:"
        )
      } else {
        background <- bg
      }

      new_color <- replace_color(line_color, 
                                 background = background,
                                 wcag = wcag,
                                 verbose = FALSE)
      
      line <- stringr::str_replace(line, line_color, new_color)
      
    }
  }
  
  return(line)
}

amend_style <- function(style_file, wcag = "AAA") {
  print(style_file)
  style <- readLines(style_file)

  bg_line <- style[grepl("chroma\\.Background", style)]
  
  if (length(bg_line) > 0) {
    if (grepl("bg\\:", bg_line)) {
      bg <- stringr::str_extract(bg_line, "bg\\:\\#[:alnum:]*") %>%
        stringr::str_remove("bg\\:")
    } else {
      bg <- stringr::str_extract(bg_line, "\\#[:alnum:]*")
    }
  } else {
    bg <- "#fff" # pygments has no background
  }

  newlines <- purrr::map_chr(style, treat_line,
                             bg = bg, wcag = wcag)
  
  if (all(newlines == style)) {
    message("The theme is already compliant!")
    return(invisible())
  }
  
  name <- style %>%
    glue::glue_collapse() %>%
    stringr::str_extract("var .* \\=") %>%
    stringr::str_remove("var") %>%
    stringr::str_remove("\\=") %>%
    stringr::str_squish()
  
  newlines <- stringr::str_replace(newlines, name, paste0(name, "-hc"))
  
  var <- style %>%
    glue::glue_collapse() %>%
    stringr::str_extract('MustNewStyle\\(".*?"') %>%
    stringr::str_remove('MustNewStyle\\("') %>%
    stringr::str_remove('"')

  newlines <- stringr::str_replace(newlines, 
                                   paste0('MustNewStyle\\("', var), 
                                   paste0('MustNewStyle\\("', var, "-hc"))

  if (wcag == "AAA") {
    new_path <- file.path("styles",
                          paste0(var, "-wcag-aaa.go"))
    
  } else {
    
    new_path <- file.path("styles",
                          paste0(tolower(name), "-wcag-aa.go"))
  }
  writeLines(newlines, new_path)

  pals::pal.bands(original = extract_colors(style_file), 
                  new = extract_colors(new_path),
                  show.names = TRUE,
                  main = name)
}



# Transform styleS --------------------
style_file <- file.path("styles", "monokai.go")

style_files <- fs::dir_ls("styles")
style_files <- style_files[style_files != file.path("styles", "api.go")]
style_files <- style_files[style_files != file.path("styles", "swapoff.go")]
style_files <- style_files[!grepl("wcag", style_files)]

# purrr::walk(style_files, amend_style, wcag = "AA")
purrr::walk(style_files, amend_style, wcag = "AAA")
