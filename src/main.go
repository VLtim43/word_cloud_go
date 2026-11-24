package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"

	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func main() {

	canvasColor := color.RGBA{0, 0, 0, 0}
	canvas := createCanvas(1000, 600, canvasColor)

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
	boldTypeface, err := opentype.Parse(gobold.TTF)
	if err != nil {
		log.Fatalf("Failed to parse bold font: %v", err)
	}

	occupiedArea := []image.Rectangle{}

	// place each word into the canvas keeping in mind the previously occupied areas
	for _, word := range words {
		wordTypeface := typeface
		if rand.Intn(2) == 0 {
			wordTypeface = boldTypeface
		}
		rectangleCoordinates, face, err := createWordRectangle(word, wordTypeface)
		if err != nil {
			log.Fatalf("Failed to measure %q: %v", word.Text, err)
		}

		canvasMaxX := canvas.Bounds().Dx() - rectangleCoordinates.Dx()
		canvasMaxY := canvas.Bounds().Dy() - rectangleCoordinates.Dy()

		if canvasMaxX < 0 || canvasMaxY < 0 {
			face.Close()
			log.Fatalf("Word %q does not fit on canvas", word.Text)
		}

		// Start at the center and move outward along a vertically squished spiral.
		centerX := canvas.Bounds().Min.X + canvasMaxX/2
		centerY := canvas.Bounds().Min.Y + canvasMaxY/2
		const verticalScale = 0.6
		maxRadius := math.Hypot(float64(canvasMaxX)/2, float64(canvasMaxY)/(2*verticalScale)) + 2
		phase := rand.Float64() * 2 * math.Pi
		var placedRect image.Rectangle
		found := false

		gap := 4

		for angle := 0.0; 1.5*angle <= maxRadius && !found; angle += 1 / math.Max(1, 1.5*angle) {
			radius := 1.5 * angle
			x := centerX + int(math.Round(radius*math.Cos(angle+phase)))
			y := centerY + int(math.Round(verticalScale*radius*math.Sin(angle+phase)))
			if x < canvas.Bounds().Min.X || x > canvas.Bounds().Min.X+canvasMaxX ||
				y < canvas.Bounds().Min.Y || y > canvas.Bounds().Min.Y+canvasMaxY {
				continue
			}
			candidate := image.Rect(x, y, x+rectangleCoordinates.Dx(), y+rectangleCoordinates.Dy())

			// add a small gap so they dont touch
			spacingRect := image.Rect(
				candidate.Min.X-gap, candidate.Min.Y-gap,
				candidate.Max.X+gap, candidate.Max.Y+gap,
			)

			free := true
			for _, occupied := range occupiedArea {
				if !spacingRect.Intersect(occupied).Empty() {
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
