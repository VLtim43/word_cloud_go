package main

import (
	"image"

	"golang.org/x/image/font"
)

type layout struct {
	face   font.Face
	bounds image.Rectangle
}

// placeImage is the entry point for positioning a measured word on the canvas.
// TODO: Choose a position that fits and does not overlap other placed words.
func placeImage(canvas *image.RGBA, wordLayout layout) *image.RGBA {
	return canvas
}
