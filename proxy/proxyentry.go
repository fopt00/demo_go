package proxy

import (
	"flag"
	"log"
)

func EntryOfProxy() {
	config := NewConfigByFlag()
	flag.Parse()
	WithAuth(config)

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetPrefix("[proxy-dump] ")
	log.Fatalln("exit: ", Run(config))
}


