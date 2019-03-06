package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ErrorStruct struct {
	Message string `json:",string"`
}

var timezones = map[string][]string{
	"ACDT": []string{
		"Australia/Adelaide",
		"Australia/Broken_Hill",
		"Australia/South",
		"Australia/Yancowinna",
	},
	"ACST": []string{
		"Australia/Darwin",
		"Australia/North",
	},
	"ACT": []string{
		"America/Eirunepe",
		"America/Porto_Acre",
		"America/Rio_Branco",
		"Brazil/Acre",
	},
	"ACWST": []string{
		"Australia/Eucla",
	},
	"AEDT": []string{
		"Australia/ACT",
		"Australia/Canberra",
		"Australia/Currie",
		"Australia/Hobart",
		"Australia/Melbourne",
		"Australia/NSW",
		"Australia/Sydney",
		"Australia/Tasmania",
		"Australia/Victoria",
	},
	"AEST": []string{
		"Australia/Brisbane",
		"Australia/Lindeman",
		"Australia/Queensland",
	},
	"AFT": []string{
		"Asia/Kabul",
	},
	"AKST": []string{
		"America/Anchorage",
		"America/Juneau",
		"America/Nome",
		"America/Sitka",
		"America/Yakutat",
		"US/Alaska",
	},

	"ALMT": []string{
		"Asia/Almaty",
	},
	"AMT": []string{
		"America/Boa_Vista",
		"America/Campo_Grande",
		"America/Cuiaba",
		"America/Manaus",
		"America/Porto_Velho",
		"Asia/Yerevan",
		"Brazil/West",
	},

	"ANAT": []string{
		"Asia/Anadyr",
	},
	"AQTT": []string{
		"Asia/Aqtau",
		"Asia/Aqtobe",
	},
	"ART": []string{
		"America/Argentina/Buenos_Aires",
		"America/Argentina/Catamarca",
		"America/Argentina/ComodRivadavia",
		"America/Argentina/Cordoba",
		"America/Argentina/Jujuy",
		"America/Argentina/La_Rioja",
		"America/Argentina/Mendoza",
		"America/Argentina/Rio_Gallegos",
		"America/Argentina/Salta",
		"America/Argentina/San_Juan",
		"America/Argentina/San_Luis",
		"America/Argentina/Tucuman",
		"America/Argentina/Ushuaia",
		"America/Buenos_Aires",
		"America/Catamarca",
		"America/Cordoba",
		"America/Jujuy",
		"America/Mendoza",
		"America/Rosario",
	},
	"AST": []string{
		"America/Anguilla",
		"America/Antigua",
		"America/Aruba",
		"America/Barbados",
		"America/Blanc-Sablon",
		"America/Curacao",
		"America/Dominica",
		"America/Glace_Bay",
		"America/Goose_Bay",
		"America/Grand_Turk",
		"America/Grenada",
		"America/Guadeloupe",
		"America/Halifax",
		"America/Kralendijk",
		"America/Lower_Princes",
		"America/Marigot",
		"America/Martinique",
		"America/Moncton",
		"America/Montserrat",
		"America/Port_of_Spain",
		"America/Puerto_Rico",
		"America/Santo_Domingo",
		"America/St_Barthelemy",
		"America/St_Kitts",
		"America/St_Lucia",
		"America/St_Thomas",
		"America/St_Vincent",
		"America/Thule",
		"America/Tortola",
		"America/Virgin",
		"Asia/Aden",
		"Asia/Baghdad",
		"Asia/Bahrain",
		"Asia/Kuwait",
		"Asia/Qatar",
		"Asia/Riyadh",
		"Atlantic/Bermuda",
		"Canada/Atlantic",
	},
	"AWST": []string{
		"Antarctica/Casey",
		"Australia/Perth",
		"Australia/West",
	},
	"AZOT": []string{
		"Atlantic/Azores",
	},
	"AZT": []string{
		"Asia/Baku",
	},
	"BDT": []string{
		"Asia/Dacca",
		"Asia/Dhaka",
	},

	"BNT": []string{
		"Asia/Brunei",
	},
	"BOT": []string{
		"America/La_Paz",
	},
	"BRT": []string{
		"America/Araguaina",
		"America/Bahia",
		"America/Belem",
		"America/Fortaleza",
		"America/Maceio",
		"America/Recife",
		"America/Santarem",
		"America/Sao_Paulo",
		"Brazil/East",
	},
	"BST": []string{
		"Pacific/Bougainville",
	},

	"BTT": []string{
		"Asia/Thimbu",
		"Asia/Thimphu",
	},
	"CAT": []string{
		"Africa/Blantyre",
		"Africa/Bujumbura",
		"Africa/Gaborone",
		"Africa/Harare",
		"Africa/Kigali",
		"Africa/Lubumbashi",
		"Africa/Lusaka",
		"Africa/Maputo",
	},
	"CCT": []string{
		"Indian/Cocos",
	},
	"CET": []string{
		"Africa/Algiers",
		"Africa/Ceuta",
		"Africa/Tunis",
		"Arctic/Longyearbyen",
		"Atlantic/Jan_Mayen",
		"CET",
		"Europe/Amsterdam",
		"Europe/Andorra",
		"Europe/Belgrade",
		"Europe/Berlin",
		"Europe/Bratislava",
		"Europe/Brussels",
		"Europe/Budapest",
		"Europe/Busingen",
		"Europe/Copenhagen",
		"Europe/Gibraltar",
		"Europe/Ljubljana",
		"Europe/Luxembourg",
		"Europe/Madrid",
		"Europe/Malta",
		"Europe/Monaco",
		"Europe/Oslo",
		"Europe/Paris",
		"Europe/Podgorica",
		"Europe/Prague",
		"Europe/Rome",
		"Europe/San_Marino",
		"Europe/Sarajevo",
		"Europe/Skopje",
		"Europe/Stockholm",
		"Europe/Tirane",
		"Europe/Vaduz",
		"Europe/Vatican",
		"Europe/Vienna",
		"Europe/Warsaw",
		"Europe/Zagreb",
		"Europe/Zurich",
		"Poland",
	},
	"CHADT": []string{
		"NZ-CHAT",
		"Pacific/Chatham",
	},
	"CHOT": []string{
		"Asia/Choibalsan",
	},
	"CHUT": []string{
		"Pacific/Chuuk",
		"Pacific/Truk",
		"Pacific/Yap",
	},
	"CKT": []string{
		"Pacific/Rarotonga",
	},
	"CLT": []string{
		"America/Santiago",
		"Antarctica/Palmer",
		"Chile/Continental",
	},
	"COT": []string{
		"America/Bogota",
	},
	"CST": []string{
		"America/Bahia_Banderas",
		"America/Belize",
		"America/Chicago",
		"America/Costa_Rica",
		"America/El_Salvador",
		"America/Guatemala",
		"America/Havana",
		"America/Indiana/Knox",
		"America/Indiana/Tell_City",
		"America/Knox_IN",
		"America/Managua",
		"America/Matamoros",
		"America/Menominee",
		"America/Merida",
		"America/Mexico_City",
		"America/Monterrey",
		"America/North_Dakota/Beulah",
		"America/North_Dakota/Center",
		"America/North_Dakota/New_Salem",
		"America/Rainy_River",
		"America/Rankin_Inlet",
		"America/Regina",
		"America/Resolute",
		"America/Swift_Current",
		"America/Tegucigalpa",
		"America/Winnipeg",
		"Asia/Chongqing",
		"Asia/Chungking",
		"Asia/Harbin",
		"Asia/Macao",
		"Asia/Macau",
		"Asia/Shanghai",
		"Asia/Taipei",
		"CST6CDT",
		"Canada/Central",
		"Canada/East-Saskatchewan",
		"Canada/Saskatchewan",
		"Cuba",
		"Mexico/General",
		"PRC",
		"ROC",
		"US/Central",
		"US/Indiana-Starke",
	},

	"CVT": []string{
		"Atlantic/Cape_Verde",
	},
	"CXT": []string{
		"Indian/Christmas",
	},
	"ChST": []string{
		"Pacific/Guam",
		"Pacific/Saipan",
	},
	"DAVT": []string{
		"Antarctica/Davis",
	},
	"DDUT": []string{
		"Antarctica/DumontDUrville",
	},
	"EAST": []string{
		"Chile/EasterIsland",
		"Pacific/Easter",
	},
	"EAT": []string{"Africa/Addis_Ababa",
		"Africa/Asmara",
		"Africa/Asmera",
		"Africa/Dar_es_Salaam",
		"Africa/Djibouti",
		"Africa/Juba",
		"Africa/Kampala",
		"Africa/Khartoum",
		"Africa/Mogadishu",
		"Africa/Nairobi",
		"Indian/Antananarivo",
		"Indian/Comoro",
		"Indian/Mayotte",
	},
	"ECT": []string{
		"America/Guayaquil",
	},
	"EET": []string{
		"Africa/Cairo",
		"Africa/Tripoli",
		"Asia/Amman",
		"Asia/Beirut",
		"Asia/Damascus",
		"Asia/Gaza",
		"Asia/Hebron",
		"Asia/Istanbul",
		"Asia/Nicosia",
		"EET",
		"Egypt",
		"Europe/Athens",
		"Europe/Bucharest",
		"Europe/Chisinau",
		"Europe/Helsinki",
		"Europe/Istanbul",
		"Europe/Kaliningrad",
		"Europe/Kiev",
		"Europe/Mariehamn",
		"Europe/Nicosia",
		"Europe/Riga",
		"Europe/Sofia",
		"Europe/Tallinn",
		"Europe/Tiraspol",
		"Europe/Uzhgorod",
		"Europe/Vilnius",
		"Europe/Zaporozhye",
		"Libya",
		"Turkey",
	},
	"EGT": []string{
		"America/Scoresbysund",
	},
	"EST": []string{
		"America/Atikokan",
		"America/Cancun",
		"America/Cayman",
		"America/Coral_Harbour",
		"America/Detroit",
		"America/Fort_Wayne",
		"America/Indiana/Indianapolis",
		"America/Indiana/Marengo",
		"America/Indiana/Petersburg",
		"America/Indiana/Vevay",
		"America/Indiana/Vincennes",
		"America/Indiana/Winamac",
		"America/Indianapolis",
		"America/Iqaluit",
		"America/Jamaica",
		"America/Kentucky/Louisville",
		"America/Kentucky/Monticello",
		"America/Louisville",
		"America/Montreal",
		"America/Nassau",
		"America/New_York",
		"America/Nipigon",
		"America/Panama",
		"America/Pangnirtung",
		"America/Port-au-Prince",
		"America/Thunder_Bay",
		"America/Toronto",
		"Canada/Eastern",
		"EST",
		"EST5EDT",
		"Jamaica",
		"US/East-Indiana",
		"US/Eastern",
		"US/Michigan",
	},
	"FJT": []string{
		"Pacific/Fiji",
	},
	"FKST": []string{
		"Atlantic/Stanley",
	},
	"FNT": []string{
		"America/Noronha",
		"Brazil/DeNoronha",
	},
	"GALT": []string{
		"Pacific/Galapagos",
	},
	"GAMT": []string{
		"Pacific/Gambier",
	},
	"GET": []string{
		"Asia/Tbilisi",
	},
	"GFT": []string{
		"America/Cayenne",
	},
	"GILT": []string{
		"Pacific/Tarawa",
	},
	"GMT": []string{
		"Africa/Abidjan",
		"Africa/Accra",
		"Africa/Bamako",
		"Africa/Banjul",
		"Africa/Bissau",
		"Africa/Conakry",
		"Africa/Dakar",
		"Africa/Freetown",
		"Africa/Lome",
		"Africa/Monrovia",
		"Africa/Nouakchott",
		"Africa/Ouagadougou",
		"Africa/Sao_Tome",
		"Africa/Timbuktu",
		"America/Danmarkshavn",
		"Atlantic/Reykjavik",
		"Atlantic/St_Helena",
		"Eire",
		"Etc/GMT",
		"Etc/GMT+0",
		"Etc/GMT-0",
		"Etc/GMT0",
		"Etc/Greenwich",
		"Europe/Belfast",
		"Europe/Dublin",
		"Europe/Guernsey",
		"Europe/Isle_of_Man",
		"Europe/Jersey",
		"Europe/London",
		"GB",
		"GB-Eire",
		"GMT",
		"GMT+0",
		"GMT-0",
		"GMT0",
		"Greenwich",
		"Iceland",
	},
	"GMT+1": []string{
		"Etc/GMT+1",
	},
	"GMT+10": []string{
		"Etc/GMT+10",
	},
	"GMT+11": []string{
		"Etc/GMT+11",
	},
	"GMT+12": []string{
		"Etc/GMT+12",
	},
	"GMT+2": []string{
		"Etc/GMT+2",
	},
	"GMT+3": []string{
		"Etc/GMT+3",
	},
	"GMT+4": []string{
		"Etc/GMT+4",
	},
	"GMT+5": []string{
		"Etc/GMT+5",
	},
	"GMT+6": []string{
		"Etc/GMT+6",
	},
	"GMT+7": []string{
		"Etc/GMT+7",
	},
	"GMT+8": []string{
		"Etc/GMT+8",
	},
	"GMT+9": []string{
		"Etc/GMT+9",
	},
	"GMT-1": []string{
		"Etc/GMT-1",
	},
	"GMT-10": []string{
		"Etc/GMT-10",
	},
	"GMT-11": []string{
		"Etc/GMT-11",
	},
	"GMT-12": []string{
		"Etc/GMT-12",
	},
	"GMT-13": []string{
		"Etc/GMT-13",
	},
	"GMT-14": []string{
		"Etc/GMT-14",
	},
	"GMT-2": []string{
		"Etc/GMT-2",
	},
	"GMT-3": []string{
		"Etc/GMT-3",
	},
	"GMT-4": []string{
		"Etc/GMT-4",
	},
	"GMT-5": []string{
		"Etc/GMT-5",
	},
	"GMT-6": []string{
		"Etc/GMT-6",
	},
	"GMT-7": []string{
		"Etc/GMT-7",
	},
	"GMT-8": []string{
		"Etc/GMT-8",
	},
	"GMT-9": []string{
		"Etc/GMT-9",
	},
	"GST": []string{
		"Asia/Dubai",
		"Asia/Muscat",
		"Atlantic/South_Georgia",
	},
	"GYT": []string{
		"America/Guyana",
	},
	"HKT": []string{
		"Asia/Hong_Kong",
		"Hongkong",
	},
	"HOVT": []string{
		"Asia/Hovd",
	},
	"HST": []string{
		"America/Adak",
		"America/Atka",
		"HST",
		"Pacific/Honolulu",
		"Pacific/Johnston",
		"US/Aleutian",
		"US/Hawaii",
	},
	"ICT": []string{
		"Asia/Bangkok",
		"Asia/Ho_Chi_Minh",
		"Asia/Phnom_Penh",
		"Asia/Saigon",
		"Asia/Vientiane",
	},
	"IOT": []string{
		"Indian/Chagos",
	},
	"IRKT": []string{
		"Asia/Chita",
		"Asia/Irkutsk",
	},
	"IRST": []string{
		"Asia/Tehran",
		"Iran",
	},
	"IST": []string{
		"Asia/Calcutta",
		"Asia/Colombo",
		"Asia/Jerusalem",
		"Asia/Kolkata",
		"Asia/Tel_Aviv",
		"Israel",
	},
	"JST": []string{
		"Asia/Tokyo",
		"Japan",
	},
	"KGT": []string{
		"Asia/Bishkek",
	},
	"KOST": []string{
		"Pacific/Kosrae",
	},
	"KRAT": []string{
		"Asia/Krasnoyarsk",
		"Asia/Novokuznetsk",
	},
	"KST": []string{
		"Asia/Pyongyang",
		"Asia/Seoul",
		"ROK",
	},
	"LHDT": []string{
		"Australia/LHI",
		"Australia/Lord_Howe",
	},
	"LINT": []string{
		"Pacific/Kiritimati",
	},
	"MAGT": []string{
		"Asia/Magadan",
	},
	"MART": []string{
		"Pacific/Marquesas",
	},
	"MAWT": []string{
		"Antarctica/Mawson",
	},
	"MET": []string{
		"MET",
	},
	"MHT": []string{
		"Kwajalein",
		"Pacific/Kwajalein",
		"Pacific/Majuro",
	},
	"MIST": []string{
		"Antarctica/Macquarie",
	},
	"MMT": []string{
		"Asia/Rangoon",
	},
	"MSK": []string{
		"Europe/Minsk",
		"Europe/Moscow",
		"Europe/Simferopol",
		"Europe/Volgograd",
		"W-SU",
	},
	"MST": []string{
		"America/Boise",
		"America/Cambridge_Bay",
		"America/Chihuahua",
		"America/Creston",
		"America/Dawson_Creek",
		"America/Denver",
		"America/Edmonton",
		"America/Fort_Nelson",
		"America/Hermosillo",
		"America/Inuvik",
		"America/Mazatlan",
		"America/Ojinaga",
		"America/Phoenix",
		"America/Shiprock",
		"America/Yellowknife",
		"Canada/Mountain",
		"MST",
		"MST7MDT",
		"Mexico/BajaSur",
		"Navajo",
		"US/Arizona",
		"US/Mountain",
	},
	"MUT": []string{
		"Indian/Mauritius",
	},
	"MVT": []string{
		"Indian/Maldives",
		"Asia/Kuala_Lumpur",
		"Asia/Kuching",
	},
	"NCT": []string{
		"Pacific/Noumea",
	},
	"NFT": []string{
		"Pacific/Norfolk",
	},
	"NOVT": []string{
		"Asia/Novosibirsk",
	},
	"NPT": []string{
		"Asia/Kathmandu",
		"Asia/Katmandu",
	},
	"NRT": []string{
		"Pacific/Nauru",
	},
	"NST": []string{
		"America/St_Johns",
		"Canada/Newfoundland",
	},
	"NUT": []string{
		"Pacific/Niue",
	},
	"NZDT": []string{
		"Antarctica/McMurdo",
		"Antarctica/South_Pole",
		"NZ",
		"Pacific/Auckland",
	},
	"OMST": []string{
		"Asia/Omsk",
	},
	"ORAT": []string{
		"Asia/Oral",
	},
	"PET": []string{
		"America/Lima",
	},
	"PETT": []string{
		"Asia/Kamchatka",
	},
	"PGT": []string{
		"Pacific/Port_Moresby",
	},
	"PHOT": []string{
		"Pacific/Enderbury",
	},
	"PHT": []string{
		"Asia/Manila",
	},
	"PKT": []string{
		"Asia/Karachi",
	},
	"PMST": []string{
		"America/Miquelon",
	},
	"PONT": []string{
		"Pacific/Pohnpei",
		"Pacific/Ponape",
	},
	"PST": []string{
		"America/Dawson",
		"America/Ensenada",
		"America/Los_Angeles",
		"America/Metlakatla",
		"America/Santa_Isabel",
		"America/Tijuana",
		"America/Vancouver",
		"America/Whitehorse",
		"Canada/Pacific",
		"Canada/Yukon",
		"Mexico/BajaNorte",
		"PST8PDT",
		"Pacific/Pitcairn",
		"US/Pacific",
		"US/Pacific-New",
	},
	"PWT": []string{
		"Pacific/Palau",
	},
	"PYST": []string{
		"America/Asuncion",
	},
	"QYZT": []string{
		"Asia/Qyzylorda",
	},
	"RET": []string{
		"Indian/Reunion",
	},
	"ROTT": []string{
		"Antarctica/Rothera",
	},
	"SAKT": []string{
		"Asia/Sakhalin",
	},
	"SAMT": []string{
		"Europe/Samara",
	},
	"SAST": []string{
		"Africa/Johannesburg",
		"Africa/Maseru",
		"Africa/Mbabane",
	},
	"SBT": []string{
		"Pacific/Guadalcanal",
	},
	"SCT": []string{
		"Indian/Mahe",
	},
	"SGT": []string{
		"Asia/Singapore",
		"Singapore",
	},
	"SRET": []string{
		"Asia/Srednekolymsk",
	},
	"SRT": []string{
		"America/Paramaribo",
	},
	"SST": []string{
		"Pacific/Midway",
		"Pacific/Pago_Pago",
		"Pacific/Samoa",
		"US/Samoa",
	},
	"SYOT": []string{
		"Antarctica/Syowa",
	},
	"TAHT": []string{
		"Pacific/Tahiti",
	},
	"TFT": []string{
		"Indian/Kerguelen",
	},
	"TJT": []string{
		"Asia/Dushanbe",
	},
	"TKT": []string{
		"Pacific/Fakaofo",
	},
	"TLT": []string{
		"Asia/Dili",
	},
	"TMT": []string{
		"Asia/Ashgabat",
		"Asia/Ashkhabad",
	},
	"TOT": []string{
		"Pacific/Tongatapu",
	},
	"TVT": []string{
		"Pacific/Funafuti",
	},
	"UCT": []string{
		"Etc/UCT",
		"UCT",
	},
	"ULAT": []string{
		"Asia/Ulaanbaatar",
		"Asia/Ulan_Bator",
	},
	"UTC": []string{
		"Antarctica/Troll",
		"Etc/UTC",
		"Etc/Universal",
		"Etc/Zulu",
		"UTC",
		"Universal",
		"Zulu",
	},
	"UYST": []string{
		"America/Montevideo",
	},
	"UZT": []string{
		"Asia/Samarkand",
		"Asia/Tashkent",
	},
	"VET": []string{
		"America/Caracas",
	},
	"VLAT": []string{
		"Asia/Ust-Nera",
		"Asia/Vladivostok",
	},
	"VOST": []string{
		"Antarctica/Vostok",
	},
	"VUT": []string{
		"Pacific/Efate",
	},
	"WAKT": []string{
		"Pacific/Wake",
	},
	"WAST": []string{
		"Africa/Windhoek",
	},
	"WAT": []string{
		"Africa/Bangui",
		"Africa/Brazzaville",
		"Africa/Douala",
		"Africa/Kinshasa",
		"Africa/Lagos",
		"Africa/Libreville",
		"Africa/Luanda",
		"Africa/Malabo",
		"Africa/Ndjamena",
		"Africa/Niamey",
		"Africa/Porto-Novo",
	},
	"WET": []string{
		"Africa/Casablanca",
		"Africa/El_Aaiun",
		"Atlantic/Canary",
		"Atlantic/Faeroe",
		"Atlantic/Faroe",
		"Atlantic/Madeira",
		"Europe/Lisbon",
		"Portugal",
		"WET",
	},
	"WFT": []string{
		"Pacific/Wallis",
	},
	"WGT": []string{
		"America/Godthab",
	},
	"WIB": []string{
		"Asia/Jakarta",
		"Asia/Pontianak",
	},
	"WIT": []string{
		"Asia/Jayapura",
	},
	"WITA": []string{
		"Asia/Makassar",
		"Asia/Ujung_Pandang",
	},
	"WSDT": []string{
		"Pacific/Apia",
	},
	"XJT": []string{
		"Asia/Kashgar",
		"Asia/Urumqi",
	},
	"YAKT": []string{
		"Asia/Khandyga",
		"Asia/Yakutsk",
	},
	"YEKT": []string{
		"Asia/Yekaterinburg",
	},
	"Local time zone must be set--see zic manual page": []string{
		"Factory",
	},
}

