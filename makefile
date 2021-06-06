w = "/mnt/d/fmb"

pre:
	clear
	@./miho_scss_compiler
	@echo "SCSS Compiled"
	@GOOS=js GOARCH=wasm go build -o ./misc/miho.wasm ./wasm
	@echo "WASM Compiled"

dev: pre
	@go build -ldflags="-X 'main.port=:8080'" -o ./_air/main ./main.go

win: pre
	@if [ -d $(w) ]; then \
		rm -rf $(w); \
	fi
	@GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.port=:5000'" -o $(w)/start.exe ./main.go
	@find ./addr -name "*.miho" | cpio -pdm $(w)
	@cp -r ./misc $(w)/