all: dicts.go dicts wasm

dicts.go: OpenCC-src
	go test -v

dicts: OpenCC-src
	(cd dicts && go test -v)

OpenCC-src:
	mkdir OpenCC-src && curl -Ls https://github.com/BYVoid/OpenCC/archive/ver.1.1.7.tar.gz | tar --strip-components 1 -C OpenCC-src -xvzf - && (cd OpenCC-src && npm install)

wasm:
	(cd wasm && GOOS=js GOARCH=wasm go build -v -o opencc.wasm)

clean:
	rm -rf OpenCC-src

.PHONY: dicts wasm
