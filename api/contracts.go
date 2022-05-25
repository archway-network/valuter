package api

import (
	"log"
	"net/http"

	"github.com/archway-network/cosmologger/block"
	"github.com/archway-network/valuter/configs"
	"github.com/archway-network/valuter/contracts"
	"github.com/archway-network/valuter/participants"
	"github.com/archway-network/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /challenges/contracts/max-net-rewards
 */
func GetMaxNetworkRewardsWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	contractsList, err := contracts.GetMaxNetworkRewardsTopContracts(
		uint64(configs.Configs.Tasks.ContractsMaxRewards.MaxWinners),
		configs.Configs.Tasks.ContractsMaxRewards.Condition.StartHight,
		configs.Configs.Tasks.ContractsMaxRewards.Condition.EndHight,
	)

	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	type Output struct {
		block.ContractRecord
		participants.ParticipantRecord
	}

	var output []Output
	for i := range contractsList {
		pRec, _ := participants.GetParticipantByAddress(contractsList[i].DeveloperAddress)

		// Some participant might use their main account address as the reward address instead of developer address
		if pRec.AccountAddress == "" && contractsList[i].DeveloperAddress != contractsList[i].RewardAddress {
			pRec, _ = participants.GetParticipantByAddress(contractsList[i].RewardAddress)
		}
		contractsList[i].MetadataJson = ""
		output = append(output,
			Output{
				ContractRecord:    contractsList[i],
				ParticipantRecord: pRec,
			})
	}

	tools.SendJSON(resp, output)

}

/*-------------*/

/*
* This function implements GET /challenges/contracts/max-net-rewards
 */
func GetContractsSubsidizeUsersFeesWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	contractsList, err := contracts.GetSubsidizeUsersFeesTopContracts(
		uint64(configs.Configs.Tasks.ContractsMaxRewards.MaxWinners),
		configs.Configs.Tasks.ContractsMaxRewards.Condition.StartHight,
		configs.Configs.Tasks.ContractsMaxRewards.Condition.EndHight,
	)

	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	type Output struct {
		block.ContractRecord
		participants.ParticipantRecord
	}

	var output []Output
	for i := range contractsList {
		pRec, _ := participants.GetParticipantByAddress(contractsList[i].DeveloperAddress)

		// Some participant might use their main account address as the reward address instead of developer address
		if pRec.AccountAddress == "" && contractsList[i].DeveloperAddress != contractsList[i].RewardAddress {
			pRec, _ = participants.GetParticipantByAddress(contractsList[i].RewardAddress)
		}
		contractsList[i].MetadataJson = ""
		output = append(output,
			Output{
				ContractRecord:    contractsList[i],
				ParticipantRecord: pRec,
			})
	}

	tools.SendJSON(resp, output)

}

/*-------------*/
