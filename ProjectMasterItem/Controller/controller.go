package Controller

import (
	"ProjectMasterItem/Model"
	"ProjectMasterItem/Node"
	"ProjectMasterItem/View"
)

func TambahItem(id int, nama string, stok int, harga int) {
	Model.InsertItem(id, nama, stok, harga)
}

func TampilkanSemuaItem() {
	items := Model.ReadAllItem()
	View.TampilkanItem(items)
}

func HapusItem(id int) bool {
	return Model.DeleteItem(id)
}

func CariItem(id int) {
	item := Model.SearchItem(id)
	View.TampilkanItemTunggal(item)
}

func EditItem(item Node.Item) bool {
	return Model.UpdateItem(item)
}
