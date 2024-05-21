package request

import "time"

type FileInfo struct {
	FileName string `json:"file_name,omitempty"`
	FileData []byte `json:"file_data,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

type UploadAttachmentRequest struct {
	Attachments []*FileInfo    `json:"attachments,omitempty"`
	ServiceName string         `json:"service_name,omitempty"`
	Acl         *string        `json:"acl,omitempty"`
	IsOverride  *bool          `json:"is_override,omitempty"`
	IsPresigned *bool          `json:"is_presigned,omitempty"`
	Duration    *time.Duration `json:"duration,omitempty"`
}

type PreSignedUrlsRequest struct {
	Urls     []string       `json:"urls,omitempty"`
	Duration *time.Duration `json:"duration,omitempty"`
}
