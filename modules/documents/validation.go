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

// UpdateRequest change set for document
type UpdateRequest struct {
	Name       string `json:"name"`
	FileSize   int    `json:"file_size"`
	FileType   string `json:"file_type"`
	ObjectName string `json:"object_name"`
	BucketName string `json:"bucket_name"`
	URLLink    string `json:"url_link"`
}
