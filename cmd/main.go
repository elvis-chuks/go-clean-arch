package main

import (
	"flag"
	"fmt"
	httpDelivery "github.com/elvis-chuks/go-clean-arch/delivery/http"
	"github.com/elvis-chuks/go-clean-arch/pkg/env"
	"github.com/elvis-chuks/go-clean-arch/pkg/logger"
	"github.com/elvis-chuks/go-clean-arch/repository/mongodb"
	"log"
)

func main() {
	l, err := logger.Init()

	if err != nil {
		panic(err)
	}

	env.LoadConfig("../")

	repo := mongodb.New(l) // can be sql, cassandra, postgresql etc

	httpConfig := httpDelivery.Config{
		UserRepo: repo.UserRepo,
	}

	app := httpDelivery.RunHttpServer(httpConfig)

	addr := flag.String("addr", fmt.Sprintf(":%s", "5001"), "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
