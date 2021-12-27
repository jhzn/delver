package main

import (
	"reflect"
	"testing"
)

func Test_getFlags(t *testing.T) {
	type args struct {
		flags []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{flags: []string{"-v", "-count=1", "-run", "'^TestSomeTestblabla$'", "./pkg/api/tests"}},
			want: []string{
				"dlv",
				"test",
				"--build-flags=./pkg/api/tests",
				"--",
				"-test.v",
				"-test.count=1",
				"-test.run",
				"'^TestSomeTestblabla$'",
				"./pkg/api/tests",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCmd(tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFlags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
