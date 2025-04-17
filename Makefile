.PHONY: setup run clean

run:
	python main.py character

clean:
	@echo "Cleaning up..."
	rm -rf .venv
	rm -rf __pycache__
	rm -rf *.pyc
	@echo "Cleanup complete!"