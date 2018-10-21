package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-pg/migrations"
	"github.com/jqs7/zwei/db"
)

func main() {
	flag.Parse()
	oldVer, newVer, err := migrations.Run(db.Instance(), flag.Args()...)
	if err != nil {
		log.Fatalln(err)
	}
	if oldVer != newVer {
		fmt.Printf("migrated from version %d to %d\n", oldVer, newVer)
	} else {
		fmt.Printf("version is %d\n", oldVer)
	}
}
