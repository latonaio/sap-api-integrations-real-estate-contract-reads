package responses

type ContractPostingTerm struct {
	Value             []struct {
		InternalRealEstateNumber  string        `json:"InternalRealEstateNumber"`
		RETermType                string        `json:"RETermType"`
		RETermNumber              string        `json:"RETermNumber"`
		ValidityStartEndDateValue string        `json:"ValidityStartEndDateValue"`
		ValidityStartDate         string        `json:"ValidityStartDate"`
		ValidityEndDate           string        `json:"ValidityEndDate"`
		RETermName                string        `json:"RETermName"`
		PaymentMethod             string        `json:"PaymentMethod"`
		REPaymentMethodCreditMemo string        `json:"REPaymentMethodCreditMemo"`
		PaymentBlockingReason     string        `json:"PaymentBlockingReason"`
		PaymentTerms              string        `json:"PaymentTerms"`
		HouseBank                 string        `json:"HouseBank"`
		HouseBankAccount          string        `json:"HouseBankAccount"`
		BankIdentification        string        `json:"BankIdentification"`
		RENoteToPayeeText         string        `json:"RENoteToPayeeText"`
		DunningArea               string        `json:"DunningArea"`
		DunningKey                string        `json:"DunningKey"`
		DunningBlockingReason     string        `json:"DunningBlockingReason"`
		REAcctDeterminationKey    string        `json:"REAcctDeterminationKey"`
		RETaxType                 string        `json:"RETaxType"`
		TaxGroup                  string        `json:"TaxGroup"`
		REIsConditionGrossAmount  bool          `json:"REIsConditionGrossAmount"`
		TaxCountry                string        `json:"TaxCountry"`
		BusinessPartner           string        `json:"BusinessPartner"`
		REAccountingObject        string        `json:"REAccountingObject"`
		TaxJurisdiction           string        `json:"TaxJurisdiction"`
		REIsConditionSplit        bool          `json:"REIsConditionSplit"`
		RECurrencyTranslationRule string        `json:"RECurrencyTranslationRule"`
		REIsPartnerBlocked        bool          `json:"REIsPartnerBlocked"`
		SEPAMandate               string        `json:"SEPAMandate"`
		SEPAMandateCreditor       string        `json:"SEPAMandateCreditor"`
	} `json:"value"`
}  
