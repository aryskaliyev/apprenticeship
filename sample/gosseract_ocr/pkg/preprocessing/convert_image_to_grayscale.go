package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ConvertImageToGrayscale(img gocv.Mat) (gocv.Mat, error) {
	if img == nil {
		return gocv.Mat{}, fmt.Errorf("input image is nil")
	}
	defer img.Release()

	gray := gocv.NewGrayImageFromImage(img)
	if gray == nil {
		return gocv.Mat{}, fmt.Errorf("failed to convert to grayscale")
	}
	return gray, nil
}
