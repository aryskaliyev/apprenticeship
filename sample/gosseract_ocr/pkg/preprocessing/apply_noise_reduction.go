package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyNoiseReduction(gray *cv.Mat) (*cv.Mat, error) {
	if gray == nil {
		return nil, fmt.Errorf("input gray is nil")
	}
	defer gray.Release()

	kernelSize := cv.Size{Width: 5, Height: 5}
	sigma := 0.0

	blurred := cv.NewMat()
	cv.GaussianBlur(gray, blurred, kernelSize, sigma)

	if blurred == nil {
		return nil, fmt.Errorf("failed to apply noise reduction")
	}

	return blurred, nil
}
