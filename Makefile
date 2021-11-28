.PHONY: all

build:
	docker build . -t go-hackathon

run:
	docker run go-hackathon