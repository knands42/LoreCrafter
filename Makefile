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

clean:
	@echo "Cleaning up..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	rm -rf .coverage
	rm -rf htmlcov
	@echo "Cleanup complete!"
