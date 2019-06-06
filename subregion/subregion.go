package subregion

import (
	"github.com/Pingflow/go-localization/region"
)

type (
	Code string
	SubRegion struct {
		Code   Code
		Name   string
		Region region.Code
	}
	SubRegions map[Code]SubRegion
)

const (
	LatinAmericaAndTheCaribbean Code = "419"
	NorthernAmerica             Code = "21"
	SubSaharanAfrica            Code = "202"
	NorthernAfrica              Code = "15"
	AustraliaAndNewZealand      Code = "53"
	Polynesia                   Code = "61"
	Melanesia                   Code = "54"
	Micronesia                  Code = "57"
	SouthernAsia                Code = "34"
	WesternAsia                 Code = "145"
	SouthEasternAsia            Code = "35"
	EasternAsia                 Code = "30"
	CentralAsia                 Code = "143"
	NorthernEurope              Code = "154"
	SouthernEurope              Code = "39"
	WesternEurope               Code = "155"
	EasternEurope               Code = "151"
)

var subRegions = SubRegions{
	LatinAmericaAndTheCaribbean: SubRegion{
		Code:   LatinAmericaAndTheCaribbean,
		Name:   "Latin America and the Caribbean",
		Region: region.Americas,
	},
	NorthernAmerica: SubRegion{
		Code:   NorthernAmerica,
		Name:   "Northern America",
		Region: region.Americas,
	},
	SubSaharanAfrica: SubRegion{
		Code:   SubSaharanAfrica,
		Name:   "Sub-Saharan Africa",
		Region: region.Africa,
	},
	NorthernAfrica: SubRegion{
		Code:   NorthernAfrica,
		Name:   "Northern Africa",
		Region: region.Africa,
	},
	AustraliaAndNewZealand: SubRegion{
		Code:   AustraliaAndNewZealand,
		Name:   "Australia and New Zealand",
		Region: region.Oceania,
	},
	Polynesia: SubRegion{
		Code:   Polynesia,
		Name:   "Polynesia",
		Region: region.Oceania,
	},
	Melanesia: SubRegion{
		Code:   Melanesia,
		Name:   "Melanesia",
		Region: region.Oceania,
	},
	Micronesia: SubRegion{
		Code:   Micronesia,
		Name:   "Micronesia",
		Region: region.Oceania,
	},
	SouthernAsia: SubRegion{
		Code:   SouthernAsia,
		Name:   "Southern Asia",
		Region: region.Asia,
	},
	WesternAsia: SubRegion{
		Code:   WesternAsia,
		Name:   "Western Asia",
		Region: region.Asia,
	},
	SouthEasternAsia: SubRegion{
		Code:   SouthEasternAsia,
		Name:   "South-eastern Asia",
		Region: region.Asia,
	},
	EasternAsia: SubRegion{
		Code:   EasternAsia,
		Name:   "Eastern Asia",
		Region: region.Asia,
	},
	CentralAsia: SubRegion{
		Code:   CentralAsia,
		Name:   "Central Asia",
		Region: region.Asia,
	},
	NorthernEurope: SubRegion{
		Code:   NorthernEurope,
		Name:   "Northern Europe",
		Region: region.Europe,
	},
	SouthernEurope: SubRegion{
		Code:   SouthernEurope,
		Name:   "Southern Europe",
		Region: region.Europe,
	},
	WesternEurope: SubRegion{
		Code:   WesternEurope,
		Name:   "Western Europe",
		Region: region.Europe,
	},
	EasternEurope: SubRegion{
		Code:   EasternEurope,
		Name:   "Eastern Europe",
		Region: region.Europe,
	},
}

func (r Code) SubRegion() SubRegion {
	return subRegions[r]
}

func List() SubRegions {
	return subRegions
}
