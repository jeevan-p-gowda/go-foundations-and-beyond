# Go foundations and beyond <img align="right" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" width="150" height="100" title='Go'/>
Learning on fundamentals of go lang.

## ⚙️ Setup
1. Install [go](https://go.dev/dl)
2. Install [VS code](https://code.visualstudio.com/download) and its recommended extensions
3. Install `htop` - process viewer 

## Why Go ?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No classes (Only structs)
    - No inheritance (Only composition)
    - No exceptions (Only errors & errors are just values)
    - No try..catch..finally
    - No implicit type conversion
- Performance
    - Equivalent to C++
    - Compiled to machine code
    - Tooling system supports cross-compilation
- Concurrrency
    - Managed concurrency model (built in scheduler)
    - Concurrent operations are represented as "goroutines"
    - Goroutines are cheap (4KB)
    - Support for concurrency is offered in the "language"
        - "go" keyword
        - "chan" data type
        - "<-" operator (channel)
        - "range" construct
        - "select-case" construct
    - Standard Library Support
        - "sync" package
        - "sync/atomic" package

## Refer
[Notes](NOTES.md)