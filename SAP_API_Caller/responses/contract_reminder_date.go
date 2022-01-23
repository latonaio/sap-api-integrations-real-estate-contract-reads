package responses  

type ContractReminderDate struct {
	Value             []struct {
		InternalRealEstateNumber string      `json:"InternalRealEstateNumber"`
		REReminderNumber         string      `json:"REReminderNumber"`
		REReminderDate           string      `json:"REReminderDate"`
		REReminderRule           string      `json:"REReminderRule"`
		REReminderReason         string      `json:"REReminderReason"`
		CreationDate             string      `json:"CreationDate"`
		CreationTime             string      `json:"CreationTime"`
		RESourceOfCreation       string      `json:"RESourceOfCreation"`
		LastChangeDate           string      `json:"LastChangeDate"`
		LastChangeTime           string      `json:"LastChangeTime"`
		RESourceOfChange         string      `json:"RESourceOfChange"`
		Responsible              string      `json:"Responsible"`
		REReminderWrkflwDate     string      `json:"REReminderWrkflwDate"`
		REReminderIsDone         bool        `json:"REReminderIsDone"`
		REReminderIsFix          bool        `json:"REReminderIsFix"`
		REReminderIsWrkflwSend   bool        `json:"REReminderIsWrkflwSend"`
		TextObjectKey            string      `json:"TextObjectKey"`
		REReminderInfoText       string      `json:"REReminderInfoText"`
	} `json:"value"`
}  
