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

	type layout struct {
		face   font.Face
		bounds image.Rectangle
	}

	wordLayout := make([]layout, len(wordsWithWeights))
	totalHeight := 0

	y := canvas.Bounds().Min.Y + (canvas.Bounds().Dy()-totalHeight)/2

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

		x := canvas.Bounds().Min.X + (canvas.Bounds().Dx()-wordLayout[i].bounds.Dx())/2

		backgroundRectangle := image.Rect(x, y, x+wordLayout[i].bounds.Dx(), y+wordLayout[i].bounds.Dy())
		draw.Draw(canvas, backgroundRectangle, image.NewUniform(color.RGBA{0, 173, 216, 255}), image.Point{}, draw.Src)

		textDrawer := font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(color.Black),
			Face: wordLayout[i].face,
			Dot:  fixed.P(x-wordLayout[i].bounds.Min.X, y-wordLayout[i].bounds.Min.Y),
		}

		textDrawer.DrawString(word.Text)
		y += wordLayout[i].bounds.Dy()
		if i < len(wordLayout)-1 {
			y++ // one pixel before the next rectangle
		}
	}

	if len(wordLayout) > 0 {
		totalHeight += len(wordLayout) - 1
	}

	return canvas, nil
}
