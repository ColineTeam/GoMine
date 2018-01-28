package interfaces

import(
	"goraklib/server"
	"gomine/items"
)

type IItemFactory interface{
	GetItemById(id int) items.Item
}

