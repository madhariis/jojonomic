package helper

import (
	"log"
	"os/exec"
)

func GenerateUUID() string {
	//required executable apps "uuidgen.exe" in same path
	out, err := exec.Command("uuidgen").Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
