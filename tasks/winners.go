package tasks

import "github.com/archway-network/valuter/winners"

func GetAllWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	w, err := GetGenesisValidatorsWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetJoinedAfterGenesisValidatorsWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetUnjailedValidatorsWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetStakingWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetGovWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetNodeUpgradeWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	w, err = GetPerformanceTestWinners()
	if err != nil {
		return winnersList, err
	}
	winnersList.MergeWithAggregateRewards(w)

	return winnersList, nil
}

// Since this function is not used frequently, let's have an easy implementation
func GetWinnerByAddress(address string) ([]winners.WinnerChallenge, error) {

	results := make([]winners.WinnerChallenge, 10)

	w, err := GetGenesisValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "validator-genesis",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetJoinedAfterGenesisValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "validator-join",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetUnjailedValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "jail-unjail",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetStakingWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "staking",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetGovWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "gov",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetNodeUpgradeWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "node-upgrade",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	w, err = GetPerformanceTestWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results = append(results, winners.WinnerChallenge{
			Challenge: "uptime",
			Rewards:   w.GetItem(index).Rewards,
		})
	}

	total := uint64(0)
	for i := range results {
		total += results[i].Rewards
	}
	results = append(results, winners.WinnerChallenge{
		Challenge: "total",
		Rewards:   total,
	})

	return results, nil
}
