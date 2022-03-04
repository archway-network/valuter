package cmd

import (
	"fmt"
	"os"

	agSigner "github.com/archway-network/augusta-testnet-signer/types"
	"github.com/archway-network/synaps-verifier/synaps"
	"github.com/archway-network/valuter/participants"
	"github.com/spf13/cobra"
)

const (
	FLAG_KYC_API_KEY   = "apikey"
	FLAG_KYC_CLIENT_ID = "client-id"
	FLAG_KYC_API_PATH  = "api-path"
)

var kycVerifyCmd = &cobra.Command{
	Use:   "kyc-verify [flags]",
	Short: "Connects to the KYC provider and import all the verification data",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		/*-----------*/

		kycApiKey, err := cmd.Flags().GetString(FLAG_KYC_API_KEY)
		if err != nil {
			return err
		}

		// Since synaps package reads KYC API info from ENV vars,
		// let's set it if user sets a custom value for them
		if kycApiKey != "" {
			os.Setenv("KYC_API_KEY", kycApiKey)
		}

		kycClientId, err := cmd.Flags().GetString(FLAG_KYC_CLIENT_ID)
		if err != nil {
			return err
		}

		// Since synaps package reads KYC API info from ENV vars,
		// let's set it if user sets a custom value for them
		if kycClientId != "" {
			os.Setenv("KYC_CLIENT_ID", kycClientId)
		}

		kycApiPath, err := cmd.Flags().GetString(FLAG_KYC_API_PATH)
		if err != nil {
			return err
		}

		// Since synaps package reads KYC API info from ENV vars,
		// let's set it if user sets a custom value for them
		if kycApiPath != "" {
			os.Setenv("KYC_API_PATH", kycApiPath)
		}

		/*-----------*/

		fmt.Printf("\nRetrieveing the participants KYC list...\n")

		// synaps.GetPendingSessions()
		ss, err := synaps.GetFinishedSessions()
		if err != nil {
			return err
		}

		fmt.Printf("[ %d ] KYC data found.\n", len(ss))

		for i := range ss {
			fmt.Printf("\nProcessing [ %s ] \twith session id [ %s ]...", ss[i].Alias, ss[i].SessionId)
			user, err := synaps.GetSessionDetails(ss[i].SessionId)
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
				continue
			}

			participant := participants.ParticipantRecord{
				ID: agSigner.ID{
					EmailAddress: user.Alias,
				},
				KycSessionId: user.SessionId,
				KycVerified:  user.IsVerified(),
			}

			if participant.KycVerified {
				fmt.Printf(" ** Verified ** ")
			}

			rowsUpdated, err := participant.UpdateKYC()
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
				continue
			}

			fmt.Printf("\t [ %d ] rows updated.", rowsUpdated)
		}

		fmt.Printf("\n\n All done.\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(kycVerifyCmd)

	kycVerifyCmd.Flags().StringP(FLAG_KYC_API_KEY, "k", "", "KYC Api-Key")
	kycVerifyCmd.Flags().StringP(FLAG_KYC_CLIENT_ID, "c", "", "KYC Client-Id")
	kycVerifyCmd.Flags().StringP(FLAG_KYC_API_PATH, "a", "", "KYC API Path")
}
