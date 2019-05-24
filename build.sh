#!/bin/sh

PROJECT_NAME="drmclient"
BIN_PATH="./bin"
DRM_FOR_MAC="drmclient_mac"
DRM_FOR_LINUX="drmclient_linux"
DRM_FOR_WINDOWS="drmclient.exe"
DRM_FOR_ARMLINUX="drmclient_armlinux"


rm -f ${BIN_PATH}/*

# for ARM-Linux
env GOOS=linux GOARCH=arm GOARM=7 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_ARMLINUX}

# for Linux
env GOOS=linux GOARCH=amd64 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_LINUX}

# for Windows
env GOOS=windows GOARCH=amd64 go build
mv ${PROJECT_NAME}.exe ${BIN_PATH}/${DRM_FOR_WINDOWS}

# for MacOS
env GOOS=darwin GOARCH=amd64 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_MAC}

