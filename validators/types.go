package validators

import "time"

type ValidatorRecord struct {
	ConsAddr string `json:"cons_address"`
	OprAddr  string `json:"opr_address"`
	AccAddr  string `json:"acc_address"`
	Moniker  string `json:"moniker"`
}

type ValidatorInfo struct {
	ValidatorRecord        `json:"validator"`
	TotalSignedBlocks      uint64  `json:"total_signed_blocks"`
	FirstSignedBlockHeight uint64  `json:"first_signed_block"`
	UpTime                 float32 `json:"up_time"`     // From the joining block
	UpTimeAll              float32 `json:"up_time_all"` // From  the first block
}

type ValidatorWithTx struct {
	ValidatorRecord `json:"validator"`
	TxHash          string    `json:"tx_hash"`
	Height          uint64    `json:"height"`
	Sender          string    `json:"sender"`
	LogTime         time.Time `json:"logtime"`
}
