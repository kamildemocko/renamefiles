# Rename Files

A simple Go utility for batch renaming files in a directory based on regex patterns.  
It automatically creates a backup of renamed files in a ZIP archive.

## Installation

```bash
go install github.com/kamildemocko/renamefiles/cmd/renamefiles@latest
```

## Usage

```bash
renamefiles -p PATTERN [DIRECTORY]
```

Where:
- `-p PATTERN`: The pattern to match for renaming files
- `DIRECTORY`: Optional directory path (defaults to current directory)

### Pattern Syntax
- `X` - matches any letter (A-Z, a-z)
- `N` - matches any digit (0-9)
- Special characters like `(` and `)` are escaped automatically

### Examples

```bash
# Remove "_XXX" suffix from all files in current directory (File_ABX.txt -> File.txt)
renamefiles -p "_XXX"

# Remove "_(NN)" pattern (where NN are digits) from files in specific directory (File_(22).txt -> File.txt)
renamefiles -p "_(NN)" /path/to/directory
```
