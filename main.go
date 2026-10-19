package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func createFile(destination string, fileName string) (*os.File, error) {

	const FILE_PERMISSION os.FileMode = 0755

	// creates the folder
	err := os.MkdirAll(destination, FILE_PERMISSION)
	if err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}

	// creates the full path with the filename and the destination path
	var destinationPath = filepath.Join(destination, fileName)

	// creates the file
	file, err := os.Create(destinationPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	fmt.Printf("Successfully created image at %v", destinationPath)

	return file, err
}

func createCanvas(canvasWidth int, canvasLength int, backgroundColor color.RGBA) *image.RGBA {

	// creates the background image
	canvas := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasLength))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	return canvas
}

func writeText(canvas *image.RGBA) *image.RGBA {

	// add text
	text := "Hello"
	face := basicfont.Face7x13

	width := font.MeasureString(face, text).Ceil()
	x := (canvas.Bounds().Dx() - width) / 2
	y := (canvas.Bounds().Dy() + face.Metrics().Height.Ceil()) / 2

	drawer := font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(color.Black),
		Face: face,
		Dot:  fixed.P(x, y),
	}
	drawer.DrawString(text)

	return canvas
}

type Word struct {
	Text   string `json:"text"`
	Weight int    `json:"weight"`
}
type Words struct {
	Words []Word `json:"words"`
}

func readJson(filename string) []Word {

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var words Words

	err = json.Unmarshal(data, &words)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return words.Words
}

func main() {

	// createdFile, err := createFile("./output", "test.png")
	// if err != nil {
	// 	log.Fatalf("Failed to create image file: %v", err)
	// }
	// defer createdFile.Close()

	// whiteColor := color.RGBA{245, 245, 245, 255}
	// canvas := writeText(createCanvas(640, 480, whiteColor))

	// // encode the image as a png
	// err = png.Encode(createdFile, canvas)
	// if err != nil {
	// 	log.Fatalf("Failed to encode image: %v", err)
	// }

	words := readJson("./word.json")

	for _, word := range words {
		fmt.Println(word.Text, word.Weight)
	}

}
