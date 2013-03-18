package main

import (
	"github.com/andyb/web"
	"log"
)

func index() string { return "hello" }

func image(ctx *web.Context, val string) {
	ctx.WriteImage("out/" + val + ".jpg")
	log.Println(val)
}

func main() {
	web.Get("/images/(.*)", image)
	web.Get("/", index)
	web.Run("0.0.0.0:8080")
}
