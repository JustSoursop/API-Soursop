// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package apis

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var LISTEN = ":3000"
var API = fiber.New()

func StartAPI() {
	API.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path())
		return c.Next()
	})

	if err := API.Listen(LISTEN); err != nil {
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

func NotEmptyS(s ...string) string {
	for _, v := range s {
		if len(v) > 0 {
			return v
		}
	}

	return ""
}

func init() {
	godotenv.Load()

	LISTEN = ":" + EnvOr("PORT", "3000")

}
