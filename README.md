# KV-Store

KV-Store is a simple key-value store command-line application built in Go. It allows users to store, retrieve, and manage key-value pairs with ease.

## Features

- **Put**: Store a key-value pair in the store.
- **Get**: Retrieve a key-value pair from the store based on the key.
- **Delete**: Remove a key-value pair from the store based on the key.
- **Job**: Can run this in the background to delete expired keys.
- **Version**: Display the current version of the KV-Store application.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.15 or later
- PostgreSQL database

## Installation

To install KV-Store, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/kv-store.git
   ```

2. Change directory to the project root:

   ```sh
   cd kv-store
   ```

3. Build the application:
   ```go
   go build -o kv-store .
   ```

## Environment Setup

Before running KV-Store, you need to configure the database connection. This is done by setting the `DATABASE_URL` environment variable in a `.env` file in the project root directory.

1. Create a `.env` file in the root directory.
2. Add the following line to the `.env` file:

   ```sh
   DATABASE_URL=postgres://username:password@localhost:5432/kvstore?sslmode=disable
   ```

   Replace `username` and `password` with your PostgreSQL username and password, respectively.

3. Save the `.env` file.

## Usage

To use KV-Store, follow these steps:

- Put a key-value pair in the store:

  ```sh
  ./kv-store PUT key value
  ```

- Retrieve a key-value pair from the store:

  ```sh
  ./kv-store GET key
  ```

- Run the job to delete expired keys:

  ```sh
  ./kv-store job
  ```

## Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.
