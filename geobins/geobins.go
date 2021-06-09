package geobins

import "github.com/pkg/errors"

// Convert region to a numerical representation
//  in order to find distances between regions
// fixme: consider modifying node state for a region field?
func GetRegion(region string) (int, error) {
	switch region {
	case "Americas":
		return Americas, nil
	case "WesternEurope", "WestEurope":
		return WesternEurope, nil
	case "CentralEurope":
		return CentralEurope, nil
	case "EasternEurope", "EastEurope":
		return EasternEurope, nil
	case "MiddleEast":
		return MiddleEast, nil
	case "Africa":
		return Africa, nil
	case "Russia":
		return Russia, nil
	case "Asia":
		return Asia, nil
	default:
		return -1, errors.Errorf("Could not parse region info ('%s')", region)
	}
}

// Convert a numerical represenation of the region to a string
// fixme: consider modifying node state for a region field?
func GetRegionFromString(region int) (string, error) {
	switch region {
	case Americas:
		return "Americas", nil
	case WesternEurope:
		return "WesternEurope", nil
	case CentralEurope:
		return "CentralEurope", nil
	case EasternEurope:
		return "EasternEurope", nil
	case MiddleEast:
		return "MiddleEast", nil
	case Africa:
		return "Africa", nil
	case Russia:
		return "Russia", nil
	case Asia:
		return "Asia", nil
	default:
		return "", errors.Errorf("Could not parse region info ('%s')", region)
	}
}

// Enumeration which links regions to numbers for legibility purposes
const (
	Americas      = 0
	WesternEurope = 1
	CentralEurope = 2
	EasternEurope = 3
	MiddleEast    = 4
	Africa        = 5
	Russia        = 6
	Asia          = 7
)

