package winners

type Winner struct {
	Address          string               // Account Address
	Rewards          uint64               // Total Reward of a winner
	Verified         bool                 // If the ID of this winner account is verified
	VerificationData VerificationDataType // When we verify the user's data, we keep a copy of the verification data here for further investigation
	// Timestamp        string               // The time of the task done, if applicable
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
