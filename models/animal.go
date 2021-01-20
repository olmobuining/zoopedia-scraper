package models

type LifeExpectancy struct {
	Male   int `json:"male"`
	Female int `json:"female"`
}

type Animal struct {
	Name                string              `json:"name"`
	Continent           []string            `json:"continent"`
	Biomes              []string            `json:"biome"`
	CanSwim             bool                `json:"can_swim"`
	Status              string              `json:"status"`
	IsExhibit           bool                `json:"exhibit"`
	Dominance           string              `json:"dominance"`
	RelationshipHuman   string              `json:"relationship_human"`
	MatingSystem        string              `json:"mating_system"`
	LifeExpectancy      LifeExpectancy      `json:"life_expectancy"`
	HabitatRequirements HabitatRequirements `json:"habitat_requirements"`
	LatinName           string              `json:"latin_name"`
	Category            string              `json:"category"`
	Description         string              `json:"description"`
	Image               string              `json:"image_url"`
	Slug                string              `json:"slug"`
	Key                 string              `json:"key"`
}

type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type SizeMinMax struct {
	Size MinMax `json:"size"`
}

type SizeMixed struct {
	Size   MinMax `json:"size"`
	Male   int    `json:"male"`
	Female int    `json:"female"`
}

type HabitatRequirements struct {
	GuestsCanEnter      bool       `json:"guest_enter"`
	IsExhibit           bool       `json:"exhibit"`
	Land                int        `json:"land"`
	LandAdditional      int        `json:"land_additional"`
	Water               int        `json:"water"`
	WaterAdditional     int        `json:"water_additional"`
	Climbable           int        `json:"climbable"`
	ClimbableAdditional int        `json:"climbable_additional"`
	Humidity            MinMax     `json:"humidity"`
	TemperatureCelsuis  MinMax     `json:"temperature"`
	Fence               FenceInfo  `json:"fence"`
	GroupMale           SizeMinMax `json:"group_male"`
	GroupFemale         SizeMinMax `json:"group_female"`
	GroupMixed          SizeMixed  `json:"group_mixed"`
}

type FenceInfo struct {
	Grade      int     `json:"grade"`
	Height     float32 `json:"height"`
	ClimbProof bool    `json:"climbproof"`
}
