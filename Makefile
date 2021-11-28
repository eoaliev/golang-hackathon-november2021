.PHONY: all

build:
	docker build . -t go-hackathon

run:
	docker container run --rm --name go-hackathon -it go-hackathon
