package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/caiguanhao/opencc"
)

func main() {
	showDicts := flag.Bool("dicts", false, "show list of dictionaries and exit")
	dict := flag.String("dict", "s2t", "specify which dictionary to use")
	flag.Parse()
	if *showDicts {
		for dict, text := range opencc.Dictionaries {
			fmt.Printf("% -5s - %s\n", dict, text)
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
