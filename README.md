# cli-dict

A lightweight command-line dictionary written in **Go**.
The tool fetches word definitions from the free **dictionaryapi.dev** API and prints them directly in the terminal.

## Features

* Look up English word definitions
* Shows part of speech
* Displays example sentences when available
* Simple and fast CLI interface

## Project Structure

```
cli-dict
│
├── cmd
│   └── dict
│       └── main.go        # CLI entrypoint
│
├── internal
│   ├── api
│   │   └── dictionary.go  # API request logic
│   └── models
│       └── response.go    # JSON response structures
│
├── go.mod
└── README.md
```

## Requirements

* Go 1.20+ (or any modern Go version)

Check your installation:

```
go version
```

## Build

Clone the repository:

```
git clone https://github.com/<your-username>/cli-dict.git
cd cli-dict
```

Build the binary:

```
go build -o dict ./cmd/dict
```

This will create the executable `dict`.

## Usage

```
./dict <word>
```

Example:

```
./dict hello
```

Example output:

```
Word: hello

noun
1. An expression of greeting.

verb
1. To greet someone.
```

## API

This project uses the free public API:

https://dictionaryapi.dev/

No API key is required.

## Purpose

This project was built as a beginner Go CLI project to practice:

* Go modules
* HTTP requests
* JSON decoding
* CLI argument handling
* project structure in Go

```
## My first go language experimental project.
```
