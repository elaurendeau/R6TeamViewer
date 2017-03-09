package domain

type Profile struct {
	Name     string
	Platform string

	Operators *Operators
	Player    *Player
	Seasons   *Seasons
}
