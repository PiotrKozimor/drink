package drink

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"
)

type Drink struct {
	HowMuch uint32
	When    time.Time
}

type DailyDrinks struct {
	Drinks []Drink
	Day    time.Time
}

const (
	kitchen   = "03:04PM"
	dayFormat = "02 Jan 06"
)

func Print(d []Drink) {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	printDrinks("", w, d)
	w.Flush()
}

func printDrinks(indent string, w io.Writer, d []Drink) {
	sum := uint32(0)
	for i, drink := range d {
		sum += drink.HowMuch
		fmt.Fprintf(w, "%s%d\t%s\t%d\n", indent, i+1, drink.When.Format(kitchen), drink.HowMuch)
	}
	fmt.Fprintf(w, "%sSUM\t\t\t\t%d\n", indent, sum)
}

func PrintAll(d []DailyDrinks) {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	for _, ddrinks := range d {
		fmt.Fprintf(w, "%s\n", ddrinks.Day.Format(dayFormat))
		printDrinks("\t\t", w, ddrinks.Drinks)
	}
	w.Flush()
}
