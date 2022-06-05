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

package display

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"strings"
	"time"
)

// Format a time.Duration into a string for the format `HH:MM.SS`.
func Format(d time.Duration) *genfuncs.Result[string] {
	var sb strings.Builder

	if d >= time.Hour {
		sb.WriteString(fmt.Sprintf("%02d:", int(d.Hours())))
	}

	if d >= time.Minute {
		d %= time.Hour
		sb.WriteString(fmt.Sprintf("%02d.", int(d.Minutes())))
	}

	d %= time.Minute
	sb.WriteString(fmt.Sprintf("%02d", int(d.Seconds())))
	return genfuncs.NewResult(sb.String())
}
