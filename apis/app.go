// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package apis

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResponseStatus string

const (
	ResponseSuccess ResponseStatus = "success"
	ResponseError   ResponseStatus = "error"
)

type Response struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
}

type Param struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Example     string `json:"example"`
	Description string `json:"description"`
}

type Contact struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Author struct {
	Name        string `json:"name"`
	Contact     []Contact
	Description string `json:"description"`
}

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

	Execute fiber.Handler `json:"-"`
}

func (a *App) Register(f *fiber.App) error {
	switch a.Method {
	case fiber.MethodGet:
		f.Get(a.Path, a.Execute)

	case fiber.MethodPost:
		f.Post(a.Path, a.Execute)

	default:
		return fmt.Errorf("method unsupported")
	}

	return nil
}

type AppList struct {
	Apps map[string]App `json:"apps"`
}

var apps = AppList{
	Apps: map[string]App{},
}

func Register(app App) error {
	if _, ok := apps.Apps[app.Path]; ok {
		return fmt.Errorf("app already exists")
	} else {
		apps.Apps[app.Path] = app
		app.Register(API)

		return nil
	}
}

func GetApp(path string) (App, bool) {
	a, ok := apps.Apps[path]

	return a, ok
}
