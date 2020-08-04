package PhotoArray

import (
	"image"
	"reflect"
	"testing"
)

func Test_createLayoutMap(t *testing.T) {
	tests := []struct {
		name                string
		imageFiles          []string
		imageHeight         int
		imageWidth          int
		rowLength           int
		wantPhotoSliceCount int
		wantOuterPoint      image.Point
		wantErr             bool
	}{
		{"3 rows of photos", []string{"one.jpg", "two.jpg", "three.jpg", "four.jpg", "five.jpg", "six.jpg"}, 100, 100, 2, 6, image.Point{X: 200, Y: 300}, false},
		{"2 rows of photos", []string{"one.jpg", "two.jpg", "three.jpg", "four.jpg", "five.jpg", "six.jpg"}, 100, 100, 3, 6, image.Point{X: 300, Y: 200}, false},
		{"no paths provided", []string{}, 100, 100, 2, 6, image.Point{X: 200, Y: 300}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCompositeMap, gotOuterPoint, err := createLayoutMap(tt.imageFiles, tt.imageHeight, tt.imageWidth, tt.rowLength)
			// check that an error exists if expected
			if err != nil {
				if !tt.wantErr {
					t.Errorf("createLayoutMap() error = %v, wantErr %v", err, tt.wantErr)
				}
				return // if there was an error, do not check the other values
			}

			// check that the composite map has the correct number of items
			if tt.wantPhotoSliceCount != len(gotCompositeMap) {
				t.Errorf("compositeMap should have %d photo slices, got %d", tt.wantPhotoSliceCount, len(gotCompositeMap))
			}
			// test outer point
			if !reflect.DeepEqual(gotOuterPoint, tt.wantOuterPoint) {
				t.Errorf("createLayoutMap() got outter point of %v but it should be %v", gotOuterPoint, tt.wantOuterPoint)
			}
		})
	}
}

func Test_maxInt(t *testing.T) {
	tests := []struct {
		name      string
		inputInts []int
		want      int
		wantErr   bool
	}{
		{name: "2 nums", inputInts: []int{100, 200}, want: 200, wantErr: false},
		{name: "5 nums", inputInts: []int{100, 200, 333, 443, 121, 0}, want: 443, wantErr: false},
		{name: "no input", inputInts: []int{}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxInt(tt.inputInts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("maxInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("maxInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPoint(t *testing.T) {
	tests := []struct {
		name        string
		inputPoints []image.Point
		wantMax     image.Point
		wantErr     bool
	}{
		{name: "no input", inputPoints: []image.Point{}, wantMax: image.Point{0, 0}, wantErr: true},
		{name: "2 points", inputPoints: []image.Point{{200, 100}, {200, 500}}, wantMax: image.Point{200, 500}, wantErr: false},
		{name: "4 points", inputPoints: []image.Point{{200, 100}, {200, 500}, {333, 333}, {100, 200}}, wantMax: image.Point{333, 500}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, err := maxPoint(tt.inputPoints...)
			if (err != nil) != tt.wantErr {
				t.Errorf("maxPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("maxPoint() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
