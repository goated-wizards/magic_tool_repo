package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func MakeGrid(sourcedir string) {
	paths, err := MakeGroups(sourcedir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(paths)
	for gridNumber, grid := range paths {
		var images []image.Image
		for _, p := range grid {
			img, err := LoadImage(p)
			if err != nil {
				fmt.Printf("Error loading %s: %v\n", p, err)
				return
			}
			images = append(images, img)
		}

		const cellSize = 300 // size of each grid cell

		// Create full canvas (3x3 grid)
		canvas := image.NewRGBA(image.Rect(0, 0, cellSize*gridSize, cellSize*gridSize))

		for i, img := range images {
			row := i / gridSize
			col := i % gridSize

			// Scale image to fit cell
			scaled := ScaleToFit(img, cellSize, cellSize)

			// Center image in its cell
			x := col*cellSize + (cellSize-scaled.Bounds().Dx())/2
			y := row*cellSize + (cellSize-scaled.Bounds().Dy())/2

			// Draw onto canvas
			draw.Draw(canvas, image.Rect(x, y, x+scaled.Bounds().Dx(), y+scaled.Bounds().Dy()), scaled, image.Point{}, draw.Over)
			output := fmt.Sprintf("output/player1grid/grid-%v.jpeg", gridNumber)
			if err := SaveAsJPEG(canvas, output); err != nil {
				fmt.Println("Error saving output:", err)
			}

			fmt.Println("Saved:", output)
		}
	}

}

func MakeGroups(dir string) ([][]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result [][]string
	var batch []string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		fullPath := filepath.Join(dir, e.Name())
		batch = append(batch, fullPath)

		if len(batch) == 9 {
			result = append(result, batch)
			batch = []string{}
		}
	}

	// leftover files (<9)
	if len(batch) > 0 {
		result = append(result, batch)
	}

	return result, nil
}

const gridSize = 3 // 3x3

// loadImage loads an image from disk.
func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}

// saveAsJPEG writes the final output image.
func SaveAsJPEG(img image.Image, outPath string) error {
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
}

// scaleToFit scales an image to fit inside targetW x targetH while preserving aspect ratio.
func ScaleToFit(src image.Image, targetW, targetH int) image.Image {
	b := src.Bounds()
	w := b.Dx()
	h := b.Dy()

	scale := float64(targetW) / float64(w)
	if float64(h)*scale > float64(targetH) {
		scale = float64(targetH) / float64(h)
	}

	newW := int(float64(w) * scale)
	newH := int(float64(h) * scale)

	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}
