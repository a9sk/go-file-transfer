#!/bin/bash

echo "[*] Creating TLS certificates..."
# Function to check if a command exists
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# Check if OpenSSL is installed
if ! command_exists openssl; then
  # Check if package manager is available and install OpenSSL
  if command_exists apt; then
    sudo apt update
    sudo apt install -y openssl
  elif command_exists yum; then
    sudo yum install -y openssl
  elif command_exists brew; then
    brew install openssl
  else
    echo "[!] Error: OpenSSL is required but could not be installed. Please install OpenSSL manually and rerun this script."; sleep 1; echo "[*] Exiting...";
    exit 1
  fi
fi

mkdir -p certificates

openssl req -x509 -nodes -newkey rsa:4096 -keyout certificates/server.key -out certificates/server.crt -days 365

echo "[*] Certificates generated successfully."; sleep 1; echo "[*] Exiting...";
