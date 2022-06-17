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

package display

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Zero",
			args:    args{duration: 0 * time.Second},
			want:    "00",
			wantErr: false,
		},
		{
			name:    "OneSecond",
			args:    args{duration: 1 * time.Second},
			want:    "01",
			wantErr: false,
		},
		{
			name:    "OneMinute",
			args:    args{duration: 1 * time.Minute},
			want:    "01.00",
			wantErr: false,
		},
		{
			name:    "OneHour",
			args:    args{duration: 1 * time.Hour},
			want:    "01:00.00",
			wantErr: false,
		},
		{
			name:    "FiftyNineSeconds",
			args:    args{duration: 59 * time.Second},
			want:    "59",
			wantErr: false,
		},
		{
			name:    "ElevenMinutesFiveSeconds",
			args:    args{duration: 11*time.Minute + 5*time.Second},
			want:    "11.05",
			wantErr: false,
		},
		{
			name:    "OneHourElevenSeconds",
			args:    args{duration: time.Hour + 11*time.Second},
			want:    "01:00.11",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Format(tt.args.duration)
			if (!got.Ok()) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", got.Error(), tt.wantErr)
				return
			}
			if got.OrElse("") != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
