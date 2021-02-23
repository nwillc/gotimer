package typeface

// FontRune is a matrix of booleans that make up a `dot matrix` character in a font.
type FontRune [][]bool

// Font a map of rune's to their FontRune that make up a font.
type Font map[rune]FontRune
