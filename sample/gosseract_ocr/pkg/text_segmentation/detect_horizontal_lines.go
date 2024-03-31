package text_segmentation

import (
	"fmt"
	"gocv.io/x/gocv"
)

func DetectHorizontalLines(binary *cv.Mat) (*cv.Mat, error) {
	if binary == nil {
		return nil, fmt.Errorf("input binary is nil")
	}
	defer binary.Release()

	edges := cv.NewMat()
	min := 50
	max := 150
	cv.Canny(binary, edges, min, max)
	if edges == nil {
		return nil, fmt.Errorf("failed to detect horizontal lines with canny")
	}
	return edges, nil
}
