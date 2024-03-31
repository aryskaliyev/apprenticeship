package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyBinarization(blurred gocv.Mat) (gocv.Mat, error) {
	if blurred.Empty() {
		return gocv.Mat{}, fmt.Errorf("input blurred is empty")
	}
	defer blurred.Close()

	thresh := 0.0
	maxval := 255.0

	binary := gocv.NewMat()
	gocv.Threshold(blurred, &binary, thresh, maxval, gocv.ThresholdBinary)

	if binary.Empty() {
		return gocv.Mat{}, fmt.Errorf("failed to apply binarization")
	}
	return binary, nil
}
