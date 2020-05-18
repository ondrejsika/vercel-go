package vercel

type ListRecordsResponse struct {
	Records []struct {
		ID        string
		Type      string
		Name      string
		Value     string
		Created   int
		Updated   int
		CreatedAt int
		UpdatedAt int
	}
}

type CreateRecordResponse struct {
	UID string
}
