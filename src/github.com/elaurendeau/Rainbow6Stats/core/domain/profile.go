package domain

//Profile is a struct to represent a user profile.
type Profile struct {
	Name     string
	Platform string

	Operators *Operators
	Player    *Player
	Seasons   *Seasons
}
