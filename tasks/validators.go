package tasks

import (
	"github.com/archway-network/valuter/winners"
)

func GetGenesisValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList
	// var listOfValidators []types.Validator
	// var err error

	// listOfValidators, _, err = GetGenesisValidators()
	// if err != nil {
	// 	return winners.WinnersList{}, err
	// }

	// for i := range listOfValidators {

	// 	newWinner := winners.Winner{
	// 		Address: listOfValidators[i].ConsAddr,
	// 		Rewards: configs.Configs.Tasks.ValidatorGenesis.Reward,
	// 	}

	// 	// if configs.Configs.IdVerification.Required {
	// 	// 	verified, err := newWinner.Verify(conn)
	// 	// 	if err != nil {
	// 	// 		return winners.WinnersList{}, err
	// 	// 	}
	// 	// 	if !verified {
	// 	// 		continue //ignore the unverified winners
	// 	// 	}
	// 	// }

	// 	winnersList.Append(newWinner)
	// 	if winnersList.Length() >= configs.Configs.Tasks.ValidatorGenesis.MaxWinners {
	// 		break // Max winners reached
	// 	}
	// }

	return winnersList, nil
}

// //TODO: We need to have a proper condition to filter in the active validators
// func GetActiveValidators(limitOffset types.DBLimitOffset) ([]validators.Validator, types.Pagination, error) {
// 	return validators.GetValidators(limitOffset)
// }

// //TODO: We need to have a proper condition to filter only the genesis validators
// func GetGenesisValidators() ([]validators.Validator, types.Pagination, error) {
// 	return validators.GetValidators(types.DBLimitOffset{Limit: 150})
// }

// //TODO: We need to have a proper condition to filter the validators that joined after the genesis
// func GetJoinedAfterGenesisValidators() ([]validators.Validator, types.Pagination, error) {
// 	return validators.GetValidators(types.DBLimitOffset{Limit: 150})
// }
