IMAGE_NAME=package-handler

build:
	docker build -t $(IMAGE_NAME) .

run: build
	docker run -d -p 8080:8080 --name $(IMAGE_NAME) $(IMAGE_NAME)
	@echo "Backend running at http://localhost:8080"

dev:
	go run .

frontend:
	cd frontend && npm install && npm run dev

stop:
	-docker stop $(IMAGE_NAME) 2>/dev/null
	-docker rm $(IMAGE_NAME) 2>/dev/null

test:
	go test ./logic
