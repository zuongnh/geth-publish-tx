## Prerequisites

solc - Solidity Compiler

```
brew update
brew tap ethereum/ethereum
brew install solidity
```

abigen

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

---

## How to create necessary files (config)

Import Zombie ABI file into `./contracts`

- Create file `Store_sol_Store.abi`
- Copy Zombie ABI into that file

Run the following command (to create `Store.go` file from the Zombie ABI)

```
abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

---

## How to run code to publish transactions

Run command: (Publish Attack Transaction)

```
go run contract-write.go
```

## References
- https://goethereumbook.org/en/
