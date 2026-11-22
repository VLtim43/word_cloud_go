package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// Measure the word and return its font face for drawing.
func createWordRectangle(word Word, typeface *opentype.Font) (image.Rectangle, font.Face, error) {
	size := 12 + 0.3*float64(max(word.Weight, 0))

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
	draw.Draw(canvas, placedRect, image.NewUniform(color.RGBA{0, 173, 216, 255}), image.Point{}, draw.Src)

	drawer := font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(color.Black),
		Face: face,
		Dot:  fixed.P(placedRect.Min.X-glyphBounds.Min.X, placedRect.Min.Y-glyphBounds.Min.Y),
	}
	drawer.DrawString(word.Text)
}
