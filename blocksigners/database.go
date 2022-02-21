package blocksigners

import (
	"time"

	"github.com/archway-network/cosmologger/database"
	"github.com/archway-network/valuter/types"
)

func DBRowToBlockSignersRecord(row database.RowType) types.BlockSignersRecord {

	if row == nil {
		return types.BlockSignersRecord{}
	}

	for i := range row {
		if row[i] == nil {
			row[i] = ""
		}
	}

	return types.BlockSignersRecord{

		BlockHeight: row[database.FIELD_BLOCK_SIGNERS_BLOCK_HEIGHT].(uint64),
		ValConsAddr: row[database.FIELD_BLOCK_SIGNERS_VAL_CONS_ADDR].(string),
		Time:        row[database.FIELD_BLOCK_SIGNERS_TIME].(time.Time),
		Signature:   row[database.FIELD_BLOCK_SIGNERS_SIGNATURE].(string),
	}
}

func DBRowsToBlockSignersRecords(row []database.RowType) []types.BlockSignersRecord {

	var res []types.BlockSignersRecord
	for i := range row {
		res = append(res, DBRowToBlockSignersRecord(row[i]))
	}

	return res
}
