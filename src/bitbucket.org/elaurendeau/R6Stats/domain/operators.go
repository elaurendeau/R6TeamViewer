package domain

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
