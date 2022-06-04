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

package generics_test

import (
	"fmt"
	"github.com/nwillc/gotimer/generics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult(t *testing.T) {
	var r *generics.Result[int]
	err := fmt.Errorf("ro ruh")
	r = generics.NewError[int](err)
	r.
		OnFailure(func(e error) {
			assert.Equal(t, err, e)
		}).
		OnSuccess(func(_ int) {
			assert.Fail(t, "success on an error")
		})

	assert.Panics(t, func() {
		r.ValueOrPanic()
	})

	assert.Equal(t, 10, r.ValueOr(10))

	r = generics.NewResult(10)
	assert.Equal(t, 10, r.ValueOrPanic())

	r = r.Then(func(i int) *generics.Result[int] {
		return generics.NewResult(i * 10)
	})
	assert.Equal(t, 100, r.ValueOrPanic())
}
