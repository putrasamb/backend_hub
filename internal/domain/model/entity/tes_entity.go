package entity

type Tes struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Umur  int    `json:"umur"`
}

func (Tes) TableName() string {
	return "tes"
}
