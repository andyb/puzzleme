package main

import (
	//"./imageslice"
	"github.com/andyb/web"
	"log"
)

func index(ctx *web.Context, val string) {
	ctx.Redirect(301, "/index.html")
}

func slice(ctx *web.Context, val string) {
	ctx.WriteString("ok")
}

func image(ctx *web.Context, val string) {
	ctx.WriteImage("out/" + val + ".jpg")
	log.Println(val)
}

func main() {
	web.Get("/images/(.*)", image)
	//web.Get("/", index)
	web.Post("/slice/(.*)", slice)
	web.Run("0.0.0.0:8080")
}
