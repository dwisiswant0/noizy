APP_NAME = noizy
VERSION  = $(shell git describe --always --tags)

GO_LDFLAGS = -s -w -X 'github.com/dwisiswant0/noizy/common.AppVersion=${VERSION}'

build:
	CGO_ENABLED=1 go build -ldflags "$(GO_LDFLAGS)" -trimpath -o ./bin/$(APP_NAME) .

build-windows: GO_LDFLAGS += -H=windowsgui
build-windows:
	set CGO_ENABLED=1 && go build -ldflags "$(GO_LDFLAGS)" -trimpath -o ./bin/$(APP_NAME) .

print-ldflags:
	@echo -n "$(GO_LDFLAGS)"

clean:
	rm -rf ./bin
