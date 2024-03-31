package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ConvertImageToGrayscale(img *cv.Mat) (*cv.Mat, error) {
	if img == nil {
		return nil, fmt.Errorf("input image is nil")
	}
	defer img.Release()

	gray := cv.NewGrayImageFromImage(img)
	if gray == nil {
		return nil, fmt.Errorf("failed to convert to grayscale")
	}
	return gray, nil
}
