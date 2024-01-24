ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: build-dynamic
build-dynamic:
	@cd lib/vid && cargo build --release
	@cp lib/vid/target/release/libvid.dylib lib/
	go build -ldflags="-r $(ROOT_DIR)lib" main.go