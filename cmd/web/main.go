package main

import "articles/internal/app"

const cfgDir = "internal/config"

func main() {
	app.Run(cfgDir)
}
