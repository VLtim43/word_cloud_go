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

func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {

	// creates the background image
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	return canvas
}

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

	whiteColor := color.RGBA{245, 245, 245, 255}
	canvas := createCanvas(640, 480, whiteColor)

	// encode the image as a png
	err = png.Encode(file, canvas)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	log.Println("Successfully created image")

}
