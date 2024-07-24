// Model/model.go
package Model

import (
	"ProjectMasterItem/Database"
	"ProjectMasterItem/Node"
)

func InsertItem(id int, nama string, stok int, harga int) {
	itemBaru := &Node.TabelItem{
		Data: Node.Item{Id: id, Nama: nama, Stok: stok, Harga: harga},
		Next: nil,
	}

	if Database.DBitem.Next == nil {
		Database.DBitem.Next = itemBaru
	} else {
		tempLL := Database.DBitem.Next
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = itemBaru
	}
}

func ReadAllItem() *Node.TabelItem {
	if Database.DBitem.Next == nil {
		return nil
	} else {
		return Database.DBitem.Next
	}
}

func DeleteItem(id int) bool {
	if Database.DBitem.Next == nil {
		return false
	}

	tempLL := &Database.DBitem
	for tempLL.Next != nil {
		if tempLL.Next.Data.Id == id {
			tempLL.Next = tempLL.Next.Next
			return true
		}
		tempLL = tempLL.Next
	}
	return false
}

func SearchItem(id int) *Node.TabelItem {
	tempLL := Database.DBitem.Next
	for tempLL != nil {
		if tempLL.Data.Id == id {
			return tempLL
		}
		tempLL = tempLL.Next
	}
	return nil
}

func UpdateItem(item Node.Item) bool {
	addr := SearchItem(item.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Nama = item.Nama
		addr.Data.Stok = item.Stok
		addr.Data.Harga = item.Harga
		return true
	}
}
