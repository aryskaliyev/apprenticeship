package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyBinarization(blurred *cv.Mat) (*cv.Mat, error) {
	if blurred == nil {
		return nil, fmt.Errorf("input blurred is nil")
	}
	defer blurred.Release()

	binary := cv.NewMat()
	cv.Threshold(blurred, binary, 0, 255, cv.ThreshBinary)

	if binary == nil {
		return nil, fmt.Errorf("failed to apply binarization")
	}
	return binary, nil
}
