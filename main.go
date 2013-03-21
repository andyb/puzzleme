package main

import (
	"./imageslice"
	"github.com/andyb/web"
	"log"
)

func slice(ctx *web.Context) {
	rows := ctx.Params["r"]
	cols := ctx.Params["c"]
	log.Println(rows)
	log.Println(cols)
	if rows != "" && cols != "" {
		img := imageslice.LoadImage("testdata/gorilla.jpg")
		imageslice.SplitImagesAndSave(img, 3, 3)
		log.Println("image slicing complete")
	} else {
		ctx.Abort(400, "Invalid parameters. Please privde r and c")
	}
}

func image(ctx *web.Context, val string) {
	ctx.WriteImage("out/" + val + ".jpg")
	log.Println(val)
}

func main() {
	web.Get("/images/(.*)", image)
	web.Post("/slice", slice)
	web.Run("0.0.0.0:8080")
}
