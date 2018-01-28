package packets

import (
	"gomine/net/info"
)

type ModalFormRequestPacket struct {
	*Packet
	FormId int32
	FormData string
}

func NewModalFormRequestPacket() * ModalFormRequestPacket{
	return &ModalFormRequestPacket{NewPacket(info.ModalFormRequestPacket), nil, ""}
}

func (pk *ModalFormRequestPacket) Encode() {
	pk.PutVarInt(pk.FormId)
	pk.PutString(pk.FormData)
}

func (pk *ModalFormRequestPacket) Decode() {

}
