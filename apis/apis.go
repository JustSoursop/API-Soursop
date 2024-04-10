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
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// This is the listen address for the API
var LISTEN = ":3000"

// This is the API instance
var API = fiber.New()

// This function starts the API
func StartAPI() {
	// Use the logger with the default configuration
	API.Use(logger.New(logger.ConfigDefault))

	// Start the API and listen on the specified listen address
	if err := API.Listen(LISTEN); err != nil {
		// If an error occurs, log it and exit
		log.Fatal(err)
	} else {
		// If successful, log the server started on the specified listen address
		log.Printf("Server started on http://0.0.0.0%s\n", LISTEN)
	}
}

// This function returns the environment variable for a given key, or the alternative given
func EnvOr(key, alt string) string {
	// Get the environment variable
	e := os.Getenv(key)
	// If the environment variable is empty, return the alternative
	if e == "" {
		return alt
	} else {
		return e
	}
}

// This function returns the first non-empty string from the given slice of strings
func NotEmptyS(s ...string) string {
	// Iterate through the given slice of strings
	for _, v := range s {
		// If the length of the string is greater than 0, return it
		if len(v) > 0 {
			return v
		}
	}

	// If no non-empty string is found, return an empty string
	return ""
}

// This function loads the environment variables from the .env file
func init() {
	// Load the environment variables from the .env file
	godotenv.Load()

	// Set the listen address to the PORT environment variable, or 3000 if not set
	LISTEN = ":" + EnvOr("PORT", "3000")

}
