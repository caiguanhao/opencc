// this file is generated by convert_test.go

package s2hk

import (
	"github.com/caiguanhao/opencc/configs"
	"github.com/caiguanhao/opencc/dicts/HKVariants"
	"github.com/caiguanhao/opencc/dicts/STCharacters"
	"github.com/caiguanhao/opencc/dicts/STPhrases"
)

const (
	Description = "Simplified Chinese to Traditional Chinese (Hong Kong standard)"
)

var (
	Dicts = configs.Dicts{
		{
			STPhrases.Dict,
			STCharacters.Dict,
		},
		{
			HKVariants.Dict,
		},
	}
)