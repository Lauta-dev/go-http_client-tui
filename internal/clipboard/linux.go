package clipboard

import (
	"fmt"
	"os"
	"os/exec"
)

func Copy(content string) {
	//xclip (X11), wl-clipboard (Wayland)

	copy := ""
	xdgSession := os.Getenv("XDG_SESSION_TYPE")
	switch xdgSession {
	case "wayland":
		copy = "wl-copy"

	case "x11":
		copy = "xclip"
	}

	cmd := exec.Command(copy)

	if copy == "xclip" {
		cmd = exec.Command(copy, "-selection", "clipboard")
	}

	in, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if _, err := in.Write([]byte(content)); err != nil {
		fmt.Println(err)
	}

	if err := in.Close(); err != nil {
		fmt.Println(err)
	}

	cmd.Wait()
}
