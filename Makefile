.PHONY: create-character search-character start-api test test-api test-cli test-cov

CHAR=Kirk

# Detect OS and set Python path accordingly
ifeq ($(OS),Windows_NT)
    PYTHON_PATH := .venv\Scripts\python
    PYTEST_PATH := .venv\Scripts\pytest
    PYTHON_ACTIVATE := .venv\Scripts\activate
else
    PYTHON_PATH := .venv/bin/python
    PYTEST_PATH := .venv/bin/pytest
    PYTHON_ACTIVATE := .venv/bin/activate
endif

# Setup commands
create-env:
	uv venv
	
install-deps:
	source $(PYTHON_ACTIVATE) | . $(PYTHON_ACTIVATE)
	uv pip install .
	uv pip install .[test]

generate-keys:
	openssl genpkey -algorithm Ed25519 -out private_key.pem
	openssl pkey -in private_key.pem -pubout -out public_key.pem

# App commands
create-character:
	$(PYTHON_PATH) main.py create-character

search-character:
	$(PYTHON_PATH) main.py search-character "$(CHAR)"

start-api:
	$(PYTHON_PATH) main.py api --host 0.0.0.0 --port 8080

# Test commands
test:
	$(PYTEST_PATH) tests

test-api:
	$(PYTEST_PATH) tests\api

test-cli:
	$(PYTEST_PATH) tests\cli

test-cov:
	$(PYTEST_PATH) --cov=src tests

# Podman Commands
podman-build:
	@echo "Building the application image..."
	podman build -t $(IMAGE_NAME) .

podman-up:
	@echo "Starting containers..."
	podman pod create --name $(CONTAINER_NAME) -p 8000:8000 -p 5432:5432
	podman run -d --pod $(CONTAINER_NAME) --name $(DB_CONTAINER) \
		-v $(DB_VOLUME):/var/lib/postgresql/data \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=app_db \
		docker.io/library/postgres:15-alpine
	podman run -d --pod $(CONTAINER_NAME) --name $(CONTAINER_NAME) \
		-e DATABASE_URL=postgresql://postgres:postgres@localhost:5432/app_db \
		$(IMAGE_NAME)

podman-down:
	@echo "Stopping and removing containers..."
	-podman stop $(CONTAINER_NAME) $(DB_CONTAINER)
	-podman rm $(CONTAINER_NAME) $(DB_CONTAINER)
	-podman pod rm $(CONTAINER_NAME)

podman-logs:
	@echo "Showing logs for $(CONTAINER_NAME)..."
	podman logs -f $(CONTAINER_NAME)

podman-db-logs:
	@echo "Showing database logs..."
	podman logs -f $(DB_CONTAINER)

podman-bash:
	@echo "Opening bash in $(CONTAINER_NAME)..."
	podman exec -it $(CONTAINER_NAME) /bin/bash

podman-restart: podman-down podman-up

podman-clean: podman-down
	@echo "Removing volumes and images..."
	-podman volume rm $(DB_VOLUME)
	-podman rmi $(IMAGE_NAME)

# Cleanup commands
clean: podman-clean
	@echo "Cleaning up Python artifacts..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	rm -rf .coverage
	rm -rf htmlcov
	@echo "Cleanup complete!"
