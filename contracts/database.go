package contracts

import (
	"github.com/archway-network/cosmologger/block"
	"github.com/archway-network/cosmologger/database"
)

func DBRowToContractRecord(row database.RowType) block.ContractRecord {

	if row == nil {
		return block.ContractRecord{}
	}

	return block.ContractRecord{
		ContractAddress:  row[database.FIELD_CONTRACTS_CONTRACT_ADDRESS].(string),
		RewardAddress:    row[database.FIELD_CONTRACTS_REWARD_ADDRESS].(string),
		DeveloperAddress: row[database.FIELD_CONTRACTS_DEVELOPER_ADDRESS].(string),
		BlockHeight:      uint64(row[database.FIELD_CONTRACTS_BLOCK_HEIGHT].(int64)),
		GasConsumed:      uint64(row[database.FIELD_CONTRACTS_GAS_CONSUMED].(int64)),
		ContractRewards: block.GasTrackerReward{
			Denom:  row[database.FIELD_CONTRACTS_REWARDS_DENOM].(string),
			Amount: row[database.FIELD_CONTRACTS_CONTRACT_REWARDS_AMOUNT].(float64),
		},
		InflationRewards: block.GasTrackerReward{
			Denom:  row[database.FIELD_CONTRACTS_REWARDS_DENOM].(string),
			Amount: row[database.FIELD_CONTRACTS_INFLATION_REWARDS_AMOUNT].(float64),
		},
		LeftoverRewards: block.GasTrackerReward{
			Denom:  row[database.FIELD_CONTRACTS_REWARDS_DENOM].(string),
			Amount: row[database.FIELD_CONTRACTS_LEFTOVER_REWARDS_AMOUNT].(float64),
		},
		CollectPremium:           row[database.FIELD_CONTRACTS_COLLECT_PREMIUM].(bool),
		GasRebateToUser:          row[database.FIELD_CONTRACTS_GAS_REBATE_TO_USER].(bool),
		PremiumPercentageCharged: uint64(row[database.FIELD_CONTRACTS_PREMIUM_PERCENTAGE_CHARGED].(int64)),
		MetadataJson:             row[database.FIELD_CONTRACTS_METADATA_JSON].(string),
	}
}

func DBRowToContractRecords(row []database.RowType) []block.ContractRecord {

	var res []block.ContractRecord
	for i := range row {
		res = append(res, DBRowToContractRecord(row[i]))
	}

	return res
}
