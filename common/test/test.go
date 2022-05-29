package id49

import (
	"reflect"
	"testing"
)

func TestID(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			"t1",
			[]string{},
			[][]string{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ""
			if !reflect.DeepEqual(tt.want, output) {
				t.Errorf("want:%v, output:%v", tt.want, output)
			}
		})
	}
}
