package viewmodels

type Home struct {
	Title string
}

func GetHome() Home {
	result := Home{
		Title: "Towed Car Locator - Home",
	}

	return result
}
