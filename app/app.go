package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação em Linha de Comando"
	app.Usage = "Busca de IP e nome de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "kevindev.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IP de servidores",
			Flags:  flags,
			Action: findIp,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIp(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
