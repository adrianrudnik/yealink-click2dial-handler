package cmd

import (
	"bufio"
	"fmt"
	"github.com/adrianrudnik/yealink-click2dial-handler/internal"
	"github.com/adrianrudnik/yealink-click2dial-handler/pkg/yealink"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var connectCmd = &cobra.Command{
	Use:  "connect",
	Example: "connect 192.168.0.100 admin [password]",
	Short: "Connects the given IP address and configures it for the scheme handler",
	Args: cobra.RangeArgs(2, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse and validate the given IP address
		ip := net.ParseIP(args[0])

		if ip == nil {
			log.Fatal().Msg("IP is not a valid textual representation and could not be parsed")
		}

		log.Info().IPAddr("phone", ip).Msg("Parsed phone IP")

		username := args[1]

		password := ""

		// Ask for password if not already given by argument
		if len(args) >= 3 {
			password = args[2]
		} else {
			p, err := askPassword()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to get password input")
			}

			password = p
		}

		// Try to resolve the API client
		client := yealink.GetApiClient()
		r, err := client.Get(fmt.Sprintf("http://%s/", ip.String()))

		if err != nil {
			log.Fatal().Err(err).Msg("Failed to validate given endpoint through HTTP")
		}

		device := yealink.Device{
			IP: ip,
			Username: username,
			Password: password,
		}

		// Detect SSL capabilities
		if r.Request.URL.Scheme == "https" {
			device.IsHTTPS = true
			log.Info().Msg("Device endpoint uses SSL")
		}

		// Trigger the first account read, lets see if we can pass through
		log.Info().Msg("Triggering config read, please watch the phone and confirm a possible access request screen, confirm it, then restart this command")

		accounts, err := yealink.ReadAccounts(&device)

		if err != nil {
			// Send notice about remote control API endpoint
			if err == yealink.ErrCommandFailedBadCredentials {
				log.Warn().Msg("Ensure you have this PC in the Action URI Allow IP List")
				u, err := device.GetRootUrl("/api#/features-remotecontrl")
				if err == nil {
					log.Warn().Msg(u.String())
				}
			}

			log.Fatal().Err(err).Msg("Failed to retrieve accounts")
		}

		var account yealink.Account

		// Lookup the default outgoing uri if given
		if len(args) == 4 {
			found := false
			// @todo look through the accounts
			for _, a := range accounts {
				if a.OutgoingURI == args[3] || a.DisplayName == args[3] || a.Label == args[3] || a.RegisterName == args[3] {
					account = a
					found = true
					break
				}
			}

			if !found {
				log.Fatal().Str("uri", args[3]).Msg("Failed to find the requested outgoing uri")
			}
		} else {
			// Let the user pick one
			account = askAccount(accounts)
		}

		device.DefaultOutgoingURI = account.OutgoingURI

		log.Info().Str("uri", account.OutgoingURI).Msg("Configured default outgoing URI")

		// Store the configuration
		path, err := internal.GetConfigFilePath()
		if err != nil {
			log.Fatal().Str("path", path).Err(err).Msg("Failed to initialize configuration file")
		}

		err = internal.StoreConfig(device)
		if err != nil {
			log.Fatal().Str("path", path).Err(err).Msg("Failed to store configuration")
		}

		log.Info().Str("path", path).Msg("Configuration stored")

		return nil
	},
}

func askPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the password:")
	fmt.Println("")

	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)

	return password, nil
}

func askAccount(accounts []yealink.Account) yealink.Account {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please pick the default outgoing URI:")
	fmt.Println("")

	for i, account := range accounts {
		fmt.Println(fmt.Sprintf("%d) %s [%s]", i, account.DisplayName, account.OutgoingURI))
	}

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		// Validate input
		in, err := strconv.Atoi(text)
		if err != nil || in < 0 || in > len(accounts)-1 {
			log.Error().Msgf("Invalid input, please choose a value between 0 and %d", len(accounts)-1)
			time.Sleep(100 * time.Millisecond)
		} else {
			return accounts[in]
		}
	}
}
