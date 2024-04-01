# Build the Docker image
docker build -t opencv-tesseract-golang .

# Run the Docker container
docker run --rm -it -v /path/to/golang_code:/app opencv-tesseract-golang


