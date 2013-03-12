package main

import "./imageslice"
import "./server"

func main() {
	img := imageslice.LoadImage("testdata/img-1.jpg")
	imageslice.SplitImagesAndSave(img)
	server.Start()
}
