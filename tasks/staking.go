package tasks

import (
	"fmt"

	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/database"
	"github.com/archway-network/valuter/tx"
	"github.com/archway-network/valuter/winners"
)

func GetStakingWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	// If someone has done the task more than once, there will be more than a record here,
	// But that's not a problem, as winners list is distinct
	SQL := fmt.Sprintf(`
		SELECT "%s", "%s"
		FROM "%s" 
		WHERE "%s" IN(
				SELECT "%s" FROM "%s" 
				WHERE "%s" IN (
							SELECT "%s" FROM "%s" 
							WHERE "%s" = $1
							GROUP BY "%s"
							HAVING COUNT( "%s") >= 2
							) 
							AND
					"%s" IN ($2, $3)
				GROUP BY "%s"
				HAVING COUNT( "%s") >= 2
			) 
			AND "%s" = $4
		ORDER BY 
			"%s" ASC`, // >= 2 is because we want the participants to delegate at least to two validators

		database.FIELD_TX_EVENTS_SENDER,
		database.FIELD_TX_EVENTS_HEIGHT,

		database.TABLE_TX_EVENTS,

		database.FIELD_TX_EVENTS_SENDER,

		database.FIELD_TX_EVENTS_SENDER,
		database.TABLE_TX_EVENTS,
		database.FIELD_TX_EVENTS_SENDER,

		database.FIELD_TX_EVENTS_SENDER,
		database.TABLE_TX_EVENTS,
		database.FIELD_TX_EVENTS_ACTION,
		database.FIELD_TX_EVENTS_SENDER,
		database.FIELD_TX_EVENTS_VALIDATOR,

		database.FIELD_TX_EVENTS_ACTION,
		database.FIELD_TX_EVENTS_SENDER,
		database.FIELD_TX_EVENTS_VALIDATOR,

		database.FIELD_TX_EVENTS_ACTION,
		database.FIELD_TX_EVENTS_HEIGHT,
	)

	rows, err := database.DB.Query(SQL,
		database.QueryParams{
			tx.ACTION_DELEGATE,
			tx.ACTION_BEGIN_REDELEGATE,
			tx.ACTION_BEGIN_UNBONDING,
			tx.ACTION_WITHDRAW_DELEGATOR_REWARD,
		})
	if err != nil {
		return winnersList, err
	}

	for i := range rows {
		newWinner := winners.Winner{
			Address: rows[i][database.FIELD_TX_EVENTS_SENDER].(string),
			Rewards: configs.Configs.Tasks.Staking.Reward,
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

		if winnersList.Length() >= configs.Configs.Tasks.Staking.MaxWinners {
			break // Max winners reached
		}

	}

	return winnersList, nil
}