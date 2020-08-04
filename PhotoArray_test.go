package PhotoArray

import (
	"testing"
)

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
