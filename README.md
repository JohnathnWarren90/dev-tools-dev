# Dev Tools CLI

A command line tool for quickly creating new development projects from predefined templates.

The goal of this project is to reduce repetitive setup when starting a new project by automatically creating folders, files, and starter code.

## Features

- Create new project structures
- Generate Go project templates
- Create folders automatically
- Generate starter files
- Built-in project templates using Go’s embed package
- Simple command line interface

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Build the tool:

```bash
go build -o dev
```

Move the executable to your PATH if you want to run it globally:

```bash
mv dev $HOME/bin/dev
```

## Usage

Create a new project:

```bash
dev new go ProjectName ProjectType
```

Example:

```bash
dev new go inventory cli
```

This will create:

```text
inventory/
├── main.go
├── models/
├── utils/
└── src/
```

## Commands

Create Project:

```bash
dev new <language> <project-name> <project-type>
```

Example:

```bash
dev new go calculator cli
```

Show Version:

```bash
dev --version
```

Example output:

```text
dev
Version: 0.0.1
```

## Supported Languages

Currently supported:

- Go

Future support planned:

- Python
- Java

## Project types

Currently available:

- cli

Coming Soon:

- tui
- server
- website

## Why This Exists

Starting a new project often requires repeating the same setup:

- creating folders
- creating files
- writing package declarations
- creating project structure

This tool automates that setup so development can start faster.

## Debugging

Currently the tool outputs submitted command arguments after project creation for troubleshooting.

Example:

```text
[dev new go inventory cli]
```

## Version

0.0.1

## License

-MIT-
This software is provided 'as-is', without any express or implied warranty. In no event shall the author be liable for any damages arising from the use of this software.
