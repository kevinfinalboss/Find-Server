package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação em Linha de Comando"
	app.Usage = "Busca de IP e nome de servidores"

	return app
}
