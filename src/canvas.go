package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// create the blank canvas
func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
	return canvas
}

// paint on the canvas
func writeText(canvas *image.RGBA, wordsWithWeights []Word, parsedFont *opentype.Font) (*image.RGBA, error) {

	wordLayout := make([]layout, len(wordsWithWeights))
	totalHeight := 0

	for i, word := range wordsWithWeights {

		// Scale weights to legible sizes without letting negative weights shrink the font.
		size := 12 + 0.3*float64(max(word.Weight, 0))
		face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{Size: size, DPI: 72})
		if err != nil {
			return nil, err
		}
		defer face.Close()

		glyphBounds, _ := font.BoundString(face, word.Text)

		wordLayout[i] = layout{
			face: face,
			bounds: image.Rect(
				glyphBounds.Min.X.Floor(), glyphBounds.Min.Y.Floor(),
				glyphBounds.Max.X.Ceil(), glyphBounds.Max.Y.Ceil(),
			),
		}
		totalHeight += wordLayout[i].bounds.Dy()
	}

	if len(wordLayout) > 0 {
		totalHeight += len(wordLayout) - 1
	}
	y := canvas.Bounds().Min.Y + (canvas.Bounds().Dy()-totalHeight)/2

	for i, word := range wordsWithWeights {
		bounds := wordLayout[i].bounds
		x := canvas.Bounds().Min.X + (canvas.Bounds().Dx()-bounds.Dx())/2

		// Draw the word's background and place its glyph bounds inside it.
		backgroundRectangle := image.Rect(x, y, x+bounds.Dx(), y+bounds.Dy())
		draw.Draw(canvas, backgroundRectangle, image.NewUniform(color.RGBA{0, 173, 216, 255}), image.Point{}, draw.Src)
		drawer := font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(color.Black),
			Face: wordLayout[i].face,
			Dot:  fixed.P(x-bounds.Min.X, y-bounds.Min.Y),
		}
		drawer.DrawString(word.Text)

		y += bounds.Dy()

		// one pixel before the next rectangle
		if i < len(wordLayout)-1 {
			y++
		}

	}

	return canvas, nil
}
