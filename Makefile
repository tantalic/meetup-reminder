build:
	go build .

run:
	go run .

docker:
	docker buildx build -t tantalic/meetup-reminder:latest --platform=linux/amd64,linux/arm64 --push .
