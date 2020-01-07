package main

import (
	"kiapanahi/mango/package1"
	"kiapanahi/mango/package2"
)

func main() {
	m := package2.NewModel1("Kia", "Panahi Rad", 31)

	package1.Log(m)
}
