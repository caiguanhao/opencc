package dicts

import (
	"bytes"
	"compress/flate"

	"github.com/liuzl/cedar-go"
)

func LoadCedarFromGOB(data []byte) *cedar.Cedar {
	c := cedar.New()
	r := flate.NewReader(bytes.NewReader(data))
	defer r.Close()
	if err := c.Load(r, "gob"); err != nil {
		panic(err)
	}
	return c
}
