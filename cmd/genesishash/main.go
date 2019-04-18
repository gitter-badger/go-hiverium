package main

import (
	"encoding/json"
	"github.com/unification-com/mainchain/cmd/utils"
	"github.com/unification-com/mainchain/core"
	"gopkg.in/urfave/cli.v1"
	"os"
)

// main is just a boring entry point to set up the CLI app.
func main() {
	app := cli.NewApp()
	app.Name = "genesishash"
	app.Usage = "Generate Hash for genesis json"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "gen",
			Usage: "path to genesis.json file",
		},
	}
	app.Action = runGenHash
	app.Run(os.Args)
}

// runWizard start the wizard and relinquish control to it.
func runGenHash(c *cli.Context) error {
	genesisPath := c.String("gen")

	if len(genesisPath) == 0 {
		utils.Fatalf("Must supply path to genesis JSON file")
	}
	file, err := os.Open(genesisPath)
	if err != nil {
		utils.Fatalf("Failed to read genesis file: %v", err)
	}
	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		utils.Fatalf("invalid genesis file: %v", err)
	}

	block := genesis.ToBlock(nil)

	println(block.Hash().Hex())

	return nil
}
