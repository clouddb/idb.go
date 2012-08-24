package main

import(
	"idb"
	"fmt"
)

func main() {
	test_path := "./test_db"

	db := idb.Idb{ test_path, idb.FILE }
	x, y := db.Setup(test_path)

	fmt.Printf("%s\n", db.Path)

	fmt.Printf("%d, %d\n", x, y )
}