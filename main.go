package main

import (
	sap_api_caller "sap-api-integrations-real-estate-contract-reads/SAP_API_Caller"
	"sap-api-integrations-real-estate-contract-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Real_Estate_Contract_Contract_Valuation_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Contract", "ContractCondition", "ContractNoticeTerm", "ContractObjectAssignment", "ContractAssignmentTerm",
			"ContractPostingTerm", "ContractReminderDate", "ContractReminderRule", "ContractRenewalTerm", "ContractRhythmTerm",
			"ContractValuationCondition", "ContractValuation",
		}
	}

	caller.AsyncGetRealEstateContract(
		inoutSDC.RealEstateContract.CompanyCode,
		inoutSDC.RealEstateContract.InternalRealEstateNumber,
		accepter,
	)
}
