package id49

import (
	"fmt"
	"testing"
)

//func TestId49(t *testing.T) {
//	tests := []struct {
//		name  string
//		input []string
//		want  [][]string
//	}{
//		{
//			"t1",
//			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
//			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
//		},
//		{
//			"t2",
//			[]string{},
//			[][]string{{}},
//		},
//		{
//			"t3",
//			[]string{"a"},
//			[][]string{{"a"}},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			output := groupAnagrams(tt.input)
//			if !reflect.DeepEqual(tt.want, output) {
//				t.Errorf("want:%v, output:%v", tt.want, output)
//			}
//		})
//	}
//}

func TestEncode(t *testing.T) {
	out := encode1("aaafjkfldfa")
	fmt.Println(out)
}
