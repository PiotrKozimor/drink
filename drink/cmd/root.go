package cmd

import (
	"log"
	"strconv"
	"time"

	"github.com/PiotrKozimor/drink"
	"github.com/spf13/cobra"
)

var dbPath string

var rootCmd = &cobra.Command{
	Use:   "drink",
	Short: "Drink",
	Long:  `Pass number of milliliters as first argument.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		much, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		s, err := drink.NewStore(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer s.DB.Close()
		s.DrinkIt(&drink.Drink{
			HowMuch: uint32(much),
			When:    time.Now(),
		})
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbPath, "database", "d", ".drink/drinks.db", "path to database file")
}
