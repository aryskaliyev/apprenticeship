package utils

import (
	"fmt"
	"gocv.io/x/gocv"
)

func SaveImage(img gocv.Mat, newImagePath string) error {
	if img.Empty() {
		return fmt.Errorf("input image is empty")
	}
	defer img.Close()

	success := gocv.IMWrite(newImagePath, img)
	if !success {
		return fmt.Errorf("Error writing to file")
	}
	fmt.Println("Success!")
	return nil
}
