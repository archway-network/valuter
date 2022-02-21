package participants

import (
	"encoding/json"
	"fmt"
	"strings"

	agSigner "github.com/archway-network/augusta-testnet-signer/types"
	"github.com/archway-network/cosmologger/database"
)

type ParticipantRecord agSigner.ID

// This function receives a json string of the signed ID,
// verifies it with the given signature and if it passes,
// the data will be added to the database
func Import(jsonStr string) error {

	container, err := getAgSignerContainer(jsonStr)
	if err != nil {
		return err
	}

	// The input string was empty
	if container == nil {
		return nil
	}

	verified, err := container.VerifySubmission()
	if err != nil {
		return err
	}
	// The data is not verified
	if !verified {
		return fmt.Errorf("the data is not verified")
	}

	// Let's add it to the database
	return AddNew(ParticipantRecord(container.ID))
}

func AddNew(participant ParticipantRecord) error {

	//Check if the record is already in the db
	queryRes, err := database.DB.Load(database.TABLE_PARTICIPANTS,
		database.RowType{
			database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS: participant.AccountAddress,
		})
	if err != nil {
		return err
	}

	// Already exist //TODO: We might want to update the record if it exist, need to decide on that
	if len(queryRes) > 0 {
		return nil
	}
	_, err = database.DB.Insert(database.TABLE_PARTICIPANTS, participant.getDBRow())
	return err
}

func getAgSignerContainer(jsonStr string) (*agSigner.Container, error) {

	jsonStr = strings.Trim(jsonStr, "\"\n\t\r' ")
	if jsonStr == "" {
		return nil, nil
	}

	var container agSigner.Container
	err := json.Unmarshal([]byte(jsonStr), &container)
	if err != nil {
		return nil, err
	}

	return &container, nil
}
