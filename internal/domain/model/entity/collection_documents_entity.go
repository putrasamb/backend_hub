package entity

import "time"

type CollectionDocument struct {
	ID              int        `json:"id"`
	DocumentType    string     `json:"document_type"`
	ReferenceNo     string     `json:"reference_no"`
	CustomerID      string     `json:"customer_id"`
	CustomerName    string     `json:"customer_name"`
	BillToID        string     `json:"bill_to_id"`
	InvoiceAmount   string     `json:"invoice_amount"`
	TukarFakturDate string     `json:"tukar_faktur_date"`
	CollectionDate  string     `json:"collection_date"`
	InvoiceDueDate  string     `json:"invoice_due_date"`
	CreatedBy       string     `json:"created_by"`
	UpdatedBy       *string    `json:"updated_by"`
	DeletedBy       *string    `json:"deleted_by,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (CollectionDocument) TableName() string {
	return "t_collection_documents"
}
