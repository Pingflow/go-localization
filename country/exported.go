package country

type (
	Alpha3 string
	Alpha3Map map[Alpha3]string
)

const (
	Andorra                              Alpha3 = "AND"
	UnitedArabEmirates                   Alpha3 = "ARE"
	Afghanistan                          Alpha3 = "AFG"
	AntiguaAndBarbuda                    Alpha3 = "ATG"
	Anguilla                             Alpha3 = "AIA"
	Albania                              Alpha3 = "ALB"
	Armenia                              Alpha3 = "ARM"
	Angola                               Alpha3 = "AGO"
	Argentina                            Alpha3 = "ARG"
	AmericanSamoa                        Alpha3 = "ASM"
	Austria                              Alpha3 = "AUT"
	Australia                            Alpha3 = "AUS"
	Aruba                                Alpha3 = "ABW"
	AlandIslands                         Alpha3 = "ALA"
	Azerbaijan                           Alpha3 = "AZE"
	BosniaAndHerzegovina                 Alpha3 = "BIH"
	Barbados                             Alpha3 = "BRB"
	Bangladesh                           Alpha3 = "BGD"
	Belgium                              Alpha3 = "BEL"
	BurkinaFaso                          Alpha3 = "BFA"
	Bulgaria                             Alpha3 = "BGR"
	Bahrain                              Alpha3 = "BHR"
	Burundi                              Alpha3 = "BDI"
	Benin                                Alpha3 = "BEN"
	SaintBarthelemy                      Alpha3 = "BLM"
	Bermuda                              Alpha3 = "BMU"
	Brunei                               Alpha3 = "BRN"
	Bolivia                              Alpha3 = "BOL"
	Brazil                               Alpha3 = "BRA"
	Bahamas                              Alpha3 = "BHS"
	Bhutan                               Alpha3 = "BTN"
	BouvetIsland                         Alpha3 = "BVT"
	Botswana                             Alpha3 = "BWA"
	Belarus                              Alpha3 = "BLR"
	Belize                               Alpha3 = "BLZ"
	Canada                               Alpha3 = "CAN"
	CocosIslands                         Alpha3 = "CCK"
	DemocraticRepublicOfTheCongo         Alpha3 = "COD"
	CentralAfricanRepublic               Alpha3 = "CAF"
	Congo                                Alpha3 = "COG"
	Switzerland                          Alpha3 = "CHE"
	CoteDIvoire                          Alpha3 = "CIV"
	CookIslands                          Alpha3 = "COK"
	Chile                                Alpha3 = "CHL"
	Cameroon                             Alpha3 = "CMR"
	China                                Alpha3 = "CHN"
	Colombia                             Alpha3 = "COL"
	CostaRica                            Alpha3 = "CRI"
	Cuba                                 Alpha3 = "CUB"
	CaboVerde                            Alpha3 = "CPV"
	Curacao                              Alpha3 = "CUW"
	ChristmasIsland                      Alpha3 = "CXR"
	Cyprus                               Alpha3 = "CYP"
	Czechia                              Alpha3 = "CZE"
	Germany                              Alpha3 = "DEU"
	Djibouti                             Alpha3 = "DJI"
	Denmark                              Alpha3 = "DNK"
	Dominica                             Alpha3 = "DMA"
	DominicanRepublic                    Alpha3 = "DOM"
	Algeria                              Alpha3 = "DZA"
	Ecuador                              Alpha3 = "ECU"
	Estonia                              Alpha3 = "EST"
	Egypt                                Alpha3 = "EGY"
	WesternSahara                        Alpha3 = "ESH"
	Eritrea                              Alpha3 = "ERI"
	Spain                                Alpha3 = "ESP"
	Ethiopia                             Alpha3 = "ETH"
	Finland                              Alpha3 = "FIN"
	Fiji                                 Alpha3 = "FJI"
	FalklandIslands                      Alpha3 = "FLK"
	Micronesia                           Alpha3 = "FSM"
	FaroeIslands                         Alpha3 = "FRO"
	France                               Alpha3 = "FRA"
	Gabon                                Alpha3 = "GAB"
	UnitedKingdom                        Alpha3 = "GBR"
	Grenada                              Alpha3 = "GRD"
	Georgia                              Alpha3 = "GEO"
	FrenchGuiana                         Alpha3 = "GUF"
	Guernsey                             Alpha3 = "GGY"
	Ghana                                Alpha3 = "GHA"
	Gibraltar                            Alpha3 = "GIB"
	Greenland                            Alpha3 = "GRL"
	Gambia                               Alpha3 = "GMB"
	Guinea                               Alpha3 = "GIN"
	Guadeloupe                           Alpha3 = "GLP"
	EquatorialGuinea                     Alpha3 = "GNQ"
	Greece                               Alpha3 = "GRC"
	SouthGeorgia                         Alpha3 = "SGS"
	Guatemala                            Alpha3 = "GTM"
	Guam                                 Alpha3 = "GUM"
	GuineaBissau                         Alpha3 = "GNB"
	Guyana                               Alpha3 = "GUY"
	HongKong                             Alpha3 = "HKG"
	HeardIsland                          Alpha3 = "HMD"
	Honduras                             Alpha3 = "HND"
	Croatia                              Alpha3 = "HRV"
	Haiti                                Alpha3 = "HTI"
	Hungary                              Alpha3 = "HUN"
	Indonesia                            Alpha3 = "IDN"
	Ireland                              Alpha3 = "IRL"
	Israel                               Alpha3 = "ISR"
	IsleOfMan                            Alpha3 = "IMN"
	India                                Alpha3 = "IND"
	BritishIndianOceanTerritory          Alpha3 = "IOT"
	Iraq                                 Alpha3 = "IRQ"
	Iran                                 Alpha3 = "IRN"
	Iceland                              Alpha3 = "ISL"
	Italy                                Alpha3 = "ITA"
	Jersey                               Alpha3 = "JEY"
	Jamaica                              Alpha3 = "JAM"
	Jordan                               Alpha3 = "JOR"
	Japan                                Alpha3 = "JPN"
	Kenya                                Alpha3 = "KEN"
	Kyrgyzstan                           Alpha3 = "KGZ"
	Cambodia                             Alpha3 = "KHM"
	Kiribati                             Alpha3 = "KIR"
	Comoros                              Alpha3 = "COM"
	SaintKittsAndNevis                   Alpha3 = "KNA"
	DemocraticPeoplesRepublicOfKorea     Alpha3 = "PRK"
	RepublicOfKorea                      Alpha3 = "KOR"
	Kuwait                               Alpha3 = "KWT"
	CaymanIslands                        Alpha3 = "CYM"
	Kazakhstan                           Alpha3 = "KAZ"
	LaoPeoplesDemocraticRepublic         Alpha3 = "LAO"
	Lebanon                              Alpha3 = "LBN"
	SaintLucia                           Alpha3 = "LCA"
	Liechtenstein                        Alpha3 = "LIE"
	SriLanka                             Alpha3 = "LKA"
	Liberia                              Alpha3 = "LBR"
	Lesotho                              Alpha3 = "LSO"
	Lithuania                            Alpha3 = "LTU"
	Luxembourg                           Alpha3 = "LUX"
	Latvia                               Alpha3 = "LVA"
	Libya                                Alpha3 = "LBY"
	Morocco                              Alpha3 = "MAR"
	Monaco                               Alpha3 = "MCO"
	RepublicOfMoldova                    Alpha3 = "MDA"
	Montenegro                           Alpha3 = "MNE"
	SaintMartinFrench                    Alpha3 = "MAF"
	Madagascar                           Alpha3 = "MDG"
	MarshallIslands                      Alpha3 = "MHL"
	TheFormerYugoslavRepublicOfMacedonia Alpha3 = "MKD"
	Mali                                 Alpha3 = "MLI"
	Myanmar                              Alpha3 = "MMR"
	Mongolia                             Alpha3 = "MNG"
	Macao                                Alpha3 = "MAC"
	NorthernMarianaIslands               Alpha3 = "MNP"
	Martinique                           Alpha3 = "MTQ"
	Mauritania                           Alpha3 = "MRT"
	Montserrat                           Alpha3 = "MSR"
	Malta                                Alpha3 = "MLT"
	Mauritius                            Alpha3 = "MUS"
	Maldives                             Alpha3 = "MDV"
	Malawi                               Alpha3 = "MWI"
	Mexico                               Alpha3 = "MEX"
	Malaysia                             Alpha3 = "MYS"
	Mozambique                           Alpha3 = "MOZ"
	Namibia                              Alpha3 = "NAM"
	NewCaledonia                         Alpha3 = "NCL"
	Niger                                Alpha3 = "NER"
	NorfolkIsland                        Alpha3 = "NFK"
	Nigeria                              Alpha3 = "NGA"
	Nicaragua                            Alpha3 = "NIC"
	Netherlands                          Alpha3 = "NLD"
	Norway                               Alpha3 = "NOR"
	Nepal                                Alpha3 = "NPL"
	Nauru                                Alpha3 = "NRU"
	Niue                                 Alpha3 = "NIU"
	NewZealand                           Alpha3 = "NZL"
	Oman                                 Alpha3 = "OMN"
	Panama                               Alpha3 = "PAN"
	Peru                                 Alpha3 = "PER"
	FrenchPolynesia                      Alpha3 = "PYF"
	PapuaNewGuinea                       Alpha3 = "PNG"
	Philippines                          Alpha3 = "PHL"
	Pakistan                             Alpha3 = "PAK"
	Poland                               Alpha3 = "POL"
	SaintPierreAndMiquelon               Alpha3 = "SPM"
	Pitcairn                             Alpha3 = "PCN"
	PuertoRico                           Alpha3 = "PRI"
	StateOfPalestine                     Alpha3 = "PSE"
	Portugal                             Alpha3 = "PRT"
	Palau                                Alpha3 = "PLW"
	Paraguay                             Alpha3 = "PRY"
	Qatar                                Alpha3 = "QAT"
	Reunion                              Alpha3 = "REU"
	Romania                              Alpha3 = "ROU"
	Serbia                               Alpha3 = "SRB"
	RussianFederation                    Alpha3 = "RUS"
	Rwanda                               Alpha3 = "RWA"
	SaudiArabia                          Alpha3 = "SAU"
	SolomonIslands                       Alpha3 = "SLB"
	Seychelles                           Alpha3 = "SYC"
	Sudan                                Alpha3 = "SDN"
	Sweden                               Alpha3 = "SWE"
	Singapore                            Alpha3 = "SGP"
	SaintHelena                          Alpha3 = "SHN"
	Slovenia                             Alpha3 = "SVN"
	SvalbardAndJanMayenIslands           Alpha3 = "SJM"
	Slovakia                             Alpha3 = "SVK"
	SierraLeone                          Alpha3 = "SLE"
	SanMarino                            Alpha3 = "SMR"
	Senegal                              Alpha3 = "SEN"
	Somalia                              Alpha3 = "SOM"
	Suriname                             Alpha3 = "SUR"
	SouthSudan                           Alpha3 = "SSD"
	SaoTomeAndPrincipe                   Alpha3 = "STP"
	ElSalvador                           Alpha3 = "SLV"
	SintMaartenDutchPart                 Alpha3 = "SXM"
	SyrianArabRepublic                   Alpha3 = "SYR"
	Swaziland                            Alpha3 = "SWZ"
	TurksAndCaicosIslands                Alpha3 = "TCA"
	Chad                                 Alpha3 = "TCD"
	FrenchSouthernTerritories            Alpha3 = "ATF"
	Togo                                 Alpha3 = "TGO"
	Thailand                             Alpha3 = "THA"
	Tajikistan                           Alpha3 = "TJK"
	Tokelau                              Alpha3 = "TKL"
	TimorLeste                           Alpha3 = "TLS"
	Turkmenistan                         Alpha3 = "TKM"
	Tunisia                              Alpha3 = "TUN"
	Tonga                                Alpha3 = "TON"
	Turkey                               Alpha3 = "TUR"
	TrinidadAndTobago                    Alpha3 = "TTO"
	Tuvalu                               Alpha3 = "TUV"
	Taiwan                               Alpha3 = "TWN"
	UnitedRepublicOfTanzania             Alpha3 = "TZA"
	Ukraine                              Alpha3 = "UKR"
	Uganda                               Alpha3 = "UGA"
	UnitedStatesMinorOutlyingIslands     Alpha3 = "UMI"
	UnitedStatesOfAmerica                Alpha3 = "USA"
	Uruguay                              Alpha3 = "URY"
	Uzbekistan                           Alpha3 = "UZB"
	HolySee                              Alpha3 = "VAT"
	SaintVincentAndTheGrenadines         Alpha3 = "VCT"
	Venezuela                            Alpha3 = "VEN"
	BritishVirginIslands                 Alpha3 = "VGB"
	UnitedStatesVirginIslands            Alpha3 = "VIR"
	VietNam                              Alpha3 = "VNM"
	Vanuatu                              Alpha3 = "VUT"
	WallisAndFutunaIslands               Alpha3 = "WLF"
	Samoa                                Alpha3 = "WSM"
	Yemen                                Alpha3 = "YEM"
	Mayotte                              Alpha3 = "MYT"
	SouthAfrica                          Alpha3 = "ZAF"
	Zambia                               Alpha3 = "ZMB"
	Zimbabwe                             Alpha3 = "ZWE"
)

