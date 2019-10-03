package country

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/Pingflow/go-localization/continent"
	"github.com/Pingflow/go-localization/language"
	"github.com/Pingflow/go-localization/region"
	"github.com/Pingflow/go-localization/subregion"
	"github.com/gobuffalo/packr/v2"
)

var box = packr.New("data", "../data")

var (
	ErrCountryNotFound = errors.New("country not found")
	ErrNoCountryRows   = errors.New("no country row")
)

type (
	Country struct {
		OfficialNameAR                string         `json:"official_name_ar"`                        // Country or Area official Arabic short name from UN Statistics Divsion
		OfficialNameCN                string         `json:"official_name_cn"`                        // Country or Area official Chinese short name from UN Statistics Divsion
		OfficialNameEN                string         `json:"official_name_en"`                        // Country or Area official English short name from UN Statistics Divsion
		OfficialNameES                string         `json:"official_name_es"`                        // Country or Area official Spanish short name from UN Statistics Divsion
		OfficialNameFR                string         `json:"official_name_fr"`                        // Country or Area official French short name from UN Statistics Divsion
		OfficialNameRU                string         `json:"official_name_ru"`                        // Country or Area official Russian short name from UN Statistics Divsion
		ISO31661Alpha2                string         `json:"ISO3166-1-Alpha-2"`                       // Alpha-2 codes from ISO 3166-1
		ISO31661Alpha3                string         `json:"ISO3166-1-Alpha-3"`                       // Alpha-3 codes from ISO 3166-1 (synonymous with World Bank Codes)
		ISO31661Numeric               string         `json:"ISO3166-1-numeric"`                       // Numeric codes from ISO 3166-1
		ISO4217CurrencyAlphabeticCode string         `json:"ISO4217-currency_alphabetic_code"`        // ISO 4217 currency alphabetic code
		ISO4217CurrencyCountryName    string         `json:"ISO4217-currency_country_name"`           // ISO 4217 country name
		ISO4217CurrencyMinorUnit      string         `json:"ISO4217-currency_minor_unit"`             // ISO 4217 currency number of minor units
		ISO4217CurrencyName           string         `json:"ISO4217-currency_name"`                   // ISO 4217 currency name
		ISO4217CurrencyNumericCode    string         `json:"ISO4217-currency_numeric_code"`           // ISO 4217 currency numeric code
		M49                           float64        `json:"M49"`                                     // UN Statistics M49 numeric codes (nearly synonymous with ISO 3166-1 numeric codes, which are based on UN M49. ISO 3166-1 does not include Channel Islands or Sark, for example)
		UNTERMArabicFormal            string         `json:"UNTERM Arabic Formal"`                    // Country's formal Arabic name from UN Protocol and Liaison Service
		UNTERMArabicShort             string         `json:"UNTERM Arabic Short"`                     // Country's short Arabic name from UN Protocol and Liaison Service
		UNTERMChineseFormal           string         `json:"UNTERM Chinese Formal"`                   // Country's formal Chinese name from UN Protocol and Liaison Service
		UNTERMChineseShort            string         `json:"UNTERM Chinese Short"`                    // Country's short Chinese name from UN Protocol and Liaison Service
		UNTERMEnglishFormal           string         `json:"UNTERM English Formal"`                   // Country's formal English name from UN Protocol and Liaison Service
		UNTERMEnglishShort            string         `json:"UNTERM English Short"`                    // Country's short English name from UN Protocol and Liaison Service
		UNTERMFrenchFormal            string         `json:"UNTERM French Formal"`                    // Country's formal French name from UN Protocol and Liaison Service
		UNTERMFrenchShort             string         `json:"UNTERM French Short"`                     // Country's short French name from UN Protocol and Liaison Service
		UNTERMRussianFormal           string         `json:"UNTERM Russian Formal"`                   // Country's formal Russian name from UN Protocol and Liaison Service
		UNTERMRussianShort            string         `json:"UNTERM Russian Short"`                    // Country's short Russian name from UN Protocol and Liaison Service
		UNTERMSpanishFormal           string         `json:"UNTERM Spanish Formal"`                   // Country's formal Spanish name from UN Protocol and Liaison Service
		UNTERMSpanishShort            string         `json:"UNTERM Spanish Short"`                    // Country's short Spanish name from UN Protocol and Liaison Service
		CLDRDisplayName               string         `json:"CLDR display name"`                       // Country's customary English short name (CLDR)
		Capital                       string         `json:"Capital"`                                 // Capital city from Geonames
		Continent                     continent.Code `json:"Continent"`                               // Continent from Geonames
		DS                            string         `json:"DS"`                                      // Distinguishing signs of vehicles in international traffic
		DevelopedDevelopingCountries  string         `json:"Developed / Developing Countries"`        // Country classification from United Nations Statistics Division
		Dial                          string         `json:"Dial"`                                    // Country code from ITU-T recommendation E.164, sometimes followed by area code
		EDGAR                         string         `json:"EDGAR"`                                   // EDGAR country code from SEC
		FIFA                          string         `json:"FIFA"`                                    // Codes assigned by the Fédération Internationale de Football Association
		FIPS                          string         `json:"FIPS"`                                    // Codes from the U.S. standard FIPS PUB 10-4
		GAUL                          string         `json:"GAUL"`                                    // Global Administrative Unit Layers from the Food and Agriculture Organization
		GeonameID                     float64        `json:"Geoname ID"`                              // Geoname ID
		GlobalCode                    string         `json:"Global Code"`                             // Country classification from United Nations Statistics Division
		GlobalName                    string         `json:"Global Name"`                             // Country classification from United Nations Statistics Division
		IOC                           string         `json:"IOC"`                                     // Codes assigned by the International Olympics Committee
		ITU                           string         `json:"ITU"`                                     // Codes assigned by the International Telecommunications Union
		IntermediateRegionCode        string         `json:"Intermediate Region Code"`                // Country classification from United Nations Statistics Division
		IntermediateRegionName        string         `json:"Intermediate Region Name"`                // Country classification from United Nations Statistics Division
		LandLockedDevelopingCountries string         `json:"Land Locked Developing Countries (LLDC)"` // Country classification from United Nations Statistics Division
		Languages                     string         `json:"Languages"`                               // Languages from Geonames
		LeastDevelopedCountries       string         `json:"Least Developed Countries (LDC)"`         // Country classification from United Nations Statistics Division
		MARC                          string         `json:"MARC"`                                    // Machine-Readable Cataloging codes from the Library of Congress
		RegionCode                    region.Code    `json:"Region Code"`                             // Country classification from United Nations Statistics Division
		RegionName                    string         `json:"Region Name"`                             // Country classification from United Nations Statistics Division
		SmallIslandDevelopingStates   string         `json:"Small Island Developing States (SIDS)"`   // Country classification from United Nations Statistics Division
		SubRegionCode                 subregion.Code `json:"Sub-region Code"`                         // Country classification from United Nations Statistics Division
		SubRegionName                 string         `json:"Sub-region Name"`                         // Country classification from United Nations Statistics Division
		TLD                           string         `json:"TLD"`                                     // Top level domain from Geonames
		WMO                           string         `json:"WMO"`                                     // Country abbreviations by the World Meteorological Organization
		IsIndependent                 string         `json:"is_independent"`                          // Country status, based on the CIA World Factbook
	}

	Countries []Country
)

