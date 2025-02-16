package liskbft

import (
	"errors"

	"github.com/LiskHQ/lisk-engine/pkg/collection/bytes"
	"github.com/LiskHQ/lisk-engine/pkg/statemachine"
)

var (
	ErrBFTParamsNotFound     = errors.New("bFT parameters does not exist")
	ErrGeneratorKeysNotFound = errors.New("generator keys does not exist")
)

func getBFTParams(paramsStore statemachine.ImmutableStore, height uint32) (*BFTParams, error) {
	start := bytes.FromUint32(0)
	end := bytes.FromUint32(height)
	kv := paramsStore.Range(start, end, 1, true)

	if len(kv) != 1 {
		return nil, ErrBFTParamsNotFound
	}
	bftParams := &BFTParams{}
	if err := bftParams.Decode(kv[0].Value()); err != nil {
		return nil, err
	}
	return bftParams, nil
}

func getGeneratorKeys(keysStore statemachine.ImmutableStore, height uint32) (*GeneratorKeys, error) {
	start := bytes.FromUint32(0)
	end := bytes.FromUint32(height)
	kv := keysStore.Range(start, end, 1, true)

	if len(kv) != 1 {
		return nil, ErrGeneratorKeysNotFound
	}
	keys := &GeneratorKeys{}
	if err := keys.Decode(kv[0].Value()); err != nil {
		return nil, err
	}
	return keys, nil
}

func deleteBFTParams(paramsStore statemachine.Store, height uint32) error {
	start := bytes.FromUint32(0)
	end := bytes.FromUint32(height)
	kv := paramsStore.Range(start, end, -1, false)

	if len(kv) <= 1 {
		return nil
	}
	for i := 0; i < len(kv)-1; i++ {
		paramsStore.Del(kv[i].Key())
	}
	return nil
}

func deleteGeneratorKeys(keysStore statemachine.Store, height uint32) error {
	start := bytes.FromUint32(0)
	end := bytes.FromUint32(height)
	kv := keysStore.Range(start, end, -1, false)

	if len(kv) <= 1 {
		return nil
	}
	for i := 0; i < len(kv)-1; i++ {
		keysStore.Del(kv[i].Key())
	}
	return nil
}
