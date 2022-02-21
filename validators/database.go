package validators

import (
	"time"

	"github.com/archway-network/cosmologger/database"
)

func DBRowToValidatorRecord(row database.RowType) ValidatorRecord {

	if row == nil {
		return ValidatorRecord{}
	}

	return ValidatorRecord{
		ConsAddr: row[database.FIELD_VALIDATORS_CONS_ADDR].(string),
		OprAddr:  row[database.FIELD_VALIDATORS_OPR_ADDR].(string),
	}
}

func DBRowToValidatorRecords(row []database.RowType) []ValidatorRecord {

	var res []ValidatorRecord
	for i := range row {
		res = append(res, DBRowToValidatorRecord(row[i]))
	}

	return res
}

func DBRowToValidatorWithTx(row database.RowType) ValidatorWithTx {

	if row == nil {
		return ValidatorWithTx{}
	}

	return ValidatorWithTx{
		ValidatorRecord: ValidatorRecord{
			ConsAddr: row[database.FIELD_VALIDATORS_CONS_ADDR].(string),
			OprAddr:  row[database.FIELD_VALIDATORS_OPR_ADDR].(string),
		},
		TxHash:  string(row[database.FIELD_TX_EVENTS_TX_HASH].([]uint8)), //char
		Height:  uint64(row[database.FIELD_TX_EVENTS_HEIGHT].(int64)),
		Sender:  row[database.FIELD_TX_EVENTS_SENDER].(string),
		LogTime: row[database.FIELD_TX_EVENTS_LOG_TIME].(time.Time),
	}
}

func DBRowToValidatorWithTxs(row []database.RowType) []ValidatorWithTx {

	var res []ValidatorWithTx
	for i := range row {
		res = append(res, DBRowToValidatorWithTx(row[i]))
	}

	return res
}
