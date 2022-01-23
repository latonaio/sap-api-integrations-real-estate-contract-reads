package responses  

type ContractRhythmTerm struct {
	Value             []struct {
		InternalRealEstateNumber       string        `json:"InternalRealEstateNumber"`
		RETermType                     string        `json:"RETermType"`
		RETermNumber                   string        `json:"RETermNumber"`
		ValidityStartEndDateValue      string        `json:"ValidityStartEndDateValue"`
		ValidityStartDate              string        `json:"ValidityStartDate"`
		ValidityEndDate                string        `json:"ValidityEndDate"`
		RETermName                     string        `json:"RETermName"`
		RENumberOfFrequencyUnits       string        `json:"RENumberOfFrequencyUnits"`
		REFrequencyUnit                string        `json:"REFrequencyUnit"`
		REStartFrequencyWeek           string        `json:"REStartFrequencyWeek"`
		REFrequencyStart               string        `json:"REFrequencyStart"`
		REConditionAmountReference     string        `json:"REConditionAmountReference"`
		REConditionAmountDiff          string        `json:"REConditionAmountDiff"`
		REProRataMethod                string        `json:"REProRataMethod"`
		REProRataMethodCalc            string        `json:"REProRataMethodCalc"`
		REPaymentForm                  string        `json:"REPaymentForm"`
		REFrequencyStartDate           string        `json:"REFrequencyStartDate"`
		REDueDateCorrectionRule        string        `json:"REDueDateCorrectionRule"`
		REDueDateNumberOfCrrtnDays     int           `json:"REDueDateNumberOfCrrtnDays"`
		REDueDateNumberOfCrrtnMonths   int           `json:"REDueDateNumberOfCrrtnMonths"`
		REDueDateNumberOfCrrtnYears    int           `json:"REDueDateNumberOfCrrtnYears"`
		REDueDateNumberOfCrrtnCalendar int           `json:"REDueDateNumberOfCrrtnCalendar"`
		REDueDateCrrtnCalendarUnit     string        `json:"REDueDateCrrtnCalendarUnit"`
		FactoryCalendar                string        `json:"FactoryCalendar"`
		REDueDateIsAtBeginning         bool          `json:"REDueDateIsAtBeginning"`
		REDueDateIsAtEnd               bool          `json:"REDueDateIsAtEnd"`
		REFixedPeriod                  string        `json:"REFixedPeriod"`
	} `json:"value"`
}  
