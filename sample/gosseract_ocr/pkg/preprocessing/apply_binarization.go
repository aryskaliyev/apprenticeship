package preprocessing

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ApplyBinarization(blurred gocv.Mat) (gocv.Mat, error) {
	if blurred == nil {
		return gocv.Mat{}, fmt.Errorf("input blurred is nil")
	}
	defer blurred.Close()

	binary := gocv.NewMat()
	gocv.Threshold(blurred, binary, 0, 255, gocv.ThresholdBinary)

	if binary == nil {
		return gocv.Mat{}, fmt.Errorf("failed to apply binarization")
	}
	return binary, nil
}
