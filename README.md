# ⏳ Sand Programming Language

Sand is a lightweight, fast, and statically-linked programming language built from scratch in Go. It's designed to be simple to read and powerful enough to handle basic scripting tasks with ease.



## ✨ Key Features
* **Zero Dependencies**: Compiled into a single, highly optimized binary.
* **Built-in Packages**: Includes a standard library (like `stdio`) for immediate I/O.
* **Variable Scoping**: Clean variable declaration with `var`.
* **Fast Execution**: Powered by a custom-built Lexer, Parser, and Evaluator.

## 🚀 Getting Started

### Installation
1.  Go to the **Releases** section of this repository.
2.  Download `Sand_Setup_v1.0.exe`.
3.  Run the installer (defaults to your Downloads folder).
4.  Open a terminal and run your first script!

### Your First Script
Create a file named `hello.snd`:
```rust
var user = "Benja";
var version = 1.0;

stdio.logln("Welcome to Sand Lang,", user);
stdio.logln("Current Version:", version);
