package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var activityLog = &cobra.Command{
	Use:     "migrate-activityLog",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating activityLog")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.ActivityLog)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model activityLog up success")
	},
}

func init() {
	rootCmd.AddCommand(activityLog)
}
