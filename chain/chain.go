package chain

import "strconv"

type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

const (
	EthereumMainnet ID = 1
	EthereumRinkeby ID = 4
	EthereumGorli   ID = 5
	EthereumKovan   ID = 42

	PolygonMainnet       ID = 137
	PolygonMumbaiTestnet ID = 80001

	BSCMainnet ID = 56
	BSCTestnet ID = 97
)
