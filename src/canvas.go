package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
	return canvas
}

func writeText(canvas *image.RGBA, wordsWithWeights []Word, parsedFont *opentype.Font) (*image.RGBA, error) {
	type wordLayout struct {
		face   font.Face
		bounds image.Rectangle
	}

	layouts := make([]wordLayout, len(wordsWithWeights))
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
		layouts[i] = wordLayout{
			face: face,
			bounds: image.Rect(
				glyphBounds.Min.X.Floor(), glyphBounds.Min.Y.Floor(),
				glyphBounds.Max.X.Ceil(), glyphBounds.Max.Y.Ceil(),
			),
		}
		totalHeight += layouts[i].bounds.Dy()
	}

	if len(layouts) > 0 {
		totalHeight += len(layouts) - 1
	}

	y := canvas.Bounds().Min.Y + (canvas.Bounds().Dy()-totalHeight)/2
	for i, word := range wordsWithWeights {
		layout := layouts[i]
		x := canvas.Bounds().Min.X + (canvas.Bounds().Dx()-layout.bounds.Dx())/2

		box := image.Rect(x, y, x+layout.bounds.Dx(), y+layout.bounds.Dy())
		draw.Draw(canvas, box, image.NewUniform(color.RGBA{0, 173, 216, 255}), image.Point{}, draw.Src)

		textDrawer := font.Drawer{
			Dst:  canvas,
			Src:  image.NewUniform(color.Black),
			Face: layout.face,
			Dot:  fixed.P(x-layout.bounds.Min.X, y-layout.bounds.Min.Y),
		}

		textDrawer.DrawString(word.Text)
		y += layout.bounds.Dy()
		if i < len(layouts)-1 {
			y++ // one pixel before the next rectangle
		}
	}

	return canvas, nil
}
