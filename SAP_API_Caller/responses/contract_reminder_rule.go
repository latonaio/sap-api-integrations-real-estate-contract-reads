package responses  

type ContractReminderRule struct {
	Value             []struct {
		InternalRealEstateNumber  string        `json:"InternalRealEstateNumber"`
		REReminderNumber          string        `json:"REReminderNumber"`
		REReminderRuleParamNumber string        `json:"REReminderRuleParamNumber"`
		REReminderRule            string        `json:"REReminderRule"`
		REReminderReason          string        `json:"REReminderReason"`
		ValidityStartDate         string        `json:"ValidityStartDate"`
		ValidityEndDate           string        `json:"ValidityEndDate"`
		REReminderParamType       string        `json:"REReminderParamType"`
		REReminderParamDate       string        `json:"REReminderParamDate"`
		REReminderParamNmbr       string        `json:"REReminderParamNmbr"`
		REReminderParamIsBoolean  bool          `json:"REReminderParamIsBoolean"`
	} `json:"value"`
}  
