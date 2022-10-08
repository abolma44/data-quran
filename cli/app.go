package main

import (
	"data-quran-cli/internal/command/quranenc"
	tanzilText "data-quran-cli/internal/command/tanzil-text"
	tanzilTrans "data-quran-cli/internal/command/tanzil-trans"

	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:      "data-quran-cli",
		Usage:     "cli for downloading Quranic data",
		UsageText: "data-quran-cli [command] [flags]",
		Commands: []*cli.Command{
			quranenc.Command(),
			tanzilTrans.Command(),
			tanzilText.Command(),
		},
	}
}
