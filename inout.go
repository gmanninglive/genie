package main

import (
	"os"
)

func WriteFile(c Command, template []byte) {
	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		err := os.MkdirAll(c.Directory, 0700)
		Check(err)
	}

	err := os.WriteFile(c.Output, []byte(template), 0644)
	Check(err)
}
