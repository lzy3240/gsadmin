package main

import (
	"gsadmin/boot"
)

////go:embed template
//var templateFs embed.FS
//
////go:embed static
//var staticFs embed.FS

func main() {
	boot.InitServer() //staticFs, templateFs
}
