// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package apis

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Initialize the application
func init() {
	// Define the App struct
	docs := App{
		Title:       "Docs API",
		Category:    "",
		Tags:        []string{},
		Author:      Author{},
		Description: "",
		Version:     "",
		Path:        "/docs/*",
		Method:      fiber.MethodGet,
		Params: []Param{
			{
				Name:     "*",
				Type:     "string",
				Required: false,
				Example:  "/base64/encode",
			},
		},
		SuccessCode:  200,
		ResponseType: "text/plain utf-8",
		Example:      "docs",
		Handler: func(c *fiber.Ctx) error {
			// Create a list of apps
			theapp := []string{}
			applist := map[string]App{}

			// Get the endpoint from the URL parameters
			endpoint := "/" + c.Params("*", "")

			// Unescape the endpoint
			if e, err := url.PathUnescape(endpoint); err == nil {
				endpoint = e
			}

			// Check if the app exists in the current list of apps
			if app, ok := GetApp(endpoint); ok {
				// If so, add it to the applist
				applist[app.Path] = app
			} else {
				// If not, use the global list of apps
				applist = apps.Apps
			}

			// Iterate through the applist and create a string for each app
			for p, a := range applist {
				params := []string{}
				for _, p := range a.Params {
					params = append(params, fmt.Sprintf("%s:%s", p.Name, p.Type))
				}

				params_text := ""

				if len(params) > 0 {
					params_text = fmt.Sprintf("\n  params :(%s)", strings.Join(params, ", "))
				}

				theapp = append(theapp,
					fmt.Sprintln(a.Method, p, params_text),
				)
			}

			// Return the list of apps
			return c.SendString(strings.Join(theapp, "\n"))
		},
	}

	// Register the docs app
	Register(docs)
	docs.Path = "/docs"
	// Register the docs app again with the updated path
	Register(docs)
}
