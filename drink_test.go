package drink

import (
	"testing"
	"time"
)

func TestPrintDrinks(t *testing.T) {
	testDrinks := []Drink{
		{
			HowMuch: 240,
			When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
		},
		{
			HowMuch: 300,
			When:    time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
		},
	}
	Print(testDrinks)
}

func TestPrintAllDrinks(t *testing.T) {
	testDrinks := map[time.Time][]Drink{
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC): {
			{
				HowMuch: 240,
				When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
			},
			{
				HowMuch: 300,
				When:    time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
			},
		},
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC): {
			{
				HowMuch: 240,
				When:    time.Date(2021, 1, 2, 12, 1, 1, 0, time.UTC),
			},
		},
	}
	PrintAll(testDrinks)

}
