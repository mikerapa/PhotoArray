package PhotoArray

import (
	"fmt"
	"image"
)

type PhotoSlice struct {
	path string
	rec  image.Rectangle
}

func maxInt(inputInts ...int) (int, error) {
	if len(inputInts) == 0 {
		return -1, fmt.Errorf("maximum int cannot be determined with no input")
	}

	max := inputInts[0]
	for _, currentInput := range inputInts[1:] {
		if currentInput > max {
			max = currentInput
		}
	}

	return max, nil
}

func maxPoint(inputPoints ...image.Point) (max image.Point, err error) {
	// return an error if there are no points
	if len(inputPoints) == 0 {
		return image.Point{X: 0, Y: 0}, fmt.Errorf("maximum point cannot be determined with no input")
	}
	// TODO this method could be simplified
	//create lists of all x and y values
	var xList, yList []int
	for _, currentPoint := range inputPoints {
		xList = append(xList, currentPoint.X)
		yList = append(yList, currentPoint.Y)
	}
	maxX, err := maxInt(xList...)
	if err != nil {
		return image.Point{X: 0, Y: 0}, err
	}

	maxY, err := maxInt(yList...)
	if err != nil {
		return image.Point{X: 0, Y: 0}, err
	}
	max = image.Point{X: maxX, Y: maxY}
	return max, nil
}

func createLayoutMap(imageFiles []string, imageHeight int, imageWidth int, rowLength int) (compositeMap []PhotoSlice, outerPoint image.Point, err error) {
	// Validate input
	if len(imageFiles) == 0 {
		err = fmt.Errorf("0 photo paths provided")
		return
	}

	if imageWidth < 1 || imageHeight < 1 {
		err = fmt.Errorf("imageHeight and imageWidth must be greater than 1. imageHeight=%d, imageWidth=%d", imageHeight, imageWidth)
		return
	}

	var photoMaxPoints []image.Point
	var row, column int
	for _, file := range imageFiles {
		rec := image.Rect(column*imageWidth, row*imageHeight, (column*imageWidth)+imageWidth, (row*imageHeight)+imageHeight)
		photoMaxPoints = append(photoMaxPoints, rec.Max)
		if column < (rowLength - 1) {
			column++
		} else {
			// advance the row
			row++
			column = 0
		}

		compositeMap = append(compositeMap, PhotoSlice{rec: rec, path: file})
	}
	// determine the outer limit of the photo
	outerPoint, err = maxPoint(photoMaxPoints...)
	return
}
