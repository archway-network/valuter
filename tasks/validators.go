package tasks

import (
	"github.com/archway-network/valuter/blocksigners"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/validators"
	"github.com/archway-network/valuter/winners"
)

func GetGenesisValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	// Those who signged the first block are considered as genesis validators
	// Since some joins might not be able to make it to the first block we change it to a higher block like 20
	listOfValidators, err := blocksigners.GetSignersByBlockHeight(20)
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
		}
		if valInfo.UpTime < configs.Configs.Tasks.ValidatorGenesis.UptimePercent {
			// Let's just ignore this validator
			continue
		}

		newWinner := winners.Winner{
			Address: listOfValidators[i].OprAddr,
			Rewards: configs.Configs.Tasks.ValidatorGenesis.Reward,
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
		if winnersList.Length() >= configs.Configs.Tasks.ValidatorGenesis.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}

func GetJoinedAfterGenesisValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	listOfValidators, err := validators.GetJoinedAfterGenesisValidators()
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
		}
		if valInfo.UpTime < configs.Configs.Tasks.ValidatorJoin.UptimePercent {
			// Let's just ignore this validator
			continue
		}

		newWinner := winners.Winner{
			Address: listOfValidators[i].OprAddr,
			Rewards: configs.Configs.Tasks.ValidatorJoin.Reward,
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
		if winnersList.Length() >= configs.Configs.Tasks.ValidatorJoin.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
