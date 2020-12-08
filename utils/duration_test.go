package utils

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
			got, err := Format(tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
