build:
	go build -o build/meetup-reminder

run:
	go run *.go

linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -tags netgo -ldflags '-w' -o build/meetup-reminder-linux-amd64

docker: linux
	docker build -t tantalic/meetup-reminder:latest .

update-ca:
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 
