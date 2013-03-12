package imageslice

import "image"
import "image/jpeg"
import "image/draw"
import "os"
import "log"
import "strconv"
import "sync"

type ImageSlice struct {
	img         image.Image
	sliceNumber int
	totalSlices int
	fileName    string
	row         int
	col         int
}

type ImageSliceSize struct {
	width  int
	height int
}

func GetSize(img image.Image, rows int, cols int) ImageSliceSize {
	rect := img.Bounds()
	size := rect.Size()
	return ImageSliceSize{(size.X / 2), (size.Y / 2)}
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
	var wg sync.WaitGroup
	rows := 2
	cols := 2
	size := GetSize(img, rows, cols)
	totalSlices := rows * cols
	curCol := 0
	curRow := 0
	for i := 1; i <= totalSlices; i++ {
		wg.Add(1)
		chopImage(ImageSlice{img, i, totalSlices, "temp", curRow, curCol}, size, &wg)
		curCol++
		if curCol%cols == 0 {
			curCol = 0
			curRow++
		}
	}

	wg.Wait()
	log.Println("Work completed")
}

/* Responsible for taking the image and the current slice and create a slice image. Will notify when complete via the channel */
func chopImage(slice ImageSlice, size ImageSliceSize, wg *sync.WaitGroup) {
	xStart := size.width * slice.col
	yStart := size.height * slice.row
	//log.Print("yStart:" + strconv.Itoa(yStart))
	log.Print("xstart:" + strconv.Itoa(xStart))
	log.Print("x:" + strconv.Itoa(xStart+size.width))
	sliceImg := image.NewRGBA(image.Rect(0, 0, size.width, size.height))
	draw.Draw(sliceImg, sliceImg.Bounds(), slice.img, image.Pt(xStart, yStart), draw.Src)
	toSave, err := os.Create("out/" + slice.fileName + strconv.Itoa(slice.sliceNumber) + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer toSave.Close()

	jpeg.Encode(toSave, sliceImg, &jpeg.Options{jpeg.DefaultQuality})
	log.Printf("Chop Chop %v", slice.sliceNumber)
	wg.Done()
}
