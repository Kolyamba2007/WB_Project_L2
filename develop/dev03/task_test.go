package main

import (
	"slices"
	"testing"
)

func Test_mansort(t *testing.T) {
	type args struct {
		line *[]string
		keys []Key
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestK",
			args: args{line: &[]string{
				"2 apple 3",
				"4 mango 5",
				"6 watermelon 2",
				"3 cherry 1",
				"1 orange 4",
				"5 banana 3",
				"5 cherry 2",
			},
				keys: []Key{&K{3}},
			},
			want: []string{
				"3 cherry 1",
				"5 cherry 2",
				"6 watermelon 2",
				"2 apple 3",
				"5 banana 3",
				"1 orange 4",
				"4 mango 5",
			},
		},
		{
			name: "TestU",
			args: args{line: &[]string{
				"23",
				"52",
				"69",
				"hoka",
				"52",
				"sadam",
				"21",
				"boba",
				"3",
				"5",
				"boba",
				"78",
			},
				keys: []Key{&U{}},
			},
			want: []string{
				"21",
				"23",
				"3",
				"5",
				"52",
				"69",
				"78",
				"boba",
				"hoka",
				"sadam",
			},
		},
		{
			name: "TestN",
			args: args{line: &[]string{
				"23",
				"2",
				"69",
				"6",
				"52",
				"21",
			},
				keys: []Key{&N{}},
			},
			want: []string{
				"2",
				"6",
				"21",
				"23",
				"52",
				"69",
			},
		},
		{
			name: "TestNU",
			args: args{line: &[]string{
				"23",
				"52",
				"69",
				"hoka",
				"52",
				"sadam",
				"21",
				"boba",
				"3",
				"5",
				"boba",
				"78",
			},
				keys: []Key{&N{}, &U{}},
			},
			want: []string{
				"boba",
				"hoka",
				"sadam",
				"3",
				"5",
				"21",
				"23",
				"52",
				"69",
				"78",
			},
		},
		{
			name: "TestR",
			args: args{line: &[]string{
				"1",
				"2",
				"3",
				"4",
			},
				keys: []Key{&R{}},
			},
			want: []string{
				"4",
				"3",
				"2",
				"1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mansort(tt.args.line, tt.args.keys...)

			if slices.Compare(*(tt.args.line), tt.want) != 0 {
				t.Errorf("mansort() error = %v, want = %v", *(tt.args.line), tt.want)
				return
			}
		})
	}
}
