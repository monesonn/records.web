.PHONY: run

APP_NAME=records-web-application
BUILD_DIR=$(PWD)/build

clean:
	rm -rf ./build

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) app.go

run:
	air
