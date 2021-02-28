package models

type Folder struct {
	ID 			string		`json:"id"`
	Name 		string		`json:"name"`
	Type 		string		`json:"type"`
	IsPublic 	bool		`json:"is_public"`
	Timestamp 	uint64		`json:"timestamp"`
	OwnerID 	uint64		`json:"owner_id"`
	Share 		[]uint64	`json:"share"`
	CompanyID 	uint64 		`json:"company_id"`
}
