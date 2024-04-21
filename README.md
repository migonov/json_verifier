# How to run

Go version used in this project - 1.22.2

### Project

You can run the whole project using

```
go run .
```

in the main directory - it will validate the *data.json* file content which should be a valid AWS:IAM:Role Policy JSON document.

### Tests

To run the tests for the whole project use

```
go test ./...
```

in the main directory.

# Overview

### Verifier package

Verifier package contains the main function **Verifier(inputJSON []byte)** which is the subject of the task.

### Utils package

Utils package contains **StringOrSlice** struct that is used to represent the *Resource* field in the input JSON document.

*Resource* field can be a single string or an array of strings. It is always parsed to a string slice.