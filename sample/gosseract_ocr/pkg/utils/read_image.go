package utils

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ReadImage(imagePath string) (gocv.Mat, error) {
	// if path == ""
	// if file exists in the path
	// if correct format
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		return gocv.Mat{}, fmt.Errorf("error reading image")
	}

	return img, nil
}
