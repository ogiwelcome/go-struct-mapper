package gostructmap

import "testing"

func TestMappingFrom(t *testing.T) {
	type args struct {
		fromObj interface{}
		toObj   interface{}
	}
	type hoge struct {
		A string `map:"B"`
	}
	type piyo struct {
		B string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success for nil",
			args: args{
				fromObj: nil,
				toObj:   nil,
			},
			wantErr: true,
		},
		{
			name: "success string",
			args: args{
				fromObj: &hoge{
					A: "test",
				},
				toObj: &piyo{
					B: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MappingFrom(tt.args.fromObj, tt.args.toObj); (err != nil) != tt.wantErr {
				t.Errorf("MappingFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
