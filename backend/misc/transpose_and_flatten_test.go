package misc

import (
	"reflect"
	"testing"
)

func TestTransposeAndFlatten(t *testing.T) {
	type args struct {
		input [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"first case",
			args{input: [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3", "b4"}, {"c1"}, {"d1", "d2"}, {"e1"}}},
			[]string{"a1", "b1", "c1", "d1", "e1", "a2", "b2", "d2", "a3", "b3", "b4"},
		},
		{
			"already flattened and transposed",
			args{input: [][]string{{"a1"}, {"b1"}, {"c1"}, {"d1"}}},
			[]string{"a1", "b1", "c1", "d1"},
		},
		{
			"empty matrix",
			args{input: [][]string{{}, {}, {}, {}}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeAndFlatten(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeAndFlatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
