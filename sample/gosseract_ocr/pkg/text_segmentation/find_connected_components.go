package text_segmentation

import (
	"fmt"
	"gocv.io/x/gocv"
)

func FindConnectedComponents(edges *gocv.Mat) (*gocv.Mat, error) {
	if edges == nil {
		return nil, fmt.Errorf("input edges is nil")
	}
	defer edges.Release()

	contours := gocv.FindContours(edges)
	if contours == nil {
		return nil, fmt.Errorf("failed to find connected components")
	}
	return contours, nil
}
