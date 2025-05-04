.PHONY: create-character search-character

CHAR=Kirk

# Detect OS and set Python path accordingly
ifeq ($(OS),Windows_NT)
    PYTHON_PATH := .venv\Scripts\python
else
    PYTHON_PATH := .venv/bin/python
endif

create-character:
	$(PYTHON_PATH) .\main.py create-character

search-character:
	$(PYTHON_PATH) .\main.py search-character "$(CHAR)"

start-api:
	$(PYTHON_PATH) .\main.py api --host 0.0.0.0 --port 8080

clean:
	@echo "Cleaning up..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	@echo "Cleanup complete!"