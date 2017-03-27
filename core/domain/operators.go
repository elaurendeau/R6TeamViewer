package domain

//Operators is a struct used to describe an operator
//It should contain the records of all operators.
type Operators struct {
	OperatorRecords []struct {
		Stats struct {
			Played   int `json:"played"`
			Wins     int `json:"wins"`
			Losses   int `json:"losses"`
			Kills    int `json:"kills"`
			Deaths   int `json:"deaths"`
			Playtime int `json:"playtime"`
			Specials struct {
				OperatorpvpRookArmorboxdeployed   string `json:"operatorpvp_rook_armorboxdeployed"`
				OperatorpvpRookArmortakenourself  string `json:"operatorpvp_rook_armortakenourself"`
				OperatorpvpRookArmortakenteammate string `json:"operatorpvp_rook_armortakenteammate"`
			} `json:"specials"`
		} `json:"stats"`
		Operator struct {
			Name   string `json:"name"`
			Ctu    string `json:"ctu"`
			Images struct {
				Figure string `json:"figure"`
				Badge  string `json:"badge"`
				Bust   string `json:"bust"`
			} `json:"images"`
		} `json:"operator"`
	} `json:"operator_records"`
}

//OperatorRepository describe the methodes used for the operator repository
type OperatorRepository interface {
	//FindByProfileNameAndPlatform is used to retrieve a struct of operators
	FindByProfileNameAndPlatform(profileName string, platform string) (*Operators, error)
}
