package PhotoArray

import "fmt"

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

func (pab *PhotoArrayBuilder) ClearPaths() error {
	pab.imageFilePaths = pab.imageFilePaths[:0]
	return nil
}

func (pab *PhotoArrayBuilder) Length() int {
	return (len(pab.imageFilePaths))
}

func (pab *PhotoArrayBuilder) GenerateArray(outputPath string) error {

	return nil
}
