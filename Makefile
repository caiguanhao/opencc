dicts.go: OpenCC-src
	go test -v

dicts: OpenCC-src
	(cd dicts && go test -v)

OpenCC-src:
	mkdir OpenCC-src && curl -Ls https://github.com/BYVoid/OpenCC/archive/ver.1.1.0.tar.gz | tar --strip-components 1 -C OpenCC-src -xvzf - && (cd OpenCC-src && npm install)

.PHONY: dicts
