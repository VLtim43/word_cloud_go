package main

import (
	"image/color"
	"image/png"
	"log"
)

func main() {
	words, err := readJson("./word.json")
	if err != nil {
		log.Fatalf("Failed to read words: %v", err)
	}

	// create the canvas file
	createdFile, err := createFile("./output", "test.png")
	if err != nil {
		log.Fatalf("Failed to create image file: %v", err)
	}
	defer createdFile.Close()

	whiteColor := color.RGBA{245, 245, 245, 255}
	canvas := writeText(createCanvas(640, 480, whiteColor), words)

	// encode the image as a png
	err = png.Encode(createdFile, canvas)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

}
