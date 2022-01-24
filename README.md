# sap-api-integrations-real-estate-contract-reads  
sap-api-integrations-real-estate-contract-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 不動産契約 データを取得するマイクロサービスです。  
sap-api-integrations-real-estate-contract-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-real-estate-contract-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_RECONTRACT_0001/overview  

## 動作環境
  sap-api-integrations-real-estate-contract-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
  sap-api-integrations-real-estate-contract-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
  sap-api-integrations-real-estate-contract-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_RECONTRACT_0001/overview  
* APIサービス名(=baseURL): api_realestatecontract/srvd_a2x/sap/ealestatecontract/0001

## 本レポジトリ に 含まれる API名
  sap-api-integrations-real-estate-contract-reads には、次の API をコールするためのリソースが含まれています。  

* A_REContract（SAP 不動産契約 - 契約）
* A_REContrCondition（SAP 不動産契約 - 契約条件）
* A_REContrNoticeTerm（SAP 不動産契約 - 契約通知期間）
* A_REContrObjAssgmt（SAP 不動産契約 - 契約対象割当）
* A_REContrPostingTerm（SAP 不動産契約 - 契約提示期間）
* A_REContrOrglAssgmtTerm（SAP 不動産契約 - 契約割当期間）
* A_REContrReminderDate（SAP 不動産契約 - 契約リマインダー日付）
* A_REContrReminderRule（SAP 不動産契約 - 契約リマインダールール）
* A_REContrRenewalTerm（SAP 不動産契約 - 契約更新期間）
* A_REContrRhythmTerm（SAP 不動産契約 - 契約リズム期間）
* A_REContrValuationCondition（SAP 不動産契約 - 契約評価条件）
* A_REContrValuation（SAP 不動産契約 - 契約評価）

## API への 値入力条件 の 初期値
  sap-api-integrations-real-estate-contract-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## SDC レイアウト

* inoutSDC.RealEstateContract.CompanyCode（会社コード）
* inoutSDC.RealEstateContract.InternalRealEstateNumber（内部不動産契約番号）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Contract" が指定されています。    
  
```
	"api_schema": "A_REContract",
	"accepter": ["Contract"],
	"real_estate_contract_no": "IS00100000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "A_REContract",
	"accepter": ["All"],
	"real_estate_contract_no": "IS00100000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 不動産契約 の 契約データ が取得された結果の JSON の例です。  
以下の項目のうち、"InternalRealEstateNumber" ～ "ValuationRelevance" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-real-estate-contract-reads/SAP_API_Caller/caller.go#L108",
	"function": "sap-api-integrations-real-estate-contract-reads/SAP_API_Caller.(*SAPAPICaller).Contract",
	"level": "INFO",
	"message": [
		{
			"InternalRealEstateNumber": "IS00100000000",
			"CompanyCode": "1010",
			"RealEstateContract": "1200000000000",
			"REStatusObject": "IS10101200000000000",
			"REInternalFinNumber": "00000001",
			"RECreationType": "2",
			"CreationDate": "2017-08-24",
			"CreationTime": "18:19:56",
			"RESourceOfCreation": "RECN",
			"LastChangeDate": "2019-10-22",
			"LastChangeTime": "14:12:32",
			"RESourceOfChange": "",
			"Responsible": "CB9980000036",
			"REUserExclusive": "",
			"REAuthorizationGroup": "",
			"REContractType": "YI50",
			"ContractStartDate": "2019-01-01",
			"ContractEndDate": "2020-01-01",
			"REContractName": "Lease-In for Power Supply Mainz",
			"REContractActivateDate": "2017-08-24",
			"RETenancyLaw": "",
			"REContractNumberOld": "",
			"REMainContractCompanyCode": "",
			"REMainContract": "",
			"REContractCurrency": "EUR",
			"REIndustrySector": "",
			"REIsSalesRelevant": false,
			"REContractDepositType": "",
			"REContractSignDate": "",
			"REContract2SignDate": "",
			"REContractCashFlowDate": "2019-01-01",
			"REContractFirstEndDate": "2020-01-01",
			"REContractNoticeDate": "",
			"REContractNoticeInDate": "",
			"REContractNoticeReason": "0",
			"REContractNoticeActivationDate": "",
			"RECashFlowArchivedToDate": "",
			"RECashFlowLockedToDate": "",
			"RECashFlowPostingFromDate": "",
			"REBusinessEntity": "",
			"REPossessionStartDate": "",
			"REPossessionEndDate": "",
			"REHasMultipleAssignments": false,
			"REObjectAvailableFromDate": "",
			"REObjectAvailableToDate": "9999-12-31",
			"ValuationRelevance": ""
		}
	],
	"time": "2022-01-23T14:04:30.261093+09:00"
}
```
