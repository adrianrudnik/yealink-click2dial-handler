package cmd

import (
	"github.com/adrianrudnik/yealink-click2dial-handler/internal"
	"github.com/adrianrudnik/yealink-click2dial-handler/pkg/yealink"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var dialCmd = &cobra.Command{
	Use:  "dial",
	Example: `dial "+49123"`,
	Short: "Dials the passed phone number",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		device, err := internal.LoadConfig()
		if err != nil {
			if err == internal.ErrNotConfigured {
				log.Fatal().Err(err).Msg(`Not configured, please run the "connect" command first`)
			}

			log.Fatal().Err(err).Msg("Failed to load configuration")
		}

		err = yealink.Call(&device, args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to initialize call")
		}
	},
}
