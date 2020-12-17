package documents

// DocumentUploadRequest request for upload and download
type DocumentUploadRequest struct {
	BucketName string `json:"bucket_name" validate:"required"`
	ObjectName string `json:"object_name" validate:"required"`
}

// DocumentDownloadRequest request for upload and download
type DocumentDownloadRequest struct {
	BucketName string `json:"bucket_name" validate:"required"`
	ObjectName string `json:"object_name" validate:"required"`
}
