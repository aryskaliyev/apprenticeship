# Use official Ubuntu image
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
    tesseract-ocr libtesseract-dev

# Clone OpenCV 4.8.1
RUN git clone --branch 4.8.1 --depth 1 https://github.com/opencv/opencv.git /opencv && \
    mkdir -p /opencv/build

# Build OpenCV
RUN cd /opencv/build && \
    cmake \
    -D CMAKE_BUILD_TYPE=RELEASE \
    -D CMAKE_INSTALL_PREFIX=/usr/local \
    -D BUILD_TESTS=OFF \
    -D BUILD_PERF_TESTS=OFF \
    -D BUILD_EXAMPLES=OFF \
    -D WITH_GTK=ON \
    .. && \
    make -j$(nproc) && \
    make install && \
    ldconfig

# Cleanup
RUN rm -rf /opencv
