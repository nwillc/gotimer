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

package generics

import "fmt"

var _ fmt.Stringer = (*Result[int])(nil)

// Result is an implementation of the Maybe pattern. This is mostly for experimentation as it is a poor fit with Go's
// traditional idiomatic error handling.
type Result[T any] struct {
	value T
	err   error
}

func NewResult[T any](t T) *Result[T] {
	return &Result[T]{value: t}
}

func NewError[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

func (r *Result[T]) Error() error {
	return r.err
}

func (r *Result[T]) Ok() bool {
	return r.err == nil
}

func (r *Result[T]) OnFailure(action func(e error)) *Result[T] {
	if !r.Ok() {
		action(r.err)
	}
	return r
}

func (r *Result[T]) OnSuccess(action func(t T)) *Result[T] {
	if r.Ok() {
		action(r.value)
	}
	return r
}

func (r *Result[T]) String() string {
	if r.Ok() {
		return fmt.Sprint(r.value)
	}

	return "error: " + r.err.Error()
}

func (r *Result[T]) Then(action func(t T) *Result[T]) *Result[T] {
	if r.Ok() {
		return action(r.value)
	}
	return r
}

func (r *Result[T]) ValueOr(v T) T {
	if r.Ok() {
		return r.value
	}
	return v
}

func (r *Result[T]) ValueOrPanic() T {
	if r.Ok() {
		return r.value
	}
	panic(r.err)
}
