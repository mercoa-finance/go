// This file was auto-generated by Fern from our API Definition.

package mercoa

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mercoa-finance/go/internal"
)

type Address struct {
	AddressLine1 string  `json:"addressLine1" url:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty" url:"addressLine2,omitempty"`
	City         string  `json:"city" url:"city"`
	// State or province code. Must be in the format XX.
	StateOrProvince string `json:"stateOrProvince" url:"stateOrProvince"`
	// Postal code. Must be in the format XXXXX or XXXXX-XXXX.
	PostalCode string  `json:"postalCode" url:"postalCode"`
	Country    *string `json:"country,omitempty" url:"country,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *Address) GetAddressLine1() string {
	if a == nil {
		return ""
	}
	return a.AddressLine1
}

func (a *Address) GetAddressLine2() *string {
	if a == nil {
		return nil
	}
	return a.AddressLine2
}

func (a *Address) GetCity() string {
	if a == nil {
		return ""
	}
	return a.City
}

func (a *Address) GetStateOrProvince() string {
	if a == nil {
		return ""
	}
	return a.StateOrProvince
}

func (a *Address) GetPostalCode() string {
	if a == nil {
		return ""
	}
	return a.PostalCode
}

func (a *Address) GetCountry() *string {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *Address) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *Address) UnmarshalJSON(data []byte) error {
	type unmarshaler Address
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Address(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *Address) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type BirthDate struct {
	Day   *string `json:"day,omitempty" url:"day,omitempty"`
	Month *string `json:"month,omitempty" url:"month,omitempty"`
	Year  *string `json:"year,omitempty" url:"year,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BirthDate) GetDay() *string {
	if b == nil {
		return nil
	}
	return b.Day
}

func (b *BirthDate) GetMonth() *string {
	if b == nil {
		return nil
	}
	return b.Month
}

func (b *BirthDate) GetYear() *string {
	if b == nil {
		return nil
	}
	return b.Year
}

