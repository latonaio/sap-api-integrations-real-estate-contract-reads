package responses

type ContractObjectAssignment struct {
	Value             []struct {
		REStatusObjectSource           string        `json:"REStatusObjectSource"`
		REObjectAssignmentType         string        `json:"REObjectAssignmentType"`
		REStatusObjectTarget           string        `json:"REStatusObjectTarget"`
		ValidityStartEndDateValue      string        `json:"ValidityStartEndDateValue"`
		InternalRealEstateNumber       string        `json:"InternalRealEstateNumber"`
		ValidityStartDate              string        `json:"ValidityStartDate"`
		ValidityEndDate                string        `json:"ValidityEndDate"`
		REObjectTypeTarget             string        `json:"REObjectTypeTarget"`
		REOnlyInfoAssgmt               bool          `json:"REOnlyInfoAssgmt"`
		REStatusObjectSourceIsArchived bool          `json:"REStatusObjectSourceIsArchived"`
		REGenerationType               string        `json:"REGenerationType"`
		REIsMainAsset                  bool          `json:"REIsMainAsset"`
		REAssignmentHasMultiple        bool          `json:"REAssignmentHasMultiple"`
		REObjectPossessionStartDate    string        `json:"REObjectPossessionStartDate"`
		REObjectPossessionEndDate      string        `json:"REObjectPossessionEndDate"`
		REGroupNumber                  string        `json:"REGroupNumber"`
		REObjectGroupName              string        `json:"REObjectGroupName"`
		REContractSubjectNumber        string        `json:"REContractSubjectNumber"`
		REContractSubjectDescription   string        `json:"REContractSubjectDescription"`
		REContractSubjectClass         string        `json:"REContractSubjectClass"`
		REContractSubjectType          string        `json:"REContractSubjectType"`
		ExternalID                     string        `json:"ExternalId"`
		REAccountingObject             string        `json:"REAccountingObject"`
		REAccountingObjectType         string        `json:"REAccountingObjectType"`
	} `json:"value"`
}  
