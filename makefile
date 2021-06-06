w = "/mnt/d/fmb"

cls:
	clear
	@rm -rf ./mscmp_init
	@rm -rf ./misc/miho.css
	@rm -rf ./misc/miho.wasm

pre: cls
	@go build -o ./mscmp_init mscmp/mscmp.go
	@./mscmp_init
	@echo "SCSS Compiled"
	@GOOS=js GOARCH=wasm go build -o ./misc/miho.wasm ./wasm
	@echo "WASM Compiled"

dev: pre
	@go build -ldflags="-X 'main.port=:8080'" -o ./_air/main ./main.go

win: pre
	@if [ -d $(w) ]; then \
		echo "Removing Legacies"; \
		rm -rf $(w); \
	fi
	@GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.port=:5000'" -o $(w)/appStart.exe ./main.go
	@find ./addr -name "*.miho" | cpio -pdm $(w)
	@cp -r ./misc $(w)/
	# Cleansing
	@rm -rf ./mscmp_init
	@rm -rf ./misc/miho.css
	@rm -rf ./misc/miho.wasm