func (b *BirthDate) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BirthDate) UnmarshalJSON(data []byte) error {
	type unmarshaler BirthDate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BirthDate(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BirthDate) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type BulkDownloadFormat string

const (
	BulkDownloadFormatCsv  BulkDownloadFormat = "CSV"
	BulkDownloadFormatJSON BulkDownloadFormat = "JSON"
)

func NewBulkDownloadFormatFromString(s string) (BulkDownloadFormat, error) {
	switch s {
	case "CSV":
		return BulkDownloadFormatCsv, nil
	case "JSON":
		return BulkDownloadFormatJSON, nil
	}
	var t BulkDownloadFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BulkDownloadFormat) Ptr() *BulkDownloadFormat {
	return &b
}

type BulkDownloadResponse struct {
	URL      string `json:"url" url:"url"`
	MimeType string `json:"mimeType" url:"mimeType"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkDownloadResponse) GetURL() string {
	if b == nil {
		return ""
	}
	return b.URL
}

func (b *BulkDownloadResponse) GetMimeType() string {
	if b == nil {
		return ""
	}
	return b.MimeType
}

func (b *BulkDownloadResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkDownloadResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkDownloadResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkDownloadResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkDownloadResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type CountryCode string

const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCv CountryCode = "CV"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeCi CountryCode = "CI"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeID CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeFm CountryCode = "FM"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeMp CountryCode = "MP"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePw CountryCode = "PW"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodePr CountryCode = "PR"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeRe CountryCode = "RE"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeVc CountryCode = "VC"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUs CountryCode = "US"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
	CountryCodeAx CountryCode = "AX"
)

func NewCountryCodeFromString(s string) (CountryCode, error) {
	switch s {
	case "AF":
		return CountryCodeAf, nil
	case "AL":
		return CountryCodeAl, nil
	case "DZ":
		return CountryCodeDz, nil
	case "AS":
		return CountryCodeAs, nil
	case "AD":
		return CountryCodeAd, nil
	case "AO":
		return CountryCodeAo, nil
	case "AI":
		return CountryCodeAi, nil
	case "AQ":
		return CountryCodeAq, nil
	case "AG":
		return CountryCodeAg, nil
	case "AR":
		return CountryCodeAr, nil
	case "AM":
		return CountryCodeAm, nil
	case "AW":
		return CountryCodeAw, nil
	case "AU":
		return CountryCodeAu, nil
	case "AT":
		return CountryCodeAt, nil
	case "AZ":
		return CountryCodeAz, nil
	case "BS":
		return CountryCodeBs, nil
	case "BH":
		return CountryCodeBh, nil
	case "BD":
		return CountryCodeBd, nil
	case "BB":
		return CountryCodeBb, nil
	case "BY":
		return CountryCodeBy, nil
	case "BE":
		return CountryCodeBe, nil
	case "BZ":
		return CountryCodeBz, nil
	case "BJ":
		return CountryCodeBj, nil
	case "BM":
		return CountryCodeBm, nil
	case "BT":
		return CountryCodeBt, nil
	case "BO":
		return CountryCodeBo, nil
	case "BQ":
		return CountryCodeBq, nil
	case "BA":
		return CountryCodeBa, nil
	case "BW":
		return CountryCodeBw, nil
	case "BV":
		return CountryCodeBv, nil
	case "BR":
		return CountryCodeBr, nil
	case "IO":
		return CountryCodeIo, nil
	case "BN":
		return CountryCodeBn, nil
	case "BG":
		return CountryCodeBg, nil
	case "BF":
		return CountryCodeBf, nil
	case "BI":
		return CountryCodeBi, nil
	case "KH":
		return CountryCodeKh, nil
	case "CM":
		return CountryCodeCm, nil
	case "CA":
		return CountryCodeCa, nil
	case "CV":
		return CountryCodeCv, nil
	case "KY":
		return CountryCodeKy, nil
	case "CF":
		return CountryCodeCf, nil
	case "TD":
		return CountryCodeTd, nil
	case "CL":
		return CountryCodeCl, nil
	case "CN":
		return CountryCodeCn, nil
	case "CX":
		return CountryCodeCx, nil
	case "CC":
		return CountryCodeCc, nil
	case "CO":
		return CountryCodeCo, nil
	case "KM":
		return CountryCodeKm, nil
	case "CG":
		return CountryCodeCg, nil
	case "CD":
		return CountryCodeCd, nil
	case "CK":
		return CountryCodeCk, nil
	case "CR":
		return CountryCodeCr, nil
	case "HR":
		return CountryCodeHr, nil
	case "CU":
		return CountryCodeCu, nil
	case "CW":
		return CountryCodeCw, nil
	case "CY":
		return CountryCodeCy, nil
	case "CZ":
		return CountryCodeCz, nil
	case "CI":
		return CountryCodeCi, nil
	case "DK":
		return CountryCodeDk, nil
	case "DJ":
		return CountryCodeDj, nil
	case "DM":
		return CountryCodeDm, nil
	case "DO":
		return CountryCodeDo, nil
	case "EC":
		return CountryCodeEc, nil
	case "EG":
		return CountryCodeEg, nil
	case "SV":
		return CountryCodeSv, nil
	case "GQ":
		return CountryCodeGq, nil
	case "ER":
		return CountryCodeEr, nil
	case "EE":
		return CountryCodeEe, nil
	case "SZ":
		return CountryCodeSz, nil
	case "ET":
		return CountryCodeEt, nil
	case "FK":
		return CountryCodeFk, nil
	case "FO":
		return CountryCodeFo, nil
	case "FJ":
		return CountryCodeFj, nil
	case "FI":
		return CountryCodeFi, nil
	case "FR":
		return CountryCodeFr, nil
	case "GF":
		return CountryCodeGf, nil
	case "PF":
		return CountryCodePf, nil
	case "TF":
		return CountryCodeTf, nil
	case "GA":
		return CountryCodeGa, nil
	case "GM":
		return CountryCodeGm, nil
	case "GE":
		return CountryCodeGe, nil
	case "DE":
		return CountryCodeDe, nil
	case "GH":
		return CountryCodeGh, nil
	case "GI":
		return CountryCodeGi, nil
	case "GR":
		return CountryCodeGr, nil
	case "GL":
		return CountryCodeGl, nil
	case "GD":
		return CountryCodeGd, nil
	case "GP":
		return CountryCodeGp, nil
	case "GU":
		return CountryCodeGu, nil
	case "GT":
		return CountryCodeGt, nil
	case "GG":
		return CountryCodeGg, nil
	case "GN":
		return CountryCodeGn, nil
	case "GW":
		return CountryCodeGw, nil
	case "GY":
		return CountryCodeGy, nil
	case "HT":
		return CountryCodeHt, nil
	case "HM":
		return CountryCodeHm, nil
	case "VA":
		return CountryCodeVa, nil
	case "HN":
		return CountryCodeHn, nil
	case "HK":
		return CountryCodeHk, nil
	case "HU":
		return CountryCodeHu, nil
	case "IS":
		return CountryCodeIs, nil
	case "IN":
		return CountryCodeIn, nil
	case "ID":
		return CountryCodeID, nil
	case "IR":
		return CountryCodeIr, nil
	case "IQ":
		return CountryCodeIq, nil
	case "IE":
		return CountryCodeIe, nil
	case "IM":
		return CountryCodeIm, nil
	case "IL":
		return CountryCodeIl, nil
	case "IT":
		return CountryCodeIt, nil
	case "JM":
		return CountryCodeJm, nil
	case "JP":
		return CountryCodeJp, nil
	case "JE":
		return CountryCodeJe, nil
	case "JO":
		return CountryCodeJo, nil
	case "KZ":
		return CountryCodeKz, nil
	case "KE":
		return CountryCodeKe, nil
	case "KI":
		return CountryCodeKi, nil
	case "KP":
		return CountryCodeKp, nil
	case "KR":
		return CountryCodeKr, nil
	case "KW":
		return CountryCodeKw, nil
	case "KG":
		return CountryCodeKg, nil
	case "LA":
		return CountryCodeLa, nil
	case "LV":
		return CountryCodeLv, nil
	case "LB":
		return CountryCodeLb, nil
	case "LS":
		return CountryCodeLs, nil
	case "LR":
		return CountryCodeLr, nil
	case "LY":
		return CountryCodeLy, nil
	case "LI":
		return CountryCodeLi, nil
	case "LT":
		return CountryCodeLt, nil
	case "LU":
		return CountryCodeLu, nil
	case "MO":
		return CountryCodeMo, nil
	case "MK":
		return CountryCodeMk, nil
	case "MG":
		return CountryCodeMg, nil
	case "MW":
		return CountryCodeMw, nil
	case "MY":
		return CountryCodeMy, nil
	case "MV":
		return CountryCodeMv, nil
	case "ML":
		return CountryCodeMl, nil
	case "MT":
		return CountryCodeMt, nil
	case "MH":
		return CountryCodeMh, nil
	case "MQ":
		return CountryCodeMq, nil
	case "MR":
		return CountryCodeMr, nil
	case "MU":
		return CountryCodeMu, nil
	case "YT":
		return CountryCodeYt, nil
	case "MX":
		return CountryCodeMx, nil
	case "FM":
		return CountryCodeFm, nil
	case "MD":
		return CountryCodeMd, nil
	case "MC":
		return CountryCodeMc, nil
	case "MN":
		return CountryCodeMn, nil
	case "ME":
		return CountryCodeMe, nil
	case "MS":
		return CountryCodeMs, nil
	case "MA":
		return CountryCodeMa, nil
	case "MZ":
		return CountryCodeMz, nil
	case "MM":
		return CountryCodeMm, nil
	case "NA":
		return CountryCodeNa, nil
	case "NR":
		return CountryCodeNr, nil
	case "NP":
		return CountryCodeNp, nil
	case "NL":
		return CountryCodeNl, nil
	case "NC":
		return CountryCodeNc, nil
	case "NZ":
		return CountryCodeNz, nil
	case "NI":
		return CountryCodeNi, nil
	case "NE":
		return CountryCodeNe, nil
	case "NG":
		return CountryCodeNg, nil
	case "NU":
		return CountryCodeNu, nil
	case "NF":
		return CountryCodeNf, nil
	case "MP":
		return CountryCodeMp, nil
	case "NO":
		return CountryCodeNo, nil
	case "OM":
		return CountryCodeOm, nil
	case "PK":
		return CountryCodePk, nil
	case "PW":
		return CountryCodePw, nil
	case "PS":
		return CountryCodePs, nil
	case "PA":
		return CountryCodePa, nil
	case "PG":
		return CountryCodePg, nil
	case "PY":
		return CountryCodePy, nil
	case "PE":
		return CountryCodePe, nil
	case "PH":
		return CountryCodePh, nil
	case "PN":
		return CountryCodePn, nil
	case "PL":
		return CountryCodePl, nil
	case "PT":
		return CountryCodePt, nil
	case "PR":
		return CountryCodePr, nil
	case "QA":
		return CountryCodeQa, nil
	case "RO":
		return CountryCodeRo, nil
	case "RU":
		return CountryCodeRu, nil
	case "RW":
		return CountryCodeRw, nil
	case "RE":
		return CountryCodeRe, nil
	case "BL":
		return CountryCodeBl, nil
	case "SH":
		return CountryCodeSh, nil
	case "KN":
		return CountryCodeKn, nil
	case "LC":
		return CountryCodeLc, nil
	case "MF":
		return CountryCodeMf, nil
	case "PM":
		return CountryCodePm, nil
	case "VC":
		return CountryCodeVc, nil
	case "WS":
		return CountryCodeWs, nil
	case "SM":
		return CountryCodeSm, nil
	case "ST":
		return CountryCodeSt, nil
	case "SA":
		return CountryCodeSa, nil
	case "SN":
		return CountryCodeSn, nil
	case "RS":
		return CountryCodeRs, nil
	case "SC":
		return CountryCodeSc, nil
	case "SL":
		return CountryCodeSl, nil
	case "SG":
		return CountryCodeSg, nil
	case "SX":
		return CountryCodeSx, nil
	case "SK":
		return CountryCodeSk, nil
	case "SI":
		return CountryCodeSi, nil
	case "SB":
		return CountryCodeSb, nil
	case "SO":
		return CountryCodeSo, nil
	case "ZA":
		return CountryCodeZa, nil
	case "GS":
		return CountryCodeGs, nil
	case "SS":
		return CountryCodeSs, nil
	case "ES":
		return CountryCodeEs, nil
	case "LK":
		return CountryCodeLk, nil
	case "SD":
		return CountryCodeSd, nil
	case "SR":
		return CountryCodeSr, nil
	case "SJ":
		return CountryCodeSj, nil
	case "SE":
		return CountryCodeSe, nil
	case "CH":
		return CountryCodeCh, nil
	case "SY":
		return CountryCodeSy, nil
	case "TW":
		return CountryCodeTw, nil
	case "TJ":
		return CountryCodeTj, nil
	case "TZ":
		return CountryCodeTz, nil
	case "TH":
		return CountryCodeTh, nil
	case "TL":
		return CountryCodeTl, nil
	case "TG":
		return CountryCodeTg, nil
	case "TK":
		return CountryCodeTk, nil
	case "TO":
		return CountryCodeTo, nil
	case "TT":
		return CountryCodeTt, nil
	case "TN":
		return CountryCodeTn, nil
	case "TR":
		return CountryCodeTr, nil
	case "TM":
		return CountryCodeTm, nil
	case "TC":
		return CountryCodeTc, nil
	case "TV":
		return CountryCodeTv, nil
	case "UG":
		return CountryCodeUg, nil
	case "UA":
		return CountryCodeUa, nil
	case "AE":
		return CountryCodeAe, nil
	case "GB":
		return CountryCodeGb, nil
	case "US":
		return CountryCodeUs, nil
	case "UM":
		return CountryCodeUm, nil
	case "UY":
		return CountryCodeUy, nil
	case "UZ":
		return CountryCodeUz, nil
	case "VU":
		return CountryCodeVu, nil
	case "VE":
		return CountryCodeVe, nil
	case "VN":
		return CountryCodeVn, nil
	case "VG":
		return CountryCodeVg, nil
	case "VI":
		return CountryCodeVi, nil
	case "WF":
		return CountryCodeWf, nil
	case "EH":
		return CountryCodeEh, nil
	case "YE":
		return CountryCodeYe, nil
	case "ZM":
		return CountryCodeZm, nil
	case "ZW":
		return CountryCodeZw, nil
	case "AX":
		return CountryCodeAx, nil
	}
	var t CountryCode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CountryCode) Ptr() *CountryCode {
	return &c
}

type DocumentResponse struct {
	// ID of the document. If not provided, this is a dynamic document that is generated on the fly.
	ID       *string      `json:"id,omitempty" url:"id,omitempty"`
	MimeType string       `json:"mimeType" url:"mimeType"`
	Type     DocumentType `json:"type" url:"type"`
	URI      string       `json:"uri" url:"uri"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DocumentResponse) GetID() *string {
	if d == nil {
		return nil
	}
	return d.ID
}

func (d *DocumentResponse) GetMimeType() string {
	if d == nil {
		return ""
	}
	return d.MimeType
}

func (d *DocumentResponse) GetType() DocumentType {
	if d == nil {
		return ""
	}
	return d.Type
}

func (d *DocumentResponse) GetURI() string {
	if d == nil {
		return ""
	}
	return d.URI
}

func (d *DocumentResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DocumentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentResponse) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DocumentType string

const (
	DocumentTypeInvoice             DocumentType = "INVOICE"
	DocumentTypePaymentConfirmation DocumentType = "PAYMENT_CONFIRMATION"
	DocumentTypeTenNinetyNine       DocumentType = "TEN_NINETY_NINE"
	DocumentTypeW9                  DocumentType = "W9"
	DocumentTypeCheck               DocumentType = "CHECK"
	DocumentTypeBankStatement       DocumentType = "BANK_STATEMENT"
	DocumentTypeContract            DocumentType = "CONTRACT"
	DocumentTypeOther               DocumentType = "OTHER"
)

func NewDocumentTypeFromString(s string) (DocumentType, error) {
	switch s {
	case "INVOICE":
		return DocumentTypeInvoice, nil
	case "PAYMENT_CONFIRMATION":
		return DocumentTypePaymentConfirmation, nil
	case "TEN_NINETY_NINE":
		return DocumentTypeTenNinetyNine, nil
	case "W9":
		return DocumentTypeW9, nil
	case "CHECK":
		return DocumentTypeCheck, nil
	case "BANK_STATEMENT":
		return DocumentTypeBankStatement, nil
	case "CONTRACT":
		return DocumentTypeContract, nil
	case "OTHER":
		return DocumentTypeOther, nil
	}
	var t DocumentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DocumentType) Ptr() *DocumentType {
	return &d
}

type FullName struct {
	FirstName  string  `json:"firstName" url:"firstName"`
	MiddleName *string `json:"middleName,omitempty" url:"middleName,omitempty"`
	LastName   string  `json:"lastName" url:"lastName"`
	Suffix     *string `json:"suffix,omitempty" url:"suffix,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *FullName) GetFirstName() string {
	if f == nil {
		return ""
	}
	return f.FirstName
}

func (f *FullName) GetMiddleName() *string {
	if f == nil {
		return nil
	}
	return f.MiddleName
}

func (f *FullName) GetLastName() string {
	if f == nil {
		return ""
	}
	return f.LastName
}

func (f *FullName) GetSuffix() *string {
	if f == nil {
		return nil
	}
	return f.Suffix
}

func (f *FullName) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FullName) UnmarshalJSON(data []byte) error {
	type unmarshaler FullName
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FullName(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FullName) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type IndividualGovernmentID struct {
	// Full Social Security Number. Must be in the format 123-45-6789.
	Ssn string `json:"ssn" url:"ssn"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *IndividualGovernmentID) GetSsn() string {
	if i == nil {
		return ""
	}
	return i.Ssn
}

func (i *IndividualGovernmentID) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *IndividualGovernmentID) UnmarshalJSON(data []byte) error {
	type unmarshaler IndividualGovernmentID
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IndividualGovernmentID(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *IndividualGovernmentID) String() string {
	if len(i.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(i.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

func NewOrderDirectionFromString(s string) (OrderDirection, error) {
	switch s {
	case "ASC":
		return OrderDirectionAsc, nil
	case "DESC":
		return OrderDirectionDesc, nil
	}
	var t OrderDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (o OrderDirection) Ptr() *OrderDirection {
	return &o
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode" url:"countryCode"`
	Number      string `json:"number" url:"number"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PhoneNumber) GetCountryCode() string {
	if p == nil {
		return ""
	}
	return p.CountryCode
}

func (p *PhoneNumber) GetNumber() string {
	if p == nil {
		return ""
	}
	return p.Number
}

func (p *PhoneNumber) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PhoneNumber) UnmarshalJSON(data []byte) error {
	type unmarshaler PhoneNumber
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PhoneNumber(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PhoneNumber) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type StringOrStringArray struct {
	String     string
	StringList []string

	typ string
}

func (s *StringOrStringArray) GetString() string {
	if s == nil {
		return ""
	}
	return s.String
}

func (s *StringOrStringArray) GetStringList() []string {
	if s == nil {
		return nil
	}
	return s.StringList
}

func (s *StringOrStringArray) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		s.typ = "String"
		s.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		s.typ = "StringList"
		s.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s StringOrStringArray) MarshalJSON() ([]byte, error) {
	if s.typ == "String" || s.String != "" {
		return json.Marshal(s.String)
	}
	if s.typ == "StringList" || s.StringList != nil {
		return json.Marshal(s.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type StringOrStringArrayVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (s *StringOrStringArray) Accept(visitor StringOrStringArrayVisitor) error {
	if s.typ == "String" || s.String != "" {
		return visitor.VisitString(s.String)
	}
	if s.typ == "StringList" || s.StringList != nil {
		return visitor.VisitStringList(s.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
