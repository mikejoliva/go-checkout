# Go Checkout

This is a simple SKU checkout implemented in Go.

## Prerequisties 

To compile and run this repo you will need to install Go >= 1.24.1, download here: https://go.dev/dl/.

## Running

To build and run the project, execute the following command from the project root directory:

```shell
go run . <SKU code(s)>
```

Where `<SKU code(s)>` is a string of products which can be found in `config/stocklist.yaml`, i.e.: `AABAACCDAA`.

## Testing

The test suite can be run with the following command:
```shell
go test ./...
```