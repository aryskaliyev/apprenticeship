package utils

import (
	"fmt"
	"gocv.io/cv"
)

func ReadImage(imagePath string) (*cv.Mat, error) {
	// if path == ""
	// if file exists in the path
	// if correct format
	img, err := cv.LoadImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Error loading image:", err)
	}
	return img, nil
}
