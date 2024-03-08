package main

import (
	"reflect"
	"testing"
)

func Test_readData(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readData(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readData() = %v, want %v", got, tt.want)
			}
		})
	}
}
