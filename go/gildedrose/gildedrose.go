package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/gildedrose/model"

/* 	- Once the sell by date has passed, Quality degrades twice as fast
	- The Quality of an item is never negative
	- "Aged Brie" actually increases in Quality the older it gets
	- The Quality of an item is never more than 50
	- "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
	- "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
	Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
	Quality drops to 0 after the concert

We have recently signed a supplier of conjured items. This requires an update to our system:

	- "Conjured" items degrade in Quality twice as fast as normal items

	Operations for all items:

	All items have a SellIn value which denotes the number of days we have to sell the item
	All items have a Quality value which denotes how valuable the item is
	At the end of each day our system lowers both values for every item
	Once the sell by date has passed, Quality degrades twice as fast
	The Quality of an item is never negative
	Operations for specific items:

	"Aged Brie" actually increases in Quality the older it gets
	The Quality of an item is never more than 50
	"Sulfuras", being a legendary item, never has to be sold or decreases in Quality
	"Backstage passes", like aged brie, increases in Quality as its SellIn value approaches; Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but Quality drops to 0 after the concert
	"Conjured" items degrade in Quality twice as fast as normal items
*/

type GildedRoseItem interface {
	UpdateQuality()
}

type GenerallItem struct {
	item *model.Item
}

type AgedBrie struct {
	item *model.Item
}

type Sulfuras struct {
	item *model.Item
}

type BackstagePasses struct {
	item *model.Item
}

type Conjured struct {
	item *model.Item
}

func NewItems(gildedRoseItems []*model.Item) []GildedRoseItem {
	items := make([]GildedRoseItem, 0, len(gildedRoseItems))
	for _, item := range gildedRoseItems {
		switch item.Name {
		case "Aged Brie":
			items = append(items, &AgedBrie{item: item})
		case "Sulfuras, Hand of Ragnaros":
			items = append(items, &Sulfuras{item: item})
		case "Backstage passes to a TAFKAL80ETC concert":
			items = append(items, &BackstagePasses{item: item})
		case "Conjured Mana Cake":
			items = append(items, &Conjured{item: item})
		default:
			items = append(items, &GenerallItem{item: item})
		}
	}
	return items
}

func (gi *GenerallItem) UpdateQuality() {
	/* 	All items have a SellIn value which denotes the number of days we have to sell the item
	   	All items have a Quality value which denotes how valuable the item is
	   	At the end of each day our system lowers both values for every item
	*/
	gi.item.SellIn--
	gi.item.Quality--
}

func (ab *AgedBrie) UpdateQuality() {
	// Will update the quality of the item
}

func (s *Sulfuras) UpdateQuality() {
	// Will update the quality of the item
}

func (bp *BackstagePasses) UpdateQuality() {
	// Will update the quality of the item
}

func (c *Conjured) UpdateQuality() {
	// Will update the quality of the item
}

/* func UpdateQuality(items []*model.Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}
} */
