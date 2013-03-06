package imageslice

import "image"
import "image/jpeg"
import "image/draw"
import "os"
import "log"
import "strconv"

type ImageSlice struct {
	img         image.Image
	sliceNumber int
	totalSlices int
	fileName    string
}

/*load the image to be sliced*/
func LoadImage(fileName string) image.Image {
	// Open the file.
	file, err := os.Open(fileName)
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
	return img
}

/* Split the image into the required number of slices and save out */
func SplitImagesAndSave(img image.Image) {
	chans := map[int]chan int{}
	totalSlices := 4
	for i := 1; i <= totalSlices; i++ {
		chans[i] = make(chan int)
		go chopImage(ImageSlice{img, i, totalSlices, "temp"}, chans[i])
	}

	//loop through the channels and block until channel for each returns
	for i := 1; i <= totalSlices; i++ {
		<-chans[i]
	}
}

/* Responsible for taking the image and the current slice and create a slice image. Will notify when complete via the channel */
func chopImage(slice ImageSlice, completed chan int) {
	//need to divide the image up by the number of slices. 
	rect := slice.img.Bounds()
	size := rect.Size()

	sliceImg := image.NewRGBA(image.Rect(0, 0, size.X/2, size.Y/2))
	draw.Draw(sliceImg, sliceImg.Bounds(), slice.img, image.ZP, draw.Src)
	toSave, err := os.Create("out/" + slice.fileName + strconv.Itoa(slice.sliceNumber) + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer toSave.Close()

	jpeg.Encode(toSave, sliceImg, &jpeg.Options{jpeg.DefaultQuality})
	log.Printf("Chop Chop %v", slice.sliceNumber)
	completed <- slice.sliceNumber
}
