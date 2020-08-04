package PhotoArray

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

type PhotoArrayBuilder struct {
	PhotoHeight    int
	PhotoWidth     int
	RowLength      int
	imageFilePaths []string
}

func NewPhotoArrayBuilder() (newBuilder PhotoArrayBuilder) {
	newBuilder = PhotoArrayBuilder{PhotoHeight: 100, PhotoWidth: 100, RowLength: 10}
	return
}

func (pab *PhotoArrayBuilder) AddPhotoPaths(paths ...string) error {
	var invalidPaths []string
	for _, newPath := range paths {
		// if the path is not valid, add it to the list of invalid paths and skip the current loop
		if !IsValidFilePath(newPath) {
			invalidPaths = append(invalidPaths, newPath)
			continue
		}
		pab.imageFilePaths = append(pab.imageFilePaths, newPath)
	}
	// if there are invalid paths, return an error
	var err error = nil
	if len(invalidPaths) > 0 {
		err = fmt.Errorf("Some of the provided paths are invalid. %v Only %d paths were added. ", invalidPaths, len(paths)-len(invalidPaths))
	}
	return err
}

func (pab *PhotoArrayBuilder) ClearPaths() {
	pab.imageFilePaths = pab.imageFilePaths[:0]
}

func (pab *PhotoArrayBuilder) Length() int {
	return len(pab.imageFilePaths)
}

func ReScaleFill(img image.Image, maxHeight int, maxWidth int) *image.NRGBA {
	newImg := imaging.Fill(img, maxWidth, maxHeight, imaging.Top, imaging.Lanczos)
	return newImg
}

func (pab *PhotoArrayBuilder) GenerateArray(outputPath string) (err error) {
	// create output image
	layout, outputSize, err := createLayoutMap(pab.imageFilePaths, pab.PhotoHeight, pab.PhotoWidth, pab.RowLength)
	if err != nil {
		return err
	}
	output := imaging.New(outputSize.X, outputSize.Y, color.NRGBA{})

	for _, currentPhotoSlice := range layout {
		newImg, err := imaging.Open(currentPhotoSlice.path)
		if err != nil {
			return err
		}
		newImg = ReScaleFill(newImg, pab.PhotoHeight, pab.PhotoWidth)
		output = imaging.Paste(output, newImg, currentPhotoSlice.rec.Min)

	}
	// Save the resulting image as JPEG.
	err = imaging.Save(output, outputPath)
	if err != nil {
		return err
	}
	return nil
}
