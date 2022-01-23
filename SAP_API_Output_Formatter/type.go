package sap_api_output_formatter

type Contract struct {
	InternalRealEstateNumber       string `json:"InternalRealEstateNumber"`
	CompanyCode                    string `json:"CompanyCode"`
	RealEstateContract             string `json:"RealEstateContract"`
	REStatusObject                 string `json:"REStatusObject"`
	REInternalFinNumber            string `json:"REInternalFinNumber"`
	RECreationType                 string `json:"RECreationType"`
	CreationDate                   string `json:"CreationDate"`
	CreationTime                   string `json:"CreationTime"`
	RESourceOfCreation             string `json:"RESourceOfCreation"`
	LastChangeDate                 string `json:"LastChangeDate"`
	LastChangeTime                 string `json:"LastChangeTime"`
	RESourceOfChange               string `json:"RESourceOfChange"`
	Responsible                    string `json:"Responsible"`
	REUserExclusive                string `json:"REUserExclusive"`
	REAuthorizationGroup           string `json:"REAuthorizationGroup"`
	REContractType                 string `json:"REContractType"`
	ContractStartDate              string `json:"ContractStartDate"`
	ContractEndDate                string `json:"ContractEndDate"`
	REContractName                 string `json:"REContractName"`
	REContractActivateDate         string `json:"REContractActivateDate"`
	RETenancyLaw                   string `json:"RETenancyLaw"`
	REContractNumberOld            string `json:"REContractNumberOld"`
	REMainContractCompanyCode      string `json:"REMainContractCompanyCode"`
	REMainContract                 string `json:"REMainContract"`
	REContractCurrency             string `json:"REContractCurrency"`
	REIndustrySector               string `json:"REIndustrySector"`
	REIsSalesRelevant              bool   `json:"REIsSalesRelevant"`
	REContractDepositType          string `json:"REContractDepositType"`
	REContractSignDate             string `json:"REContractSignDate"`
	REContract2SignDate            string `json:"REContract2SignDate"`
	REContractCashFlowDate         string `json:"REContractCashFlowDate"`
	REContractFirstEndDate         string `json:"REContractFirstEndDate"`
	REContractNoticeDate           string `json:"REContractNoticeDate"`
	REContractNoticeInDate         string `json:"REContractNoticeInDate"`
	REContractNoticeReason         string `json:"REContractNoticeReason"`
	REContractNoticeActivationDate string `json:"REContractNoticeActivationDate"`
	RECashFlowArchivedToDate       string `json:"RECashFlowArchivedToDate"`
	RECashFlowLockedToDate         string `json:"RECashFlowLockedToDate"`
	RECashFlowPostingFromDate      string `json:"RECashFlowPostingFromDate"`
	REBusinessEntity               string `json:"REBusinessEntity"`
	REPossessionStartDate          string `json:"REPossessionStartDate"`
	REPossessionEndDate            string `json:"REPossessionEndDate"`
	REHasMultipleAssignments       bool   `json:"REHasMultipleAssignments"`
	REObjectAvailableFromDate      string `json:"REObjectAvailableFromDate"`
	REObjectAvailableToDate        string `json:"REObjectAvailableToDate"`
	ValuationRelevance             string `json:"ValuationRelevance"`
}

type ContractCondition struct {
	REConditionUUID                string  `json:"REConditionUUID"`
	InternalRealEstateNumber       string  `json:"InternalRealEstateNumber"`
	REStatusObjectCalculation      string  `json:"REStatusObjectCalculation"`
	REStatusObjectIDCalculation    string  `json:"REStatusObjectIDCalculation"`
	REConditionType                string  `json:"REConditionType"`
	REExtConditionPurpose          string  `json:"REExtConditionPurpose"`
	ValidityStartEndDateValue      string  `json:"ValidityStartEndDateValue"`
	ValidityStartDate              string  `json:"ValidityStartDate"`
	REStatusObjectDistribution     string  `json:"REStatusObjectDistribution"`
	REStatusObjectIDDistribution   string  `json:"REStatusObjectIDDistribution"`
	REObjectTypeDistribution       string  `json:"REObjectTypeDistribution"`
	REObjectTypePosting            string  `json:"REObjectTypePosting"`
	ValidityEndDate                string  `json:"ValidityEndDate"`
	CreationDateTime               string  `json:"CreationDateTime"`
	RESourceOfCreation             string  `json:"RESourceOfCreation"`
	LastChangeDateTime             string  `json:"LastChangeDateTime"`
	RESourceOfChange               string  `json:"RESourceOfChange"`
	REIsOneTimeCondition           bool    `json:"REIsOneTimeCondition"`
	REConditionIsStatistical       bool    `json:"REConditionIsStatistical"`
	REPostingTerm                  string  `json:"REPostingTerm"`
	RERhythmTerm                   string  `json:"RERhythmTerm"`
	REAdjustmentNumber             string  `json:"REAdjustmentNumber"`
	REOrglAssignmentTerm           string  `json:"REOrglAssignmentTerm"`
	RESalesTerm                    string  `json:"RESalesTerm"`
	REPeakSalesTerm                string  `json:"REPeakSalesTerm"`
	RESrvcChrgSettlementPostingTrm string  `json:"RESrvcChrgSettlementPostingTrm"`
	REWithholdingTaxTerm           string  `json:"REWithholdingTaxTerm"`
	RECalculationRule              string  `json:"RECalculationRule"`
	REUnitPrice                    float64 `json:"REUnitPrice"`
	REConditionCurrency            string  `json:"REConditionCurrency"`
	RECalculationRuleParam1        string  `json:"RECalculationRuleParam1"`
	RECalculationRuleParam2        string  `json:"RECalculationRuleParam2"`
	REDistributionRule             string  `json:"REDistributionRule"`
	REDistributionRuleParam1       string  `json:"REDistributionRuleParam1"`
	REDistributionRuleParam2       string  `json:"REDistributionRuleParam2"`
	REReasonForChange              string  `json:"REReasonForChange"`
	REStsObjectParamCalculation    string  `json:"REStsObjectParamCalculation"`
	REStsObjectParamDistribution   string  `json:"REStsObjectParamDistribution"`
}

