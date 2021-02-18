package viewmodels

type Home struct {
	Title string
}

func GetHome() Home {
	result := Home{
		Title: "HomePageOfNexflix",
	}

	return result
}
