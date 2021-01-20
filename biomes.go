package main

const (
	AQUATIC   string = "Aquatic"
	DESERT    string = "Desert"
	GRASSLAND string = "Grassland"
	TAIGA     string = "Taiga"
	TEMPERATE string = "Temperate"
	TROPICAL  string = "Tropical"
	TUNDRA    string = "Tundra"
)

func getBiomes(row []interface{}) []string {
	var biomes []string
	if row[6] == "x" {
		biomes = append(biomes, AQUATIC)
	}
	if row[7] == "x" {
		biomes = append(biomes, DESERT)
	}
	if row[8] == "x" {
		biomes = append(biomes, GRASSLAND)
	}
	if row[9] == "x" {
		biomes = append(biomes, TEMPERATE)
	}
	if row[10] == "x" {
		biomes = append(biomes, TROPICAL)
	}
	if row[11] == "x" {
		biomes = append(biomes, TAIGA)
	}
	if row[12] == "x" {
		biomes = append(biomes, TUNDRA)
	}

	return biomes
}
