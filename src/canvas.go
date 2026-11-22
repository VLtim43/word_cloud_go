package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
)

func createFile(destination string, fileName string) (*os.File, error) {
	if err := os.MkdirAll(destination, 0755); err != nil {
		return nil, fmt.Errorf("create output folder: %w", err)
	}

	destinationPath := filepath.Join(destination, fileName)
	file, err := os.Create(destinationPath)
	if err != nil {
		return nil, fmt.Errorf("create output file: %w", err)
	}

	return file, nil
}

// create the blank canvas
func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
	return canvas
}
