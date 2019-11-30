package main

import (
	"os/exec"
)

func main() {
	err := exec.Command("open", "https://www.yahoo.co.jp/").Start()

	if err != nil {
		return err
	}
}
