package main

import "image"
import "image/jpeg"
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

	chopChan := make(chan int, 3)

	for i := 1; i <= 4; i++ {
		//need to work out how to block whie these complete? Buffered channel??
		go ChopImage(img, i, chopChan)
	}

	<-chopChan
	WriteImageToFile()
}

func ChopImage(img image.Image, sliceNumber int, completed chan int) {
	//img.Bounds(image.Rectangle)
	time.Sleep(3 * time.Second)
	log.Println("Chop chop ")
	completed <- sliceNumber
}

func WriteImageToFile() {
	log.Println("Writing image to file")
}
