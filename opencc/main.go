package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/caiguanhao/opencc"
)

func main() {
	showDicts := flag.Bool("dicts", false, "show list of dictionaries and exit")
	dict := flag.String("dict", "s2t", "specify which dictionary to use")
	flag.Parse()
	if *showDicts {
		names := []string{}
		nameMaxLen := 0
		for name := range opencc.Dictionaries {
			names = append(names, name)
			if len(name) > nameMaxLen {
				nameMaxLen = len(name)
			}
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf(fmt.Sprintf("%% -%ds - %%s\n", nameMaxLen), name, opencc.Dictionaries[name])
		}
		return
	}
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		content, _ := ioutil.ReadAll(os.Stdin)
		fmt.Println(opencc.Convert(*dict, string(content)))
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			fmt.Println(opencc.Convert(*dict, text))
		}
	}
}
