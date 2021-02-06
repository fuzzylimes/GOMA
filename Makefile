build_css:
	npm run css-build

build: build_css
	go build -o ./dist/goma_site cmd/app/main.go 

run:
	go run .