.PHONY: setup run clean

# Detect OS and set Python path accordingly
ifeq ($(OS),Windows_NT)
    PYTHON_PATH := .venv\Scripts\python
else
    PYTHON_PATH := .venv/bin/python
endif

run:
	$(PYTHON_PATH) .\main.py

clean:
	@echo "Cleaning up..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	@echo "Cleanup complete!"