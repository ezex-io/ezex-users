# ezeX Users Service

A service for managing user profiles and settings on the ezeX platform.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **[Go](https://go.dev/doc/install/)**: The Go programming language.
- **Development Tools**: run `make devtools` to install necessary tools for development.

### Build

To build the project, use:

```bash
make build
```

## Test

### Unit Tests

To run unit tests, use:

```bash
make test
```

### Integration Tests

The project includes integration tests that require a PostgreSQL database. To run them:

```bash
make test-integration
```

Integration tests use environment variables defined in `.env.example`:

```
EZEX_USERS_DB_ADDRESS=localhost:5432  # Host:port for test database
EZEX_USERS_DB_DATABASE=postgres       # Database name
EZEX_USERS_DB_USERNAME=postgres       # Database username 
EZEX_USERS_DB_PASSWORD=postgres       # Database password
EZEX_USERS_DB_PORT=5432               # Port for PostgreSQL container
```

To set up your own test configuration:
- Modify `.env.example` with your custom settings

When running integration tests, the system will:
- Create a temporary test database with a unique name
- Apply all migrations
- Run the tests against it
- Drop the database when tests complete

## Code Quality and Formatting

To automatically format the code, run:

```bash
make fmt
```

Run the linter to catch common mistakes and improve code quality:

```bash
make check
```

## Contributing

Contributions are most welcome!
Whether it's code, documentation, or ideas, every contribution makes a difference.
Please read the [Contributing](CONTRIBUTING.md) guide to get started.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