var alpha3 = Alpha3Map{
	Andorra:                              "AND",
	UnitedArabEmirates:                   "ARE",
	Afghanistan:                          "AFG",
	AntiguaAndBarbuda:                    "ATG",
	Anguilla:                             "AIA",
	Albania:                              "ALB",
	Armenia:                              "ARM",
	Angola:                               "AGO",
	Argentina:                            "ARG",
	AmericanSamoa:                        "ASM",
	Austria:                              "AUT",
	Australia:                            "AUS",
	Aruba:                                "ABW",
	AlandIslands:                         "ALA",
	Azerbaijan:                           "AZE",
	BosniaAndHerzegovina:                 "BIH",
	Barbados:                             "BRB",
	Bangladesh:                           "BGD",
	Belgium:                              "BEL",
	BurkinaFaso:                          "BFA",
	Bulgaria:                             "BGR",
	Bahrain:                              "BHR",
	Burundi:                              "BDI",
	Benin:                                "BEN",
	SaintBarthelemy:                      "BLM",
	Bermuda:                              "BMU",
	Brunei:                               "BRN",
	Bolivia:                              "BOL",
	Brazil:                               "BRA",
	Bahamas:                              "BHS",
	Bhutan:                               "BTN",
	BouvetIsland:                         "BVT",
	Botswana:                             "BWA",
	Belarus:                              "BLR",
	Belize:                               "BLZ",
	Canada:                               "CAN",
	CocosIslands:                         "CCK",
	DemocraticRepublicOfTheCongo:         "COD",
	CentralAfricanRepublic:               "CAF",
	Congo:                                "COG",
	Switzerland:                          "CHE",
	CoteDIvoire:                          "CIV",
	CookIslands:                          "COK",
	Chile:                                "CHL",
	Cameroon:                             "CMR",
	China:                                "CHN",
	Colombia:                             "COL",
	CostaRica:                            "CRI",
	Cuba:                                 "CUB",
	CaboVerde:                            "CPV",
	Curacao:                              "CUW",
	ChristmasIsland:                      "CXR",
	Cyprus:                               "CYP",
	Czechia:                              "CZE",
	Germany:                              "DEU",
	Djibouti:                             "DJI",
	Denmark:                              "DNK",
	Dominica:                             "DMA",
	DominicanRepublic:                    "DOM",
	Algeria:                              "DZA",
	Ecuador:                              "ECU",
	Estonia:                              "EST",
	Egypt:                                "EGY",
	WesternSahara:                        "ESH",
	Eritrea:                              "ERI",
	Spain:                                "ESP",
	Ethiopia:                             "ETH",
	Finland:                              "FIN",
	Fiji:                                 "FJI",
	FalklandIslands:                      "FLK",
	Micronesia:                           "FSM",
	FaroeIslands:                         "FRO",
	France:                               "FRA",
	Gabon:                                "GAB",
	UnitedKingdom:                        "GBR",
	Grenada:                              "GRD",
	Georgia:                              "GEO",
	FrenchGuiana:                         "GUF",
	Guernsey:                             "GGY",
	Ghana:                                "GHA",
	Gibraltar:                            "GIB",
	Greenland:                            "GRL",
	Gambia:                               "GMB",
	Guinea:                               "GIN",
	Guadeloupe:                           "GLP",
	EquatorialGuinea:                     "GNQ",
	Greece:                               "GRC",
	SouthGeorgia:                         "SGS",
	Guatemala:                            "GTM",
	Guam:                                 "GUM",
	GuineaBissau:                         "GNB",
	Guyana:                               "GUY",
	HongKong:                             "HKG",
	HeardIsland:                          "HMD",
	Honduras:                             "HND",
	Croatia:                              "HRV",
	Haiti:                                "HTI",
	Hungary:                              "HUN",
	Indonesia:                            "IDN",
	Ireland:                              "IRL",
	Israel:                               "ISR",
	IsleOfMan:                            "IMN",
	India:                                "IND",
	BritishIndianOceanTerritory:          "IOT",
	Iraq:                                 "IRQ",
	Iran:                                 "IRN",
	Iceland:                              "ISL",
	Italy:                                "ITA",
	Jersey:                               "JEY",
	Jamaica:                              "JAM",
	Jordan:                               "JOR",
	Japan:                                "JPN",
	Kenya:                                "KEN",
	Kyrgyzstan:                           "KGZ",
	Cambodia:                             "KHM",
	Kiribati:                             "KIR",
	Comoros:                              "COM",
	SaintKittsAndNevis:                   "KNA",
	DemocraticPeoplesRepublicOfKorea:     "PRK",
	RepublicOfKorea:                      "KOR",
	Kuwait:                               "KWT",
	CaymanIslands:                        "CYM",
	Kazakhstan:                           "KAZ",
	LaoPeoplesDemocraticRepublic:         "LAO",
	Lebanon:                              "LBN",
	SaintLucia:                           "LCA",
	Liechtenstein:                        "LIE",
	SriLanka:                             "LKA",
	Liberia:                              "LBR",
	Lesotho:                              "LSO",
	Lithuania:                            "LTU",
	Luxembourg:                           "LUX",
	Latvia:                               "LVA",
	Libya:                                "LBY",
	Morocco:                              "MAR",
	Monaco:                               "MCO",
	RepublicOfMoldova:                    "MDA",
	Montenegro:                           "MNE",
	SaintMartinFrench:                    "MAF",
	Madagascar:                           "MDG",
	MarshallIslands:                      "MHL",
	TheFormerYugoslavRepublicOfMacedonia: "MKD",
	Mali:                                 "MLI",
	Myanmar:                              "MMR",
	Mongolia:                             "MNG",
	Macao:                                "MAC",
	NorthernMarianaIslands:               "MNP",
	Martinique:                           "MTQ",
	Mauritania:                           "MRT",
	Montserrat:                           "MSR",
	Malta:                                "MLT",
	Mauritius:                            "MUS",
	Maldives:                             "MDV",
	Malawi:                               "MWI",
	Mexico:                               "MEX",
	Malaysia:                             "MYS",
	Mozambique:                           "MOZ",
	Namibia:                              "NAM",
	NewCaledonia:                         "NCL",
	Niger:                                "NER",
	NorfolkIsland:                        "NFK",
	Nigeria:                              "NGA",
	Nicaragua:                            "NIC",
	Netherlands:                          "NLD",
	Norway:                               "NOR",
	Nepal:                                "NPL",
	Nauru:                                "NRU",
	Niue:                                 "NIU",
	NewZealand:                           "NZL",
	Oman:                                 "OMN",
	Panama:                               "PAN",
	Peru:                                 "PER",
	FrenchPolynesia:                      "PYF",
	PapuaNewGuinea:                       "PNG",
	Philippines:                          "PHL",
	Pakistan:                             "PAK",
	Poland:                               "POL",
	SaintPierreAndMiquelon:               "SPM",
	Pitcairn:                             "PCN",
	PuertoRico:                           "PRI",
	StateOfPalestine:                     "PSE",
	Portugal:                             "PRT",
	Palau:                                "PLW",
	Paraguay:                             "PRY",
	Qatar:                                "QAT",
	Reunion:                              "REU",
	Romania:                              "ROU",
	Serbia:                               "SRB",
	RussianFederation:                    "RUS",
	Rwanda:                               "RWA",
	SaudiArabia:                          "SAU",
	SolomonIslands:                       "SLB",
	Seychelles:                           "SYC",
	Sudan:                                "SDN",
	Sweden:                               "SWE",
	Singapore:                            "SGP",
	SaintHelena:                          "SHN",
	Slovenia:                             "SVN",
	SvalbardAndJanMayenIslands:           "SJM",
	Slovakia:                             "SVK",
	SierraLeone:                          "SLE",
	SanMarino:                            "SMR",
	Senegal:                              "SEN",
	Somalia:                              "SOM",
	Suriname:                             "SUR",
	SouthSudan:                           "SSD",
	SaoTomeAndPrincipe:                   "STP",
	ElSalvador:                           "SLV",
	SintMaartenDutchPart:                 "SXM",
	SyrianArabRepublic:                   "SYR",
	Swaziland:                            "SWZ",
	TurksAndCaicosIslands:                "TCA",
	Chad:                                 "TCD",
	FrenchSouthernTerritories:            "ATF",
	Togo:                                 "TGO",
	Thailand:                             "THA",
	Tajikistan:                           "TJK",
	Tokelau:                              "TKL",
	TimorLeste:                           "TLS",
	Turkmenistan:                         "TKM",
	Tunisia:                              "TUN",
	Tonga:                                "TON",
	Turkey:                               "TUR",
	TrinidadAndTobago:                    "TTO",
	Tuvalu:                               "TUV",
	Taiwan:                               "TWN",
	UnitedRepublicOfTanzania:             "TZA",
	Ukraine:                              "UKR",
	Uganda:                               "UGA",
	UnitedStatesMinorOutlyingIslands:     "UMI",
	UnitedStatesOfAmerica:                "USA",
	Uruguay:                              "URY",
	Uzbekistan:                           "UZB",
	HolySee:                              "VAT",
	SaintVincentAndTheGrenadines:         "VCT",
	Venezuela:                            "VEN",
	BritishVirginIslands:                 "VGB",
	UnitedStatesVirginIslands:            "VIR",
	VietNam:                              "VNM",
	Vanuatu:                              "VUT",
	WallisAndFutunaIslands:               "WLF",
	Samoa:                                "WSM",
	Yemen:                                "YEM",
	Mayotte:                              "MYT",
	SouthAfrica:                          "ZAF",
	Zambia:                               "ZMB",
	Zimbabwe:                             "ZWE",
}

func (alpha3 Alpha3) String() string {
	return string(alpha3)
}

func (alpha3 Alpha3) Country() Country {
	c, _ := Find(alpha3.String())
	return *c
}

func HasCountry(value string) bool {
	if _, ok := alpha3[Alpha3(value)]; ok {
		return true
	}

	return false

}
