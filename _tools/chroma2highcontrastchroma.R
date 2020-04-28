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

contrast <- function(color, background) {
  # https://www.w3.org/TR/WCAG20-TECHS/G17.html
  l1 <- calculate_luminance(color)
  l2 <- calculate_luminance(background)
  
  
  (min(l1, l2) + 0.05) / (max(l1, l2) + 0.05)
  

}

background <- "#f69289"
color <- "#ffffff"
contrast(color, background)