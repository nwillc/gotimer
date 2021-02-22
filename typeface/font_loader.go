package typeface

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

const blackPixel = 35

//go:embed bitmaps/medium/*.txt
var bitmaps embed.FS

// Medium sized font
var Medium Font

func init() {
	m, err := readBitmaps(bitmaps, "bitmaps/medium")
	if err != nil {
		panic("Could not load bitmaps")
	}
	Medium = m
}
func readBitmaps(fs embed.FS, path string) (Font, error) {
	runes := make(map[rune]FontRune)
	files, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		r, err := toFontRune(fs, file.Name())
		if err != nil {
			return nil, err
		}
		name, err := toCharName(file.Name())
		if err != nil {
			return nil, err
		}
		runes[name] = r
	}
	return runes, nil
}


func toFontRune(fs embed.FS, name string) (FontRune, error) {
	txt, err := fs.ReadFile("bitmaps" + "/medium/" + name)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(txt),"\n")
	var fr [][]bool
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		var l []bool
		for _, c := range line {
			l = append(l, c == blackPixel)
		}
		fr = append(fr, l)
	}
	return fr, nil
}

func toCharName(path string) (rune, error) {
	parts := strings.Split(path, ".")
	if len(parts) != 2 {
		return rune(-1), fmt.Errorf("malformed parts %d", len(parts))
	}
	name := (parts[0])[1:]
	ascii, err := strconv.Atoi(name)
	if err != nil {
		return rune(-1), err
	}
	r := rune(ascii)
	return r, nil
}
