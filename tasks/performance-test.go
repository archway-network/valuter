package tasks

import (
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/validators"
	"github.com/archway-network/valuter/winners"
)

func GetPerformanceTestWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	for i := range configs.Configs.Tasks.UpTime.Conditions {

		roundWinnerList, err := GetPerformanceTestWinnersPerLoadBurst(i)
		if err != nil {
			return winnersList, err
		}

		winnersList.MergeWithAggregateRewards(roundWinnerList)
	}

	return winnersList, nil
}

// This function receives a load index which is just the index in the config file
// for the load burst which starts from zero
func GetPerformanceTestWinnersPerLoadBurst(loadIndex int) (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.UpTime.MaxWinners == 0 {
		return winnersList, nil
	}

	listOfValidators, err := GetValidatorsSortedByUpTimeInBlockHeightRange(
		configs.Configs.Tasks.UpTime.Conditions[loadIndex].StartHight,
		configs.Configs.Tasks.UpTime.Conditions[loadIndex].EndHight,
	)
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		// Filter out those who could not maintaine the required uptime
		if listOfValidators[i].UpTime < configs.Configs.Tasks.UpTime.Conditions[loadIndex].UptimePercent {
			continue
		}

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
			Rewards:         configs.Configs.Tasks.UpTime.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)
		if winnersList.Length() >= configs.Configs.Tasks.UpTime.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}

func GetValidatorsSortedByUpTimeInBlockHeightRange(beginHeight, endHeight uint64) ([]validators.ValidatorInfo, error) {

	listOfValidators, err := validators.GetAllValidators()
	if err != nil {
		return nil, err
	}

	// get validators info which includes uptime in the given range
	var listOfValidatorsInfo []validators.ValidatorInfo
	for i := range listOfValidators {

		vInfo, err := listOfValidators[i].GetValidatorInfoByBlockHeightRange(beginHeight, endHeight)
		if err != nil {
			return nil, err
		}
		listOfValidatorsInfo = append(listOfValidatorsInfo, vInfo)
	}

	// Sorting based on uptime
	// Note: since this is not a performance critical,
	// we implemented a simple bubble sort
	for j := 0; j < len(listOfValidatorsInfo); j++ {
		modified := false
		for i := 0; i < len(listOfValidatorsInfo)-1; i++ {
			if listOfValidatorsInfo[i].UpTime < listOfValidatorsInfo[i+1].UpTime {
				// swap
				tmp := listOfValidatorsInfo[i]
				listOfValidatorsInfo[i] = listOfValidatorsInfo[i+1]
				listOfValidatorsInfo[i+1] = tmp
				modified = true
			}
		}
		// Already sorted, just for having a bit better performance
		if !modified {
			break
		}
	}

	return listOfValidatorsInfo, nil
}
