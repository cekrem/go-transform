# go-transform

A flexible Go-based text transformation tool that uses a plugin architecture to process input streams. This project follows Clean Architecture principles in general and the Dependency Inversion Principle in particular and emphasizes modularity through its plugin system.

## Features

- Plugin-based architecture for extensible text transformations
- Clean separation of concerns following Clean Architecture principles
- Simple CLI interface
- Stream-based processing for efficient handling of large inputs

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
├── cmd/            # Application entrypoint
├── internal/       # Internal application code
├── pkg/           # Public packages providing high level, stable abstractions
├── plugins/       # Transformer plugins: this is where you add your own plugins
└── build/         # Compiled binaries and plugins
```

## Development

### Creating New Plugins

Plugins should be created in the `plugins` directory and implement the transformer interface. See the `passthrough` plugin for an example implementation.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
