package cmd

import "flag"

type CliOptions struct {
	EnvFilePath string // Ruta del archivo `.env`
	Help        bool   // Bool para mostrar la ayuda
	ActHistory  bool
}

func Launch() CliOptions {
	envFilePath := flag.String("env-file", "", ".env")
	activateHistory := flag.Bool("activate-history", false, "history .db")
	help := flag.Bool("help", false, "Ayuda")
	helpShort := flag.Bool("h", false, "Ayuda")
	flag.Parse()

	if *help || *helpShort {
		Help()
		return CliOptions{Help: true}
	}

	return CliOptions{
		EnvFilePath: *envFilePath,
		Help:        *help,
		ActHistory:  *activateHistory,
	}
}
