// this file is generated by convert_test.go

package s2twp

import (
	"github.com/caiguanhao/opencc/configs"
	"github.com/caiguanhao/opencc/dicts/STCharacters"
	"github.com/caiguanhao/opencc/dicts/STPhrases"
	"github.com/caiguanhao/opencc/dicts/TWPhrases"
	"github.com/caiguanhao/opencc/dicts/TWVariants"
)

const (
	Description = "Simplified Chinese to Traditional Chinese (Taiwan standard, with phrases)"
)

var (
	Dicts = configs.Dicts{
		{
			STPhrases.Dict,
			STCharacters.Dict,
		},
		{
			TWPhrases.Dict,
		},
		{
			TWVariants.Dict,
		},
	}
)