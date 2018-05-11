package main

import (
	"fmt"

	"github.com/user/sql"
	"github.com/user/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("Connected to:", sql.Connect())
}
