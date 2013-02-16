package main

//import "image"
import "os"
import "log"

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
}
