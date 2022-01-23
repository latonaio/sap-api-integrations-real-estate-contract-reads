package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-real-estate-contract-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToContract(raw []byte, l *logger.Logger) ([]Contract, error) {
	pm := &responses.Contract{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Contract. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contract := make([]Contract, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contract = append(contract, Contract{
    InternalRealEstateNumber:       data.InternalRealEstateNumber,
    CompanyCode:                    data.CompanyCode,
    RealEstateContract:             data.RealEstateContract,
    REStatusObject:                 data.REStatusObject,
    REInternalFinNumber:            data.REInternalFinNumber,
    RECreationType:                 data.RECreationType,
    CreationDate:                   data.CreationDate,
    CreationTime:                   data.CreationTime,
    RESourceOfCreation:             data.RESourceOfCreation,
    LastChangeDate:                 data.LastChangeDate,
    LastChangeTime:                 data.LastChangeTime,
    RESourceOfChange:               data.RESourceOfChange,
    Responsible:                    data.Responsible,
    REUserExclusive:                data.REUserExclusive,
    REAuthorizationGroup:           data.REAuthorizationGroup,
    REContractType:                 data.REContractType,
    ContractStartDate:              data.ContractStartDate,
    ContractEndDate:                data.ContractEndDate,
    REContractName:                 data.REContractName,
    REContractActivateDate:         data.REContractActivateDate,
    RETenancyLaw:                   data.RETenancyLaw,
    REContractNumberOld:            data.REContractNumberOld,
    REMainContractCompanyCode:      data.REMainContractCompanyCode,
    REMainContract:                 data.REMainContract,
    REContractCurrency:             data.REContractCurrency,
    REIndustrySector:               data.REIndustrySector,
    REIsSalesRelevant:              data.REIsSalesRelevant,
    REContractDepositType:          data.REContractDepositType,
    REContractSignDate:             data.REContractSignDate,
    REContract2SignDate:            data.REContract2SignDate,
    REContractCashFlowDate:         data.REContractCashFlowDate,
    REContractFirstEndDate:         data.REContractFirstEndDate,
    REContractNoticeDate:           data.REContractNoticeDate,
    REContractNoticeInDate:         data.REContractNoticeInDate,
    REContractNoticeReason:         data.REContractNoticeReason,
    REContractNoticeActivationDate: data.REContractNoticeActivationDate,
    RECashFlowArchivedToDate:       data.RECashFlowArchivedToDate,
    RECashFlowLockedToDate:         data.RECashFlowLockedToDate,
    RECashFlowPostingFromDate:      data.RECashFlowPostingFromDate,
    REBusinessEntity:               data.REBusinessEntity,
    REPossessionStartDate:          data.REPossessionStartDate,
    REPossessionEndDate:            data.REPossessionEndDate,
    REHasMultipleAssignments:       data.REHasMultipleAssignments,
    REObjectAvailableFromDate:      data.REObjectAvailableFromDate,
    REObjectAvailableToDate:        data.REObjectAvailableToDate,
    ValuationRelevance:             data.ValuationRelevance,
		})
	}

	return contract, nil
}

func ConvertToContractCondition(raw []byte, l *logger.Logger) ([]ContractCondition, error) {
	pm := &responses.ContractCondition{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractCondition. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractCondition := make([]ContractCondition, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractCondition = append(contractCondition, ContractCondition{
    REConditionUUID:                data.REConditionUUID,
    InternalRealEstateNumber:       data.InternalRealEstateNumber,
    REStatusObjectCalculation:      data.REStatusObjectCalculation,
    REStatusObjectIDCalculation:    data.REStatusObjectIDCalculation,
    REConditionType:                data.REConditionType,
    REExtConditionPurpose:          data.REExtConditionPurpose,
    ValidityStartEndDateValue:      data.ValidityStartEndDateValue,
    ValidityStartDate:              data.ValidityStartDate,
    REStatusObjectDistribution:     data.REStatusObjectDistribution,
    REStatusObjectIDDistribution:   data.REStatusObjectIDDistribution,
    REObjectTypeDistribution:       data.REObjectTypeDistribution,
    REObjectTypePosting:            data.REObjectTypePosting,
    ValidityEndDate:                data.ValidityEndDate,
    CreationDateTime:               data.CreationDateTime,
    RESourceOfCreation:             data.RESourceOfCreation,
    LastChangeDateTime:             data.LastChangeDateTime,
    RESourceOfChange:               data.RESourceOfChange,
    REIsOneTimeCondition:           data.REIsOneTimeCondition,
    REConditionIsStatistical:       data.REConditionIsStatistical,
    REPostingTerm:                  data.REPostingTerm,
    RERhythmTerm:                   data.RERhythmTerm,
    REAdjustmentNumber:             data.REAdjustmentNumber,
    REOrglAssignmentTerm:           data.REOrglAssignmentTerm,
    RESalesTerm:                    data.RESalesTerm,
    REPeakSalesTerm:                data.REPeakSalesTerm,
    RESrvcChrgSettlementPostingTrm: data.RESrvcChrgSettlementPostingTrm,
    REWithholdingTaxTerm:           data.REWithholdingTaxTerm,
    RECalculationRule:              data.RECalculationRule,
    REUnitPrice:                    data.REUnitPrice,
    REConditionCurrency:            data.REConditionCurrency,
    RECalculationRuleParam1:        data.RECalculationRuleParam1,
    RECalculationRuleParam2:        data.RECalculationRuleParam2,
    REDistributionRule:             data.REDistributionRule,
    REDistributionRuleParam1:       data.REDistributionRuleParam1,
    REDistributionRuleParam2:       data.REDistributionRuleParam2,
    REReasonForChange:              data.REReasonForChange,
    REStsObjectParamCalculation:    data.REStsObjectParamCalculation,
    REStsObjectParamDistribution:   data.REStsObjectParamDistribution,
		})
	}

	return contractCondition, nil
}

func ConvertToContractNoticeTerm(raw []byte, l *logger.Logger) ([]ContractNoticeTerm, error) {
	pm := &responses.ContractNoticeTerm{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractNoticeTerm. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractNoticeTerm := make([]ContractNoticeTerm, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractNoticeTerm = append(contractNoticeTerm, ContractNoticeTerm{
    InternalRealEstateNumber:   data.InternalRealEstateNumber,
    RETermType:                 data.RETermType,
    RETermNumber:               data.RETermNumber,
    RENoticeRule:               data.RENoticeRule,
    RENoticeSequenceNo:         data.RENoticeSequenceNo,
    RETermName:                 data.RETermName,
    RENoticeType:               data.RENoticeType,
    RENoticeProcedure:          data.RENoticeProcedure,
    RENoticeGivingParty:        data.RENoticeGivingParty,
    RENoticeRuleType:           data.RENoticeRuleType,
    RENoticeRuleDescription:    data.RENoticeRuleDescription,
    RETermPeriodInYears:        data.RETermPeriodInYears,
    RETermPeriodInMonths:       data.RETermPeriodInMonths,
    RETermPeriodInDays:         data.RETermPeriodInDays,
    REPeriodEndRhythmType:      data.REPeriodEndRhythmType,
    RENoticePeriodInMonths:     data.RENoticePeriodInMonths,
    RENoticePeriodInWeeks:      data.RENoticePeriodInWeeks,
    RENoticePeriodInDays:       data.RENoticePeriodInDays,
    RENoticeGracePeriod:        data.RENoticeGracePeriod,
    REGracePeriodCalOrWorkDays: data.REGracePeriodCalOrWorkDays,
    RENoticePeriodCalendar:     data.RENoticePeriodCalendar,
    RENoticeYear:               data.RENoticeYear,
    RENoticeMonth:              data.RENoticeMonth,
    RENoticeDay:                data.RENoticeDay,
    RENoticeReceiptYear:        data.RENoticeReceiptYear,
    RENoticeReceiptMonth:       data.RENoticeReceiptMonth,
    RENoticeReceiptDay:         data.RENoticeReceiptDay,
    RENoticeRcptCalOrWorkDays:  data.RENoticeRcptCalOrWorkDays,
    RENoticeReceiptCalendar:    data.RENoticeReceiptCalendar,
		})
	}

	return contractNoticeTerm, nil
}

func ConvertToContractObjectAssignment(raw []byte, l *logger.Logger) ([]ContractObjectAssignment, error) {
	pm := &responses.ContractObjectAssignment{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractObjectAssignment. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractObjectAssignment := make([]ContractObjectAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractObjectAssignment = append(contractObjectAssignment, ContractObjectAssignment{
    REStatusObjectSource:           data.REStatusObjectSource,
    REObjectAssignmentType:         data.REObjectAssignmentType,
    REStatusObjectTarget:           data.REStatusObjectTarget,
    ValidityStartEndDateValue:      data.ValidityStartEndDateValue,
    InternalRealEstateNumber:       data.InternalRealEstateNumber,
    ValidityStartDate:              data.ValidityStartDate,
    ValidityEndDate:                data.ValidityEndDate,   
    REObjectTypeTarget:             data.REObjectTypeTarget,
    REOnlyInfoAssgmt:               data.REOnlyInfoAssgmt,
    REStatusObjectSourceIsArchived: data.REStatusObjectSourceIsArchived,
    REGenerationType:               data.REGenerationType,
    REIsMainAsset:                  data.REIsMainAsset,
    REAssignmentHasMultiple:        data.REAssignmentHasMultiple,
    REObjectPossessionStartDate:    data.REObjectPossessionStartDate,
    REObjectPossessionEndDate:      data.REObjectPossessionEndDate,
    REGroupNumber:                  data.REGroupNumber,
    REObjectGroupName:              data.REObjectGroupName,
    REContractSubjectNumber:        data.REContractSubjectNumber,
    REContractSubjectDescription:   data.REContractSubjectDescription,
    REContractSubjectClass:         data.REContractSubjectClass,
    REContractSubjectType:          data.REContractSubjectType,
    ExternalID:                     data.ExternalID,
    REAccountingObject:             data.REAccountingObject,
    REAccountingObjectType:         data.REAccountingObjectType,
		})
	}

	return contractObjectAssignment, nil
}

func ConvertToContractAssignmentTerm(raw []byte, l *logger.Logger) ([]ContractAssignmentTerm, error) {
	pm := &responses.ContractAssignmentTerm{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractAssignmentTerm. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractAssignmentTerm := make([]ContractAssignmentTerm, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractAssignmentTerm = append(contractAssignmentTerm, ContractAssignmentTerm{
    InternalRealEstateNumber:  data.InternalRealEstateNumber,
    RETermType:                data.RETermType,
    RETermNumber:              data.RETermNumber,
    ValidityStartEndDateValue: data.ValidityStartEndDateValue,
    ValidityStartDate:         data.ValidityStartDate,
    ValidityEndDate:           data.ValidityEndDate,
    RETermName:                data.RETermName,
    BusinessArea:              data.BusinessArea,
    ProfitCenter:              data.ProfitCenter,
    REStatusObject:            data.REStatusObject,
    TaxJurisdiction:           data.TaxJurisdiction,
    Fund:                      data.Fund,
    FundsCenter:               data.FundsCenter,
    CommitmentItem:            data.CommitmentItem,
    FunctionalArea:            data.FunctionalArea,
    BudgetPeriod:              data.BudgetPeriod,
    ControllingArea:           data.ControllingArea,
    TaxCalculationProcedure:   data.TaxCalculationProcedure,
    FinancialManagementArea:   data.FinancialManagementArea,
		})
	}

	return contractAssignmentTerm, nil
}

func ConvertToContractPostingTerm(raw []byte, l *logger.Logger) ([]ContractPostingTerm, error) {
	pm := &responses.ContractPostingTerm{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractPostingTerm. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractPostingTerm := make([]ContractPostingTerm, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractPostingTerm = append(contractPostingTerm, ContractPostingTerm{
    InternalRealEstateNumber:  data.InternalRealEstateNumber,
    RETermType:                data.RETermType,
    RETermNumber:              data.RETermNumber,
    ValidityStartEndDateValue: data.ValidityStartEndDateValue,
    ValidityStartDate:         data.ValidityStartDate,
    ValidityEndDate:           data.ValidityEndDate,
    RETermName:                data.RETermName,
    PaymentMethod:             data.PaymentMethod,
    REPaymentMethodCreditMemo: data.REPaymentMethodCreditMemo,
    PaymentBlockingReason:     data.PaymentBlockingReason,
    PaymentTerms:              data.PaymentTerms,
    HouseBank:                 data.HouseBank,
    HouseBankAccount:          data.HouseBankAccount,
    BankIdentification:        data.BankIdentification,
    RENoteToPayeeText:         data.RENoteToPayeeText,
    DunningArea:               data.DunningArea,
    DunningKey:                data.DunningKey,
    DunningBlockingReason:     data.DunningBlockingReason,
    REAcctDeterminationKey:    data.REAcctDeterminationKey,
    RETaxType:                 data.RETaxType,
    TaxGroup:                  data.TaxGroup,
    REIsConditionGrossAmount:  data.REIsConditionGrossAmount,
    TaxCountry:                data.TaxCountry,
    BusinessPartner:           data.BusinessPartner,
    REAccountingObject:        data.REAccountingObject,
    TaxJurisdiction:           data.TaxJurisdiction,
    REIsConditionSplit:        data.REIsConditionSplit,
    RECurrencyTranslationRule: data.RECurrencyTranslationRule,
    REIsPartnerBlocked:        data.REIsPartnerBlocked,
    SEPAMandate:               data.SEPAMandate,
    SEPAMandateCreditor:       data.SEPAMandateCreditor,
		})
	}

	return contractPostingTerm, nil
}

func ConvertToContractReminderDate(raw []byte, l *logger.Logger) ([]ContractReminderDate, error) {
	pm := &responses.ContractReminderDate{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractReminderDate. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractReminderDate := make([]ContractReminderDate, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractReminderDate = append(contractReminderDate, ContractReminderDate{
    InternalRealEstateNumber: data.InternalRealEstateNumber,
    REReminderNumber:         data.REReminderNumber,
    REReminderDate:           data.REReminderDate,
    REReminderRule:           data.REReminderRule,
    REReminderReason:         data.REReminderReason,
    CreationDate:             data.CreationDate,
    CreationTime:             data.CreationTime,
    RESourceOfCreation:       data.RESourceOfCreation,
    LastChangeDate:           data.LastChangeDate,
    LastChangeTime:           data.LastChangeTime,
    RESourceOfChange:         data.RESourceOfChange,
    Responsible:              data.Responsible,
    REReminderWrkflwDate:     data.REReminderWrkflwDate,
    REReminderIsDone:         data.REReminderIsDone,
    REReminderIsFix:          data.REReminderIsFix,
    REReminderIsWrkflwSend:   data.REReminderIsWrkflwSend,
    TextObjectKey:            data.TextObjectKey,
    REReminderInfoText:       data.REReminderInfoText,
		})
	}

	return contractReminderDate, nil
}

func ConvertToContractReminderRule(raw []byte, l *logger.Logger) ([]ContractReminderRule, error) {
	pm := &responses.ContractReminderRule{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractReminderRule. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractReminderRule := make([]ContractReminderRule, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractReminderRule = append(contractReminderRule, ContractReminderRule{
    InternalRealEstateNumber:  data.InternalRealEstateNumber,
    REReminderNumber:          data.REReminderNumber,
    REReminderRuleParamNumber: data.REReminderRuleParamNumber,
    REReminderRule:            data.REReminderRule,
    REReminderReason:          data.REReminderReason,
    ValidityStartDate:         data.ValidityStartDate,
    ValidityEndDate:           data.ValidityEndDate,
    REReminderParamType:       data.REReminderParamType,
    REReminderParamDate:       data.REReminderParamDate,
    REReminderParamNmbr:       data.REReminderParamNmbr,
    REReminderParamIsBoolean:  data.REReminderParamIsBoolean,
		})
	}

	return contractReminderRule, nil
}

func ConvertToContractRenewalTerm(raw []byte, l *logger.Logger) ([]ContractRenewalTerm, error) {
	pm := &responses.ContractRenewalTerm{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractRenewalTerm. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractRenewalTerm := make([]ContractRenewalTerm, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractRenewalTerm = append(contractRenewalTerm, ContractRenewalTerm{
    InternalRealEstateNumber:       data.InternalRealEstateNumber,
    RETermType:                     data.RETermType,
    RETermNumber:                   data.RETermNumber,
    RERenewalType:                  data.RERenewalType,
    RERenewalSequenceNumber:        data.RERenewalSequenceNumber,
    RERenewalRuleType:              data.RERenewalRuleType,
    RETermName:                     data.RETermName,
    RERenewalRule:                  data.RERenewalRule,
    REAutomaticRenewalType:         data.REAutomaticRenewalType,
    RENumberOfRenewals:             data.RENumberOfRenewals,
    RERenewalPeriodInYears:         data.RERenewalPeriodInYears,
    RERenewalPeriodInMonths:        data.RERenewalPeriodInMonths,
    RERenewalPeriodInDays:          data.RERenewalPeriodInDays,
    RERenewalRoundingDateRule:      data.RERenewalRoundingDateRule,
    RENotificationPeriodInYears:    data.RENotificationPeriodInYears,
    RENotificationPeriodInMonths:   data.RENotificationPeriodInMonths,
    RENotificationPeriodInWeeks:    data.RENotificationPeriodInWeeks,
    RENotificationPeriodInDays:     data.RENotificationPeriodInDays,
    RENotificationRoundingDateRule: data.RENotificationRoundingDateRule,
		})
	}

	return contractRenewalTerm, nil
}

func ConvertToContractRhythmTerm(raw []byte, l *logger.Logger) ([]ContractRhythmTerm, error) {
	pm := &responses.ContractRhythmTerm{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractRhythmTerm. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractRhythmTerm := make([]ContractRhythmTerm, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractRhythmTerm = append(contractRhythmTerm, ContractRhythmTerm{
    InternalRealEstateNumber:       data.InternalRealEstateNumber,
    RETermType:                     data.RETermType,
    RETermNumber:                   data.RETermNumber,
    ValidityStartEndDateValue:      data.ValidityStartEndDateValue,
    ValidityStartDate:              data.ValidityStartDate,
    ValidityEndDate:                data.ValidityEndDate,
    RETermName:                     data.RETermName,
    RENumberOfFrequencyUnits:       data.RENumberOfFrequencyUnits,
    REFrequencyUnit:                data.REFrequencyUnit,
    REStartFrequencyWeek:           data.REStartFrequencyWeek,
    REFrequencyStart:               data.REFrequencyStart,
    REConditionAmountReference:     data.REConditionAmountReference,
    REConditionAmountDiff:          data.REConditionAmountDiff,
    REProRataMethod:                data.REProRataMethod,
    REProRataMethodCalc:            data.REProRataMethodCalc,
    REPaymentForm:                  data.REPaymentForm,
    REFrequencyStartDate:           data.REFrequencyStartDate,
    REDueDateCorrectionRule:        data.REDueDateCorrectionRule,
    REDueDateNumberOfCrrtnDays:     data.REDueDateNumberOfCrrtnDays,
    REDueDateNumberOfCrrtnMonths:   data.REDueDateNumberOfCrrtnMonths,
    REDueDateNumberOfCrrtnYears:    data.REDueDateNumberOfCrrtnYears,
    REDueDateNumberOfCrrtnCalendar: data.REDueDateNumberOfCrrtnCalendar,
    REDueDateCrrtnCalendarUnit:     data.REDueDateCrrtnCalendarUnit,
    FactoryCalendar:                data.FactoryCalendar,
    REDueDateIsAtBeginning:         data.REDueDateIsAtBeginning,
    REDueDateIsAtEnd:               data.REDueDateIsAtEnd,
    REFixedPeriod:                  data.REFixedPeriod,
		})
	}

	return contractRhythmTerm, nil
}

func ConvertToContractValuationCondition(raw []byte, l *logger.Logger) ([]ContractValuationCondition, error) {
	pm := &responses.ContractValuationCondition{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractValuationCondition. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractValuationCondition := make([]ContractValuationCondition, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractValuationCondition = append(contractValuationCondition, ContractValuationCondition{
    InternalRealEstateNumber:     data.InternalRealEstateNumber,
    RETermNumber:                 data.RETermNumber,
    ValidityStartEndDateValue:    data.ValidityStartEndDateValue,
    REConditionType:              data.REConditionType,
    REConditionValidityStartDate: data.REConditionValidityStartDate,
    REExtConditionPurpose:        data.REExtConditionPurpose,
    REStatusObjectCalculation:    data.REStatusObjectCalculation,
    RETermName:                   data.RETermName,
    ValidityStartDate:            data.ValidityStartDate,
    ValidityEndDate:              data.ValidityEndDate,
    REConditionValidityEndDate:   data.REConditionValidityEndDate,
    REValuationCndnProperty:      data.REValuationCndnProperty,
    REValuationCndnConsdtn:       data.REValuationCndnConsdtn,
    REIsValuationCndnConsdtn:     data.REIsValuationCndnConsdtn,
    REValuationCndnSharePercent:  data.REValuationCndnSharePercent,
    REValuationCndnShareAbsltAmt: data.REValuationCndnShareAbsltAmt,
    REValuationCurrency:          data.REValuationCurrency,
    REValuationCndnStatus:        data.REValuationCndnStatus,
    REInfoText:                   data.REInfoText,
		})
	}

	return contractValuationCondition, nil
}

func ConvertToContractValuation(raw []byte, l *logger.Logger) ([]ContractValuation, error) {
	pm := &responses.ContractValuation{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContractValuation. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	contractValuation := make([]ContractValuation, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		contractValuation = append(contractValuation, ContractValuation{
    InternalRealEstateNumber:     data.InternalRealEstateNumber,
    RETermNumber:                 data.RETermNumber,
    ValidityStartEndDateValue:    data.ValidityStartEndDateValue,
    RETermType:                   data.RETermType,
    RETermName:                   data.RETermName,
    REValuationRule:              data.REValuationRule,
    ValidityStartDate:            data.ValidityStartDate,
    ValidityEndDate:              data.ValidityEndDate,
    REStatusObject:               data.REStatusObject,
    RealEstateObjectType:         data.RealEstateObjectType,
    REIdentification:             data.REIdentification,
    ObjectValidFrom:              data.ObjectValidFrom,
    REConsiderationStartDate:     data.REConsiderationStartDate,
    RECashFlowPostingFromDate:    data.RECashFlowPostingFromDate,
    REStatusObjectAsset:          data.REStatusObjectAsset,
    REValuationClassification:    data.REValuationClassification,
    REInterestRate:               data.REInterestRate,
    REFrequencyTerm:              data.REFrequencyTerm,
    REDistributionRule:           data.REDistributionRule,
    REDistributionRuleParam1:     data.REDistributionRuleParam1,
    REDistributionRuleParam2:     data.REDistributionRuleParam2,
    REProbableEndDate:            data.REProbableEndDate,
    REAssetRightOfUseEndDate:     data.REAssetRightOfUseEndDate,
    REValuationRuleStatus:        data.REValuationRuleStatus,
    REValuationStatus:            data.REValuationStatus,
    REValuationStatusReason:      data.REValuationStatusReason,
    REValuationBehavior:          data.REValuationBehavior,
    RETaxType:                    data.RETaxType,
    TaxGroup:                     data.TaxGroup,
    REAccountingObject:           data.REAccountingObject,
    REInfoText:                   data.REInfoText,
    REValuationFactorNumerator:   data.REValuationFactorNumerator,
    REValuationFactorDenominator: data.REValuationFactorDenominator,
    REValuationCurrency:          data.REValuationCurrency,
    REValuationQuestionnaireUUID: data.REValuationQuestionnaireUUID,
    Country:                      data.Country,
		})
	}

	return contractValuation, nil
}

