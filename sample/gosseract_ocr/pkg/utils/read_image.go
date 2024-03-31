package utils

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ReadImage(imagePath string) (*gocv.Mat, error) {
	// if path == ""
	// if file exists in the path
	// if correct format
	img, err := gocv.IMRead(imagePath, gocv.IMReadColor)
	if err != nil {
		return nil, fmt.Errorf("Error loading image:", err)
	}
	defer img.Close()
	return img, nil
}
