package entities

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

type Slot int

const (
	HEAD  Slot = 0
	TORSE Slot = 1
	LEGS  Slot = 2
	FEET  Slot = 3
	ARMS  Slot = 4
	HANDS Slot = 5
	BACK  Slot = 6

	OFF_HAND  Slot = 7
	MAIN_HAND Slot = 8
)

type EquippedGear struct {
	slots map[Slot]*Item
}

func (gear EquippedGear) GetItemFromSlot(slot Slot) *Item { return gear.slots[slot] }

func (gear *EquippedGear) Equip(item *Item) error {
	if item.Equipped() {
		err := errors.New("item is already equipped")
		log.WithError(err).WithField("itemName", item.Name())
		return err
	}

	itemSlots := item.Slots()

	if len(itemSlots) == 0 {
		err := errors.New("cannot equip this item")
		log.WithError(err).WithField("itemName", item.Name())
		return err
	}

	slotsOpen := true
	for _, slot := range itemSlots {
		if gear.slots[slot] != nil {
			slotsOpen = false
		}
	}

	if !slotsOpen {
		err := errors.New("slot(s) not available")
		log.WithError(err).WithFields(log.Fields{"itemName": item.Name(), "itemSlots": itemSlots})
		return err
	}

	for _, slot := range itemSlots {
		gear.slots[slot] = item
	}

	item.equipped = true
	return nil
}

func (gear *EquippedGear) Unequip(item *Item) error {
	if !item.Equipped() {
		err := errors.New("item was already not equipped")
		log.WithError(err).WithField("itemName", item.Name())
		return err
	}

	for _, slot := range item.Slots() {
		gear.slots[slot] = nil
	}

	item.equipped = false
	return nil
}
