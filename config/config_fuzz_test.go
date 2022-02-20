//go:build go1.18
// +build go1.18

package config

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func FuzzLoad(f *testing.F) {
	testcases := []string{
		"testdata/advanced.yml",
		"testdata/database.yml",
		"testdata/default.yml",
		"testdata/deprecated.yml",
	}

	for _, tc := range testcases {
		b, _ := ioutil.ReadFile(tc)
		f.Add(b)
	}

	f.Fuzz(func(t *testing.T, in []byte) {
		_, err := loadReader(bytes.NewReader(in))
		if err != nil {
			t.Skip()
		}
	})
}
