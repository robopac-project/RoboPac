package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/pagu-project/Pagu/internal/entity"

	"github.com/pagu-project/Pagu/internal/engine"
	"github.com/pagu-project/Pagu/pkg/log"

	"github.com/pactus-project/pactus/crypto"
	pagu "github.com/pagu-project/Pagu"
	pCmd "github.com/pagu-project/Pagu/cmd"
	"github.com/pagu-project/Pagu/config"
	cobra "github.com/spf13/cobra"
)

var configPath string

const PROMPT = "\n>> "

func run(cmd *cobra.Command, args []string) {
	configs, err := config.Load(configPath)
	pCmd.ExitOnError(cmd, err)

	log.InitGlobalLogger(configs.Logger)

	if configs.Network == "Localnet" {
		crypto.AddressHRP = "tpc"
	}

	botEngine, err := engine.NewBotEngine(configs)
	pCmd.ExitOnError(cmd, err)

	botEngine.RegisterAllCommands()
	botEngine.Start()

	reader := bufio.NewReader(os.Stdin)

	for {
		cmd.Print(PROMPT)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if strings.ToLower(input) == "exit" {
			cmd.Println("exiting from cli")

			return
		}

		inputs := strings.Split(input, " ")

		response := botEngine.Run(entity.AppIdCLI, "0", inputs)

		cmd.Printf("%v\n%v", response.Title, response.Message)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "pagu-cli",
		Version: pagu.StringVersion(),
		Run:     run,
	}

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config.yml", "config path ./config.yml")
	err := rootCmd.Execute()
	pCmd.ExitOnError(rootCmd, err)
}
