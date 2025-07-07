package clipboard

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Copy(content string) {
	xdgSession := strings.ToLower(os.Getenv("XDG_SESSION_TYPE"))
	var copyCmd *exec.Cmd

	switch xdgSession {
	case "wayland":
		copyCmd = exec.Command("wl-copy")
	case "x11":
		copyCmd = exec.Command("xclip", "-selection", "clipboard")
	default:
		fmt.Println("Sesión desconocida o no soportada:", xdgSession)
		return
	}

	stdin, err := copyCmd.StdinPipe()
	if err != nil {
		fmt.Println("Error obteniendo stdin:", err)
		return
	}

	if err := copyCmd.Start(); err != nil {
		fmt.Println("Error al iniciar comando:", err)
		return
	}

	if _, err := stdin.Write([]byte(content)); err != nil {
		fmt.Println("Error al escribir al stdin:", err)
	}

	if err := stdin.Close(); err != nil {
		fmt.Println("Error al cerrar stdin:", err)
	}

	if err := copyCmd.Wait(); err != nil {
		fmt.Println("Error esperando al comando:", err)
	}
}
