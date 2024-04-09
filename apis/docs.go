// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package apis

import (
	"fmt"
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
		Path:        "/docs",
		Method:      fiber.MethodGet,
		Params: []Param{
			{
				Name:     "app",
				Type:     "string",
				Required: false,
				Example:  "/b64/encode",
			},
		},
		SuccessCode:  200,
		ResponseType: "text/plain utf-8",
		Example:      "docs",
		Execute: func(c *fiber.Ctx) error {
			theapp := []string{}

			// endpoint := c.Params("app")
			// if app, ok := GetApp(endpoint); ok {
			// 	theapp = append(theapp,
			// 		fmt.Sprintln(app.Method, app.Path, "params", len(app.Params)),
			// 	)
			// }

			for p, a := range apps.Apps {
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
