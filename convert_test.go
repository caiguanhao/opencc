package opencc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestAllCases(t *testing.T) {
	files, err := filepath.Glob("OpenCC-src/test/testcases/*.in")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		base := filepath.Base(file)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		fmt.Println("testing", name)
		expected, err := ioutil.ReadFile("OpenCC-src/test/testcases/" + name + ".ans")
		if err != nil {
			panic(err)
		}
		actual := Convert(name, string(content))
		if actual != string(expected) {
			t.Fatal("bad case for", name)
		}
	}
}
