#!/bin/sh

BIN_PATH="./bin"
DRM_FOR_MAC="drmclient_mac"
DRM_FOR_LINUX="drmclient_linux"
DRM_FOR_WINDOWS="drmclient.exe"
DRM_FOR_ARMLINIX="drmclient_armlinux"


rm -f ${BIN_PATH}/*

# for ARM-Linux
env GOOS=linux GOARCH=arm GOARM=7 go build
mv drmclient ${BIN_PATH}/drmclient_armlinux

# for Linux
env GOOS=linux GOARCH=amd64 go build
mv drmclient ${BIN_PATH}/drmclient_linux

# for Windows
env GOOS=windows GOARCH=amd64 go build
mv drmclient.exe ${BIN_PATH}/

# for MacOS
env GOOS=darwin GOARCH=amd64 go build
mv drmclient ${BIN_PATH}/drmclient_mac

