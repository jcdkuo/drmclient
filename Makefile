PROJECT_NAME = drmclient
PLATFORM_MACOS = mac
PLATFORM_LINUX = linux
PLATFORM_WINDOWS = exe
PLATFORM_ARM = armlinux
BIN_PATH = ${PWD}/bin

all: compile

build:
	@go build

clean:
	@rm -rf ${PWD}/${PROJECT_NAME}
	@rm -rf ${BIN_PATH}/*

compile:
	GOOS=linux GOARCH=arm GOARM=7 go build -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_ARM}
	GOOS=linux GOARCH=amd64 go build -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_LINUX}
	GOOS=windows GOARCH=amd64 go build -o ${BIN_PATH}/${PROJECT_NAME}.${PLATFORM_WINDOWS}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_PATH}/${PROJECT_NAME}_${PLATFORM_MACOS}

