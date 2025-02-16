// Code generated by github.com/LiskHQ/lisk-engine/pkg/codec/gen; DO NOT EDIT.

package generator

import (
	"github.com/LiskHQ/lisk-engine/pkg/codec"
)

func (e *Keypair) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.address)
	writer.WriteBytes(2, e.generationPublicKey)
	writer.WriteBytes(3, e.generationPrivateKey)
	writer.WriteBytes(4, e.blsPrivateKey)
	return writer.Result()
}

func (e *Keypair) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *Keypair) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *Keypair) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *Keypair) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.address = val
	}
	{
		val, err := reader.ReadBytes(2, false)
		if err != nil {
			return err
		}
		e.generationPublicKey = val
	}
	{
		val, err := reader.ReadBytes(3, false)
		if err != nil {
			return err
		}
		e.generationPrivateKey = val
	}
	{
		val, err := reader.ReadBytes(4, false)
		if err != nil {
			return err
		}
		e.blsPrivateKey = val
	}
	return nil
}

func (e *Keypair) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.address = val
	}
	{
		val, err := reader.ReadBytes(2, true)
		if err != nil {
			return err
		}
		e.generationPublicKey = val
	}
	{
		val, err := reader.ReadBytes(3, true)
		if err != nil {
			return err
		}
		e.generationPrivateKey = val
	}
	{
		val, err := reader.ReadBytes(4, true)
		if err != nil {
			return err
		}
		e.blsPrivateKey = val
	}
	return nil
}

func (e *GeneratorInfo) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteUInt32(1, e.Height)
	writer.WriteUInt32(2, e.MaxHeightPrevoted)
	writer.WriteUInt32(3, e.MaxHeightGenerated)
	return writer.Result()
}

func (e *GeneratorInfo) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *GeneratorInfo) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *GeneratorInfo) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *GeneratorInfo) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadUInt32(1, false)
		if err != nil {
			return err
		}
		e.Height = val
	}
	{
		val, err := reader.ReadUInt32(2, false)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	{
		val, err := reader.ReadUInt32(3, false)
		if err != nil {
			return err
		}
		e.MaxHeightGenerated = val
	}
	return nil
}

func (e *GeneratorInfo) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadUInt32(1, true)
		if err != nil {
			return err
		}
		e.Height = val
	}
	{
		val, err := reader.ReadUInt32(2, true)
		if err != nil {
			return err
		}
		e.MaxHeightPrevoted = val
	}
	{
		val, err := reader.ReadUInt32(3, true)
		if err != nil {
			return err
		}
		e.MaxHeightGenerated = val
	}
	return nil
}

func (e *Keys) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.Address)
	writer.WriteString(2, e.Type)
	writer.WriteBytes(3, e.Data)
	return writer.Result()
}

func (e *Keys) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *Keys) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *Keys) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *Keys) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.Address = val
	}
	{
		val, err := reader.ReadString(2, false)
		if err != nil {
			return err
		}
		e.Type = val
	}
	{
		val, err := reader.ReadBytes(3, false)
		if err != nil {
			return err
		}
		e.Data = val
	}
	return nil
}

func (e *Keys) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.Address = val
	}
	{
		val, err := reader.ReadString(2, true)
		if err != nil {
			return err
		}
		e.Type = val
	}
	{
		val, err := reader.ReadBytes(3, true)
		if err != nil {
			return err
		}
		e.Data = val
	}
	return nil
}

func (e *PlainKeys) Encode() []byte {
	writer := codec.NewWriter()
	writer.WriteBytes(1, e.GeneratorKey)
	writer.WriteBytes(2, e.GeneratorPrivateKey)
	writer.WriteBytes(3, e.BLSKey)
	writer.WriteBytes(4, e.BLSPrivateKey)
	return writer.Result()
}

func (e *PlainKeys) Decode(data []byte) error {
	reader := codec.NewReader(data)
	return e.DecodeFromReader(reader)
}

func (e *PlainKeys) MustDecode(data []byte) {
	if err := e.Decode(data); err != nil {
		panic(err)
	}
}

func (e *PlainKeys) DecodeStrict(data []byte) error {
	reader := codec.NewReader(data)
	if err := e.DecodeStrictFromReader(reader); err != nil {
		return err
	}
	if reader.HasUnreadBytes() {
		return codec.ErrUnreadBytes
	}
	return nil
}

func (e *PlainKeys) DecodeFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, false)
		if err != nil {
			return err
		}
		e.GeneratorKey = val
	}
	{
		val, err := reader.ReadBytes(2, false)
		if err != nil {
			return err
		}
		e.GeneratorPrivateKey = val
	}
	{
		val, err := reader.ReadBytes(3, false)
		if err != nil {
			return err
		}
		e.BLSKey = val
	}
	{
		val, err := reader.ReadBytes(4, false)
		if err != nil {
			return err
		}
		e.BLSPrivateKey = val
	}
	return nil
}

func (e *PlainKeys) DecodeStrictFromReader(reader *codec.Reader) error {
	{
		val, err := reader.ReadBytes(1, true)
		if err != nil {
			return err
		}
		e.GeneratorKey = val
	}
	{
		val, err := reader.ReadBytes(2, true)
		if err != nil {
			return err
		}
		e.GeneratorPrivateKey = val
	}
	{
		val, err := reader.ReadBytes(3, true)
		if err != nil {
			return err
		}
		e.BLSKey = val
	}
	{
		val, err := reader.ReadBytes(4, true)
		if err != nil {
			return err
		}
		e.BLSPrivateKey = val
	}
	return nil
}
