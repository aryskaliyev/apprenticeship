package text_segmentation

import (
	"fmt"
	"gocv.io/x/gocv"
)

func DetectHorizontalLines(binary *gocv.Mat) (*gocv.Mat, error) {
	if binary == nil {
		return nil, fmt.Errorf("input binary is nil")
	}
	defer binary.Release()

	edges := gocv.NewMat()
	min := 50
	max := 150
	gocv.Canny(binary, edges, min, max)
	if edges == nil {
		return nil, fmt.Errorf("failed to detect horizontal lines with canny")
	}
	return edges, nil
}
