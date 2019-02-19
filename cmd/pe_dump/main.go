package main

import (
	"flag"
	"log"

	"github.com/kr/pretty"
	"github.com/mewmew/pe"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, pePath := range flag.Args() {
		if err := parse(pePath); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func parse(pePath string) error {
	file, err := pe.ParseFile(pePath)
	if err != nil {
		return errors.WithStack(err)
	}
	file.Content = nil
	pretty.Println(file)
	return nil
}
