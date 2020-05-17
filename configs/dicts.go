package configs

import (
	"strings"

	"github.com/liuzl/da"
)

type Dicts [][]*da.Dict

func (dicts Dicts) Convert(in string) string {
	var token string
	for _, group := range dicts {
		r := []rune(in)
		var tokens []string
		for i := 0; i < len(r); {
			s := r[i:]
			max := 0
			for _, dict := range group {
				ret, err := dict.PrefixMatch(string(s))
				if err != nil {
					continue
				}
				if len(ret) > 0 {
					o := ""
					for k, v := range ret {
						if len(k) > max {
							max = len(k)
							token = v[0]
							o = k
						}
					}
					i += len([]rune(o))
					break
				}
			}
			if max == 0 {
				token = string(r[i])
				i++
			}
			tokens = append(tokens, token)
		}
		in = strings.Join(tokens, "")
	}
	return in
}
