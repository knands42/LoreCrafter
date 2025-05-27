# Use an official Python runtime as a parent image
FROM python:3.11-slim

# Set environment variable for unbuffered output
ENV PYTHONUNBUFFERED 1

ARG GOOGLE_API_KEY
ARG OPENAI_API_KEY
ARG DATABASE_URL
ARG PASETO_PRIVATE_KEY
ARG PASETO_PUBLIC_KEY

ENV GOOGLE_API_KEY=${GOOGLE_API_KEY}
ENV OPENAI_API_KEY=${OPENAI_API_KEY}
ENV DATABASE_URL=${DATABASE_URL}
ENV PASETO_PRIVATE_KEY=${PASETO_PRIVATE_KEY}
ENV PASETO_PUBLIC_KEY=${PASETO_PUBLIC_KEY}

# Set work directory
WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc python3-dev libpq-dev \
    build-essential cmake libboost-all-dev libeigen3-dev \
    # WeasyPrint dependencies
    python3-cffi python3-brotli \
    libpango-1.0-0 libpangoft2-1.0-0 libharfbuzz0b \
    # Additional system libraries
    libgirepository-1.0-1 gir1.2-pango-1.0 \
    # Fonts
    fonts-dejavu fonts-freefont-ttf \
    # Clean up
    && rm -rf /var/lib/apt/lists/* \
    && pip install --no-cache-dir uv

# Install Python dependencies
COPY pyproject.toml .

# Create virtual environment and install dependencies
RUN python -m venv /opt/venv \
    && . /opt/venv/bin/activate \
    && pip install --upgrade pip \
    && pip install uv \
    && uv pip install -e .


# Generate encryption keys
RUN openssl genpkey -algorithm Ed25519 -out private_key.pem \
    && openssl pkey -in private_key.pem -pubout -out public_key.pem

# Set the virtual environment for subsequent commands
ENV PATH="/opt/venv/bin:$PATH"

# Copy project
COPY . .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the application
CMD ["uvicorn", "main_http:app", "--host", "0.0.0.0", "--port", "8000"]
