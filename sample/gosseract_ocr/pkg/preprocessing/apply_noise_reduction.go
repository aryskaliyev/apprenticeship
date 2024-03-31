package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyNoiseReduction(gray gocv.Mat) (gocv.Mat, error) {
	if gray.Empty() {
		return gocv.Mat{}, fmt.Errorf("input gray is empty")
	}
	defer gray.Close()

	sigmaX := 0.0
	sigmaY := 0.0

	blurred := gocv.NewMat()
	defer blurred.Close()

	gocv.GaussianBlur(gray, &blurred, gocv.Point{X: 5, Y: 5}, sigmaX, sigmaY, gocv.BorderDefault)

	if blurred.Empty() {
		return gocv.Mat{}, fmt.Errorf("failed to apply noise reduction")
	}
	return blurred, nil
}
