package configs

type Configuration struct {
	// GRPC struct {
	// 	Server       string `json:"server"`
	// 	TLS          bool   `json:"tls"`
	// 	APICallRetry int    `json:"api-call-retry"`
	// 	CallTimeout  int    `json:"call-timeout"`
	// } `json:"grpc"`

	Tasks struct {
		Gov struct {
			MaxWinners int      `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64   `json:"reward"`      // Reward for each winner
			Proposals  []uint64 `json:"proposals"`   // The list of Proposal Ids to be investigated
		} `json:"gov"`

		ValidatorGenesis struct { // Validators who joined at genesis
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"validators-genesis"`

		ValidatorJoin struct { // Validators who joined after genesis
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"validators-joined"`

		JailUnjail struct {
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"jail-unjail"`

		Staking struct {
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"staking"`

		NodeUpgrade struct {
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner

			Condition struct {
				UpgradeHight uint64 `json:"upgrade-hight"` // The Block hight that the upgraded validator must sign

			} `json:"condition"` // Winner Condition

		} `json:"node-upgrade"`

		UpTime struct {
			MaxWinners int    `json:"max-winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner

			// There might be multiple load bursts
			Conditions []struct {
				StartHight    uint64  `json:"start-hight"`    // The Block hight that the load burst starts
				EndHight      uint64  `json:"end-hight"`      // The Block hight that the load burst ends
				UptimePercent float32 `json:"uptime-percent"` // The percentage of blocks that the winners must not miss to sign
			} `json:"conditions"` // Uptime Conditions

		} `json:"uptime"`
	} `json:"tasks"`

	Bech32Prefix struct {
		Account struct {
			Address string `json:"address"`
			PubKey  string `json:"pubkey"`
		} `json:"account"`

		Validator struct {
			Address string `json:"address"`
			PubKey  string `json:"pubkey"`
		} `json:"validator"`

		Consensus struct {
			Address string `json:"address"`
			PubKey  string `json:"pubkey"`
		} `json:"consensus"`
	} `json:"bech32-prefix"`

	BlockExplorer struct {
		TxHash    string `json:"tx-hash"`
		Account   string `json:"account"`
		Validator string `json:"validator"`
	} `json:"block-explorer"`

	IdVerification struct {
		Required  bool `json:"required"` // If it is required to do an ID verification and filter out the not-verified users
		InputFile struct {
			Path   string `json:"path"` // Path to the CSV file containing the verification data
			Fields struct {
				Email string `json:"email"`
				KYCId string `json:"kyc-id"`
			} `json:"fields"`
		} `json:"input-file"`
		VerifierAccount string `json:"verifier-account"` // An account that all ID verification tx is sent to (in its Memo field)
	} `json:"id-verification"`

	API struct {
		RowsPerPage uint64 `json:"rows-per-page"`
	} `json:"api"`
}
