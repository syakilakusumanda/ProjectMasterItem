// main.go
package main

import (
	"ProjectMasterItem/Controller"
	"ProjectMasterItem/Node"
	"fmt"
)

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Item")
		fmt.Println("2. Tampilkan Semua Item")
		fmt.Println("3. Cari Item")
		fmt.Println("4. Edit Item")
		fmt.Println("5. Hapus Item")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi: ")

		var opsi string
		fmt.Scanln(&opsi)

		switch opsi {
		case "1":
			tambahItem()
		case "2":
			Controller.TampilkanSemuaItem()
		case "3":
			cariItem()
		case "4":
			editItem()
		case "5":
			hapusItem()
		case "6":
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func tambahItem() {
	var id, stok, harga int
	var nama string

	fmt.Print("Masukkan ID: ")
	fmt.Scanln(&id)

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Stok: ")
	fmt.Scanln(&stok)

	fmt.Print("Masukkan Harga: ")
	fmt.Scanln(&harga)

	Controller.TambahItem(id, nama, stok, harga)
	fmt.Println("Item berhasil ditambahkan.")
}

func cariItem() {
	var id int
	fmt.Print("Masukkan ID item yang dicari: ")
	fmt.Scanln(&id)

	Controller.CariItem(id)
}

func editItem() {
	var id, stok, harga int
	var nama string

	fmt.Print("Masukkan ID item yang ingin diedit: ")
	fmt.Scanln(&id)

	fmt.Print("Masukkan Nama Baru: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Stok Baru: ")
	fmt.Scanln(&stok)

	fmt.Print("Masukkan Harga Baru: ")
	fmt.Scanln(&harga)

	itemBaru := Node.Item{Id: id, Nama: nama, Stok: stok, Harga: harga}
	if Controller.EditItem(itemBaru) {
		fmt.Println("Item berhasil diperbarui.")
	} else {
		fmt.Println("Item tidak ditemukan.")
	}
}

func hapusItem() {
	var id int
	fmt.Print("Masukkan ID item yang ingin dihapus: ")
	fmt.Scanln(&id)

	if Controller.HapusItem(id) {
		fmt.Println("Item berhasil dihapus.")
	} else {
		fmt.Println("Item tidak ditemukan.")
	}
}
