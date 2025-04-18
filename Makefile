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

create-character-with-reference:
	$(PYTHON_PATH) .\main.py create-character-with-reference "$(CHAR)"

clean:
	@echo "Cleaning up..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	@echo "Cleanup complete!"