package utils

import (
	"fmt"
	"gocv.io/x/gocv"
)

func SaveImage(img *cv.Mat, newImagePath string) error {
	if img == nil {
		return fmt.Errorf("input image is nil")
	}
	defer img.Release()

	err := cv.SaveImage(newImagePath, img)
	if err != nil {
		return fmt.Errorf("Error saving image:", err)
	}
}