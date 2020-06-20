package main

import (
	"flag"
	"os"

	"github.com/maateen/dockohealer/internal/watcher"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

var buildTime string
var gitSHA string
var versionString string

func main() {
	versionPtr := flag.Bool("version", false, "Show version information.")
	flag.Parse()

	if *versionPtr {
		log.WithFields(log.Fields{
			"buildTime": buildTime,
			"gitSHA":    gitSHA,
			"version":   versionString,
		}).Info()
		os.Exit(0)
	}
	watcher.Watch()
}
