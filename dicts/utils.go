package dicts

import (
	"bytes"

	"github.com/liuzl/cedar-go"
)

func loadCedarFromGOB(data []byte) *cedar.Cedar {
	c := cedar.New()
	r := bytes.NewReader(data)
	if err := c.Load(r, "gob"); err != nil {
		panic(err)
	}
	return c
}
