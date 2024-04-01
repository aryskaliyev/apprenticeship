# Use Ubuntu 20.04 as the base image
FROM ubuntu:20.04

# Set environment variable to prevent tzdata from prompting for user input
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    wget \
    unzip \
    pkg-config \
    libgtk2.0-dev \
    libtbb2 \
    libtbb-dev \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libdc1394-22-dev \
    libleptonica-dev \
    libpango1.0-dev \
    libcairo2-dev \
    tesseract-ocr \
    libtesseract-dev

# Download and install Go version 1.22
RUN wget -q https://dl.google.com/go/go1.22.linux-amd64.tar.gz && \
    echo "3bc921a7fd74cf1d1343e54aaea30f5b50eb195f6d78fb17f01d589b6df4ef19  go1.22.linux-amd64.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf go1.22.linux-amd64.tar.gz && \
    rm go1.22.linux-amd64.tar.gz

# Set environment variables for Go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOROOT="/usr/local/go"
ENV GOPATH="/go"

# Install GoCV
RUN apt-get install -y libopencv-dev
RUN go get -u -d gocv.io/x/gocv

# Set the working directory
WORKDIR /app

# Copy the Golang code and go.mod/go.sum files into the container
COPY ./golang_code/* ./
COPY ./golang_code/go.mod ./
COPY ./golang_code/go.sum ./

# Download and install Go dependencies
RUN go mod download

# Build the Golang code
RUN go build -o app .

# Set the entrypoint to run the Golang code
CMD ["./app"]
