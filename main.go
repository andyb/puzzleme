package main

import "./imageslice"

func main() {
	img := imageslice.LoadImage("testdata/img-1.jpg")
	imageslice.SplitImagesAndSave(img)
}
