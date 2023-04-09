package desk

type Desk struct {
	ID int `json:"id"`
	TableName string `json:"table_name"`
	IsOccupied bool `json:"is_occupied"`
	CorporateId string `json:"corporate_id"`
}
