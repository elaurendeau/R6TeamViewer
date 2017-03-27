package domain

import "time"

//Player struct represent the statistic for a specific player
type Player struct {
	Player struct {
		Username  string    `json:"username"`
		Platform  string    `json:"platform"`
		UbisoftID string    `json:"ubisoft_id"`
		IndexedAt time.Time `json:"indexed_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Stats struct {
			Ranked struct {
				HasPlayed bool    `json:"has_played"`
				Wins      int     `json:"wins"`
				Losses    int     `json:"losses"`
				Wlr       float64 `json:"wlr"`
				Kills     int     `json:"kills"`
				Deaths    int     `json:"deaths"`
				Kd        float64 `json:"kd"`
				Playtime  int     `json:"playtime"`
			} `json:"ranked"`
			Casual struct {
				HasPlayed bool    `json:"has_played"`
				Wins      int     `json:"wins"`
				Losses    int     `json:"losses"`
				Wlr       float64 `json:"wlr"`
				Kills     int     `json:"kills"`
				Deaths    int     `json:"deaths"`
				Kd        float64 `json:"kd"`
				Playtime  int     `json:"playtime"`
			} `json:"casual"`
			Overall struct {
				Revives                int `json:"revives"`
				Suicides               int `json:"suicides"`
				ReinforcementsDeployed int `json:"reinforcements_deployed"`
				BarricadesBuilt        int `json:"barricades_built"`
				StepsMoved             int `json:"steps_moved"`
				BulletsFired           int `json:"bullets_fired"`
				BulletsHit             int `json:"bullets_hit"`
				Headshots              int `json:"headshots"`
				MeleeKills             int `json:"melee_kills"`
				PenetrationKills       int `json:"penetration_kills"`
				Assists                int `json:"assists"`
			} `json:"overall"`
			Progression struct {
				Level int `json:"level"`
				Xp    int `json:"xp"`
			} `json:"progression"`
		} `json:"stats"`
	} `json:"player"`
}

//PlayerRepository describe the methodes used for the player repository
type PlayerRepository interface {
	//FindByProfileNameAndPlatform is used to retrieve a struct of player
	FindByProfileNameAndPlatform(profileName string, platform string) (*Player, error)
}
