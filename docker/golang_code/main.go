package main

import (
    "fmt"
    "log"
    "github.com/otiai10/gosseract/v2"
    "gocv.io/x/gocv"
)

func main() {
    // Test OpenCV
    img := gocv.NewMat()
    defer img.Close()
    if img.Empty() {
        log.Fatal("Error: OpenCV is not working properly")
    }

    fmt.Println("OpenCV is working properly")

    // Test Tesseract
    client := gosseract.NewClient()
    defer client.Close()

    // Set image for OCR
    imgPath := "image.jpg" // Path to your test image
    err := client.SetImage(imgPath)
    if err != nil {
        log.Fatalf("Error setting image for OCR: %v", err)
    }

    // Perform OCR
    text, err := client.Text()
    if err != nil {
        log.Fatalf("Error performing OCR: %v", err)
    }

    fmt.Println("Detected text:", text)

    fmt.Println("Tesseract is working properly")
}
