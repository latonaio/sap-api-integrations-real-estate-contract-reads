package responses

type ContractValuationCondition struct {
	Value []struct {
		InternalRealEstateNumber     string  `json:"InternalRealEstateNumber"`
		RETermNumber                 string  `json:"RETermNumber"`
		ValidityStartEndDateValue    string  `json:"ValidityStartEndDateValue"`
		REConditionType              string  `json:"REConditionType"`
		REConditionValidityStartDate string  `json:"REConditionValidityStartDate"`
		REExtConditionPurpose        string  `json:"REExtConditionPurpose"`
		REStatusObjectCalculation    string  `json:"REStatusObjectCalculation"`
		RETermName                   string  `json:"RETermName"`
		ValidityStartDate            string  `json:"ValidityStartDate"`
		ValidityEndDate              string  `json:"ValidityEndDate"`
		REConditionValidityEndDate   string  `json:"REConditionValidityEndDate"`
		REValuationCndnProperty      string  `json:"REValuationCndnProperty"`
		REValuationCndnConsdtn       string  `json:"REValuationCndnConsdtn"`
		REIsValuationCndnConsdtn     bool    `json:"REIsValuationCndnConsdtn"`
		REValuationCndnSharePercent  float64 `json:"REValuationCndnSharePercent"`
		REValuationCndnShareAbsltAmt int     `json:"REValuationCndnShareAbsltAmt"`
		REValuationCurrency          string  `json:"REValuationCurrency"`
		REValuationCndnStatus        string  `json:"REValuationCndnStatus"`
		REInfoText                   string  `json:"REInfoText"`
	} `json:"value"`
}
