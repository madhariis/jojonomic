package models

// Documents ...
type Documents struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`
	FolderID  string                 `json:"folder_id"`
	Content   map[string]interface{} `json:"content"`
	OwnerID   uint64                 `json:"owner_id"`
	Share     []uint64               `json:"share"`
	Timestamp uint64                 `json:"timestamp"`
	CompanyID uint64                 `json:"company_id"`
}
