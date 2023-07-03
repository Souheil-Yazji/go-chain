# go-chain

## Simple Blockchain with Golang

This is a simple blockchain implementation using Golang. It demonstrates the basic concepts of blockchain technology, including block creation, hashing, validation, and persistence using BadgerDB. More to come soon.

## Features

- Block creation with index, timestamp, data, and hash.
- Hash calculation using SHA256 algorithm.
- Block validation based on the previous block's hash.
- Persistence of the blockchain using BadgerDB.

## Prerequisites

- Go programming language (version 1.20.x)
- BadgerDB v4

## Installation

1. Clone the repository: 
```git clone https://github.com/Souheil-Yazji/go-chain.git```
2. Change to the project directory: 
```cd go-chain```
3. Install the project dependencies: 
```go mod download```


## Usage

1. Run the application: 
``` go run main.go```

2. No real way to interact with the application yet, coming soon!

## Testing

The project includes tests to ensure the correctness of the blockchain functionality. To run the tests, use the following command:

```go test ./...```


## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Submit a pull request to the main repository.

## License

This project is licensed under the [MIT License](LICENSE).


## Contact

For any questions or feedback, please make a github issue.








