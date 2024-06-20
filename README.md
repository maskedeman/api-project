# API Project

## Description

This project is a RESTful API built with Go using the Fiber framework. It uses MySQL for data storage and Redis for caching.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/maskedeman/api-project.git
    ```
2. Navigate to the project directory:
    ```bash
    cd api-project
    ```
3. Initialize:
    ```bash
    go mod init api-project
    ```
4. Install dependencies:
    ```bash
    go mod tidy
    ```

## Configuration

The project is configured via a `config.json` file. Here's a brief explanation of each setting:

- `server`: The domain and port where the server will run.
- `database`: The configuration for the MySQL database.
- `jwt`: The secrets and expiry times for JWT tokens.
- `redis`: The configuration for the Redis cache.
- `security`: The entropy for password hashing.
- `verbose`: Whether to run the server in verbose mode.
- `logger`: Whether to enable logging.

Adjust these settings as per the given config.json.example file. Make a new file config.json in internal/config and copy the contents of config.json.example file present in the same directory before running the project.

## Running the Project

You can run the project in two modes:

- Normal mode:
    ```bash
    make run
    ```
- Watch mode (the server will restart on file changes):
    ```bash
    make run watch-mode
    ```

## Other Commands

- To tidy up dependencies:
    ```bash
    make tidy
    ```
- To run database migrations:
    ```bash
    make migrate
    ```
- To lint the code:
    ```bash
    make lint
    ```

## API-Testing

- Postman collections along with environment are available in postman-collection folder for convenience .
