package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/jackxunkun163/XrayR/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
