package tasks

import (
	"sort"

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

		// If the participant is not verified by KYC provider, just ignore it
		if !pRecord.KycVerified {
			continue
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

	sort.Slice(listOfValidatorsInfo, func(i, j int) bool {
		return listOfValidatorsInfo[i].UpTime > listOfValidatorsInfo[j].UpTime
	})

	return listOfValidatorsInfo, nil
}
