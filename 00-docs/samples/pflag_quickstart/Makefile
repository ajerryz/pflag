GO := go
APP := app
BLD_DIR := .
#SRCS := $(wildcard *.go)  # 只有当前目录
SRCS := $(shell find . -name '*.go')

.PHONY: build clean


build:
	$(GO) build -o $(BLD_DIR)/$(APP) $(SRCS)

clean:
	rm $(BLD_DIR)/$(APP)
