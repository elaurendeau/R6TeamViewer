package domain

//Seasons structure represent the season statistics for a player
type Seasons struct {
	Seasons struct {
		Num4 struct {
			Ncsa struct {
				Wins     int    `json:"wins"`
				Losses   int    `json:"losses"`
				Abandons int    `json:"abandons"`
				Season   int    `json:"season"`
				Region   string `json:"region"`
				Ranking  struct {
					Rating     float64 `json:"rating"`
					NextRating int     `json:"next_rating"`
					PrevRating int     `json:"prev_rating"`
					Mean       float64 `json:"mean"`
					Stdev      int     `json:"stdev"`
					Rank       int     `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins     int    `json:"wins"`
				Losses   int    `json:"losses"`
				Abandons int    `json:"abandons"`
				Season   int    `json:"season"`
				Region   string `json:"region"`
				Ranking  struct {
					Rating     float64 `json:"rating"`
					NextRating int     `json:"next_rating"`
					PrevRating int     `json:"prev_rating"`
					Mean       float64 `json:"mean"`
					Stdev      int     `json:"stdev"`
					Rank       int     `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"4"`
		Num5 struct {
			Ncsa struct {
				Wins     int    `json:"wins"`
				Losses   int    `json:"losses"`
				Abandons int    `json:"abandons"`
				Season   int    `json:"season"`
				Region   string `json:"region"`
				Ranking  struct {
					Rating     float64 `json:"rating"`
					NextRating int     `json:"next_rating"`
					PrevRating int     `json:"prev_rating"`
					Mean       float64 `json:"mean"`
					Stdev      int     `json:"stdev"`
					Rank       int     `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins     int    `json:"wins"`
				Losses   int    `json:"losses"`
				Abandons int    `json:"abandons"`
				Season   int    `json:"season"`
				Region   string `json:"region"`
				Ranking  struct {
					Rating     float64 `json:"rating"`
					NextRating int     `json:"next_rating"`
					PrevRating int     `json:"prev_rating"`
					Mean       float64 `json:"mean"`
					Stdev      int     `json:"stdev"`
					Rank       int     `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"5"`
	} `json:"seasons"`
}

//OperatorRepository describe the methods used to retrieve the Season structure
type SeasonRepository interface {
	//FindByProfileNameAndPlatform is used to retrieve a struct of seasons
	FindByProfileNameAndPlatform(profileName string, platform string) (*Seasons, error)
}
