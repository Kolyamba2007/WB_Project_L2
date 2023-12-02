package main

import (
	"reflect"
	"testing"
)

func TestGrepContext(t *testing.T) {
	type args struct {
		line *[]string
		key  ContextKey
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestA2",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &After{n: 2, str: "watermelon"},
			},
			want: []string{
				"watermelon",
				"cherry",
				"orangeapp",
			},
		},
		{
			name: "TestB2",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &Before{n: 2, str: "watermelon"},
			},
			want: []string{
				"apple",
				"mangoApP",
				"watermelon",
			},
		},
		{
			name: "TestC2",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &Context{n: 2, str: "watermelon"},
			},
			want: []string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrepContext(tt.args.line, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrepContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepCount(t *testing.T) {
	type args struct {
		line *[]string
		key  CountKey
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestCount",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &Count{"cherry"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrepCount(tt.args.line, tt.args.key); got != tt.want {
				t.Errorf("GrepCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepText(t *testing.T) {
	type args struct {
		line *[]string
		key  TextKey
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestI",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &IgnoreCase{str: "app"},
			},
			want: []string{
				"apple",
				"mangoApP",
				"orangeapp",
			},
		},
		{
			name: "TestV",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &Invert{str: "app"},
			},
			want: []string{
				"mangoApP",
				"watermelon",
				"cherry",
				"cherry",
				"banana",
			},
		},
		{
			name: "TestF",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &Fixed{str: "app"},
			},
			want: []string{
				"apple",
				"orangeapp",
			},
		},
		{
			name: "TestN",
			args: args{line: &[]string{
				"apple",
				"mangoApP",
				"watermelon",
				"cherry",
				"orangeapp",
				"cherry",
				"banana",
			},
				key: &LineNum{str: "app"},
			},
			want: []string{
				"1:apple",
				"5:orangeapp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrepText(tt.args.line, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrepText() = %v, want %v", got, tt.want)
			}
		})
	}
}
