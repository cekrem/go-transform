# go-transform

A flexible Go-based text transformation tool that uses a plugin architecture to process input streams. This project demonstrates the Dependency Inversion Principle with three distinct layers, emphasizing modularity through its plugin system.

## Features

- Plugin-based architecture for extensible text transformations
- Clean separation of concerns following Clean Architecture principles
- Simple CLI interface

## Prerequisites

- Go 1.23 or later
- Make

## Installation

```bash
# Clone the repository
git clone https://github.com/cekrem/go-transform.git
cd go-transform

# Build the main application and plugins
make
```

## Usage

The tool reads from standard input and writes to standard output:

```bash
# Using the default passthrough transformer
echo "Hello, World!" | ./build/transform

# Using a specific transformer
echo "Hello, World!" | ./build/transform -transformer=passthrough
```

## Project Structure

```
.
├── cmd/                 # Application entrypoint
├── pkg/
│   └── domain/         # Core business rules and interfaces
├── internal/
│   └── app/           # Application logic
│       └── processor/ # Transformation orchestration
└── plugins/           # Infrastructure implementations (plugins)
```

## Architecture

This project demonstrates Clean Architecture principles with three layers:

### Domain Layer (`pkg/domain`)

- Contains core business rules and interfaces
- Has no external dependencies
- Defines what transformers should do
- `Transformer` and `Plugin` interfaces

### Application Layer (`internal/app`)

- Contains core application logic
- Depends only on domain interfaces
- Coordinates the transformation process
- `Processor` that manages plugins and executes transformations

### Infrastructure Layer (`plugins`)

- Contains concrete implementations
- Depends on domain interfaces
- Implements specific transformation strategies
- `passthrough` plugin

The project follows the Dependency Inversion Principle by:

1. Defining abstractions in the domain layer
2. Having both application and infrastructure layers depend on domain interfaces
3. Ensuring all dependencies point toward the domain layer

## Development

### Creating New Plugins

Plugins should be created in the `plugins` directory and implement the transformer interface from the domain layer. See the `passthrough` plugin for an example implementation.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
