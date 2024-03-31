package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ConvertImageToGrayscale(img gocv.Mat) (gocv.Mat, error) {
	if img.Empty() {
		return gocv.Mat{}, fmt.Errorf("input image is empty")
	}
	defer img.Close()

	gray := gocv.NewGrayImageFromImage(img)
	if gray.Empty() {
		return gocv.Mat{}, fmt.Errorf("failed to convert to grayscale")
	}
	return gray, nil
}
