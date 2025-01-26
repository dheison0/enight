BUILD_DIR = build
GO_LDFLAGS = -w -s
GOFLAGS = -v

create_build_dir:
	mkdir -p "$(BUILD_DIR)"
	-false

system: create_build_dir
ifdef DISABLE_CGO
		export CGO_ENABLE=0
else
		export CGO_ENABLE=1
endif
	cd server && go build --ldflags="$(GO_LDFLAGS)" -o "../$(BUILD_DIR)/server" $(GOFLAGS)

www: create_build_dir
	cd web && npm run build -- --outDir "../$(BUILD_DIR)/www"

clean:
	rm -rf "$(BUILD_DIR)"

all: clean system www

run: all
	cd "$(BUILD_DIR)" && ./server
