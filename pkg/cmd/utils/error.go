package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
	os.Exit(1)
}