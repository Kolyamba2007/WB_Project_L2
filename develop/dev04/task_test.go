package main

import (
	"slices"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		words *[]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "По возрастанию",
			args: args{words: &[]string{
				"слиток",
				"тяпка",
				"листок",
				"пятак",
			},
			},
			want: map[string][]string{
				"слиток": {"листок", "слиток"},
				"тяпка":  {"пятак", "тяпка"},
			},
		},
		{
			name: "Из одного элемента",
			args: args{words: &[]string{
				"слиток",
				"тяпка",
				"слиток",
				"пятак",
			},
			},
			want: map[string][]string{
				"тяпка": {"пятак", "тяпка"},
			},
		},
		{
			name: "Нижний регистр у слов множества",
			args: args{words: &[]string{
				"СЛиток",
				"тяпка",
				"листок",
				"пЯтАк",
			},
			},
			want: map[string][]string{
				"СЛиток": {"листок", "слиток"},
				"тяпка":  {"пятак", "тяпка"},
			},
		},
		{
			name: "Слово встречается один раз",
			args: args{words: &[]string{
				"банан",
				"БАНАН",
				"НаБаН",
				"набан",
			},
			},
			want: map[string][]string{
				"банан": {"банан", "набан"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.args.words)

			for key, value := range *got {
				if slices.Compare(tt.want[key], *value) != 0 {
					t.Errorf("Anagrams with %v key = %v, want = %v", key, *value, tt.want[key])
				}
			}
		})
	}
}
