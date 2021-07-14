module perun.network/go-perun

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.1
	github.com/miguelmota/go-ethereum-hdwallet v0.0.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace github.com/ethereum/go-ethereum => github.com/ggwpez/go-ethereum v1.10.2-0.20210614094901-426e0588f3e2
