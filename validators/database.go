package validators

import (
	"github.com/archway-network/valuter/database"
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
