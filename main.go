// Copyright (C) 2024 Soursop
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at <https://mozilla.org/MPL/2.0/>.

package main

import (
	_ "api-soursop/addons"
	"api-soursop/apis"
)

func main() {

	apis.StartAPI()
}
