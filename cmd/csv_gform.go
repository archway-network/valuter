package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/archway-network/valuter/participants"
	"github.com/spf13/cobra"
)

const (
	FLAG_CSV_DELIMITER = "delimiter"
	FLAG_JSON_KEYWORD  = "json"

	NAME_COLUMN    = "legal name" // Part of the column name is enough (must be lowercase)
	EMAIL_COLUMN   = "username"   // Part of the column name is enough (must be lowercase)
	COUNTRY_COLUMN = "country"    // Part of the column name is enough (must be lowercase)
)

var importGFormCSV = &cobra.Command{
	Use:   "add-gform-csv [csv-file-path]",
	Short: "import a CSV file from the google form to verify and store participant's id data",
	Long:  `This command looks for ID JSON field of the given CSV file and uses "augusta-testnet-signer" to verify the submitted data.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		csvFilePath := args[0]

		delimiter, err := cmd.Flags().GetString(FLAG_CSV_DELIMITER)
		if err != nil {
			return err
		}

		jsonKeyword, err := cmd.Flags().GetString(FLAG_JSON_KEYWORD)
		if err != nil {
			return err
		}

		f, err := os.Open(csvFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		csvReader.Comma = rune(delimiter[0])

		fmt.Printf("\nProcessing the form data...\n")

		var jsonCols []int
		var emailCols []int
		var nameCols []int
		var countryCols []int
		for rowCounter := 0; ; rowCounter++ {
			rec, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			// The header
			if rowCounter == 0 {
				for i := range rec {
					if strings.Contains(strings.ToLower(rec[i]), jsonKeyword) {
						jsonCols = append(jsonCols, i)

					} else if strings.Contains(strings.ToLower(rec[i]), EMAIL_COLUMN) {
						emailCols = append(emailCols, i)

					} else if strings.Contains(strings.ToLower(rec[i]), NAME_COLUMN) {
						nameCols = append(nameCols, i)

					} else if strings.Contains(strings.ToLower(rec[i]), COUNTRY_COLUMN) {
						countryCols = append(countryCols, i)
					}
				}
				continue
			}

			// Processing the rows
			for i := range jsonCols {

				// fmt.Printf("\trec[jsonCols[i]]: %v\n\n", rec[jsonCols[i]])

				err := participants.ImportBySignature(rec[jsonCols[i]])
				if err != nil {
					fmt.Printf("\n====> Error on importing by Signature: %s \n%v\n", err, rec[jsonCols[i]])
				}
				fmt.Printf("\r\tProcessing record %5d", rowCounter)
			}

			/*------*/

			emailAddr := ""
			fullName := ""
			country := ""

			for i := range emailCols {
				emailAddr = rec[emailCols[i]]
				if emailAddr != "" {
					break
				}
			}

			for i := range nameCols {
				fullName = rec[nameCols[i]]
				if fullName != "" {
					break
				}
			}

			for i := range countryCols {
				country = rec[countryCols[i]]
				if country != "" {
					break
				}
			}

			err = participants.ImportByEmail(emailAddr, fullName, country)
			if err != nil {
				fmt.Printf("\n====> Error on importing by Email: %s \n%v\n", err, emailAddr)
			}
			fmt.Printf("\r\tProcessing record %5d", rowCounter)
		}
		fmt.Printf("\nDone.\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(importGFormCSV)
	importGFormCSV.Flags().StringP(FLAG_CSV_DELIMITER, "d", ",", "CSV delimiter")
	importGFormCSV.Flags().StringP(FLAG_JSON_KEYWORD, "j", "json", "Keyword search for JSON ID")
}
