package image

import "mime/multipart"

type Image interface {
	Connect(Config) error
	UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error
	BucketCreate(name_bucket string) error
	BucketDelete(bucket_name string) error
	BucketExist(bucket_name string) (bool, error)
}

type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}
