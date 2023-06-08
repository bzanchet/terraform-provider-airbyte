// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceOpenweatherUpdateLanguage - You can use lang parameter to get the output in your language. The contents of the description field will be translated. See <a href="https://openweathermap.org/api/one-call-api#multi">here</a> for the list of supported languages.
type SourceOpenweatherUpdateLanguage string

const (
	SourceOpenweatherUpdateLanguageAf   SourceOpenweatherUpdateLanguage = "af"
	SourceOpenweatherUpdateLanguageAl   SourceOpenweatherUpdateLanguage = "al"
	SourceOpenweatherUpdateLanguageAr   SourceOpenweatherUpdateLanguage = "ar"
	SourceOpenweatherUpdateLanguageAz   SourceOpenweatherUpdateLanguage = "az"
	SourceOpenweatherUpdateLanguageBg   SourceOpenweatherUpdateLanguage = "bg"
	SourceOpenweatherUpdateLanguageCa   SourceOpenweatherUpdateLanguage = "ca"
	SourceOpenweatherUpdateLanguageCz   SourceOpenweatherUpdateLanguage = "cz"
	SourceOpenweatherUpdateLanguageDa   SourceOpenweatherUpdateLanguage = "da"
	SourceOpenweatherUpdateLanguageDe   SourceOpenweatherUpdateLanguage = "de"
	SourceOpenweatherUpdateLanguageEl   SourceOpenweatherUpdateLanguage = "el"
	SourceOpenweatherUpdateLanguageEn   SourceOpenweatherUpdateLanguage = "en"
	SourceOpenweatherUpdateLanguageEu   SourceOpenweatherUpdateLanguage = "eu"
	SourceOpenweatherUpdateLanguageFa   SourceOpenweatherUpdateLanguage = "fa"
	SourceOpenweatherUpdateLanguageFi   SourceOpenweatherUpdateLanguage = "fi"
	SourceOpenweatherUpdateLanguageFr   SourceOpenweatherUpdateLanguage = "fr"
	SourceOpenweatherUpdateLanguageGl   SourceOpenweatherUpdateLanguage = "gl"
	SourceOpenweatherUpdateLanguageHe   SourceOpenweatherUpdateLanguage = "he"
	SourceOpenweatherUpdateLanguageHi   SourceOpenweatherUpdateLanguage = "hi"
	SourceOpenweatherUpdateLanguageHr   SourceOpenweatherUpdateLanguage = "hr"
	SourceOpenweatherUpdateLanguageHu   SourceOpenweatherUpdateLanguage = "hu"
	SourceOpenweatherUpdateLanguageID   SourceOpenweatherUpdateLanguage = "id"
	SourceOpenweatherUpdateLanguageIt   SourceOpenweatherUpdateLanguage = "it"
	SourceOpenweatherUpdateLanguageJa   SourceOpenweatherUpdateLanguage = "ja"
	SourceOpenweatherUpdateLanguageKr   SourceOpenweatherUpdateLanguage = "kr"
	SourceOpenweatherUpdateLanguageLa   SourceOpenweatherUpdateLanguage = "la"
	SourceOpenweatherUpdateLanguageLt   SourceOpenweatherUpdateLanguage = "lt"
	SourceOpenweatherUpdateLanguageMk   SourceOpenweatherUpdateLanguage = "mk"
	SourceOpenweatherUpdateLanguageNo   SourceOpenweatherUpdateLanguage = "no"
	SourceOpenweatherUpdateLanguageNl   SourceOpenweatherUpdateLanguage = "nl"
	SourceOpenweatherUpdateLanguagePl   SourceOpenweatherUpdateLanguage = "pl"
	SourceOpenweatherUpdateLanguagePt   SourceOpenweatherUpdateLanguage = "pt"
	SourceOpenweatherUpdateLanguagePtBr SourceOpenweatherUpdateLanguage = "pt_br"
	SourceOpenweatherUpdateLanguageRo   SourceOpenweatherUpdateLanguage = "ro"
	SourceOpenweatherUpdateLanguageRu   SourceOpenweatherUpdateLanguage = "ru"
	SourceOpenweatherUpdateLanguageSv   SourceOpenweatherUpdateLanguage = "sv"
	SourceOpenweatherUpdateLanguageSe   SourceOpenweatherUpdateLanguage = "se"
	SourceOpenweatherUpdateLanguageSk   SourceOpenweatherUpdateLanguage = "sk"
	SourceOpenweatherUpdateLanguageSl   SourceOpenweatherUpdateLanguage = "sl"
	SourceOpenweatherUpdateLanguageSp   SourceOpenweatherUpdateLanguage = "sp"
	SourceOpenweatherUpdateLanguageEs   SourceOpenweatherUpdateLanguage = "es"
	SourceOpenweatherUpdateLanguageSr   SourceOpenweatherUpdateLanguage = "sr"
	SourceOpenweatherUpdateLanguageTh   SourceOpenweatherUpdateLanguage = "th"
	SourceOpenweatherUpdateLanguageTr   SourceOpenweatherUpdateLanguage = "tr"
	SourceOpenweatherUpdateLanguageUa   SourceOpenweatherUpdateLanguage = "ua"
	SourceOpenweatherUpdateLanguageUk   SourceOpenweatherUpdateLanguage = "uk"
	SourceOpenweatherUpdateLanguageVi   SourceOpenweatherUpdateLanguage = "vi"
	SourceOpenweatherUpdateLanguageZhCn SourceOpenweatherUpdateLanguage = "zh_cn"
	SourceOpenweatherUpdateLanguageZhTw SourceOpenweatherUpdateLanguage = "zh_tw"
	SourceOpenweatherUpdateLanguageZu   SourceOpenweatherUpdateLanguage = "zu"
)

