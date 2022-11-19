package utils

import (
	"io"
	"os"
)

func CopyIconInStartup() {
	dirname, err := os.UserHomeDir()
	CheckErr(err)
	in, err := os.Open(dirname + `\Desktop\kodee.lnk`)
	CheckErr(err)
	defer in.Close()

	out, err := os.Create(dirname + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\kodee.lnk`)
	CheckErr(err)
	defer out.Close()

	_, err = io.Copy(out, in)
	CheckErr(err)
}