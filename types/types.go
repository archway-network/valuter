package types

import "time"

type BlockSignersRecord struct {
	BlockHeight uint64
	ValConsAddr string
	Time        time.Time
	Signature   string
}

type BlockRecord struct {
	BlockHash string
	Height    uint64
	NumOfTxs  uint64
	Time      time.Time
	Signers   []BlockSignersRecord
}

type TxRecord struct {
	TxHash      string
	Height      uint64
	Module      string
	Sender      string
	Receiver    string
	Validator   string
	Action      string
	Amount      string
	TxAccSeq    string
	TxSignature string
	ProposalId  uint64
	TxMemo      string
	Json        string
	LogTime     time.Time
}

type DBLimitOffset struct {
	Limit  uint64
	Offset uint64
	Page   uint64
}

type Pagination struct {
	CurrentPage uint64 `json:"current_page"`
	TotalPages  uint64 `json:"total_pages"`
	TotalRows   uint64 `json:"total_entries"`
}
