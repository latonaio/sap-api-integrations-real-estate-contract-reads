package responses

type Contract struct {
	Value             []struct {
		InternalRealEstateNumber       string        `json:"InternalRealEstateNumber"`
		CompanyCode                    string        `json:"CompanyCode"`
		RealEstateContract             string        `json:"RealEstateContract"`
		REStatusObject                 string        `json:"REStatusObject"`
		REInternalFinNumber            string        `json:"REInternalFinNumber"`
		RECreationType                 string        `json:"RECreationType"`
		CreationDate                   string        `json:"CreationDate"`
		CreationTime                   string        `json:"CreationTime"`
		RESourceOfCreation             string        `json:"RESourceOfCreation"`
		LastChangeDate                 string        `json:"LastChangeDate"`
		LastChangeTime                 string        `json:"LastChangeTime"`
		RESourceOfChange               string        `json:"RESourceOfChange"`
		Responsible                    string        `json:"Responsible"`
		REUserExclusive                string        `json:"REUserExclusive"`
		REAuthorizationGroup           string        `json:"REAuthorizationGroup"`
		REContractType                 string        `json:"REContractType"`
		ContractStartDate              string        `json:"ContractStartDate"`
		ContractEndDate                string        `json:"ContractEndDate"`
		REContractName                 string        `json:"REContractName"`
		REContractActivateDate         string        `json:"REContractActivateDate"`
		RETenancyLaw                   string        `json:"RETenancyLaw"`
		REContractNumberOld            string        `json:"REContractNumberOld"`
		REMainContractCompanyCode      string        `json:"REMainContractCompanyCode"`
		REMainContract                 string        `json:"REMainContract"`
		REContractCurrency             string        `json:"REContractCurrency"`
		REIndustrySector               string        `json:"REIndustrySector"`
		REIsSalesRelevant              bool          `json:"REIsSalesRelevant"`
		REContractDepositType          string        `json:"REContractDepositType"`
		REContractSignDate             string        `json:"REContractSignDate"`
		REContract2SignDate            string        `json:"REContract2SignDate"`
		REContractCashFlowDate         string        `json:"REContractCashFlowDate"`
		REContractFirstEndDate         string        `json:"REContractFirstEndDate"`
		REContractNoticeDate           string        `json:"REContractNoticeDate"`
		REContractNoticeInDate         string        `json:"REContractNoticeInDate"`
		REContractNoticeReason         string        `json:"REContractNoticeReason"`
		REContractNoticeActivationDate string        `json:"REContractNoticeActivationDate"`
		RECashFlowArchivedToDate       string        `json:"RECashFlowArchivedToDate"`
		RECashFlowLockedToDate         string        `json:"RECashFlowLockedToDate"`
		RECashFlowPostingFromDate      string        `json:"RECashFlowPostingFromDate"`
		REBusinessEntity               string        `json:"REBusinessEntity"`
		REPossessionStartDate          string        `json:"REPossessionStartDate"`
		REPossessionEndDate            string        `json:"REPossessionEndDate"`
		REHasMultipleAssignments       bool          `json:"REHasMultipleAssignments"`
		REObjectAvailableFromDate      string        `json:"REObjectAvailableFromDate"`
		REObjectAvailableToDate        string        `json:"REObjectAvailableToDate"`
		ValuationRelevance             string        `json:"ValuationRelevance"`
	} `json:"value"`
}  
