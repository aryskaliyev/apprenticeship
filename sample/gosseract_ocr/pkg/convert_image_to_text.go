package pkg

import (
	"gocv.io/x/gocv"
	"ocrx/pkg/preprocessing"
	"ocrx/pkg/utils"
)

func ConvertImageToText(readFilePath, writeFilePath string) error {
	img := gocv.NewMat()
	defer img.Close()
	if img, err := utils.ReadImage(readFilePath); err != nil {
		return err
	}

	if img, err := preprocessing.ConvertImageToGrayScale(img); err != nil {
		return err
	}

	if img, err := preprocessing.ApplyNoiseReduction(img); err != nil {
		return err
	}

	if img, err := preprocessing.ApplyBinarization(img); err != nil {
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
}
