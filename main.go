package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/namsral/flag"
)

var (
	db  *bolt.DB
	cfg Config
)

func main() {
	var (
		version bool
		config  string
		dbpath  string
		title   string
		fqdn    string
		bind    string
		url     string
	)

	flag.BoolVar(&version, "v", false, "display version information")

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&dbpath, "dbpath", "search.db", "Database path")
	flag.StringVar(&title, "title", "Search", "OpenSearch Title")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&fqdn, "fqdn", "gadget", "FQDN for public access")
	flag.StringVar(&url, "url", DefaultURL, "default url to redirect to")

	flag.Parse()

	if version {
		fmt.Println(FullVersion())
		os.Exit(0)
	}

	cfg.Title = title
	cfg.FQDN = fqdn
	cfg.URL = url

	var err error
	db, err = bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = EnsureDefaultBookmarks()
	if err != nil {
		log.Fatal(err)
	}

	NewServer(bind, cfg).ListenAndServe()
}
