package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func main() {

	whiteColor := color.RGBA{245, 245, 245, 255}

	// read the words list
	wordsWithWeights, err := readJson("./word.json")
	if err != nil {
		log.Fatalf("Failed to read words: %v", err)
	}

	// creates the fonts
	parsedFont, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}

	// writes on canvas
	canvas, err := writeText(createCanvas(640, 480, whiteColor), wordsWithWeights, parsedFont)
	if err != nil {
		log.Fatalf("Failed to draw words: %v", err)
	}

	// create the file
	createdFile, err := createFile("./output", "test.png")
	if err != nil {
		log.Fatalf("Failed to create image file: %v", err)
	}

	// encode the image as a png
	err = png.Encode(createdFile, canvas)
	if err != nil {
		createdFile.Close()
		log.Fatalf("Failed to encode image: %v", err)
	}
	if err := createdFile.Close(); err != nil {
		log.Fatalf("Failed to close image file: %v", err)
	}

	fmt.Println("Successfully created image at output/test.png")
}
