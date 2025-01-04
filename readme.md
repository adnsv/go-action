# gitstat

A command-line utility that extracts and formats Git repository version information into JSON format. It's particularly useful for CI/CD pipelines and build systems that need to access version information programmatically.

## Features

- Extracts comprehensive Git repository metadata:
  - Current branch name
  - Full and short commit hashes
  - Author date
  - Latest tag information
  - Number of commits since last tag
- Provides version information in multiple formats:
  - Semantic versioning (major.minor.patch...)
  - Quad dot-separated version quad
  - NNNN comma-separated version quad
- JSON output for easy parsing and integration

## Installation

```bash
go install github.com/adnsv/gitstat@latest
```

## Usage

Basic usage (outputs to stdout):
```bash
gitstat
```

Save output to a file:
```bash
gitstat -o version.json
```

Enable verbose output:
```bash
gitstat --verbose
```

## Output Format

The tool outputs JSON with the following structure:

```json
{
  "branch": "main",
  "hash": "a1b2c3d4e5f6...",
  "short-hash": "a1b2c3d",
  "author-date": "2024-03-21T15:30:00",
  "last-tag": "v1.2.3",
  "additional-commits": 0,
  "ver-semantic": "1.2.3",
  "ver-quad": "1.2.3.50000",
  "ver-nnnn": "1,2,3,50000",
  "ver-major": 1,
  "ver-minor": 2,
  "ver-patch": 3,
}
```

## Command Line Options

| Flag | Description |
|------|-------------|
| `--verbose` | Enable detailed output logging |
| `-o, --output <file>` | Write JSON output to specified file instead of stdout |
