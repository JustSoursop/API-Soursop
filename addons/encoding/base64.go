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

const (
	errDataRequired = "field `data` is required"
	errEmptyResult  = "empty result"
)

var author = apis.Author{
	Name: "Soursop",
	Contact: []apis.Contact{
		{
			Type: "github",
			Data: "github.com/JustSoursop",
		},
	},
}

func init() {
	apis.Register(base64Decoder())
	apis.Register(base64Encoder())
}

// base64Decoder returns a fiber.App that decodes a base64 encoded string.
func base64Decoder() apis.App {
	// Create a new apis.App with the provided parameters.
	return apis.App{
		Title:       "Base64 Decoder",
		Category:    "encoding",
		Tags:        []string{""},
		Author:      author,
		Description: "Decode base64 encoded into string",
		Version:     "1.0.0",
		Path:        "/b64_decode/:data?",
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
		Example:      "curl -X GET 'http://localhost:3000/b64_decode/SGVsbG8gd29ybGQgISE='",
		Handler: func(c *fiber.Ctx) error {

			// Create a new apis.Response to store the result.
			resp := apis.Response{
				Status: "error",
			}

			// Get the query parameters and store them in the datas slice.
			datas := []string{
				c.Query("data", ""),
			}

			// If there is a query parameter, add it to the datas slice.
			if data_q, err := url.QueryUnescape(c.Params("data", "")); err == nil {
				datas = append(datas, data_q)
			}

			// Get the data from the datas slice.
			data := apis.NotEmptyS(datas...)

			// If there is data, decode it.
			if len(data) > 0 {

				// Decode the base64 string.
				decoded, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					resp.Message = err.Error()
				} else {
					// If successful, set the status and response data.
					resp.Status = apis.ResponseSuccess
					resp.Data = string(decoded)
				}
			} else {
				// If no data was provided, return an error message.
				resp.Message = errDataRequired
			}

			// Return the JSON response.
			return c.JSON(resp)
		},
	}
}

// base64Encoder returns a base64 encoder App.
func base64Encoder() apis.App {
	// Create an App with the given parameters.
	return apis.App{
		Title:       "Base64 Encoder",
		Category:    "encoding",
		Tags:        []string{},
		Author:      author,
		Description: "Encode string into base64 encoded string",
		Version:     "1.0.0",
		Path:        "/b64_encode/:data?",
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
		Example:      "curl -X GET 'http://localhost:3000/b64_encode/Hello%20world%20!!'",
		Handler: func(c *fiber.Ctx) error {

			// Create a Response with the given parameters.
			resp := apis.Response{
				Status: apis.ResponseError,
			}

			// Get the query parameters.
			datas := []string{
				c.Query("data", ""),
			}

			// Get the path parameters.
			if data_q, err := url.QueryUnescape(c.Params("data", "")); err == nil {
				datas = append(datas, data_q)
			}

			// Combine the parameters.
			data := apis.NotEmptyS(datas...)

			// Check if the data isn't empty.
			if len(data) > 0 {

				// Encode the data.
				encoded := base64.StdEncoding.EncodeToString([]byte(data))
				// Check if the encoding was successful.
				if len(encoded) <= 0 {
					// If not, set the response message.
					resp.Message = errEmptyResult
				} else {
					// If so, set the response status and data.
					resp.Status = apis.ResponseSuccess
					resp.Data = string(encoded)
				}
			} else {
				// If the data is empty, set the response message.
				resp.Message = errDataRequired
			}

			// Return the JSON response.
			return c.JSON(resp)
		},
	}

}
