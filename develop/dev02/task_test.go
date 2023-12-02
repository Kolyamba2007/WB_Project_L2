package main

import "testing"

func Test_unpack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "a4bc2d5e",
			args:    args{str: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "abcd",
			args:    args{str: "abcd"},
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "3abc",
			args:    args{str: "3abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "45",
			args:    args{str: "45"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "aaa10b",
			args:    args{str: "aaa10b"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "--5|3",
			args:    args{str: "--5|3"},
			want:    "------|||",
			wantErr: false,
		},
		{
			name:    "aaa0b",
			args:    args{str: "aaa0b"},
			want:    "aab",
			wantErr: false,
		},
		{
			name:    "",
			args:    args{str: ""},
			want:    "",
			wantErr: false,
		},
		{
			name:    "d\n5abc",
			args:    args{str: "d\n5abc"},
			want:    "d\n\n\n\n\nabc",
			wantErr: false,
		},
		{
			name:    "-",
			args:    args{str: "-"},
			want:    "-",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpack(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
