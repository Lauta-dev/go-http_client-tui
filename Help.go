package main

import "fmt"

func HelpTrigger() {

	dialog := `

	Programa desarrollado por: https://github.com/lauta-dev

--help | -help: Mostrar esta ayuda
--env-file | -env-file: Cargar archivo .env
--active-history | -active-history: Activar historial de request
	`

	fmt.Print(dialog)

}
