package files

import (
	"fmt"
	"testing"
)

func TestIterDirectory(t *testing.T) {
	type args struct {
		dirPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := IterDirectory("/tmp/stargate.lock")
			fmt.Println(b)
		})
	}
}
