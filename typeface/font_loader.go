/*
 *  Copyright (c) 2021,  nwillc@gmail.com
 *
 *  Permission to use, copy, modify, and/or distribute this software for any
 *  purpose with or without fee is hereby granted, provided that the above
 *  copyright notice and this permission notice appear in all copies.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 *  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 *  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 *  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 *  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 *  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 *  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package typeface

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

const blackPixel = 35

//go:embed bitmaps/*/*.txt
var bitmaps embed.FS

// AvailableFonts is a map of available Font by name.
var AvailableFonts map[string]Font

// FontNames available in the app
var FontNames []string

func init() {
	FontNames = fontNames(bitmaps)
	AvailableFonts = make(map[string]Font)
	for _, name := range FontNames {
		m, err := readBitmaps(bitmaps, "bitmaps/"+name)
		if err != nil {
			panic("Could not load bitmaps " + err.Error())
		}
		AvailableFonts[name] = m
	}
}

func readBitmaps(fs embed.FS, path string) (Font, error) {
	runes := make(map[rune]FontRune)
	files, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		r, err := toFontRune(fs, path, file.Name())
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

func toFontRune(fs embed.FS, fontName string, name string) (FontRune, error) {
	txt, err := fs.ReadFile(fontName + "/" + name)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(txt), "\n")
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

func fontNames(fs embed.FS) []string {
	var names []string
	entries, _ := fs.ReadDir("bitmaps")
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		names = append(names, entry.Name())
	}
	return names
}
