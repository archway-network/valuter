package tx

import (
	"time"

	"github.com/archway-network/valuter/database"
	"github.com/archway-network/valuter/types"
)

func DBRowToTxRecord(row database.RowType) types.TxRecord {

	return types.TxRecord{

		TxHash:      string(row[database.FIELD_TX_EVENTS_TX_HASH].([]uint8)),
		Height:      row[database.FIELD_TX_EVENTS_HEIGHT].(uint64),
		Module:      row[database.FIELD_TX_EVENTS_MODULE].(string),
		Sender:      row[database.FIELD_TX_EVENTS_SENDER].(string),
		Receiver:    row[database.FIELD_TX_EVENTS_RECEIVER].(string),
		Validator:   row[database.FIELD_TX_EVENTS_VALIDATOR].(string),
		Action:      row[database.FIELD_TX_EVENTS_ACTION].(string),
		Amount:      row[database.FIELD_TX_EVENTS_AMOUNT].(string),
		TxAccSeq:    row[database.FIELD_TX_EVENTS_TX_ACCSEQ].(string),
		TxSignature: row[database.FIELD_TX_EVENTS_TX_SIGNATURE].(string),
		ProposalId:  uint64(row[database.FIELD_TX_EVENTS_PROPOSAL_ID].(int64)),
		TxMemo:      row[database.FIELD_TX_EVENTS_TX_MEMO].(string),
		Json:        row[database.FIELD_TX_EVENTS_JSON].(string),
		LogTime:     row[database.FIELD_TX_EVENTS_LOG_TIME].(time.Time),
	}
}

func DBRowsToTxRecords(row []database.RowType) []types.TxRecord {

	var res []types.TxRecord
	for i := range row {
		res = append(res, DBRowToTxRecord(row[i]))
	}

	return res

}
