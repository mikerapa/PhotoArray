package PhotoArray

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func randomizedFilePath(pathTemplate string) string {
	pathTemplate = strings.Replace(pathTemplate, "#", "%d", 1)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rndNum := r1.Intn(1000)
	return fmt.Sprintf(pathTemplate, rndNum)
}

func TestNewPhotoArrayBuilder(t *testing.T) {
	newBuilder := NewPhotoArrayBuilder()
	const (
		defaultRowLength   = 10
		defaultPhotoHeight = 100
		defaultPhotoWidth  = 100
	)
	if newBuilder.RowLength != defaultRowLength {
		t.Errorf("NewPhotoArrayBuilder() RowLength default should be %v, got %v", defaultRowLength, newBuilder.RowLength)
	}

	if newBuilder.PhotoHeight != defaultPhotoHeight {
		t.Errorf("NewPhotoArrayBuilder() PhotoHeight default should be %v, got %v", defaultPhotoHeight, newBuilder.PhotoHeight)
	}

	if newBuilder.PhotoWidth != defaultPhotoWidth {
		t.Errorf("NewPhotoArrayBuilder() PhotoWidth default should be %v, got %v", defaultPhotoWidth, newBuilder.PhotoWidth)
	}

}

func TestPABAddAndClear(t *testing.T) {
	// create new PhotoArrayBuilder
	newBuilder := NewPhotoArrayBuilder()
	if newBuilder.Length() != 0 {
		t.Errorf("PhotoArrayBuilder Length should be 0, but it is %v", newBuilder.Length())
	}

	// Add paths
	newBuilder.AddPhotoPaths("./testimages/soccer1.jpg", "./testimages/soccer2.jpg")
	if newBuilder.Length() != 2 {
		t.Errorf("PhotoArrayBuilder Length should be 2, but it is %v", newBuilder.Length())
	}

	// clear the list and make sure it works
	newBuilder.ClearPaths()
	if newBuilder.Length() != 0 {
		t.Errorf("PhotoArrayBuilder Length should be 0, but it is %v", newBuilder.Length())
	}

}

func TestPhotoArrayBuilder_AddPhotoPaths(t *testing.T) {

	tests := []struct {
		name           string
		paths          []string
		wantErr        bool
		expectedLength int
	}{
		{name: "one valid path", paths: []string{"./testimages/soccer3.jpg"}, wantErr: false, expectedLength: 1},
		{name: "one invalid path", paths: []string{"./testimages/empty.jpg"}, wantErr: true, expectedLength: 0},
		{name: "one invalid path, 2 valid paths", paths: []string{"./testimages/soccer3.jpg", "./testimages/soccerx.jpg", "./testimages/soccer1.jpg"}, wantErr: true, expectedLength: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pab := NewPhotoArrayBuilder()
			err := pab.AddPhotoPaths(tt.paths...)

			// check length
			if tt.expectedLength != pab.Length() {
				t.Errorf("AddPhotoPath() Length expected to be %v, got %v", tt.expectedLength, pab.Length())
			}

			// check for error
			if tt.wantErr && err == nil {
				t.Errorf("AddPhotoPath() error expected")
			}

		})
	}
}

func TestPhotoArrayBuilder_GenerateArray(t *testing.T) {
	tests := []struct {
		name           string
		PhotoHeight    int
		PhotoWidth     int
		RowLength      int
		imageFilePaths []string
		wantErr        bool
	}{
		{name: "no paths", PhotoHeight: 100, PhotoWidth: 100, RowLength: 10, imageFilePaths: []string{}, wantErr: true},
		{name: "all 6 sample images", PhotoHeight: 100, PhotoWidth: 100, RowLength: 3, imageFilePaths: []string{"./testimages/soccer1.jpg", "./testimages/soccer2.jpg", "./testimages/soccer3.jpg", "./testimages/soccer4.jpg", "./testimages/soccer5.jpg", "./testimages/soccer6.jpg"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pab := NewPhotoArrayBuilder()
			pab.PhotoWidth = tt.PhotoWidth
			pab.PhotoHeight = tt.PhotoHeight
			pab.RowLength = tt.RowLength
			pab.AddPhotoPaths(tt.imageFilePaths...)

			// create an output path
			outputPath := randomizedFilePath("./testoutputimages/outputimage#.jpg")
			err := pab.GenerateArray(outputPath)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("GenerateArray() error = %v, wantErr %v", err, tt.wantErr)
				}
				return // if there's an error, do not check any other values
			}

			// check for the existence of the image, if there was no error
			if !IsValidFilePath(outputPath) {
				t.Errorf("file %s was not created.", outputPath)
			}
		})
	}
}
