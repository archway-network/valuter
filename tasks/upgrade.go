package tasks

import (
	"github.com/archway-network/valuter/blocksigners"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/winners"
)

func GetNodeUpgradeWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	listOfValidators, err := blocksigners.GetSignersByBlockHeight(configs.Configs.Tasks.NodeUpgrade.Condition.UpgradeHight)
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		newWinner := winners.Winner{
			Address: listOfValidators[i].OprAddr,
			Rewards: configs.Configs.Tasks.NodeUpgrade.Reward,
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
		if winnersList.Length() >= configs.Configs.Tasks.NodeUpgrade.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
