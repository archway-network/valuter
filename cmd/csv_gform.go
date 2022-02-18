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
		fmt.Println(err)
		if err != nil {
			return err
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		csvReader.Comma = rune(delimiter[0])

		var jsonCols []int
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
					}
				}
				continue
			}

			// Processing the rows
			for i := range jsonCols {
				err := participants.Import(rec[jsonCols[i]])
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(importGFormCSV)
	importGFormCSV.Flags().StringP(FLAG_CSV_DELIMITER, "d", ",", "CSV delimiter")
	importGFormCSV.Flags().StringP(FLAG_JSON_KEYWORD, "j", "json", "Keyword search for JSON ID")
}
