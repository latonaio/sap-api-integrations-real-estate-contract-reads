package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-real-estate-contract-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetRealEstateContract(companyCode, internalRealEstateNumber string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Contract":
			func() {
				c.Contract(companyCode, internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractCondition":
			func() {
				c.ContractCondition(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractNoticeTerm":
			func() {
				c.ContractNoticeTerm(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractObjectAssignment":
			func() {
				c.ContractObjectAssignment(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractAssignmentTerm":
			func() {
				c.ContractAssignmentTerm(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractPostingTerm":
			func() {
				c.ContractPostingTerm(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractReminderDate":
			func() {
				c.ContractReminderDate(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractReminderRule":
			func() {
				c.ContractReminderRule(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractRenewalTerm":
			func() {
				c.ContractRenewalTerm(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractRhythmTerm":
			func() {
				c.ContractRhythmTerm(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractValuationCondition":
			func() {
				c.ContractValuationCondition(internalRealEstateNumber)
				wg.Done()
			}()
		case "ContractValuation":
			func() {
				c.ContractValuation(internalRealEstateNumber)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Contract(companyCode, internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContract("A_REContract", companyCode, internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContract(api, companyCode, internalRealEstateNumber string) ([]sap_api_output_formatter.Contract, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContract(req, companyCode, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContract(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractCondition(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractCondition("A_REContrCondition", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractCondition(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractCondition, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractCondition(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractCondition(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractNoticeTerm(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractNoticeTerm("A_REContrNoticeTerm", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractNoticeTerm(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractNoticeTerm, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractNoticeTerm(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractNoticeTerm(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractObjectAssignment(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractObjectAssignment("A_REContrObjAssgmt", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractObjectAssignment(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractObjectAssignment, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractObjectAssignment(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractObjectAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractAssignmentTerm(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractAssignmentTerm("A_REContrOrglAssgmtTerm", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractAssignmentTerm(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractAssignmentTerm, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractAssignmentTerm(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractAssignmentTerm(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractPostingTerm(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractPostingTerm("A_REContrPostingTerm", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractPostingTerm(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractPostingTerm, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractPostingTerm(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractPostingTerm(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractReminderDate(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractReminderDate("A_REContrReminderDate", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractReminderDate(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractReminderDate, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithOperationContractReminderDate(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractReminderDate(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractReminderRule(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractReminderRule("A_REContrReminderRule", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractReminderRule(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractReminderRule, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractReminderRule(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractReminderRule(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractRenewalTerm(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractRenewalTerm("A_REContrRenewalTerm", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractRenewalTerm(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractRenewalTerm, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractRenewalTerm(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractRenewalTerm(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractRhythmTerm(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractRhythmTerm("A_REContrRhythmTerm", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractRhythmTerm(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractRhythmTerm, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractRhythmTerm(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractRhythmTerm(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractValuationCondition(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractValuationCondition("A_REContrValuationCondition", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractValuationCondition(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractValuationCondition, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractValuationCondition(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractValuationCondition(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContractValuation(internalRealEstateNumber string) {
	data, err := c.callRealEstateContractSrvAPIRequirementContractValuation("A_REContrValuation", internalRealEstateNumber)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callRealEstateContractSrvAPIRequirementContractValuation(api, internalRealEstateNumber string) ([]sap_api_output_formatter.ContractValuation, error) {
	url := strings.Join([]string{c.baseURL, "api_real_estate_contract/srvd_a2x/sap/api_recontract/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContractValuation(req, internalRealEstateNumber)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContractValuation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithContract(req *http.Request, companyCode, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CompanyCode eq '%s' and InternalRealEstateNumber eq '%s'", companyCode, internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractCondition(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractNoticeTerm(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractObjectAssignment(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractAssignmentTerm(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractPostingTerm(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithOperationContractReminderDate(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractReminderRule(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractRenewalTerm(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractRhythmTerm(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractValuationCondition(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContractValuation(req *http.Request, internalRealEstateNumber string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InternalRealEstateNumber eq '%s'", internalRealEstateNumber))
	req.URL.RawQuery = params.Encode()
}
