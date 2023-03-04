package entities

import log "github.com/sirupsen/logrus"

type Item struct {
	name      string
	level     int
	itemType  int // TODO: Add enum values?
	maxHealth int
	health    int
	equipped  bool
	slots     []Slot
}

func NewObject(name string, level int, maxHealth int, damage int) Character {
	return Character{maxHealth: maxHealth, level: level, health: maxHealth, damage: damage}
}

func (i Item) Name() string   { return i.name }
func (i Item) Level() int     { return i.level }
func (i Item) Health() int    { return i.health }
func (i Item) ItemType() int  { return i.itemType }
func (i Item) Equipped() bool { return i.equipped }
func (i Item) Slots() []Slot  { return i.slots }
func (i Item) PrintHealth()   { log.Infof("%s has %d/%d health", i.name, i.health, i.maxHealth) }

func (i *Item) TakeDamage(damage int) int {
	maxDamage := i.health
	if damage > maxDamage {
		damage = maxDamage
	}
	i.health -= damage

	return damage
}

func (i *Item) TakeHealing(healing int) int {
	maxHealing := i.maxHealth - i.health
	if healing > maxHealing {
		healing = maxHealing
	}
	i.health += healing

	return healing
}
