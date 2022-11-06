package gostructmap

import (
	"reflect"
	"testing"
)

func TestMappingFrom(t *testing.T) {
	type args struct {
		c interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "success for nil",
			args: args{
				c: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MappingFrom(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MappingFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
