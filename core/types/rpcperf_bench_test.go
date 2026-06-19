package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func benchLogs(n int) []*Log {
	logs := make([]*Log, n)
	for i := range logs {
		logs[i] = &Log{
			Address:        common.Address{0x11, 0x22, byte(i)},
			Topics:         []common.Hash{{0x01, byte(i)}, {0x02, byte(i)}, {0x03, byte(i)}},
			Data:           make([]byte, 64),
			BlockNumber:    uint64(18000000 + i),
			TxHash:         common.Hash{0xaa, byte(i)},
			TxIndex:        uint(i),
			BlockHash:      common.Hash{0xbb, byte(i)},
			BlockTimestamp: uint64(1700000000 + i),
			Index:          uint(i),
		}
	}
	return logs
}

func benchHeader() *Header {
	return &Header{
		ParentHash:  common.Hash{0x01},
		UncleHash:   common.Hash{0x02},
		Coinbase:    common.Address{0x03},
		Root:        common.Hash{0x04},
		TxHash:      common.Hash{0x05},
		ReceiptHash: common.Hash{0x06},
		Bloom:       Bloom{0x07, 0x08, 0x09},
		Difficulty:  big.NewInt(1),
		Number:      big.NewInt(18000000),
		GasLimit:    30000000,
		GasUsed:     15000000,
		Time:        1700000000,
		Extra:       make([]byte, 32),
		MixDigest:   common.Hash{0x0a},
	}
}

func BenchmarkMarshalGetLogs(b *testing.B) {
	logs := benchLogs(200)
	b.ReportAllocs()
	for b.Loop() {
		if _, err := json.Marshal(logs); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalHeader(b *testing.B) {
	h := benchHeader()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := json.Marshal(h); err != nil {
			b.Fatal(err)
		}
	}
}
