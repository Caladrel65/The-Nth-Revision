package entities

import log "github.com/sirupsen/logrus"

type Character struct {
	name      string
	level     int
	maxHealth int
	health    int
	damage    int
	gear      *EquippedGear
}

func NewCharacter(name string, level int, maxHealth int, damage int) Character {
	return Character{
		name:      name,
		level:     level,
		maxHealth: maxHealth,
		health:    maxHealth,
		damage:    damage,
		gear:      &EquippedGear{},
	}
}

func (c Character) Name() string            { return c.name }
func (c Character) Level() int              { return c.level }
func (c Character) Health() int             { return c.health }
func (c Character) Damage() int             { return c.damage }
func (c Character) Gear() *EquippedGear     { return c.gear }
func (c Character) PrintHealth()            { log.Infof("%s has %d/%d health", c.name, c.health, c.maxHealth) }
func (c Character) canPerformActions() bool { return c.health > 0 }

func (c *Character) TakeDamage(damage int) int {
	maxDamage := c.health
	if damage > maxDamage {
		damage = maxDamage
	}
	c.health -= damage

	return damage
}

func (c *Character) TakeHealing(healing int) int {
	maxHealing := c.maxHealth - c.health
	if healing > maxHealing {
		healing = maxHealing
	}
	c.health += healing

	return healing
}

func (c Character) Attack(target Entity) int {
	if !c.canPerformActions() {
		return 0
	}

	return target.TakeDamage(c.damage)
}

func (c Character) UseItem(item Item) {

}
