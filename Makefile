build_css:
	npm run css-build

build: build_css
	go build -o ./dist/goma_site cmd/app/main.go 

run:
	go run .

# Must include file as arg i.e. make add_page file=test
add_page:
	go run tools/generator.go -f $(file)