func (e SourceOpenweatherUpdateLanguage) ToPointer() *SourceOpenweatherUpdateLanguage {
	return &e
}

func (e *SourceOpenweatherUpdateLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "af":
		fallthrough
	case "al":
		fallthrough
	case "ar":
		fallthrough
	case "az":
		fallthrough
	case "bg":
		fallthrough
	case "ca":
		fallthrough
	case "cz":
		fallthrough
	case "da":
		fallthrough
	case "de":
		fallthrough
	case "el":
		fallthrough
	case "en":
		fallthrough
	case "eu":
		fallthrough
	case "fa":
		fallthrough
	case "fi":
		fallthrough
	case "fr":
		fallthrough
	case "gl":
		fallthrough
	case "he":
		fallthrough
	case "hi":
		fallthrough
	case "hr":
		fallthrough
	case "hu":
		fallthrough
	case "id":
		fallthrough
	case "it":
		fallthrough
	case "ja":
		fallthrough
	case "kr":
		fallthrough
	case "la":
		fallthrough
	case "lt":
		fallthrough
	case "mk":
		fallthrough
	case "no":
		fallthrough
	case "nl":
		fallthrough
	case "pl":
		fallthrough
	case "pt":
		fallthrough
	case "pt_br":
		fallthrough
	case "ro":
		fallthrough
	case "ru":
		fallthrough
	case "sv":
		fallthrough
	case "se":
		fallthrough
	case "sk":
		fallthrough
	case "sl":
		fallthrough
	case "sp":
		fallthrough
	case "es":
		fallthrough
	case "sr":
		fallthrough
	case "th":
		fallthrough
	case "tr":
		fallthrough
	case "ua":
		fallthrough
	case "uk":
		fallthrough
	case "vi":
		fallthrough
	case "zh_cn":
		fallthrough
	case "zh_tw":
		fallthrough
	case "zu":
		*e = SourceOpenweatherUpdateLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOpenweatherUpdateLanguage: %v", v)
	}
}

// SourceOpenweatherUpdateUnits - Units of measurement. standard, metric and imperial units are available. If you do not use the units parameter, standard units will be applied by default.
type SourceOpenweatherUpdateUnits string

const (
	SourceOpenweatherUpdateUnitsStandard SourceOpenweatherUpdateUnits = "standard"
	SourceOpenweatherUpdateUnitsMetric   SourceOpenweatherUpdateUnits = "metric"
	SourceOpenweatherUpdateUnitsImperial SourceOpenweatherUpdateUnits = "imperial"
)

func (e SourceOpenweatherUpdateUnits) ToPointer() *SourceOpenweatherUpdateUnits {
	return &e
}

func (e *SourceOpenweatherUpdateUnits) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "metric":
		fallthrough
	case "imperial":
		*e = SourceOpenweatherUpdateUnits(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOpenweatherUpdateUnits: %v", v)
	}
}

type SourceOpenweatherUpdate struct {
	// Your OpenWeather API Key. See <a href="https://openweathermap.org/api">here</a>. The key is case sensitive.
	Appid string `json:"appid"`
	// You can use lang parameter to get the output in your language. The contents of the description field will be translated. See <a href="https://openweathermap.org/api/one-call-api#multi">here</a> for the list of supported languages.
	Lang *SourceOpenweatherUpdateLanguage `json:"lang,omitempty"`
	// Latitude for which you want to get weather condition from. (min -90, max 90)
	Lat string `json:"lat"`
	// Longitude for which you want to get weather condition from. (min -180, max 180)
	Lon string `json:"lon"`
	// Units of measurement. standard, metric and imperial units are available. If you do not use the units parameter, standard units will be applied by default.
	Units *SourceOpenweatherUpdateUnits `json:"units,omitempty"`
}