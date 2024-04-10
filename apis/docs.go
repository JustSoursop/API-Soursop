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

func init() {
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
			theapp := []string{}
			applist := map[string]App{}

			endpoint := "/" + c.Params("*", "")

			if e, err := url.PathUnescape(endpoint); err == nil {
				endpoint = e
			}

			if app, ok := GetApp(endpoint); ok {
				applist[app.Path] = app
			} else {
				applist = apps.Apps
			}

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

			// c.Type("text/html", "utf-8")
			return c.SendString(strings.Join(theapp, "\n"))
		},
	}

	Register(docs)
	docs.Path = "/docs"
	Register(docs)
}
