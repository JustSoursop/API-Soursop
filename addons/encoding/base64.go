// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package encoding

import (
	"api-soursop/apis"
	"encoding/base64"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func init() {
	apis.Register(base64Decoder())
	apis.Register(base64Encoder())
}

func base64Decoder() apis.App {
	return apis.App{
		Title:       "Base64 Decoder",
		Category:    "encoding",
		Tags:        []string{},
		Author:      apis.Author{},
		Description: "",
		Version:     "1.0.0",
		Path:        "/b64/decode/:data?",
		Method:      fiber.MethodGet,
		Params: []apis.Param{
			{
				Name:     "data",
				Type:     "string",
				Required: true,
				Example:  "SGVsbG8gd29ybGQgISE=",
			},
		},
		SuccessCode:  200,
		ResponseType: "application/json",
		Example:      "",
		Execute: func(c *fiber.Ctx) error {

			resp := apis.Response{
				Status: "error",
			}

			datas := []string{
				c.Query("data", ""),
			}

			if data_q, err := url.QueryUnescape(c.Params("data", "")); err == nil {
				datas = append(datas, data_q)
			}

			data := apis.NotEmptyS(datas...)

			if len(data) > 0 {

				decoded, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					resp.Message = err.Error()
				} else {
					resp.Status = "success"
					resp.Data = string(decoded)
				}
			} else {
				resp.Message = "Data is required"
			}

			return c.JSON(resp)
		},
	}
}

func base64Encoder() apis.App {
	return apis.App{
		Title:       "Base64 Encoder",
		Category:    "encoding",
		Tags:        []string{},
		Author:      apis.Author{},
		Description: "",
		Version:     "1.0.0",
		Path:        "/b64/encode/:data?",
		Method:      fiber.MethodGet,
		Params: []apis.Param{
			{
				Name:     "data",
				Type:     "string",
				Required: true,
				Example:  "Hello world !!",
			},
		},
		SuccessCode:  200,
		ResponseType: "application/json",
		Example:      "",
		Execute: func(c *fiber.Ctx) error {

			resp := apis.Response{
				Status: apis.ResponseError,
			}

			datas := []string{
				c.Query("data", ""),
			}

			if data_q, err := url.QueryUnescape(c.Params("data", "")); err == nil {
				datas = append(datas, data_q)
			}

			data := apis.NotEmptyS(datas...)

			if len(data) > 0 {

				encoded := base64.StdEncoding.EncodeToString([]byte(data))
				if len(encoded) <= 0 {
					resp.Message = "Encoding failed"
				} else {
					resp.Status = apis.ResponseSuccess
					resp.Data = string(encoded)
				}
			} else {
				resp.Message = "Field `data` is required"
			}

			return c.JSON(resp)
		},
	}
}
