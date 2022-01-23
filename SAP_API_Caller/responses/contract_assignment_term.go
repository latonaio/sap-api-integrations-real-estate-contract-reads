package responses
  
type ContractAssignmentTerm struct {
	Value             []struct {
		InternalRealEstateNumber  string        `json:"InternalRealEstateNumber"`
		RETermType                string        `json:"RETermType"`
		RETermNumber              string        `json:"RETermNumber"`
		ValidityStartEndDateValue string        `json:"ValidityStartEndDateValue"`
		ValidityStartDate         string        `json:"ValidityStartDate"`
		ValidityEndDate           string        `json:"ValidityEndDate"`
		RETermName                string        `json:"RETermName"`
		BusinessArea              string        `json:"BusinessArea"`
		ProfitCenter              string        `json:"ProfitCenter"`
		REStatusObject            string        `json:"REStatusObject"`
		TaxJurisdiction           string        `json:"TaxJurisdiction"`
		Fund                      string        `json:"Fund"`
		FundsCenter               string        `json:"FundsCenter"`
		CommitmentItem            string        `json:"CommitmentItem"`
		FunctionalArea            string        `json:"FunctionalArea"`
		BudgetPeriod              string        `json:"BudgetPeriod"`
		ControllingArea           string        `json:"ControllingArea"`
		TaxCalculationProcedure   string        `json:"TaxCalculationProcedure"`
		FinancialManagementArea   string        `json:"FinancialManagementArea"`
	} `json:"value"`
}  
