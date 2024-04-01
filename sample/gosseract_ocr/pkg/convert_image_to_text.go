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

	gray, err := preprocessing.ConvertImageToGrayscale(img)
	if err != nil {
		return err
	}

/*	blurred, err := preprocessing.ApplyNoiseReduction(gray)
	if err != nil {
		return err
	}
*/
	binary, err := preprocessing.ApplyBinarization(gray)
	if err != nil {
		return err
	}
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
