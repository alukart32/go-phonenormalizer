package main

import (
	"flag"
	"log"

	"alukart32.com/phoneNormalizer/config"
	"alukart32.com/phoneNormalizer/internal/app"
)

var (
	cfg = flag.String("cfg", "../../config/config.yaml", "The database connection and normalizer config")
)

func main() {
	flag.Parse()

	fail := func(err error) {
		log.Fatalf("phone normalizer :: %v", err)
	}

	// Get config
	cfg, err := config.New(*cfg)
	if err != nil {
		fail(err)
	}

	// Run app
	if err = app.Run(cfg); err != nil {
		fail(err)
	}
}
