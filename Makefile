IMAGE_NAME=package-handler

build:
	docker build -t $(IMAGE_NAME) .

run: build
	docker run -p 8080:8080 $(IMAGE_NAME)

stop:
	docker stop $(shell docker ps -q --filter ancestor=$(IMAGE_NAME))