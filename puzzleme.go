package main

import "image"
import "image/jpeg"
import "image/draw"
import "os"
import "log"
import "time"

func main() {
	LoadImage()
}

func LoadImage() {
	// Open the file.
	file, err := os.Open("testdata/img-1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("Load Image Completed")
	// Decode the image.
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	chans := map[int]chan int{}
	totalSlices := 4
	for i := 1; i <= totalSlices; i++ {
		//need to work out how to block whie these complete? Buffered channel??
		chans[i] = make(chan int)
		go ChopImage(img, i, totalSlices, chans[i])
	}

	//loop through the channels and block until channel for each returns
	for i := 1; i <= totalSlices; i++ {
		<-chans[i]
	}

	WriteImageToFile()
}

func ChopImage(img image.Image, sliceNumber int, totalSlices int, completed chan int) {
	//need to divide the image up by the number of slices. 

	sliceImg := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(sliceImg, sliceImg.Bounds(), img, image.ZP, draw.Src)
	toSave, _ := os.Create("out/temp.jpg")
	jpeg.Encode(toSave, sliceImg, nil)
	log.Printf("Chop Chop %v", sliceNumber)
	completed <- sliceNumber
}

func WriteImageToFile() {
	log.Println("Writing image to file")
}
