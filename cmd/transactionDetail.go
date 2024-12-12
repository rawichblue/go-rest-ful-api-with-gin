package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var transactionDetail = &cobra.Command{
	Use:     "migrate-transactionDetail",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating transactionDetail")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.TransactionsDetail)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model transactionDetail up success")
	},
}

func init() {
	rootCmd.AddCommand(transactionDetail)
}
