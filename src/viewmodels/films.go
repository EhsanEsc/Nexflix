package viewmodels

type Films struct {
	Title               string
	PropertiesNames     []string
	Filtered_Film_score int
	Films               []Film
}

// Add Director Name and stuff to bussinus model not here
type Film struct {
	ID               int
	Name             string
	Score            int
	PropertiesValues []string
}

func GetFilms() Films {
	result := Films{
		Title: "ListOfFilms",
		PropertiesNames: []string{
			"Name",
			"DirectorName",
			"Score",
		},
	}

	result.Films = append(result.Films, Film{
		ID:    1,
		Name:  "HunterXHunter",
		Score: 10,
		PropertiesValues: []string{
			"Hiroshi Kōjina",
			"10",
		}})

	result.Films = append(result.Films, Film{
		ID:    2,
		Name:  "Attack On Titan",
		Score: 22,
		PropertiesValues: []string{
			"Hajime Isayama",
			"22",
		}})

	return result
}
