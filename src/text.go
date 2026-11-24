package main

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var wordColors = [...]color.RGBA{
	{0, 82, 155, 255},
	{0, 119, 182, 255},
	{0, 150, 199, 255},
	{72, 207, 218, 255},
}

// Measure the word and return its font face for drawing.
func createWordRectangle(word Word, typeface *opentype.Font) (image.Rectangle, font.Face, error) {
	size := 1.5 * (12 + 0.5*float64(max(word.Weight, 0)))

	face, err := opentype.NewFace(typeface, &opentype.FaceOptions{
		Size: size,
		DPI:  72,
	})
	if err != nil {
		return image.Rectangle{}, nil, err
	}

	glyphBounds, _ := font.BoundString(face, word.Text)
	bounds := image.Rect(
		glyphBounds.Min.X.Floor(), glyphBounds.Min.Y.Floor(),
		glyphBounds.Max.X.Ceil(), glyphBounds.Max.Y.Ceil(),
	)
	return bounds, face, nil
}

func paintRectangle(canvas *image.RGBA, word Word, face font.Face, glyphBounds, placedRect image.Rectangle) {
	draw.Draw(canvas, placedRect, image.NewUniform(color.RGBA{0, 0, 0, 0}), image.Point{}, draw.Over)

	drawer := font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(wordColors[rand.Intn(len(wordColors))]),
		Face: face,
		Dot:  fixed.P(placedRect.Min.X-glyphBounds.Min.X, placedRect.Min.Y-glyphBounds.Min.Y),
	}
	drawer.DrawString(word.Text)
}
