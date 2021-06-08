# VFA-314 FULL METAL BITCHES
> Progressive Web Application that using Web Assembly and SCSS from Go-lang.
> Back-end has built with Go-Fiber Web Framework, Ubuntu 20.04 LTS on WSL 2.

## How to Clone?

1. Make Go-lang workspace directories.
```
    mkdir -p go/bin go/pkg go/src
```
2. Move to Go-lang development source directory.
```
    cd go/src
```
3. Clone this repository.
```
    git clone https://www.github.com/kim-ahri/fmb.git
```
4. Make Go-lang module manager and download dependencies.
```
    go mod init
    go mod tidy
```
5. Bring Web Assembly Controller.
```
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./misc
```
6. Check Makefile for build. If you want using `make dev` you need **cosmtrek/air.**
```
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
7. Check `.air.toml` to setting of **cosmtrek/air.** when you once ready,
```
	air
```
8. If you want build to windows environment, check `w` variable value in Makefile. and then,
```
	make win
```
9. You can change port in build time with using parameters in Makefile.
```
	-ldflags="-X 'main.port=:5000'"
``` 

### Happy Hack