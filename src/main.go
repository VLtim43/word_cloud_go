package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func main() {

	canvasColor := color.RGBA{245, 245, 245, 255}
	canvas := createCanvas(640, 480, canvasColor)

	// read the words list
	words, err := readJson("./word.json")
	if err != nil {
		log.Fatalf("Failed to read words: %v", err)
	}

	// creates the fonts
	typeface, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}

	occupiedeArea := []image.Rectangle{}

	// place each word into the canvas keeping in mind the previously occupied areas
	for _, word := range words {
		rect, err := placeWordRectagle(canvas, word, typeface, occupiedeArea)
		if err != nil {
			log.Fatalf("Failed to place %q: %v", word.Text, err)
		}
		occupiedeArea = append(occupiedeArea, rect)

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
