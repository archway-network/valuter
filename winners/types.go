package winners

import (
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/validators"
)

type Winner struct {
	Address         string                         // Account Address
	Rewards         uint64                         // Total Reward of a winner
	ParticipantData participants.ParticipantRecord // If the information of this winner account if verified
	ValidatorInfo   validators.ValidatorInfo       // If the winner is a validator, the data will be set here
}

type hashMapType map[string]int // This map is used for quick search { string: Winner.Address, int: index to the item in the WinnersList.list slice}

type WinnersList struct {
	list    []Winner // To keep the order of items we use the int index
	hashMap hashMapType
}

type VerificationDataType struct { //map[email]...
	Email string
	KYCId string
}

type WinnerChallenge struct {
	Challenge string
	Rewards   uint64
}
