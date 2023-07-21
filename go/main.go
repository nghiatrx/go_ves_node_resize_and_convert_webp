package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
)

func handle(fileName string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	outputFile := fmt.Sprintf("./output/output-%d.webp", index)
	resizedWidth := 1000.0

	// Load the input image
	inputImage, err := vips.NewImageFromFile(fileName)
	if err != nil {
		log.Fatal(fmt.Errorf("error loading input image: %w", err))
	}

	// Resize the image
	err = inputImage.Resize(resizedWidth/(float64(inputImage.Width())), vips.KernelAuto)
	b, _, _ := inputImage.ExportWebp(&vips.WebpExportParams{Quality: 80})

	err = ioutil.WriteFile(outputFile, b, 755)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	os.Setenv("VIPS_DEBUG", "0")
	t1 := time.Now().UnixMilli()

	vips.Startup(nil)

	defer vips.Shutdown()

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
