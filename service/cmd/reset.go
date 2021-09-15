package cmd

import (
	"github.com/adrianrudnik/yealink-url-scheme-handler/service/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:  "reset",
	Short: "Deletes local configuration",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.DeleteConfig()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to delete config file")
		}

		log.Info().Msg("Config file deleted")
	},
}
