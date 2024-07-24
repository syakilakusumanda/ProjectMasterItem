// View/view.go
package View

import (
	"ProjectMasterItem/Node"
	"fmt"
)

func TampilkanItem(item *Node.TabelItem) {
	if item == nil {
		fmt.Println("Tidak ada item yang ditemukan.")
		return
	}
	for item != nil {
		fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %d\n", item.Data.Id, item.Data.Nama, item.Data.Stok, item.Data.Harga)
		item = item.Next
	}
}

func TampilkanItemTunggal(item *Node.TabelItem) {
	if item == nil {
		fmt.Println("Tidak ada item yang ditemukan.")
	} else {
		fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %d\n", item.Data.Id, item.Data.Nama, item.Data.Stok, item.Data.Harga)
	}
}
