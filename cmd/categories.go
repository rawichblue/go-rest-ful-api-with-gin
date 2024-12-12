package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var categories = &cobra.Command{
	Use:     "migrate-categories",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating categories")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Categories)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model categories up success")
	},
}

func init() {
	rootCmd.AddCommand(categories)
}
