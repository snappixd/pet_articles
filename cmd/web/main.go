package main

import "articles/internal/app"

const cfgDir = "configs"

func main() {
	app.Run(cfgDir)
}
