package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"

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

	occupiedArea := []image.Rectangle{}

	// place each word into the canvas keeping in mind the previously occupied areas
	for _, word := range words {
		rectangleCoordinates, face, err := createWordRectangle(word, typeface)
		if err != nil {
			log.Fatalf("Failed to measure %q: %v", word.Text, err)
		}

		canvasMaxX := canvas.Bounds().Dx() - rectangleCoordinates.Dx()
		canvasMaxY := canvas.Bounds().Dy() - rectangleCoordinates.Dy()

		if canvasMaxX < 0 || canvasMaxY < 0 {
			face.Close()
			log.Fatalf("Word %q does not fit on canvas", word.Text)
		}

		// Try each possible position, starting at a random one.
		positions := (canvasMaxX + 1) * (canvasMaxY + 1)
		start := rand.Intn(positions)
		var placedRect image.Rectangle
		found := false

		for attempt := 0; attempt < positions && !found; attempt++ {
			position := (start + attempt) % positions
			x := canvas.Bounds().Min.X + position%(canvasMaxX+1)
			y := canvas.Bounds().Min.Y + position/(canvasMaxX+1)
			candidate := image.Rect(x, y, x+rectangleCoordinates.Dx(), y+rectangleCoordinates.Dy())

			free := true
			for _, occupied := range occupiedArea {
				if !candidate.Intersect(occupied).Empty() {
					free = false
					break
				}
			}
			if free {
				placedRect = candidate
				found = true
			}
		}
		if !found {
			face.Close()
			log.Fatalf("No space on canvas for %q", word.Text)
		}

		paintRectangle(canvas, word, face, rectangleCoordinates, placedRect)
		face.Close()
		occupiedArea = append(occupiedArea, placedRect)
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
