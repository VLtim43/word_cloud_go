package main

import (
	"image"

	"golang.org/x/image/font"
)

type layout struct {
	face   font.Face
	bounds image.Rectangle
}

func placeImage(canvas *image.RGBA, wordLayout layout) *image.RGBA {
	// place images on a spiral and handle collisions

	return canvas
}
