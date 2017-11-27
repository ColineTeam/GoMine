package handlers

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
	"goraklib/server"
)

type LoginHandler struct {
	*players.PacketHandler
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{players.NewPacketHandler(info.LoginPacket)}
}

func (handler LoginHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if loginPacket, ok := packet.(packets.LoginPacket); ok {
		var player = players.NewPlayer(server, loginPacket.Username, loginPacket.ClientUUID, loginPacket.ClientXUID, loginPacket.ClientId)
		player.SetLanguage(loginPacket.Language)

		pk := packets.NewPlayStatusPacket()
		pk.Status = 0
		server.GetRakLibAdapter().SendPacket(pk, session)

		pk3 := packets.NewResourcePackInfoPacket()
		server.GetRakLibAdapter().SendPacket(pk3, session)

		server.GetPlayerFactory().AddPlayer(player, session)
	}

	return true
}