type ContractNoticeTerm struct {
	InternalRealEstateNumber   string `json:"InternalRealEstateNumber"`
	RETermType                 string `json:"RETermType"`
	RETermNumber               string `json:"RETermNumber"`
	RENoticeRule               string `json:"RENoticeRule"`
	RENoticeSequenceNo         string `json:"RENoticeSequenceNo"`
	RETermName                 string `json:"RETermName"`
	RENoticeType               string `json:"RENoticeType"`
	RENoticeProcedure          string `json:"RENoticeProcedure"`
	RENoticeGivingParty        string `json:"RENoticeGivingParty"`
	RENoticeRuleType           string `json:"RENoticeRuleType"`
	RENoticeRuleDescription    string `json:"RENoticeRuleDescription"`
	RETermPeriodInYears        string `json:"RETermPeriodInYears"`
	RETermPeriodInMonths       string `json:"RETermPeriodInMonths"`
	RETermPeriodInDays         string `json:"RETermPeriodInDays"`
	REPeriodEndRhythmType      string `json:"REPeriodEndRhythmType"`
	RENoticePeriodInMonths     string `json:"RENoticePeriodInMonths"`
	RENoticePeriodInWeeks      string `json:"RENoticePeriodInWeeks"`
	RENoticePeriodInDays       string `json:"RENoticePeriodInDays"`
	RENoticeGracePeriod        string `json:"RENoticeGracePeriod"`
	REGracePeriodCalOrWorkDays string `json:"REGracePeriodCalOrWorkDays"`
	RENoticePeriodCalendar     string `json:"RENoticePeriodCalendar"`
	RENoticeYear               string `json:"RENoticeYear"`
	RENoticeMonth              string `json:"RENoticeMonth"`
	RENoticeDay                string `json:"RENoticeDay"`
	RENoticeReceiptYear        string `json:"RENoticeReceiptYear"`
	RENoticeReceiptMonth       string `json:"RENoticeReceiptMonth"`
	RENoticeReceiptDay         string `json:"RENoticeReceiptDay"`
	RENoticeRcptCalOrWorkDays  string `json:"RENoticeRcptCalOrWorkDays"`
	RENoticeReceiptCalendar    string `json:"RENoticeReceiptCalendar"`
}

