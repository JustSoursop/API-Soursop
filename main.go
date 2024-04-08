// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var LISTEN = ":3000"
var APP = fiber.New()

func StartAPI() {
	APP.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path())
		return c.Next()
	})

	if err := APP.Listen(LISTEN); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on http://0.0.0.0%s\n", LISTEN)
	}
}

func EnvOr(key, alt string) string {
	e := os.Getenv(key)
	if e == "" {
		return alt
	} else {
		return e
	}
}

func init() {
	godotenv.Load()
	LISTEN = ":" + EnvOr("PORT", "3000")
}

func main() {

	StartAPI()
}
