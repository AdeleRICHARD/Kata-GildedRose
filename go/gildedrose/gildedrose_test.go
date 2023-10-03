package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/model"
)

func Test_Foo(t *testing.T) {
	var items = []*model.Item{
		{
			Name:    "Warglaive of Azzinoth",
			SellIn:  10,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
