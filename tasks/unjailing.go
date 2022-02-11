package tasks

import (
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/validators"
	"github.com/archway-network/valuter/winners"
)

func GetUnjailedValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	activeValidators, err := validators.GetUnjailedValidators()
	if err != nil {
		return winners.WinnersList{}, err
	}

	for i := range activeValidators {

		newWinner := winners.Winner{
			Address: activeValidators[i].OprAddr,
			Rewards: configs.Configs.Tasks.JailUnjail.Reward,
		}

		// if configs.Configs.IdVerification.Required {
		// 	verified, err := newWinner.Verify(conn)
		// 	if err != nil {
		// 		return winners.WinnersList{}, err
		// 	}
		// 	if !verified {
		// 		continue //ignore the unverified winners
		// 	}
		// }

		winnersList.Append(newWinner)

		if winnersList.Length() >= configs.Configs.Tasks.JailUnjail.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
