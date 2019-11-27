run:
	go run cmd/main.go && open image.png

test:
	go test raytracer/*_test.go -v