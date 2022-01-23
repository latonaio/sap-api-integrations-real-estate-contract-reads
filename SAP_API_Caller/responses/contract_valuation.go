package responses

type ContractValuation struct {
	Value []struct {
		InternalRealEstateNumber     string  `json:"InternalRealEstateNumber"`
		RETermNumber                 string  `json:"RETermNumber"`
		ValidityStartEndDateValue    string  `json:"ValidityStartEndDateValue"`
		RETermType                   string  `json:"RETermType"`
		RETermName                   string  `json:"RETermName"`
		REValuationRule              string  `json:"REValuationRule"`
		ValidityStartDate            string  `json:"ValidityStartDate"`
		ValidityEndDate              string  `json:"ValidityEndDate"`
		REStatusObject               string  `json:"REStatusObject"`
		RealEstateObjectType         string  `json:"RealEstateObjectType"`
		REIdentification             string  `json:"REIdentification"`
		ObjectValidFrom              string  `json:"ObjectValidFrom"`
		REConsiderationStartDate     string  `json:"REConsiderationStartDate"`
		RECashFlowPostingFromDate    string  `json:"RECashFlowPostingFromDate"`
		REStatusObjectAsset          string  `json:"REStatusObjectAsset"`
		REValuationClassification    string  `json:"REValuationClassification"`
		REInterestRate               float64 `json:"REInterestRate"`
		REFrequencyTerm              string  `json:"REFrequencyTerm"`
		REDistributionRule           string  `json:"REDistributionRule"`
		REDistributionRuleParam1     string  `json:"REDistributionRuleParam1"`
		REDistributionRuleParam2     string  `json:"REDistributionRuleParam2"`
		REProbableEndDate            string  `json:"REProbableEndDate"`
		REAssetRightOfUseEndDate     string  `json:"REAssetRightOfUseEndDate"`
		REValuationRuleStatus        string  `json:"REValuationRuleStatus"`
		REValuationStatus            string  `json:"REValuationStatus"`
		REValuationStatusReason      string  `json:"REValuationStatusReason"`
		REValuationBehavior          string  `json:"REValuationBehavior"`
		RETaxType                    string  `json:"RETaxType"`
		TaxGroup                     string  `json:"TaxGroup"`
		REAccountingObject           string  `json:"REAccountingObject"`
		REInfoText                   string  `json:"REInfoText"`
		REValuationFactorNumerator   int     `json:"REValuationFactorNumerator"`
		REValuationFactorDenominator int     `json:"REValuationFactorDenominator"`
		REValuationCurrency          string  `json:"REValuationCurrency"`
		REValuationQuestionnaireUUID string  `json:"REValuationQuestionnaireUUID"`
		Country                      string  `json:"Country"`
	} `json:"value"`
}
