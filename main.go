package main

import (
	"app-mahasiswa-api2/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Exec()
}
