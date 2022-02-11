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
func GetWinnerByAddress(address string) (map[string]uint64, error) {

	results := make(map[string]uint64)

	w, err := GetGenesisValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["validator_genesis"] = w.GetItem(index).Rewards
	}

	w, err = GetJoinedAfterGenesisValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["validator_join"] = w.GetItem(index).Rewards
	}

	w, err = GetUnjailedValidatorsWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["jail_unjail"] = w.GetItem(index).Rewards
	}

	w, err = GetStakingWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["staking"] = w.GetItem(index).Rewards
	}

	w, err = GetGovWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["gov"] = w.GetItem(index).Rewards
	}

	w, err = GetNodeUpgradeWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["node_upgrade"] = w.GetItem(index).Rewards
	}

	w, err = GetPerformanceTestWinners()
	if err != nil {
		return nil, err
	}
	if index := w.FindByAddress(address); index != -1 {
		results["uptime"] = w.GetItem(index).Rewards
	}

	total := uint64(0)
	for i := range results {
		total += results[i]
	}
	results["total"] = total

	return results, nil
}
