// main_test.go
package main

import (
	"testing"
)

func TestArgs(t *testing.T) {
	if err := argcheck("", "", ""); err == nil {
		t.Errorf("引数なし")
	}
	if err := argcheck("infile", "", ""); err == nil {
		t.Errorf("アウトファイルの指定なし")
	}
	if err := argcheck("", "", "outfile"); err == nil {
		t.Errorf("ファイル、ディレクトリの指定なし")
	}
	if err := argcheck("infile", "dir", "outfile"); err == nil {
		t.Errorf("ファイル、ディレクトリ両方が指定")
	}
}
