package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyBinarization(img *cv.Mat) (*cv.Mat, error) {
	if img == nil {
		return nil, fmt.Errorf("input image is nil")
	}
	defer img.Release()

	binary := cv.NewMat()
	cv.Threshold(img, binary, 0, 255, cv.ThreshBinary)

	if binary == nil {
		return nil, fmt.Errorf("failed to apply binarization")
	}
	return binary, nil
}
