package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyNoiseReduction(img *cv.Mat) (*cv.Mat, error) {
	if img == nil {
		return nil, fmt.Errorf("input image is nil")
	}
	defer img.Release()

	kernelSize := cv.Size{Width: 5, Height: 5}
	sigma := 0.0

	blurred := cv.NewMat()
	cv.GaussianBlur(img, blurred, kernelSize, sigma)

	if blurred == nil {
		return nil, fmt.Errorf("failed to apply noise reduction")
	}

	return blurred, nil
}
