docker:
	go build main.go
	docker build -t hello .
	rm main

