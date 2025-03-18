package converter

import (
	"backend_hub/internal/adapter/dto/response"
	"backend_hub/internal/domain/model/entity"
)

func ConvertCollectionDocumentResponse(data *entity.CollectionDocument) *response.CollectionDocument {
	if data == nil {
		return nil
	}

	// Handle nil pointers safely using default values
	return &response.CollectionDocument{
		ID:              data.ID,
		DocumentType:    data.DocumentType,
		ReferenceNo:     data.ReferenceNo,
		CustomerID:      data.CustomerID,
		CustomerName:    data.CustomerName,
		BillToID:        data.BillToID,
		InvoiceAmount:   data.InvoiceAmount,
		TukarFakturDate: data.TukarFakturDate,
		CollectionDate:  data.CollectionDate,
		InvoiceDueDate:  data.InvoiceDueDate,
	}
}

func ConvertCollectionDocumentListResponse(l *[]entity.CollectionDocument) *[]*response.CollectionDocument {
	if l == nil {
		return nil
	}

	dtos := make([]*response.CollectionDocument, 0, len(*l))
	for _, data := range *l {
		dtos = append(dtos, ConvertCollectionDocumentResponse(&data))
	}
	return &dtos
}
