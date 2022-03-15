package main

import (
	"flag"
	"math/rand"
	"time"

	"22dojo-online/pkg/server"

	"github.com/pkg/profile"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	server.Serve(addr)
}
