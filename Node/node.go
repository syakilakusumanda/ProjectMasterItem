// Node/node.go
package Node

type Item struct {
	Id    int
	Nama  string
	Stok  int
	Harga int
}

type TabelItem struct {
	Data Item
	Next *TabelItem
}
