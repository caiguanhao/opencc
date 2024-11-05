all: dicts dicts.go wasm/opencc.wasm

dicts.go: OpenCC-src
	go test -v

dicts: OpenCC-src
	(cd dicts && go test -v)

OpenCC-src:
	mkdir OpenCC-src && \
		curl -Ls https://github.com/BYVoid/OpenCC/archive/ver.1.1.9.tar.gz | \
		tar --strip-components 1 -C OpenCC-src -xvzf - && (cd OpenCC-src && npm install)

wasm/opencc.wasm:
	(cd wasm && GOOS=js GOARCH=wasm go build -v -o opencc.wasm)

clean:
	rm -rf OpenCC-src wasm/opencc.wasm

.PHONY: all dicts clean
