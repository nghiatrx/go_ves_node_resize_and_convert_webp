package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func handle(fileName string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	outputFile := fmt.Sprintf("./output-%d.webp", index)
	resizedWidth := 1000

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Resize the image
	resizedImg := resize.Resize(uint(resizedWidth), 0, img, resize.NearestNeighbor)

	var buf bytes.Buffer

	if err := jpeg.Encode(&buf, resizedImg, nil); err != nil {
		log.Fatal(err)
	}

	// Convert the resized JPG to WebP format
	outputWebpFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outputWebpFile.Close()

	if err := webp.Encode(outputWebpFile, resizedImg, &webp.Options{Quality: 80}); err != nil {
		log.Fatal(err)
	}

}

func main() {
	t1 := time.Now().UnixMilli()

	input := []string{
		"./input/input-0.jpg",
		"./input/input-1.jpg",
		"./input/input-2.jpg",
		"./input/input-3.jpg",
		"./input/input-4.jpg",
		"./input/input-5.jpg",
	}

	var wg sync.WaitGroup

	for index, file := range input {
		wg.Add(1)
		go handle(file, index, &wg)
	}

	wg.Wait()

	t2 := time.Now().UnixMilli()

	fmt.Println("time: ", t2-t1)
}
