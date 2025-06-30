package main

import (
	"os/exec"
)

func Clip() error {
	//xclip (X11), wl-clipboard (Wayland)

	copy := "wl-copy"
	text := "Esto fue copiado desde GOlang"

	cmd := exec.Command(copy)
	in, err := cmd.StdinPipe()

	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
