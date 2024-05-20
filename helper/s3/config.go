package s3

type BucketS3Config struct {
	Prefix     string `json:"prefix,default=inapp"`
	Region     string `json:"bucket-region,default=ap-southeast-1"`
	BucketName string `json:"bucket-name"`
	AccessKey  string `json:"bucket-access-key"`
	SecretKey  string `json:"bucket-secret-key"`
	CDNUrl     string `json:"cdn-url"`
}
