package region

type (
	Code string
	Region struct {
		Code Code
		Name string
	}
	Regions map[Code]Region
)

const (
	Africa   Code = "2"
	Americas Code = "19"
	Asia     Code = "142"
	Europe   Code = "150"
	Oceania  Code = "9"
)

var regions = Regions{
	Africa: Region{
		Code: Africa,
		Name: "Africa",
	},
	Americas: Region{
		Code: Americas,
		Name: "Americas",
	},
	Asia: Region{
		Code: Asia,
		Name: "Asia",
	},
	Europe: Region{
		Code: Europe,
		Name: "Europe",
	},
	Oceania: Region{
		Code: Oceania,
		Name: "Oceania",
	},
}

func (r Code) Region() Region {
	return regions[r]
}

func List() Regions {
	return regions
}
