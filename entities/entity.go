package entities

type Entity interface {
	Name() string
	TakeDamage(int) int
	TakeHealing(int) int
	Health() int
	PrintHealth()
}
