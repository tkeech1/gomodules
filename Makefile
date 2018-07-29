init-module:
	docker run -it --rm -v $(PWD):/opt --workdir /opt golang:1.11-rc go mod -init -module github.com/tkeech1/gomodules

sync:
	docker run -it --rm -v $(PWD):/opt --workdir /opt golang:1.11-rc go mod -sync