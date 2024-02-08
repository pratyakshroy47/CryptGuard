# CryptGuard: Secure File Encryption for Go

This project provides a simple file encryption utility written in Go, allowing users to encrypt and decrypt files using AES-GCM encryption.

## Features

- Encrypt files with a password
- Decrypt encrypted files using the same password
- Command-line interface for ease of use
- Strong encryption using AES-GCM algorithm

## Usage

To use the file encryption utility, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone <repository_url>


2. Navigate to the project directory:

   ```bash
   cd <project_directory>

3. Run the desired command to encrypt or decrypt a file:

   ```bash
   # Encrypt a file
   go run . encrypt /path/to/your/file

   # Decrypt a file
   go run . decrypt /path/to/your/encrypted/file

4. Follow the prompts to enter your password and confirm it.


## Commands
encrypt: Encrypts a file given a password
decrypt: Tries to decrypt a file using a password
help: Displays help text 
