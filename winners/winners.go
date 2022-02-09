package winners

import (
	"fmt"
)

/*------------------*/

func (w *WinnersList) GetItem(index int) Winner {
	return w.list[index]
}

/*------------------*/

// This function merges two lists of winners if a winner already exist, it is not added
func (w *WinnersList) Merge(newList WinnersList) {
	for i := range newList.list {
		w.Append(newList.list[i])
	}
}

/*------------------*/

// This function merges two lists of winners and aggregates the rewards
// If a winner already exist, her/his rewards will be aggregated
func (w *WinnersList) MergeWithAggregateRewards(newList WinnersList) {
	for i := range newList.list {
		w.AppendWithAggregateRewards(newList.list[i])
	}
}

/*------------------*/

// This function cuts the end of a list
func (w WinnersList) Trim(length int) WinnersList {

	if w.Length() == length {
		return w
	}

	var trimmedList WinnersList
	for i := range w.list {
		trimmedList.Append(w.list[i])
		if i >= length-1 {
			break
		}
	}
	return trimmedList
}

/*------------------*/

func (w WinnersList) Length() int {
	return len(w.list)
}

/*------------------*/

// Return result:
// 		-1 : Not found
// 		>-1: The item index
func (w WinnersList) FindByAddress(address string) int {

	if index, ok := w.hashMap[address]; ok {
		return index
	}
	return -1
}

/*------------------*/

func (w *WinnersList) Append(item Winner) {

	if index := w.FindByAddress(item.Address); index == -1 {
		newIndex := w.Length()
		w.list = append(w.list, item)

		if w.hashMap == nil {
			w.hashMap = make(hashMapType)
		}
		w.hashMap[item.Address] = newIndex
	}
}

/*------------------*/

func (w *WinnersList) AppendWithAggregateRewards(item Winner) {

	if index := w.FindByAddress(item.Address); index != -1 {
		w.list[index].Rewards += item.Rewards
	} else {
		newIndex := w.Length()
		w.list = append(w.list, item)

		if w.hashMap == nil {
			w.hashMap = make(hashMapType)
		}
		w.hashMap[item.Address] = newIndex
	}
}

/*------------------*/

func (w WinnersList) Print() {

	for i := range w.list {
		fmt.Printf("%d \t%s\tRewards: %d\n",
			i+1,
			w.list[i].Address,
			w.list[i].Rewards,
		)
	}
}

/*------------------*/

func (w *WinnersList) GetVerifiedOnly() WinnersList {

	var output WinnersList
	for i := range w.list {
		if w.list[i].Verified {
			output.Append(w.list[i])
		}
	}
	return output
}
