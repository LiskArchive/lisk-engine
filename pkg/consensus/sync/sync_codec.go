// Code generated by github.com/LiskHQ/lisk-engine/pkg/codec/gen; DO NOT EDIT.

package sync

import (
	"github.com/LiskHQ/lisk-engine/pkg/blockchain"
	"github.com/LiskHQ/lisk-engine/pkg/codec"
)

func (e *GetHighestCommonBlockRequest) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytesArray(1, e.IDs)
	return writer.Result()
}

func (e *GetHighestCommonBlockRequest) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GetHighestCommonBlockRequest) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GetHighestCommonBlockRequest) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GetHighestCommonBlockRequest) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytesArray(1)
		if err != nil {
			return err
		}
		e.IDs = val
	}
	return nil
}

func (e *GetHighestCommonBlockRequest) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytesArray(1)
		if err != nil {
			return err
		}
		e.IDs = val
	}
	return nil
}

func (e *GetHighestCommonBlockResponse) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.ID)
	return writer.Result()
}

func (e *GetHighestCommonBlockResponse) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GetHighestCommonBlockResponse) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GetHighestCommonBlockResponse) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GetHighestCommonBlockResponse) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.ID = val
	}
	return nil
}

func (e *GetHighestCommonBlockResponse) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.ID = val
	}
	return nil
}

func (e *GetBlocksFromIDRequest) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.ID)
	return writer.Result()
}

func (e *GetBlocksFromIDRequest) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GetBlocksFromIDRequest) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GetBlocksFromIDRequest) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GetBlocksFromIDRequest) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.ID = val
	}
	return nil
}

func (e *GetBlocksFromIDRequest) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.ID = val
	}
	return nil
}

func (e *GetBlocksFromIDResponse) Encode() []byte {
	writer := codec.NewWriter()
	{
		for _, val := range e.Blocks {
			if val != nil {
				writer.WriteEncodable(1, val)
			}
		}
	}
	return writer.Result()
}

func (e *GetBlocksFromIDResponse) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GetBlocksFromIDResponse) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GetBlocksFromIDResponse) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GetBlocksFromIDResponse) DecodeFromReader(reader *codec.Reader) error {
	{
		vals, err := reader.ReadDecodables(1, func() codec.DecodableReader { return new(blockchain.Block) })
		if err != nil {
			return err
		}
		r := make([]*blockchain.Block, len(vals))
		for i, v := range vals {
			r[i] = v.(*blockchain.Block)
		}
		e.Blocks = r
	}
	return nil
}

func (e *GetBlocksFromIDResponse) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		vals, err := reader.ReadDecodables(1, func() codec.DecodableReader { return new(blockchain.Block) })
		if err != nil {
			return err
		}
		r := make([]*blockchain.Block, len(vals))
		for i, v := range vals {
			r[i] = v.(*blockchain.Block)
		}
		e.Blocks = r
	}
	return nil
}
