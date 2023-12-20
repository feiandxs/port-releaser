# Port-Releaser

Port-Releaser is a cross-platform command-line tool designed to identify and optionally terminate processes that are occupying specific ports. It supports Windows, macOS, and Linux, making it a versatile tool for developers and system administrators.

## Features

- **Cross-platform Support**: Runs on Windows, macOS, and Linux.
- **Easy to Use**: Quickly find processes using a specific port with a simple command.
- **Safe Process Termination**: Offers the option to safely terminate the process after identification.

## Installation

To install Port-Releaser, clone this repository and build the tool using Go:

```bash
git clone https://github.com/feiandxs/port-releaser.git
cd port-releaser
go build
```

## Usage

To use Port-Releaser, simply run the following command with the desired port number:

```bash
./port-releaser <port-number>
```

For example, to find the process using port 8080:

```bash
./port-releaser 8080
```

The tool will list any processes using the specified port and prompt you to terminate them if desired.

## Contributing

Contributions to Port-Releaser are welcome and appreciated. If you have suggestions for improvements or bug fixes, please submit a pull request or open an issue.

## License

Forget about it, all codes in `main.go` are from ChatGPT.
