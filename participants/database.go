package participants

import (
	agSigner "github.com/archway-network/augusta-testnet-signer/types"
	"github.com/archway-network/cosmologger/database"
)

func DBRowToParticipantRecord(row database.RowType) ParticipantRecord {

	if row == nil {
		return ParticipantRecord{}
	}

	for i := range row {
		if row[i] == nil {
			row[i] = ""
		}
	}

	return ParticipantRecord{

		ID: agSigner.ID{
			AccountAddress: row[database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS].(string),
			FullLegalName:  row[database.FIELD_PARTICIPANTS_FULL_LEGAL_NAME].(string),
			GithubHandle:   row[database.FIELD_PARTICIPANTS_GITHUB_HANDLE].(string),
			EmailAddress:   row[database.FIELD_PARTICIPANTS_EMAIL_ADDRESS].(string),
			PubKey:         row[database.FIELD_PARTICIPANTS_PUBKEY].(string),
		},
		KycSessionId: row[database.FIELD_PARTICIPANTS_KYC_SESSION_ID].(string),
		KycVerified:  row[database.FIELD_PARTICIPANTS_KYC_VERIFIED].(bool),
	}
}

func DBRowToParticipantRecords(row []database.RowType) []ParticipantRecord {

	var res []ParticipantRecord
	for i := range row {
		res = append(res, DBRowToParticipantRecord(row[i]))
	}

	return res
}

func (p ParticipantRecord) getDBRow() database.RowType {
	return database.RowType{

		database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS: p.AccountAddress,
		database.FIELD_PARTICIPANTS_FULL_LEGAL_NAME: p.FullLegalName,
		database.FIELD_PARTICIPANTS_GITHUB_HANDLE:   p.GithubHandle,
		database.FIELD_PARTICIPANTS_EMAIL_ADDRESS:   p.EmailAddress,
		database.FIELD_PARTICIPANTS_PUBKEY:          p.PubKey,
		database.FIELD_PARTICIPANTS_KYC_SESSION_ID:  p.KycSessionId,
		database.FIELD_PARTICIPANTS_KYC_VERIFIED:    p.KycVerified,
	}
}
