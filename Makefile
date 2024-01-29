INCLUDE_PATH := $(abspath .)
LIBRARY_PATH := $(abspath .)

all: modules build

modules:
	go mod tidy

build:
	mkdir -p bin && \
	@C_INCLUDE_PATH=${INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} go build ./cmd/go-whisper-subtitles/main.go
