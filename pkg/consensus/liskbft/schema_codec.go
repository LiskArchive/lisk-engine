// Code generated by github.com/LiskHQ/lisk-engine/pkg/codec/gen; DO NOT EDIT.

package liskbft

import (
	"github.com/LiskHQ/lisk-engine/pkg/codec"
)

func (e *GetValidatorInfoResponse) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteUInt32(1, e.MaxHeightPrevoted)
	return writer.Result()
}

func (e *GetValidatorInfoResponse) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GetValidatorInfoResponse) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GetValidatorInfoResponse) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GetValidatorInfoResponse) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadUInt32(1, false)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	return nil
}

func (e *GetValidatorInfoResponse) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadUInt32(1, true)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	return nil
}

func (e *IsBFTComplientRequest) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.GeneratorPublicKey)
	writer.WriteUInt32(2, e.Height)
	writer.WriteUInt32(3, e.MaxHeightGenerated)
	writer.WriteUInt32(4, e.MaxHeightPrevoted)
	return writer.Result()
}

func (e *IsBFTComplientRequest) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *IsBFTComplientRequest) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *IsBFTComplientRequest) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *IsBFTComplientRequest) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.GeneratorPublicKey = val
	}
	{
		val, err := reader.ReadUInt32(2, false)
		if err != nil {
			return err
		}
		e.Height = val
	}
	{
		val, err := reader.ReadUInt32(3, false)
		if err != nil {
			return err
		}
		e.MaxHeightGenerated = val
	}
	{
		val, err := reader.ReadUInt32(4, false)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	return nil
}

func (e *IsBFTComplientRequest) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.GeneratorPublicKey = val
	}
	{
		val, err := reader.ReadUInt32(2, true)
		if err != nil {
			return err
		}
		e.Height = val
	}
	{
		val, err := reader.ReadUInt32(3, true)
		if err != nil {
			return err
		}
		e.MaxHeightGenerated = val
	}
	{
		val, err := reader.ReadUInt32(4, true)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	return nil
}

func (e *IsBFTComplientResponse) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBool(1, e.Valid)
	return writer.Result()
}

func (e *IsBFTComplientResponse) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *IsBFTComplientResponse) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *IsBFTComplientResponse) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *IsBFTComplientResponse) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBool(1, false)
		if err != nil {
			return err
		}
		e.Valid = val
	}
	return nil
}

func (e *IsBFTComplientResponse) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBool(1, true)
		if err != nil {
			return err
		}
		e.Valid = val
	}
	return nil
}
