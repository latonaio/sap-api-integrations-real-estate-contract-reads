package responses

type ContractNoticeTerm struct {
	Value             []struct {
		InternalRealEstateNumber   string        `json:"InternalRealEstateNumber"`
		RETermType                 string        `json:"RETermType"`
		RETermNumber               string        `json:"RETermNumber"`
		RENoticeRule               string        `json:"RENoticeRule"`
		RENoticeSequenceNo         string        `json:"RENoticeSequenceNo"`
		RETermName                 string        `json:"RETermName"`
		RENoticeType               string        `json:"RENoticeType"`
		RENoticeProcedure          string        `json:"RENoticeProcedure"`
		RENoticeGivingParty        string        `json:"RENoticeGivingParty"`
		RENoticeRuleType           string        `json:"RENoticeRuleType"`
		RENoticeRuleDescription    string        `json:"RENoticeRuleDescription"`
		RETermPeriodInYears        string        `json:"RETermPeriodInYears"`
		RETermPeriodInMonths       string        `json:"RETermPeriodInMonths"`
		RETermPeriodInDays         string        `json:"RETermPeriodInDays"`
		REPeriodEndRhythmType      string        `json:"REPeriodEndRhythmType"`
		RENoticePeriodInMonths     string        `json:"RENoticePeriodInMonths"`
		RENoticePeriodInWeeks      string        `json:"RENoticePeriodInWeeks"`
		RENoticePeriodInDays       string        `json:"RENoticePeriodInDays"`
		RENoticeGracePeriod        string        `json:"RENoticeGracePeriod"`
		REGracePeriodCalOrWorkDays string        `json:"REGracePeriodCalOrWorkDays"`
		RENoticePeriodCalendar     string        `json:"RENoticePeriodCalendar"`
		RENoticeYear               string        `json:"RENoticeYear"`
		RENoticeMonth              string        `json:"RENoticeMonth"`
		RENoticeDay                string        `json:"RENoticeDay"`
		RENoticeReceiptYear        string        `json:"RENoticeReceiptYear"`
		RENoticeReceiptMonth       string        `json:"RENoticeReceiptMonth"`
		RENoticeReceiptDay         string        `json:"RENoticeReceiptDay"`
		RENoticeRcptCalOrWorkDays  string        `json:"RENoticeRcptCalOrWorkDays"`
		RENoticeReceiptCalendar    string        `json:"RENoticeReceiptCalendar"`
	} `json:"value"`
}  
