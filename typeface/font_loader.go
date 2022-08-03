/*
 *  Copyright (c) 2022,  nwillc@gmail.com
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
	"github.com/nwillc/genfuncs/container/gslices"
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/nwillc/genfuncs/results"
	"io/fs"
	"strconv"
	"strings"

	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

const blackPixel = 35

var (
	//go:embed bitmaps/*/*.txt
	bitmaps embed.FS

	// AvailableFonts is a map of available Font by name.
	AvailableFonts container.GMap[string, Font]

	// FontNames available in the app
	FontNames container.GSlice[string]

	entryIsDir   genfuncs.Function[fs.DirEntry, bool]   = func(e fs.DirEntry) bool { return e.IsDir() }
	entryName    genfuncs.Function[fs.DirEntry, string] = func(e fs.DirEntry) string { return e.Name() }
	hasData      genfuncs.Function[string, bool]        = func(l string) bool { return len(l) > 2 }
	toRuneSlice  genfuncs.Function[string, []rune]      = func(s string) []rune { return []rune(s) }
	toPixel      genfuncs.Function[rune, bool]          = func(r rune) bool { return r == blackPixel }
	toPixelSlice genfuncs.Function[[]rune, []bool]      = func(rs []rune) []bool { return gslices.Map(rs, toPixel) }
	toFont       maps.ValueFor[string, Font]            = func(n string) *genfuncs.Result[Font] { return readBitmaps(bitmaps, "bitmaps/"+n) }
)

func init() {
	FontNames = fontNames(bitmaps)
	AvailableFonts = sequences.AssociateWith[string, Font](FontNames, toFont).MustGet()
}

func readBitmaps(embedFs embed.FS, path string) *genfuncs.Result[Font] {
	result := genfuncs.NewResultError(embedFs.ReadDir(path))
	var files container.GSlice[fs.DirEntry] = result.MustGet()
	fMap := sequences.Associate[fs.DirEntry, rune, FontRune](files, func(f fs.DirEntry) *genfuncs.Result[*maps.Entry[rune, FontRune]] {
		return genfuncs.NewResult[*maps.Entry[rune, FontRune]](&maps.Entry[rune, FontRune]{Key: toCharName(f.Name()), Value: toFontRune(embedFs, path, f.Name()).MustGet()})
	})
	return genfuncs.NewResult(Font(fMap.MustGet()))
}

func toFontRune(fs embed.FS, fontName string, name string) *genfuncs.Result[FontRune] {
	return results.Map(
		results.Map(
			genfuncs.NewResultError(fs.ReadFile(fontName+"/"+name)),
			func(b []byte) *genfuncs.Result[container.GSlice[string]] {
				var lll container.GSlice[string] = strings.Split(string(b), "\n")
				fl := lll.Filter(hasData)
				return genfuncs.NewResult(fl)
			}),
		func(lines container.GSlice[string]) *genfuncs.Result[FontRune] {
			return genfuncs.NewResult[FontRune]([][]bool(gslices.Map(gslices.Map(lines, toRuneSlice), toPixelSlice)))
		},
	)
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

func fontNames(efs embed.FS) container.GSlice[string] {
	var entries container.GSlice[fs.DirEntry]
	entries, _ = efs.ReadDir("bitmaps")
	return gslices.Map(entries.Filter(entryIsDir), entryName)
}
