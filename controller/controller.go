package controller

import (
	"api/model"
	"api/services"
	"encoding/json"
	"net/http"
	"strconv"
)

// func GetAllHandler(w http.ResponseWriter, r *http.Request) {
// 	ads, err := services.GetALlAd()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(ads)
// }

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var ad model.Ad
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 參數驗證Age <1~100>、Gender <enum：M、F>、Country <enum：TW、JP 等符合 https://zh.wikipedia.org/wiki/ISO_3166-1 >、Platform <enum：android, ios, web>
	for idx, condition := range ad.Conditions {
		if condition.AgeStart != nil && condition.AgeEnd != nil {
			//如果 condition.AgeStart或condition.AgeEnd不是1到100之間的數字，則返回錯誤
			if *condition.AgeStart < 1 || *condition.AgeStart > 100 || *condition.AgeEnd < 1 || *condition.AgeEnd > 100 {
				http.Error(w, "Age must be between 1 and 100", http.StatusBadRequest)
				return
			}
			//如果 condition.AgeStart和condition.AgeEnd相反，則返回錯誤
			if *condition.AgeStart > *condition.AgeEnd {
				http.Error(w, "AgeStart must be less than AgeEnd", http.StatusBadRequest)
				return
			}
		} else if !(condition.AgeStart == nil && condition.AgeEnd == nil) { //當只設定其中之一時，返回錯誤
			http.Error(w, "AgeStart and AgeEnd must be set together", http.StatusBadRequest)
			return
		} else {
			ageStart := 1
			ageEnd := 100
			ad.Conditions[idx].AgeStart = &ageStart
			ad.Conditions[idx].AgeEnd = &ageEnd
		}

		//如果 condition.Gender不是M或F，則返回錯誤
		if condition.Gender != nil {
			for _, gender := range *condition.Gender {
				if gender != "M" && gender != "F" {
					http.Error(w, "Gender must be M or F", http.StatusBadRequest)
					return
				}
			}
		} else {
			// 如果 condition.Gender為空，則將其設置為M和F
			gender := []string{"M", "F"}
			ad.Conditions[idx].Gender = (*model.StringSlice)(&gender)
		}

		if condition.Country != nil {
			//如果 condition.Country不是正規的ISO 3166-1，則返回錯誤。
			for _, country := range *condition.Country {
				if !isVaidCountryCode(country) {
					http.Error(w, "Country must be a valid ISO 3166-1 code", http.StatusBadRequest)
					return
				}
			}
		} else {
			country := []string{"ALL"}
			ad.Conditions[idx].Country = (*model.StringSlice)(&country)
		}

		if condition.Platform != nil {
			//如果 condition.Platform不是android、ios或web，則返回錯誤
			for _, platform := range *condition.Platform {
				if platform != "android" && platform != "ios" && platform != "web" {
					http.Error(w, "Platform must be android, ios or web", http.StatusBadRequest)
					return
				}
			}
		} else {
			//如果 condition.Platform為空，則將其設置為三種平台都有
			platform := []string{"android", "ios", "web"}
			ad.Conditions[idx].Platform = (*model.StringSlice)(&platform)
		}
	}

	err = services.CreateAd(&ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	adJson, err := json.Marshal(ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(adJson)
}

func SearchAdsHandler(w http.ResponseWriter, r *http.Request) {
	// 解析查詢參數
	ageStr := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageStr) //轉為整數
	if err != nil {
		age = 0
	}
	gender := r.URL.Query().Get("gender")
	country := r.URL.Query().Get("country")
	platform := r.URL.Query().Get("platform")
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		http.Error(w, "Missing offset parameter", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "Missing offset parameter", http.StatusBadRequest)
		return
	}

	// 設置預設值
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > 100 {
		limit = 5
	}
	// 調用 SearchAds 服務
	ads, err := services.SearchAds(gender, country, platform, age, offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 設置header
	w.Header().Set("Content-Type", "application/json")

	// 編碼查詢結果為 JSON 格式
	json.NewEncoder(w).Encode(ads)
}

func isVaidCountryCode(code string) bool {
	// 檢查 country 是否是正規的 ISO 3166-1 國家代碼
	switch code {
	case "AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG",
		"AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB",
		"BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW",
		"BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM",
		"CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM",
		"CD", "CG", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ",
		"DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE",
		"SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF",
		"GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP",
		"GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN",
		"HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL",
		"IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR",
		"KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT",
		"LU", "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ",
		"MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS",
		"MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI",
		"NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS",
		"PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA",
		"MK", "RO", "RU", "RW", "RE", "BL", "SH", "KN", "LC", "MF",
		"PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL",
		"SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES",
		"LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TW", "TJ", "TZ",
		"TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC",
		"TV", "UG", "UA", "AE", "GB", "UM", "US", "UY", "UZ", "VU",
		"VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW":
		return true
	default:
		return false
	}
}
