## Go wrapper for G.729

This package provides Go bindings for the BelledonneCommunications C library bcg729.

The C libraries and docs are provided in https://github.com/BelledonneCommunications/bcg729.
This package just handles the wrapping in Go, and is unaffiliated with BelledonneCommunications.

### Prerequisites

Install the C library bcg729 in  https://github.com/BelledonneCommunications/bcg729.
You have to install the library, and add the path library installed to environment variable.

### Examples

See the examples in the examples folder.
You can test encoder, decoder example with sample file.

```bash
go run examples/main.go
```