var Geobins = map[string]int{
	"AI": Americas,
	"AQ": Americas,
	"AG": Americas,
	"AR": Americas,
	"BS": Americas,
	"UY": Americas,
	"BB": Americas,
	"BR": Americas,
	"PE": Americas,
	"BZ": Americas,
	"CL": Americas,
	"BM": Americas,
	"EC": Americas,
	"BO": Americas,
	"GY": Americas,
	"BQ": Americas,
	"GF": Americas,
	"SR": Americas,
	"BW": Americas,
	"PY": Americas,
	"MX": Americas,
	"CA": Americas,
	"US": Americas,
	"KY": Americas,
	"CO": Americas,
	"NI": Americas,
	"PR": Americas,
	"CR": Americas,
	"CU": Americas,
	"CW": Americas,
	"DM": Americas,
	"DO": Americas,
	"SV": Americas,
	"GT": Americas,
	"FK": Americas,
	"JM": Americas,
	"PA": Americas,
	"VE": Americas,
	"VG": Americas,
	"VI": Americas,
	"HT": Americas,
	"HN": Americas,
	"GL": Americas,
	"GD": Americas,
	"TC": Americas,
	"TT": Americas,
	"GS": Americas,
	"MS": Americas,
	"SX": Americas,
	"GP": Americas,
	"MQ": Americas,
	"LC": Americas,
	"KN": Americas,
	"PM": Americas,
	"BL": Americas,
	"VC": Americas,
	"AW": Americas,
	"MF": Americas,
	"AD": WesternEurope,
	"GB": WesternEurope,
	"ES": WesternEurope,
	"PT": WesternEurope,
	"NL": WesternEurope,
	"FR": WesternEurope,
	"IE": WesternEurope,
	"FO": WesternEurope,
	"IS": WesternEurope,
	"GI": WesternEurope,
	"MC": WesternEurope,
	"GG": WesternEurope,
	"JE": WesternEurope,
	"IM": WesternEurope,
	"BE": WesternEurope,
	"LU": WesternEurope,
	"AT": CentralEurope,
	"PL": CentralEurope,
	"NO": CentralEurope,
	"MT": CentralEurope,
	"DK": CentralEurope,
	"SE": CentralEurope,
	"CH": CentralEurope,
	"SK": CentralEurope,
	"SI": CentralEurope,
	"IT": CentralEurope,
	"DE": CentralEurope,
	"HU": CentralEurope,
	"LI": CentralEurope,
	"SJ": CentralEurope,
	"SM": CentralEurope,
	"VA": CentralEurope,
	"CZ": CentralEurope,
	"AX": EasternEurope,
	"AL": EasternEurope,
	"BY": EasternEurope,
	"UA": EasternEurope,
	"BA": EasternEurope,
	"LT": EasternEurope,
	"LV": EasternEurope,
	"BG": EasternEurope,
	"MD": EasternEurope,
	"HR": EasternEurope,
	"FI": EasternEurope,
	"CY": EasternEurope,
	"RO": EasternEurope,
	"EE": EasternEurope,
	"GR": EasternEurope,
	"ME": EasternEurope,
	"RS": EasternEurope,
	"MK": EasternEurope,
	"GE": EasternEurope,
	"AF": MiddleEast,
	"AM": MiddleEast,
	"AZ": MiddleEast,
	"BH": MiddleEast,
	"TR": MiddleEast,
	"SA": MiddleEast,
	"AE": MiddleEast,
	"IQ": MiddleEast,
	"QA": MiddleEast,
	"YE": MiddleEast,
	"TM": MiddleEast,
	"UZ": MiddleEast,
	"PK": MiddleEast,
	"IL": MiddleEast,
	"JO": MiddleEast,
	"KW": MiddleEast,
	"PS": MiddleEast,
	"SY": MiddleEast,
	"IR": MiddleEast,
	"LB": MiddleEast,
	"OM": MiddleEast,
	"KZ": MiddleEast,
	"KG": MiddleEast,
	"TJ": MiddleEast,
	"DZ": Africa,
	"AO": Africa,
	"BJ": Africa,
	"BV": Africa,
	"BF": Africa,
	"BI": Africa,
	"ZA": Africa,
	"CV": Africa,
	"CM": Africa,
	"CF": Africa,
	"TD": Africa,
	"RW": Africa,
	"KM": Africa,
	"CG": Africa,
	"CD": Africa,
	"SN": Africa,
	"SO": Africa,
	"CI": Africa,
	"ZM": Africa,
	"ZW": Africa,
	"ML": Africa,
	"NE": Africa,
	"NG": Africa,
	"UG": Africa,
	"DJ": Africa,
	"EG": Africa,
	"ET": Africa,
	"GA": Africa,
	"GM": Africa,
	"GH": Africa,
	"GQ": Africa,
	"ER": Africa,
	"SZ": Africa,
	"KE": Africa,
	"LR": Africa,
	"LY": Africa,
	"MW": Africa,
	"SL": Africa,
	"TN": Africa,
	"TG": Africa,
	"SS": Africa,
	"EH": Africa,
	"SD": Africa,
	"NA": Africa,
	"MA": Africa,
	"MZ": Africa,
	"TZ": Africa,
	"MR": Africa,
	"MU": Africa,
	"SC": Africa,
	"GN": Africa,
	"GW": Africa,
	"MG": Africa,
	"LS": Africa,
	"YT": Africa,
	"ST": Africa,
	"RE": Africa,
	"SH": Africa,
	"RU": Russia,
	"AS": Asia,
	"AU": Asia,
	"BD": Asia,
	"TH": Asia,
	"BT": Asia,
	"PH": Asia,
	"IO": Asia,
	"JP": Asia,
	"BN": Asia,
	"KH": Asia,
	"CN": Asia,
	"CX": Asia,
	"CC": Asia,
	"TL": Asia,
	"ID": Asia,
	"PG": Asia,
	"SB": Asia,
	"CK": Asia,
	"FJ": Asia,
	"NC": Asia,
	"MY": Asia,
	"VN": Asia,
	"LK": Asia,
	"NP": Asia,
	"IN": Asia,
	"KP": Asia,
	"KR": Asia,
	"TW": Asia,
	"NZ": Asia,
	"SG": Asia,
	"MM": Asia,
	"MN": Asia,
	"HK": Asia,
	"VU": Asia,
	"GU": Asia,
	"WF": Asia,
	"TV": Asia,
	"TO": Asia,
	"TK": Asia,
	"NR": Asia,
	"NF": Asia,
	"KI": Asia,
	"MH": Asia,
	"MV": Asia,
	"MO": Asia,
	"LA": Asia,
	"FM": Asia,
	"PW": Asia,
	"NU": Asia,
	"MP": Asia,
	"WS": Asia,
	"PF": Asia,
	"PN": Asia,
	"UM": Asia,
	"HM": Asia,
	"TF": Asia,
}
