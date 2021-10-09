package drink

import (
	"bytes"
	"encoding/binary"
	"time"

	"go.etcd.io/bbolt"
)

type Store struct {
	*bbolt.DB
}

func NewStore(path string) (*Store, error) {

	db, err := bbolt.Open(path, 0644, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(t *bbolt.Tx) error {
		if getBucket(t) == nil {
			_, err := t.CreateBucket([]byte{1})
			return err
		}
		return nil
	})
	return &Store{
		DB: db,
	}, err
}

func (s *Store) DrinkIt(d *Drink) error {
	k, v := marshall(d)
	return s.DB.Update(func(t *bbolt.Tx) error {
		return getBucket(t).Put(k, v)
	})
}

func (s *Store) RemoveLastDrink() error {
	return s.DB.Update(func(t *bbolt.Tx) error {
		b := getBucket(t)
		c := b.Cursor()
		k, _ := c.Last()
		return b.Delete(k)
	})
}

func (s *Store) GetTodaysDrinks(when time.Time) ([]Drink, error) {
	since := []byte(when.Truncate(24 * time.Hour).Format(time.RFC3339))
	to := []byte(when.Truncate(24 * time.Hour).Add(24 * time.Hour).Format(time.RFC3339))
	drinks := make([]Drink, 0, 2)
	return drinks, s.DB.View(func(t *bbolt.Tx) error {
		b := getBucket(t)
		c := b.Cursor()
		for k, v := c.Seek(since); k != nil && bytes.Compare(k, to) <= 0; k, v = c.Next() {
			drink, err := unmarshall(k, v)
			if err != nil {
				return err
			}
			drinks = append(drinks, *drink)
		}
		return nil
	})
}

func (s *Store) GetAllDrinks() ([]DailyDrinks, error) {
	var dailyDrinks []DailyDrinks
	var day *time.Time
	var dailyDrink DailyDrinks
	return dailyDrinks, s.DB.View(func(t *bbolt.Tx) error {
		b := getBucket(t)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			drink, err := unmarshall(k, v)
			if err != nil {
				return err
			}
			nextDay := drink.When.Truncate(24 * time.Hour)
			if day == nil {
				dailyDrink.Day = nextDay
				day = &nextDay
			}
			if nextDay != *day {
				dailyDrinks = append(dailyDrinks, dailyDrink)
				dailyDrink.Day = nextDay
				dailyDrink.Drinks = nil
				day = &nextDay
			}
			dailyDrink.Drinks = append(dailyDrink.Drinks, *drink)
		}
		dailyDrinks = append(dailyDrinks, dailyDrink)
		return nil
	})
}

func getBucket(t *bbolt.Tx) *bbolt.Bucket {
	return t.Bucket([]byte{1})
}

func marshall(d *Drink) (k, v []byte) {
	key := d.When.Format(time.RFC3339)
	var val [4]byte
	binary.LittleEndian.PutUint32(val[:], d.HowMuch)
	k = []byte(key)
	v = val[:]
	return
}

func unmarshall(k, v []byte) (*Drink, error) {
	key, err := time.Parse(time.RFC3339, string(k))
	if err != nil {
		return nil, err
	}
	milliliters := binary.LittleEndian.Uint32(v)
	return &Drink{
		HowMuch: milliliters,
		When:    key,
	}, nil
}
