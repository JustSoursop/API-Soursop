// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package apis

import (
	"fmt"
	"path"

	"github.com/gofiber/fiber/v2"
)

// ResponseStatus is a string representing the status of a response
type ResponseStatus string

const (
	ResponseSuccess ResponseStatus = "success"
	ResponseError   ResponseStatus = "error"
)

// Response is a struct representing a response
type Response struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Type    string         `json:"type,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
}

// Param is a struct representing a parameter
type Param struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Example     string `json:"example"`
	Description string `json:"description"`
}

// Contact is a struct representing contact information
type Contact struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// Author is a struct representing an author
type Author struct {
	Name        string `json:"name"`
	Contact     []Contact
	Description string `json:"description"`
}

// App is a struct representing an API
type App struct {
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`

	Author      Author `json:"author"`
	Description string `json:"description"`
	Version     string `json:"version"`

	Path   string  `json:"path"`
	Method string  `json:"method"`
	Params []Param `json:"params"`

	SuccessCode  int    `json:"success_code"`
	ResponseType string `json:"response_type"`

	Example string `json:"example"`

	Handler fiber.Handler `json:"-"`
}

// Register registers an API with a given fiber app
func (a *App) Register(f *fiber.App) error {
	switch a.Method {
	case fiber.MethodGet:
		f.Get(a.Path, a.Handler)

	case fiber.MethodPost:
		f.Post(a.Path, a.Handler)

	case fiber.MethodDelete:
		f.Delete(a.Path, a.Handler)

	case fiber.MethodPut:
		f.Put(a.Path, a.Handler)

	case fiber.MethodPatch:
		f.Patch(a.Path, a.Handler)

	default:
		return fmt.Errorf("method unsupported")
	}

	return nil
}

// AppList is a struct representing a list of APIs
type AppList struct {
	Apps map[string]App `json:"apps"`
}

var apps = AppList{
	Apps: map[string]App{},
}

// Register registers an API with the global list of APIs
func Register(app App) error {

	app_path := path.Join(app.Category, app.Path)
	if _, ok := apps.Apps[app_path]; ok {
		return fmt.Errorf("app already exists")
	} else {
		apps.Apps[app_path] = app
		app.Register(API)

		return nil
	}
}

// GetApp retrieves an API from the global list of APIs
func GetApp(path string) (App, bool) {
	a, ok := apps.Apps[path]

	return a, ok
}
