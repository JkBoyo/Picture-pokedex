package pokeapi

type pokemon struct {
	Abilities      []Abilities `json:"abilities"`
	BaseExperience int         `json:"base_experience"`
	ID             int         `json:"id"`
	IsDefault      bool        `json:"is_default"`
	Name           string      `json:"name"`
	Order          int         `json:"order"`
	Species        Species     `json:"species"`
	Sprites        Sprites     `json:"sprites"`
	Stats          []Stats     `json:"stats"`
	Types          []Types     `json:"types"`
	Weight         int         `json:"weight"`
}

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Other struct {
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}

type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type Sprites struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
	Other            Other  `json:"other"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
