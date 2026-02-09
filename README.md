# Browser Cleaner

**Browser Cleaner** is a privacy-focused tool built in Go for cleaning cookies, cache, and browsing data from major browsers like Chrome, Firefox, Edge, and Safari. It supports Linux, macOS, and Windows only tested on Linux tho.

## Disclaimer

This project is primarily **original code**, but inspiration has been taken from various other open-source projects. It is a **learning project** created to explore Golang and the development of CLI tools.

### Important Notes

- **Not production-ready**: This tool is not meant for use in any serious or sensitive contexts.
- **May contain bugs**: Expect potential errors, inefficiencies, and unoptimized code.
- **Run at your own risk**: It has not been rigorously tested and could delete unintended data.

## How to Use

### Install

1. Clone the repository:
    ```
    git clone https://github.com/rs-jensen/browser-cleaner.git
    cd browser-cleaner
    ```

2. Build the project:
    ```
    go build -o browser-cleaner
    ```

### Commands

| Command                | Description                                   |
|------------------------|-----------------------------------------------|
| `browser-cleaner scan` | Scans browser data without removing it.       |
| `browser-cleaner clean`| Cleans cookies, cache, sessions, and history. |
| `--dry-run`            | Preview what will be deleted without cleaning.|
| `--browser`            | Specify specific browser (chrome, firefox).   |
| `--verbose`            | Detailed file-by-file output.                 |

### Examples

```bash
# Scan all installed browsers
browser-cleaner scan

# Perform a dry-run clean for Chrome
browser-cleaner clean --browser chrome --dry-run --verbose

# Clean everything with verbose output
browser-cleaner clean --verbose
```

---

## Acknowledgements

Parts of this project take **inspiration from other open-source codebases** and examples available online. The goal was to learn and experiment, so some concepts and patterns were adapted or extended.

**Thank you to the Go community** for providing helpful examples and libraries like `cobra` and `fatih/color`. 

---
**NOTE**: This project is provided "as-is", with no guarantees. Feedback and suggestions are welcome!
