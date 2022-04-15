package main

import (
	"log"
	"os"
	"path"
	"time"

	"pocketsubs/cli/helpers"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Pocket Fansubs CLI Helper"
	app.Version = "v0.1.0 (alpha)"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "Douglas Martins",
			Email: "douglas.martins.m099@gmail.com",
		},
	}

	app.Usage = "Ferramenta para auxiliar a tradução de fansubs"
	app.HideHelp = true
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "backup",
			Value:   false,
			Usage:   "fazer backup do arquivo original caso novo arquivo não seja especificado",
			Aliases: []string{"b"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "dialogs",
			Aliases: []string{"d"},
			Usage:   "gerencia a lista de falas do arquivo de legendas",
			Subcommands: []*cli.Command{
				{
					Name:      "extract",
					Aliases:   []string{"e"},
					Usage:     "extrair lista de falas do arquivo `ORIGINAL` para o arquivo `NEW`",
					ArgsUsage: "[ORIGINAL] [NEW]",
					Action: func(c *cli.Context) error {
						return helpers.Extract(c.Args().Get(0), c.Args().Get(1), c.String("ext"), c.Bool("backup"))
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "ext",
							Value:   ".txt",
							Usage:   "extensão/formato do arquivo de tradução",
							Aliases: []string{"e"},
						},
					},
				},
				{
					Name:      "restore",
					Aliases:   []string{"r"},
					Usage:     "substitui as falas do arquivo de legendas `ORIGINAL` com as da lista de falas no arquivo `FILE`",
					ArgsUsage: "[ORIGINAL] [FILE]",
					Action: func(c *cli.Context) error {
						return helpers.Restore(c.Args().Get(0), c.Args().Get(1), c.Bool("backup"))
					},
				},
			},
		},
		{
			Name:    "symbol",
			Aliases: []string{"s"},
			Usage:   "troca nomes próprios e outras expressões não traduzíveis para facilitar a tradução automática",
			Subcommands: []*cli.Command{
				{
					Name:      "replace",
					Aliases:   []string{"s"},
					ArgsUsage: "[ORIGINAL] [NEW]",
					Action: func(c *cli.Context) error {
						return helpers.Tokenize(c.Args().Get(0), c.Args().Get(1), c.String("glossary"), c.Bool("backup"))
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "glossary",
							Value:   "dictionaries",
							Usage:   "pasta onde se encontram os dicionários",
							Aliases: []string{"g"},
						},
					},
				},
				{
					Name:      "restore",
					Aliases:   []string{"r"},
					Usage:     "substitui as falas do arquivo de legendas `ORIGINAL` com as da lista de falas no arquivo `NEW`",
					ArgsUsage: "[ORIGINAL] [NEW]",
					Action: func(c *cli.Context) error {
						return helpers.Untokenize(c.Args().Get(0), c.Args().Get(1), c.String("glossary"), c.Bool("backup"))
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "glossary",
							Value:   "dictionaries",
							Usage:   "pasta onde se encontram os dicionários",
							Aliases: []string{"g"},
						},
					},
				},
			},
		},
		{
			Name:      "translate",
			Aliases:   []string{"t"},
			Usage:     "traduz o arquivo de falas e salva o resultado em um novo arquivo",
			ArgsUsage: "[DIALOGS] [NEW]",
			Action: func(c *cli.Context) error {
				return helpers.Translate(c.Args().Get(0), c.Args().Get(1), c.String("lang"), c.Int("chunk"), c.String("ext"), c.Bool("backup"))
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "lang",
					Value:   "pt",
					Usage:   "idioma de destino da tradução",
					Aliases: []string{"l"},
				},
				&cli.IntFlag{
					Name:    "chunk",
					Value:   30,
					Usage:   "nº de linhas a traduzir por vez",
					Aliases: []string{"c"},
				},
				&cli.StringFlag{
					Name:    "ext",
					Value:   ".tls",
					Usage:   "extensão/formato do arquivo de tradução",
					Aliases: []string{"e"},
				},
			},
		},
		{
			Name:      "stylish",
			Usage:     "adiciona estilos ao arquivo de legendas",
			ArgsUsage: "[ORIGINAL] [NEW]",
			Action: func(c *cli.Context) error {
				return helpers.Stylish(c.Args().Get(0), c.Args().Get(1), c.String("stylesheet"), c.String("delimitter"), c.Bool("backup"))
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "stylesheet",
					Value:   "stylesheet.txt",
					Usage:   "arquivo com os estilos a aplicar",
					Aliases: []string{"s"},
				},
				&cli.StringFlag{
					Name:    "delimitter",
					Value:   "[V4+ Styles]",
					Usage:   "identificador de estilos no arquivo de legenda",
					Aliases: []string{"d"},
					Hidden:  true,
				},
			},
		},
		{
			Name:      "format",
			Usage:     "adiciona estilos de tipos ao arquivo de legendas/falas",
			ArgsUsage: "[ORIGINAL] [NEW]",
			Action: func(c *cli.Context) error {
				return helpers.Format(c.Args().Get(0), c.Args().Get(1), c.String("moves"), c.String("styles"), c.Bool("backup"))
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "moves",
					Usage:   "localização da lista de golpes/ataques e seu tipo",
					Value:   path.Join("dictionaries", "moves.json"),
					Aliases: []string{"m"},
				},
				&cli.StringFlag{
					Name:    "styles",
					Usage:   "formatação para cada tipo",
					Value:   path.Join("dictionaries", "types.json"),
					Aliases: []string{"s"},
				},
			},
			Subcommands: []*cli.Command{
				{
					Name:                   "clean",
					Aliases:                []string{"c"},
					ArgsUsage:              "[ORIGINAL] [NEW]",
					Usage:                  "remove todas as tags e quebras de linha",
					UseShortOptionHandling: true,
					Action: func(c *cli.Context) error {
						return helpers.Clean(c.Args().Get(0), c.Args().Get(1), c.Bool("tags"), c.Bool("lines"), c.Bool("backup"))
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "tags",
							Usage:   "remover tags ASS do arquivo",
							Value:   false,
							Aliases: []string{"t"},
						},
						&cli.BoolFlag{
							Name:    "lines",
							Usage:   "Remover quebra de linhas (\\N) do arquivo",
							Value:   false,
							Aliases: []string{"l"},
						},
					},
				},
			},
		},
		{
			Name:      "delay",
			Usage:     "adia os intervalos de tempo do arquivo de legendas",
			ArgsUsage: "[ORIGINAL] [NEW]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "time",
					Usage:    "intervalo de tempo a adiar [h:mm:ss:ms]",
					Aliases:  []string{"t"},
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				return helpers.Delay(c.Args().Get(0), c.Args().Get(0), c.String("time"), c.Bool("backup"))
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
