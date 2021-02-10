package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Document is filled with information of the file stored in the bucket
type Document struct {
	ID         string     `xorm:"id" json:"id"`
	Name       string     `xorm:"name" json:"name"`
	FileSize   int        `xorm:"file_size" json:"file_size"`
	FileType   string     `xorm:"file_type" json:"file_type"`
	ObjectName string     `xorm:"object_name" json:"object_name"`
	BucketName string     `xorm:"bucket_name" json:"bucket_name"`
	URLLink    string     `xorm:"url_link" json:"url_link"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// DocumentChangeSet change set for document
type DocumentChangeSet struct {
	Name       string `xorm:"name" json:"name"`
	FileSize   int    `xorm:"file_size" json:"file_size"`
	FileType   string `xorm:"file_type" json:"file_type"`
	ObjectName string `xorm:"object_name" json:"object_name"`
	BucketName string `xorm:"bucket_name" json:"bucket_name"`
	URLLink    string `xorm:"url_link" json:"url_link"`
}

// NewDocument used to create a new document
func NewDocument(name, objectName, bucketName, fileType string, fileSize int) (document Document, err error) {
	id := uuid.NewV4().String()

	document = Document{
		ID:         id,
		Name:       name,
		FileSize:   fileSize,
		FileType:   fileType,
		ObjectName: objectName,
		BucketName: bucketName,
	}
	return
}
