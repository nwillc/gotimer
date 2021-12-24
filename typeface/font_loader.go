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
	"github.com/nwillc/genfuncs"
	"io/fs"
	"strconv"
	"strings"
)

const blackPixel = 35

var (
	//go:embed bitmaps/*/*.txt
	bitmaps embed.FS

	// AvailableFonts is a map of available Font by name.
	AvailableFonts map[string]Font

	// FontNames available in the app
	FontNames []string
)

func init() {
	FontNames = fontNames(bitmaps)
	AvailableFonts = genfuncs.AssociateWith(FontNames, func(n string) Font { return readBitmaps(bitmaps, "bitmaps/"+n) })
}

func readBitmaps(embedFs embed.FS, path string) Font {
	files, err := embedFs.ReadDir(path)
	if err != nil {
		panic(err)
	}
	return genfuncs.Associate(files, func(f fs.DirEntry) (rune, FontRune) {
		return toCharName(f.Name()), toFontRune(embedFs, path, f.Name())
	})
}

func toFontRune(fs embed.FS, fontName string, name string) FontRune {
	txt, err := fs.ReadFile(fontName + "/" + name)
	if err != nil {
		panic(err)
	}
	return genfuncs.Map(
		genfuncs.Map(
			genfuncs.Filter(
				strings.Split(string(txt), "\n"),
				func(l string) bool { return len(l) > 2 }),
			func(s string) []rune { return []rune(s) }),
		func(rs []rune) []bool { return genfuncs.Map(rs, func(r rune) bool { return r == blackPixel }) })
}

func toCharName(path string) rune {
	parts := strings.Split(path, ".")
	if len(parts) != 2 {
		panic(fmt.Errorf("malformed parts %d", len(parts)))
	}
	name := (parts[0])[1:]
	ascii, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}
	r := rune(ascii)
	return r
}

func fontNames(efs embed.FS) []string {
	entries, _ := efs.ReadDir("bitmaps")
	return genfuncs.Map(
		genfuncs.Filter(
			entries, func(e fs.DirEntry) bool { return e.IsDir() }),
		func(e fs.DirEntry) string { return e.Name() })
}
