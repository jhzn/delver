package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCmd(t *testing.T) {
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
			name: "Tests that 'go test' input is correctly translated into the dlv equivalent",
			args: args{flags: []string{"-v", "-count=1", "-run", "'^TestSomeTestblabla$'", "./pkg/api/tests"}},
			want: []string{
				"dlv",
				"test",
				fmt.Sprintf(`--build-flags='%s' %s`, `-gcflags=-N`, "./pkg/api/tests"),
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
				t.Errorf("getCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
