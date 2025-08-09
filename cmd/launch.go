package cmd

import "flag"

type CliOptions struct {
	EnvFilePath string // Ruta del archivo `.env`
	Help        bool   // Bool para mostrar la ayuda
}

func Launch() CliOptions {
	envFilePath := flag.String("env-file", "", ".env")
	envFilePathShort := flag.String("ef", "", ".env")
	help := flag.Bool("help", false, "Ayuda")
	helpShort := flag.Bool("h", false, "Ayuda")

	flag.Parse()

	if *help || *helpShort {
		Help()
		return CliOptions{Help: true}
	}

	if *envFilePath != "" {
		return CliOptions{EnvFilePath: *envFilePath}
	}

	if *envFilePathShort != "" {
		return CliOptions{EnvFilePath: *envFilePathShort}
	}

	return CliOptions{}
}
