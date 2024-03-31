package utils

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ReadImage(imagePath string) (*gocv.Mat, error) {
	// if path == ""
	// if file exists in the path
	// if correct format
	img, err := gocv.LoadImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Error loading image:", err)
	}
	return img, nil
}
