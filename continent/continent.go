package continent

type (
	Code string
	Continent struct {
		Code Code
		Name string
	}
	Continents map[Code]Continent
)

const (
	Africa       Code = "AF"
	NorthAmerica Code = "NA"
	Oceania      Code = "OC"
	Antarctica   Code = "AN"
	Asia         Code = "AS"
	Europe       Code = "EU"
	SouthAmerica Code = "SA"
)

var continents = Continents{
	Africa: Continent{
		Code: Africa,
		Name: "Africa",
	},
	NorthAmerica: Continent{
		Code: NorthAmerica,
		Name: "North America",
	},
	Oceania: Continent{
		Code: Oceania,
		Name: "Oceania",
	},
	Antarctica: Continent{
		Code: Antarctica,
		Name: "Antarctica",
	},
	Asia: Continent{
		Code: Asia,
		Name: "Asia",
	},
	Europe: Continent{
		Code: Europe,
		Name: "Europe",
	},
	SouthAmerica: Continent{
		Code: SouthAmerica,
		Name: "South America",
	},
}

func (c Code) Continent() Continent {
	return continents[c]
}

func List() Continents {
	return continents
}

func Has(c string) bool {
	if _, ok := continents[Code(c)]; ok {
		return true
	}
	return false
}
