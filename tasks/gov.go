package tasks

import (
	"fmt"

	"github.com/archway-network/cosmologger/database"
	cosmoLogTx "github.com/archway-network/cosmologger/tx"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/winners"
)

func GetGovWinnersPerProposal(proposalId uint64) (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.Gov.MaxWinners == 0 {
		return winnersList, nil
	}

	// If someone has done the task more than once, there will be more than a record here,
	// But that's not a problem, as winners list is distinct
	SQL := fmt.Sprintf(`
		SELECT "%s", "%s" 
		FROM "%s" 
		WHERE 
			"%s" = $1 AND 
			"%s" = $2
		ORDER BY "%s" ASC`, // >= 2 is because we want the participants to delegate at least to two validators

		database.FIELD_TX_EVENTS_SENDER,
		database.FIELD_TX_EVENTS_HEIGHT,

		database.TABLE_TX_EVENTS,

		database.FIELD_TX_EVENTS_ACTION,
		database.FIELD_TX_EVENTS_PROPOSAL_ID,

		database.FIELD_TX_EVENTS_HEIGHT,
	)

	rows, err := database.DB.Query(SQL,
		database.QueryParams{
			cosmoLogTx.ACTION_VOTE,
			proposalId,
		})
	if err != nil {
		return winnersList, err
	}

	for i := range rows {

		address := rows[i][database.FIELD_TX_EVENTS_SENDER].(string)

		pRecord, err := participants.GetParticipantByAddress(address)
		if err != nil {
			return winnersList, err
		}

		// If the participant is not verified by KYC provider, just ignore it
		if !pRecord.KycVerified {
			continue
		}

		newWinner := winners.Winner{
			Address:         address,
			Rewards:         configs.Configs.Tasks.Gov.Reward,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)

		if winnersList.Length() >= configs.Configs.Tasks.Gov.MaxWinners {
			break // Max winners reached
		}

	}

	return winnersList, nil
}

func GetGovWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	for i := range configs.Configs.Tasks.Gov.Proposals {

		proposalId := uint64(configs.Configs.Tasks.Gov.Proposals[i])
		proposalWinnerList, err := GetGovWinnersPerProposal(proposalId)
		if err != nil {
			return winnersList, err
		}

		winnersList.MergeWithAggregateRewards(proposalWinnerList)
	}

	return winnersList, nil
}
