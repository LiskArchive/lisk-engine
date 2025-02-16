package sync

import (
	"bytes"
	"errors"
	"math/rand"
)

func getBestNodeInfo(nodeInfos []*NodeInfo) (*NodeInfo, error) {
	if len(nodeInfos) == 0 {
		return nil, errors.New("peer does not exist to select")
	}

	highestMaxPrevotedGroup := getLargestMaxHeightPrevotedNodeInfo(nodeInfos)
	highestHeightGroup := getLargestHeightNodeInfo(highestMaxPrevotedGroup)
	selectedInfo := getMostFrequesntBlockIDNodeInfo(highestHeightGroup)
	index := rand.Intn(len(selectedInfo))
	return selectedInfo[index], nil
}

func getLargestMaxHeightPrevotedNodeInfo(info []*NodeInfo) []*NodeInfo {
	if len(info) == 0 {
		return []*NodeInfo{}
	}
	result := []*NodeInfo{}
	maxValue := info[0].maxHeightPrevoted
	for _, val := range info {
		if val.maxHeightPrevoted > maxValue {
			maxValue = val.maxHeightPrevoted
			result = []*NodeInfo{}
			result = append(result, val)
			continue
		}
		if val.maxHeightPrevoted == maxValue {
			result = append(result, val)
		}
	}
	return result
}

func getLargestHeightNodeInfo(info []*NodeInfo) []*NodeInfo {
	if len(info) == 0 {
		return []*NodeInfo{}
	}
	result := []*NodeInfo{}
	maxValue := info[0].height
	for _, val := range info {
		if val.height > maxValue {
			maxValue = val.height
			result = []*NodeInfo{}
			result = append(result, val)
			continue
		}
		if val.height == maxValue {
			result = append(result, val)
		}
	}
	return result
}

func getMostFrequesntBlockIDNodeInfo(info []*NodeInfo) []*NodeInfo {
	frequency := map[string]int{}
	for _, val := range info {
		count, exist := frequency[string(val.lastBlockID)]
		if exist {
			frequency[string(val.lastBlockID)] = count + 1
		} else {
			frequency[string(val.lastBlockID)] = 1
		}
	}
	max := 0
	var blockID []byte
	for idStr, count := range frequency {
		if count > max {
			blockID = []byte(idStr)
		}
	}
	if blockID == nil {
		return []*NodeInfo{}
	}
	result := []*NodeInfo{}
	for _, val := range info {
		if bytes.Equal(val.lastBlockID, blockID) {
			result = append(result, val)
		}
	}
	return result
}
