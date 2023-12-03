package main

import (
	"reflect"
	"testing"
)

func TestMancut(t *testing.T) {
	type args struct {
		line      *[]string
		fields    []int
		delimiter string
		separated bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
		wantErr    bool
	}{
		{
			name: "Test1",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about politeness in others",
				},
				fields:    []int{1, 4},
				delimiter: " ",
				separated: false,
			},
			wantResult: []string{
				"He particular",
				"about others",
			},
			wantErr: false,
		},
		{
			name: "Test2",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about politeness in others",
				},
				fields:    []int{5},
				delimiter: " ",
				separated: false,
			},
			wantResult: nil,
			wantErr:    false,
		},
		{
			name: "Test3",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about politeness in others",
				},
				fields:    []int{2, 5},
				delimiter: " ",
				separated: false,
			},
			wantResult: []string{
				"was",
				"politeness",
			},
			wantErr: false,
		},
		{
			name: "Test4",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{2, 5},
				delimiter: " ",
				separated: false,
			},
			wantResult: []string{
				"was",
				"about6politeness6in6others",
			},
			wantErr: false,
		},
		{
			name: "Test5",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{5},
				delimiter: " ",
				separated: false,
			},
			wantResult: []string{
				"about6politeness6in6others",
			},
			wantErr: false,
		},
		{
			name: "Test6",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{5},
				delimiter: " ",
				separated: true,
			},
			wantResult: nil,
			wantErr:    false,
		},
		{
			name: "Test7",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{2},
				delimiter: "!",
				separated: true,
			},
			wantResult: nil,
			wantErr:    false,
		},
		{
			name: "Test8",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{2},
				delimiter: "!",
				separated: false,
			},
			wantResult: []string{
				"He was extraordinarily particular",
				"about6politeness6in6others",
			},
			wantErr: false,
		},
		{
			name: "Test9",
			args: args{
				line: &[]string{
					"He was extraordinarily particular",
					"about6politeness6in6others",
				},
				fields:    []int{3},
				delimiter: " ",
				separated: true,
			},
			wantResult: []string{
				"extraordinarily",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Mancut(tt.args.line, tt.args.fields, tt.args.delimiter, tt.args.separated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mancut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Mancut() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
