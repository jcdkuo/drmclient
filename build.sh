#!/bin/sh

PROJECT_NAME="drmclient"
BIN_PATH="./bin"
DRM_FOR_MAC="drmclient_mac"
DRM_FOR_LINUX="drmclient_linux"
DRM_FOR_WINDOWS="drmclient.exe"
DRM_FOR_ARMLINUX="drmclient_armlinux"

START_MSG()
{
    echo "\033[1;31m Starting to build ${PROJECT_NAME} for ${1}...... \033[0m"
}

END_MSG()
{
    echo "\033[1;37m -> \033[1;44m ${1} \033[0m \033[1;37m is now built in bin/ \033[0m"
}

rm -f ${BIN_PATH}/*

# for ARM-Linux
START_MSG "arm-linux"
env GOOS=linux GOARCH=arm GOARM=7 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_ARMLINUX}
END_MSG ${DRM_FOR_ARMLINUX}

# for Linux
START_MSG "linux"
env GOOS=linux GOARCH=amd64 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_LINUX}
END_MSG ${DRM_FOR_LINUX}

# for Windows
START_MSG "Windows"
env GOOS=windows GOARCH=amd64 go build
mv ${PROJECT_NAME}.exe ${BIN_PATH}/${DRM_FOR_WINDOWS}
END_MSG ${DRM_FOR_WINDOWS}

# for MacOS
START_MSG "MacOS"
env GOOS=darwin GOARCH=amd64 go build
mv ${PROJECT_NAME} ${BIN_PATH}/${DRM_FOR_MAC}
END_MSG ${DRM_FOR_MAC}

