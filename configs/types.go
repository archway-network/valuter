package configs

type Configuration struct {
	// GRPC struct {
	// 	Server       string `json:"server"`
	// 	TLS          bool   `json:"tls"`
	// 	APICallRetry int    `json:"api_call_retry"`
	// 	CallTimeout  int    `json:"call_timeout"`
	// } `json:"grpc"`

	Tasks struct {
		Gov struct {
			MaxWinners int      `json:"max_winners"` // Max number of winners for this tasks
			Proposals  []uint64 `json:"proposals"`   // The list of Proposal Ids to be investigated
			Reward     uint64   `json:"reward"`      // Reward for each winner
		} `json:"gov"`

		ValidatorGenesis struct { // Validators who joined at genesis
			MaxWinners int    `json:"max_winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
			Condition  struct {
				JoinHight   uint64 `json:"join_hight"`   // The Block hight that the validator is suppose to join
				SignPercent int    `json:"sign_percent"` // The percentage of blocks that should be signed by the validator
			} `json:"condition"` // Winner Condition

		} `json:"validator_genesis"`

		ValidatorJoin struct { // Validators who joined after genesis
			MaxWinners int    `json:"max_winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
			Condition  struct {
				JoinHight   uint64 `json:"join_hight"`   // The Block hight that the validator is suppose to join
				SignPercent int    `json:"sign_percent"` // The percentage of blocks that should be signed by the validator
			} `json:"condition"` // Winner Condition

		} `json:"validator_join"`

		JailUnjail struct {
			MaxWinners int    `json:"max_winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"jail_unjail"`

		Staking struct {
			MaxWinners int    `json:"max_winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
		} `json:"staking"`

		UpTime struct {
			MaxWinners int    `json:"max_winners"` // Max number of winners for this tasks
			Reward     uint64 `json:"reward"`      // Reward for each winner
			Conditions []struct {
				StartHight    uint64 `json:"start_hight"`    // The Block hight that the load burst starts
				EndHight      uint64 `json:"end_hight"`      // The Block hight that the load burst ends
				UptimePercent int    `json:"uptime_percent"` // The percentage of blocks that the winners must not miss to sign
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
	} `json:"bech32_prefix"`

	BlockExplorer struct {
		TxHash    string `json:"tx_hash"`
		Account   string `json:"account"`
		Validator string `json:"validator"`
	} `json:"block_explorer"`

	Report struct {
		OutputDir string `json:"output_dir"`
	} `json:"report"`

	IdVerification struct {
		Required   bool `json:"required"`    // If it is required to do an ID verification and filter out the not-verified users
		HTMLReport bool `json:"html_report"` // If the ID verification data should be shown in the HTML report
		InputFile  struct {
			Path   string `json:"path"` // Path to the CSV file containing the verification data
			Fields struct {
				Email string `json:"email"`
				KYCId string `json:"kyc_id"`
			} `json:"fields"`
		} `json:"input_file"`
		VerifierAccount string `json:"verifier_account"` // An account that all ID verification tx is sent to (in its Memo field)
	} `json:"id_verification"`

	API struct {
		RowsPerPage uint64 `json:"rows_per_page"`
	} `json:"api"`
}
