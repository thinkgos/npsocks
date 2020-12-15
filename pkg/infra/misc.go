package infra

import (
	"log"
)

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
