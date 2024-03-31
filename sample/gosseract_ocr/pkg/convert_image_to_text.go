package pkg

import (
	"fmt"
	"ocrx/pkg/preprocessing"
	"ocrx/pkg/utils"
)

func ConvertImageToText(readFilePath, writeFilePath string) error {
	img, err := utils.ReadImage(readFilePath)
	if err != nil {
		return err
	}
	defer img.Close()

	gray, err := preprocessing.ConvertImageToGrayscale(img)
	if err != nil {
		return err
	}
	defer gray.Close()

	blurred, err := preprocessing.ApplyNoiseReduction(img)
	if err != nil {
		return err
	}
	defer blurred.Close()

	binary, err := preprocessing.ApplyBinarization(img)
	if err != nil {
		return err
	}
	defer binary.Close()

	if writeFilePath == "" {
		fmt.Println(binary)
	} else {
		err := utils.SaveImage(binary, writeFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}
