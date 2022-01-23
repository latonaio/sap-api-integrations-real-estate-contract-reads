package responses

type ContractCondition struct {
	Value             []struct {
		REConditionUUID                string        `json:"REConditionUUID"`
		InternalRealEstateNumber       string        `json:"InternalRealEstateNumber"`
		REStatusObjectCalculation      string        `json:"REStatusObjectCalculation"`
		REStatusObjectIDCalculation    string        `json:"REStatusObjectIDCalculation"`
		REConditionType                string        `json:"REConditionType"`
		REExtConditionPurpose          string        `json:"REExtConditionPurpose"`
		ValidityStartEndDateValue      string        `json:"ValidityStartEndDateValue"`
		ValidityStartDate              string        `json:"ValidityStartDate"`
		REStatusObjectDistribution     string        `json:"REStatusObjectDistribution"`
		REStatusObjectIDDistribution   string        `json:"REStatusObjectIDDistribution"`
		REObjectTypeDistribution       string        `json:"REObjectTypeDistribution"`
		REObjectTypePosting            string        `json:"REObjectTypePosting"`
		ValidityEndDate                string        `json:"ValidityEndDate"`
		CreationDateTime               string        `json:"CreationDateTime"`
		RESourceOfCreation             string        `json:"RESourceOfCreation"`
		LastChangeDateTime             string        `json:"LastChangeDateTime"`
		RESourceOfChange               string        `json:"RESourceOfChange"`
		REIsOneTimeCondition           bool          `json:"REIsOneTimeCondition"`
		REConditionIsStatistical       bool          `json:"REConditionIsStatistical"`
		REPostingTerm                  string        `json:"REPostingTerm"`
		RERhythmTerm                   string        `json:"RERhythmTerm"`
		REAdjustmentNumber             string        `json:"REAdjustmentNumber"`
		REOrglAssignmentTerm           string        `json:"REOrglAssignmentTerm"`
		RESalesTerm                    string        `json:"RESalesTerm"`
		REPeakSalesTerm                string        `json:"REPeakSalesTerm"`
		RESrvcChrgSettlementPostingTrm string        `json:"RESrvcChrgSettlementPostingTrm"`
		REWithholdingTaxTerm           string        `json:"REWithholdingTaxTerm"`
		RECalculationRule              string        `json:"RECalculationRule"`
		REUnitPrice                    float64       `json:"REUnitPrice"`
		REConditionCurrency            string        `json:"REConditionCurrency"`
		RECalculationRuleParam1        string        `json:"RECalculationRuleParam1"`
		RECalculationRuleParam2        string        `json:"RECalculationRuleParam2"`
		REDistributionRule             string        `json:"REDistributionRule"`
		REDistributionRuleParam1       string        `json:"REDistributionRuleParam1"`
		REDistributionRuleParam2       string        `json:"REDistributionRuleParam2"`
		REReasonForChange              string        `json:"REReasonForChange"`
		REStsObjectParamCalculation    string        `json:"REStsObjectParamCalculation"`
		REStsObjectParamDistribution   string        `json:"REStsObjectParamDistribution"`
	} `json:"value"`
}  
