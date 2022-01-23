package responses  

type ContractRenewalTerm struct {
	Value             []struct {
		InternalRealEstateNumber       string        `json:"InternalRealEstateNumber"`
		RETermType                     string        `json:"RETermType"`
		RETermNumber                   string        `json:"RETermNumber"`
		RERenewalType                  string        `json:"RERenewalType"`
		RERenewalSequenceNumber        string        `json:"RERenewalSequenceNumber"`
		RERenewalRuleType              string        `json:"RERenewalRuleType"`
		RETermName                     string        `json:"RETermName"`
		RERenewalRule                  string        `json:"RERenewalRule"`
		REAutomaticRenewalType         string        `json:"REAutomaticRenewalType"`
		RENumberOfRenewals             string        `json:"RENumberOfRenewals"`
		RERenewalPeriodInYears         string        `json:"RERenewalPeriodInYears"`
		RERenewalPeriodInMonths        string        `json:"RERenewalPeriodInMonths"`
		RERenewalPeriodInDays          string        `json:"RERenewalPeriodInDays"`
		RERenewalRoundingDateRule      string        `json:"RERenewalRoundingDateRule"`
		RENotificationPeriodInYears    string        `json:"RENotificationPeriodInYears"`
		RENotificationPeriodInMonths   string        `json:"RENotificationPeriodInMonths"`
		RENotificationPeriodInWeeks    string        `json:"RENotificationPeriodInWeeks"`
		RENotificationPeriodInDays     string        `json:"RENotificationPeriodInDays"`
		RENotificationRoundingDateRule string        `json:"RENotificationRoundingDateRule"`
	} `json:"value"`
}  
