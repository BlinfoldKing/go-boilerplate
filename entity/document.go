package entity

import "github.com/satori/uuid"

// Document is filled with information of the file stored in the bucket
type Document struct {
	ID         string `xorm:"id" json:"id"`
	ObjectName string `xorm:"object_name" json:"object_name"`
	BucketName string `xorm:"bucket_name" json:"bucket_name"`
}

type DocumentCreateResponse struct {
	Document Document `json:"document"`
	URL      string   `json:"url"`
}

// NewDocument used to create a new document
func NewDocument(object, bucket string) (document Document, err error) {
	id := uuid.NewV4().String()

	document = Document{
		ID:         id,
		ObjectName: object,
		BucketName: bucket,
	}
	return
}
