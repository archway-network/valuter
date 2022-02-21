package tasks

import (
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/validators"
	"github.com/archway-network/valuter/winners"
)

func GetUnjailedValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	listOfValidators, err := validators.GetUnjailedValidators()
	if err != nil {
		return winners.WinnersList{}, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
		}

		pRecord, err := participants.GetParticipantByAddress(listOfValidators[i].AccAddr)
		if err != nil {
			return winnersList, err
		}

		newWinner := winners.Winner{
			Address:         listOfValidators[i].AccAddr,
			Rewards:         configs.Configs.Tasks.JailUnjail.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)

		if winnersList.Length() >= configs.Configs.Tasks.JailUnjail.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
