package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyNoiseReduction(gray gocv.Mat) (gocv.Mat, error) {
	if gray == nil {
		return gocv.Mat{}, fmt.Errorf("input gray is nil")
	}
	defer gray.Close()

	kernelSize := gocv.Size{Width: 5, Height: 5}
	sigma := 0.0

	blurred := gocv.NewMat()
	gocv.GaussianBlur(gray, blurred, kernelSize, sigma)

	if blurred == nil {
		return gocv.Mat{}, fmt.Errorf("failed to apply noise reduction")
	}

	return blurred, nil
}
