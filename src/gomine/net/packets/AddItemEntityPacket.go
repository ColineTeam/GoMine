package packets

import (
	"gomine/vectors"
	"gomine/entities"
	"gomine/net/info"
	"gomine/entities/math"
	"gomine/items"
)

type AddItemEntityPacket struct {
	*Packet
	UniqueId int64
	RuntimeId  uint64
	Item items.Item
	Position   vectors.TripleVector
	Motion     vectors.TripleVector
	Rotation   math.Rotation

	EntityData map[uint32][]interface{}
}

func NewAddItemEntityPacket() *AddItemEntityPacket {
	return &AddItemEntityPacket{NewPacket(info.AddItemEntityPacket), 0, 0, nil, vectors.TripleVector{}, vectors.TripleVector{}, math.Rotation{}, nil}
}

func (pk *AddItemEntityPacket) Encode() {
	pk.PutUniqueId(pk.UniqueId)
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutSlot(pk.Item)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotationObject(pk.Rotation, false)
	pk.PutEntityData(pk.EntityData)
	pk.PutUnsignedVarInt(0)
}

func (pk *AddItemEntityPacket) Decode() {
}