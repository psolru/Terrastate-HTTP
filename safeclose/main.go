package safeclose

import (
	"io"
	"log"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
