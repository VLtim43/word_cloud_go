package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {

	destination := "./output"
	fileName := "test.png"

	// creates the folder
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}

	// creates the full path with the filename and the destination path
	var destinationPath = filepath.Join(destination, fileName)

	// creates the file
	file, err := os.Create(destinationPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Printf("Successfully created image at %v", destinationPath)

	// set canvas attributes
	canvasWidth := 640
	canvasLength := 480

	// creates the image
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	// encode the image as a png
	err = png.Encode(file, canvas)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	log.Println("Successfully created image")

}
