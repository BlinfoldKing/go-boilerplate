package documents

// DocumentRequest request for upload and download
type DocumentRequest struct {
	BucketName string `json:"bucket_name" validate:"required"`
	ObjectName string `json:"object_name" validate:"required"`
}
