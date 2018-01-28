package items

import (
	"gomine/interfaces"
)


type ItemFactory struct{
	mainServer interfaces.IServer
	creativeItems map[int]Item
}

func NewItemFactory(mainserv interfaces.IServer) *ItemFactory{
	return &ItemFactory{mainserv, make(map[int]Item)}
}

func (ifactory *ItemFactory) Initialize(realItem *Item) {
	ifactory.creativeItems[*realItem.itemId] = *realItem
}

func (ifactory *ItemFactory) GetCreativeItemByItemIt(id int) Item {
	return ifactory.creativeItems[id]
}




