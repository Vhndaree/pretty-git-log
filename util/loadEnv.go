package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ParseAndSetEnv - set env variables
func ParseAndSetEnv(fName string) {
	if fName == "" {
		fName = ".env"
	}
	f := readFile(fName)
	s := strings.Split(f, "\n")

	for _, v := range s {
		e := strings.Split(v, "=")
		os.Setenv(e[0], e[1])
	}
}

func readFile(fName string) string {
	fHandler, err := os.Open(fName)
	defer fHandler.Close()

	if err != nil {
		fmt.Printf("Error while opening %s file: ", fName)
		log.Fatal(err)
	}

	f, err := ioutil.ReadAll(fHandler)
	if err != nil {
		fmt.Printf("Error while reading %s file: ", fName)
		log.Fatal(err)
	}

	return string(f)
}