type ContractObjectAssignment struct {
	REStatusObjectSource           string `json:"REStatusObjectSource"`
	REObjectAssignmentType         string `json:"REObjectAssignmentType"`
	REStatusObjectTarget           string `json:"REStatusObjectTarget"`
	ValidityStartEndDateValue      string `json:"ValidityStartEndDateValue"`
	InternalRealEstateNumber       string `json:"InternalRealEstateNumber"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	REObjectTypeTarget             string `json:"REObjectTypeTarget"`
	REOnlyInfoAssgmt               bool   `json:"REOnlyInfoAssgmt"`
	REStatusObjectSourceIsArchived bool   `json:"REStatusObjectSourceIsArchived"`
	REGenerationType               string `json:"REGenerationType"`
	REIsMainAsset                  bool   `json:"REIsMainAsset"`
	REAssignmentHasMultiple        bool   `json:"REAssignmentHasMultiple"`
	REObjectPossessionStartDate    string `json:"REObjectPossessionStartDate"`
	REObjectPossessionEndDate      string `json:"REObjectPossessionEndDate"`
	REGroupNumber                  string `json:"REGroupNumber"`
	REObjectGroupName              string `json:"REObjectGroupName"`
	REContractSubjectNumber        string `json:"REContractSubjectNumber"`
	REContractSubjectDescription   string `json:"REContractSubjectDescription"`
	REContractSubjectClass         string `json:"REContractSubjectClass"`
	REContractSubjectType          string `json:"REContractSubjectType"`
	ExternalID                     string `json:"ExternalId"`
	REAccountingObject             string `json:"REAccountingObject"`
	REAccountingObjectType         string `json:"REAccountingObjectType"`
}

type ContractAssignmentTerm struct {
	InternalRealEstateNumber  string `json:"InternalRealEstateNumber"`
	RETermType                string `json:"RETermType"`
	RETermNumber              string `json:"RETermNumber"`
	ValidityStartEndDateValue string `json:"ValidityStartEndDateValue"`
	ValidityStartDate         string `json:"ValidityStartDate"`
	ValidityEndDate           string `json:"ValidityEndDate"`
	RETermName                string `json:"RETermName"`
	BusinessArea              string `json:"BusinessArea"`
	ProfitCenter              string `json:"ProfitCenter"`
	REStatusObject            string `json:"REStatusObject"`
	TaxJurisdiction           string `json:"TaxJurisdiction"`
	Fund                      string `json:"Fund"`
	FundsCenter               string `json:"FundsCenter"`
	CommitmentItem            string `json:"CommitmentItem"`
	FunctionalArea            string `json:"FunctionalArea"`
	BudgetPeriod              string `json:"BudgetPeriod"`
	ControllingArea           string `json:"ControllingArea"`
	TaxCalculationProcedure   string `json:"TaxCalculationProcedure"`
	FinancialManagementArea   string `json:"FinancialManagementArea"`
}

type ContractPostingTerm struct {
	InternalRealEstateNumber  string `json:"InternalRealEstateNumber"`
	RETermType                string `json:"RETermType"`
	RETermNumber              string `json:"RETermNumber"`
	ValidityStartEndDateValue string `json:"ValidityStartEndDateValue"`
	ValidityStartDate         string `json:"ValidityStartDate"`
	ValidityEndDate           string `json:"ValidityEndDate"`
	RETermName                string `json:"RETermName"`
	PaymentMethod             string `json:"PaymentMethod"`
	REPaymentMethodCreditMemo string `json:"REPaymentMethodCreditMemo"`
	PaymentBlockingReason     string `json:"PaymentBlockingReason"`
	PaymentTerms              string `json:"PaymentTerms"`
	HouseBank                 string `json:"HouseBank"`
	HouseBankAccount          string `json:"HouseBankAccount"`
	BankIdentification        string `json:"BankIdentification"`
	RENoteToPayeeText         string `json:"RENoteToPayeeText"`
	DunningArea               string `json:"DunningArea"`
	DunningKey                string `json:"DunningKey"`
	DunningBlockingReason     string `json:"DunningBlockingReason"`
	REAcctDeterminationKey    string `json:"REAcctDeterminationKey"`
	RETaxType                 string `json:"RETaxType"`
	TaxGroup                  string `json:"TaxGroup"`
	REIsConditionGrossAmount  bool   `json:"REIsConditionGrossAmount"`
	TaxCountry                string `json:"TaxCountry"`
	BusinessPartner           string `json:"BusinessPartner"`
	REAccountingObject        string `json:"REAccountingObject"`
	TaxJurisdiction           string `json:"TaxJurisdiction"`
	REIsConditionSplit        bool   `json:"REIsConditionSplit"`
	RECurrencyTranslationRule string `json:"RECurrencyTranslationRule"`
	REIsPartnerBlocked        bool   `json:"REIsPartnerBlocked"`
	SEPAMandate               string `json:"SEPAMandate"`
	SEPAMandateCreditor       string `json:"SEPAMandateCreditor"`
}

type ContractReminderDate struct {
	InternalRealEstateNumber string `json:"InternalRealEstateNumber"`
	REReminderNumber         string `json:"REReminderNumber"`
	REReminderDate           string `json:"REReminderDate"`
	REReminderRule           string `json:"REReminderRule"`
	REReminderReason         string `json:"REReminderReason"`
	CreationDate             string `json:"CreationDate"`
	CreationTime             string `json:"CreationTime"`
	RESourceOfCreation       string `json:"RESourceOfCreation"`
	LastChangeDate           string `json:"LastChangeDate"`
	LastChangeTime           string `json:"LastChangeTime"`
	RESourceOfChange         string `json:"RESourceOfChange"`
	Responsible              string `json:"Responsible"`
	REReminderWrkflwDate     string `json:"REReminderWrkflwDate"`
	REReminderIsDone         bool   `json:"REReminderIsDone"`
	REReminderIsFix          bool   `json:"REReminderIsFix"`
	REReminderIsWrkflwSend   bool   `json:"REReminderIsWrkflwSend"`
	TextObjectKey            string `json:"TextObjectKey"`
	REReminderInfoText       string `json:"REReminderInfoText"`
}

type ContractReminderRule struct {
	InternalRealEstateNumber  string `json:"InternalRealEstateNumber"`
	REReminderNumber          string `json:"REReminderNumber"`
	REReminderRuleParamNumber string `json:"REReminderRuleParamNumber"`
	REReminderRule            string `json:"REReminderRule"`
	REReminderReason          string `json:"REReminderReason"`
	ValidityStartDate         string `json:"ValidityStartDate"`
	ValidityEndDate           string `json:"ValidityEndDate"`
	REReminderParamType       string `json:"REReminderParamType"`
	REReminderParamDate       string `json:"REReminderParamDate"`
	REReminderParamNmbr       string `json:"REReminderParamNmbr"`
	REReminderParamIsBoolean  bool   `json:"REReminderParamIsBoolean"`
}

type ContractRenewalTerm struct {
	InternalRealEstateNumber       string `json:"InternalRealEstateNumber"`
	RETermType                     string `json:"RETermType"`
	RETermNumber                   string `json:"RETermNumber"`
	RERenewalType                  string `json:"RERenewalType"`
	RERenewalSequenceNumber        string `json:"RERenewalSequenceNumber"`
	RERenewalRuleType              string `json:"RERenewalRuleType"`
	RETermName                     string `json:"RETermName"`
	RERenewalRule                  string `json:"RERenewalRule"`
	REAutomaticRenewalType         string `json:"REAutomaticRenewalType"`
	RENumberOfRenewals             string `json:"RENumberOfRenewals"`
	RERenewalPeriodInYears         string `json:"RERenewalPeriodInYears"`
	RERenewalPeriodInMonths        string `json:"RERenewalPeriodInMonths"`
	RERenewalPeriodInDays          string `json:"RERenewalPeriodInDays"`
	RERenewalRoundingDateRule      string `json:"RERenewalRoundingDateRule"`
	RENotificationPeriodInYears    string `json:"RENotificationPeriodInYears"`
	RENotificationPeriodInMonths   string `json:"RENotificationPeriodInMonths"`
	RENotificationPeriodInWeeks    string `json:"RENotificationPeriodInWeeks"`
	RENotificationPeriodInDays     string `json:"RENotificationPeriodInDays"`
	RENotificationRoundingDateRule string `json:"RENotificationRoundingDateRule"`
}

type ContractRhythmTerm struct {
	InternalRealEstateNumber       string `json:"InternalRealEstateNumber"`
	RETermType                     string `json:"RETermType"`
	RETermNumber                   string `json:"RETermNumber"`
	ValidityStartEndDateValue      string `json:"ValidityStartEndDateValue"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	RETermName                     string `json:"RETermName"`
	RENumberOfFrequencyUnits       string `json:"RENumberOfFrequencyUnits"`
	REFrequencyUnit                string `json:"REFrequencyUnit"`
	REStartFrequencyWeek           string `json:"REStartFrequencyWeek"`
	REFrequencyStart               string `json:"REFrequencyStart"`
	REConditionAmountReference     string `json:"REConditionAmountReference"`
	REConditionAmountDiff          string `json:"REConditionAmountDiff"`
	REProRataMethod                string `json:"REProRataMethod"`
	REProRataMethodCalc            string `json:"REProRataMethodCalc"`
	REPaymentForm                  string `json:"REPaymentForm"`
	REFrequencyStartDate           string `json:"REFrequencyStartDate"`
	REDueDateCorrectionRule        string `json:"REDueDateCorrectionRule"`
	REDueDateNumberOfCrrtnDays     int    `json:"REDueDateNumberOfCrrtnDays"`
	REDueDateNumberOfCrrtnMonths   int    `json:"REDueDateNumberOfCrrtnMonths"`
	REDueDateNumberOfCrrtnYears    int    `json:"REDueDateNumberOfCrrtnYears"`
	REDueDateNumberOfCrrtnCalendar int    `json:"REDueDateNumberOfCrrtnCalendar"`
	REDueDateCrrtnCalendarUnit     string `json:"REDueDateCrrtnCalendarUnit"`
	FactoryCalendar                string `json:"FactoryCalendar"`
	REDueDateIsAtBeginning         bool   `json:"REDueDateIsAtBeginning"`
	REDueDateIsAtEnd               bool   `json:"REDueDateIsAtEnd"`
	REFixedPeriod                  string `json:"REFixedPeriod"`
}

type ContractValuationCondition struct {
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
}

type ContractValuation struct {
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
}
