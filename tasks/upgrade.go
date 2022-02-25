package tasks

import (
	"github.com/archway-network/valuter/blocksigners"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/winners"
)

func GetNodeUpgradeWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.NodeUpgrade.MaxWinners == 0 {
		return winnersList, nil
	}

	listOfValidators, err := blocksigners.GetSignersByBlockHeight(configs.Configs.Tasks.NodeUpgrade.Condition.UpgradeHight)
	if err != nil {
		return winnersList, err
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
			Rewards:         configs.Configs.Tasks.NodeUpgrade.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)
		if winnersList.Length() >= configs.Configs.Tasks.NodeUpgrade.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
