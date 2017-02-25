package domain

import "time"

type Profile struct {
	Name string
	Platform string

	Operators *Operators
	Player *Player
	Seasons *Seasons
}

type Operators struct {
	OperatorRecords []struct {
		Stats struct {
			Played int `json:"played"`
			Wins int `json:"wins"`
			Losses int `json:"losses"`
			Kills int `json:"kills"`
			Deaths int `json:"deaths"`
			Playtime int `json:"playtime"`
			Specials struct {
				OperatorpvpRookArmorboxdeployed string `json:"operatorpvp_rook_armorboxdeployed"`
				OperatorpvpRookArmortakenourself string `json:"operatorpvp_rook_armortakenourself"`
				OperatorpvpRookArmortakenteammate string `json:"operatorpvp_rook_armortakenteammate"`
			} `json:"specials"`
		} `json:"stats"`
		Operator struct {
			Name string `json:"name"`
			Ctu string `json:"ctu"`
			Images struct {
				Figure string `json:"figure"`
				Badge string `json:"badge"`
				Bust string `json:"bust"`
			} `json:"images"`
		} `json:"operator"`
	} `json:"operator_records"`
}

type Player struct {
	Player struct {
		Username string `json:"username"`
		Platform string `json:"platform"`
		UbisoftID string `json:"ubisoft_id"`
		IndexedAt time.Time `json:"indexed_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Stats struct {
			Ranked struct {
				HasPlayed bool `json:"has_played"`
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Wlr float64 `json:"wlr"`
				Kills int `json:"kills"`
				Deaths int `json:"deaths"`
				Kd float64 `json:"kd"`
				Playtime int `json:"playtime"`
			} `json:"ranked"`
			Casual struct {
				HasPlayed bool `json:"has_played"`
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Wlr float64 `json:"wlr"`
				Kills int `json:"kills"`
				Deaths int `json:"deaths"`
				Kd float64 `json:"kd"`
				Playtime int `json:"playtime"`
			} `json:"casual"`
			Overall struct {
				Revives int `json:"revives"`
				Suicides int `json:"suicides"`
				ReinforcementsDeployed int `json:"reinforcements_deployed"`
				BarricadesBuilt int `json:"barricades_built"`
				StepsMoved int `json:"steps_moved"`
				BulletsFired int `json:"bullets_fired"`
				BulletsHit int `json:"bullets_hit"`
				Headshots int `json:"headshots"`
				MeleeKills int `json:"melee_kills"`
				PenetrationKills int `json:"penetration_kills"`
				Assists int `json:"assists"`
			} `json:"overall"`
			Progression struct {
				Level int `json:"level"`
				Xp int `json:"xp"`
			} `json:"progression"`
		} `json:"stats"`
	} `json:"player"`
}

type Seasons struct {
	Seasons struct {
		Num4 struct {
			Ncsa struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"4"`
		Num5 struct {
			Ncsa struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"5"`
	} `json:"seasons"`
}
