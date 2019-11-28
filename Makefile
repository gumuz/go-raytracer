run-projectile:
	go run cmd/projectile/projectile.go && open projectile.png

run-ball:
	go run cmd/ball/ball.go && open ball.png

test:
	go test raytracer/*_test.go -v