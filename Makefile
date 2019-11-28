run-projectile:
	go run cmd/projectile/projectile.go && open projectile.png

test:
	go test raytracer/*_test.go -v