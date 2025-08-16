.PHONY: check-env up run down test

check-env:
	@if [ ! -f .env ]; then \
		echo "ERROR: The .env file is missing!"; \
		echo "Please create a .env file with your secrets. Check the README.md for more information."; \
		exit 1; \
	fi
	@echo "Found .env file. Proceeding with Docker Compose..."
up: check-env
	docker-compose up -d
run: check-env
	docker-compose up -d --build
down:
	docker-compose down
test:
	go test