// TODO: this file is getting bit. Maybe parts of it can be moved
package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ionut-maxim/notarry/internal"
	"github.com/ionut-maxim/notarry/pkg/message"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Slack struct {
		Webhook  string `mapstructure:"webhook"`
		Template string `mapstructure:"template"`
	} `mapstructure:"slack"`
}

const (
	appName = "notarry"
)

var (
	//go:embed assets/template.yaml
	defaultGrabTemplate string

	config Config

	rootCmd = &cobra.Command{
		Use:   "notarry",
		Short: "A notifier for *arr applications",
		Long: `notarry is a CLI application for *arr notifications.
This application can send Slack messages when user
with any *arr application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			webhook := config.Slack.Webhook
			template := config.Slack.Template

			if err := message.SendSlackMessage(webhook, template); err != nil {
				log.Print(err)
				return err
			}

			return nil
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	envs := internal.GetEnvVariables("radarr")

	log.Info().Str("environment_variables", fmt.Sprint(strings.Join(envs, "\n"))).Msg("Event received")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(appName)

	ex, err := os.Executable()
	if err != nil {
		log.Print(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Print(err)
		}
	}

	viper.SetEnvPrefix(appName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("slack.template", defaultGrabTemplate)
	// FIXME: Currently using workaround #2 from this issue
	// https://github.com/spf13/viper/issues/761 for problems
	// with Environment Variables and Unmarshalling the config
	viper.SetDefault("slack.webhook", "")

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
	}
}