func getTimeZone(w http.ResponseWriter, r *http.Request)  {
	urlZones := r.URL.Query().Get("zones")

	result := map[string]string{}

	errors := false

	zones := strings.Split(urlZones, ",")

	for _, key := range zones {
		if _,zoneExist := result[key]; !zoneExist {
			zoneTime, err :=  getTime(key)

			if err != nil {
				jsonErr, _ := json.Marshal(&ErrorStruct{Message: err.Error()})
				errors = true

				fmt.Fprintln(w, string(jsonErr))
				break
			}

			result[key] = zoneTime
		}
	}

	if len(result) != 0 && !errors {
		jsonString, err := json.Marshal(result)

		if err != nil {
			panic(err)
		}

		fmt.Fprintln(w, string(jsonString))
	}
}

func getTime(timeZoneName string) (string, error) {
	if _, keyExists := timezones[timeZoneName]; keyExists {
		now := time.Now()
		loc, _ := time.LoadLocation(timezones[timeZoneName][0])
		utc := now.In(loc)
		return utc.Format("15:04:05"), nil
	}

	return "", errors.New("Can't find timezone: " + timeZoneName)
}


func main() {
	http.HandleFunc("/getTime", getTimeZone)


	fmt.Println("Starting http server on port :8080")
	http.ListenAndServe(":8080", nil)
}