func List() (Countries, error) {
	var c Countries

	j, e := box.Find("country-codes.json")
	if e != nil {
		return nil, e
	}

	if e := json.Unmarshal(j, &c); e != nil {
		return nil, e
	}

	if len(c) == 0 {
		return nil, ErrNoCountryRows
	}

	return c, nil
}

func ListByContinent(query continent.Code) (Countries, error) {

	var c Countries

	l, e := List()
	if e != nil {
		return nil, e
	}

	for _, v := range l {
		if v.Continent == query {
			c = append(c, v)
		}
	}

	if len(c) == 0 {
		return nil, ErrNoCountryRows
	}

	return c, nil
}

func ListByRegion(query region.Code) (Countries, error) {

	var c Countries

	l, e := List()
	if e != nil {
		return nil, e
	}

	for _, v := range l {
		if v.RegionCode == query {
			c = append(c, v)
		}
	}

	if len(c) == 0 {
		return nil, ErrNoCountryRows
	}

	return c, nil
}

func ListBySubRegion(query subregion.Code) (Countries, error) {

	var c Countries

	l, e := List()
	if e != nil {
		return nil, e
	}

	for _, v := range l {
		if v.SubRegionCode == query {
			c = append(c, v)
		}
	}

	if len(c) == 0 {
		return nil, ErrNoCountryRows
	}

	return c, nil
}

func Find(query string) (*Country, error) {

	l, e := List()
	if e != nil {
		return nil, e
	}

	for _, c := range l {
		if c.ISO31661Alpha3 == query ||
			c.ISO31661Alpha2 == query ||
			c.ISO31661Numeric == query ||
			c.CLDRDisplayName == query ||
			c.OfficialNameAR == query ||
			c.OfficialNameCN == query ||
			c.OfficialNameEN == query ||
			c.OfficialNameES == query ||
			c.OfficialNameFR == query ||
			c.OfficialNameRU == query {
			return &c, nil
		}
	}

	return nil, ErrCountryNotFound
}

func (c Country) LanguagesList() language.Languages {

	var lgs language.Languages

	r, _ := regexp.Compile(`^(^[a-zA-Z]{2})(-)?([a-zA-Z]{2})?$`)

	for _, v := range strings.Split(c.Languages, ",") {
		buf := r.FindStringSubmatch(v)
		if len(buf) > 2 {
			l, e := language.Find(buf[1])
			if e == nil {
				lgs = append(lgs, *l)
			}
		}
	}

	return lgs
}

func (c Country) MainLanguage() language.Language {
	return c.LanguagesList()[0]
}
