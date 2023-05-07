package main

import (
	_ "github.com/devmeireles/go-vue-image-evaluate/app/docs"
)

// @title Image Evaluate API
// @version 1.0
// @description A quickly API setup to evaulate avatar.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
func main() {
	setupDatabase()
	setupRoutes()
}
