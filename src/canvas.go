package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
	return canvas
}

func writeText(canvas *image.RGBA, words []Word) *image.RGBA {
	face := basicfont.Face7x13
	lineHeight := face.Metrics().Height.Ceil()
	y := (canvas.Bounds().Dy()-len(words)*lineHeight)/2 + face.Metrics().Ascent.Ceil()

	for _, word := range words {
		x := (canvas.Bounds().Dx() - font.MeasureString(face, word.Text).Ceil()) / 2
		drawer := font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(color.Black),
			Face: face,
			Dot:  fixed.P(x, y),
		}
		drawer.DrawString(word.Text)
		y += lineHeight
	}

	return canvas
}
