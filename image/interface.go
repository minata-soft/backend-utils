package image

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type Image interface {
	Connect(Config) (Image, error)
	UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error
	GetImage(objectName string, bucket_name string) ([]byte, error)
	ObjectDelete(bucket_name string, object_name string) error
	ObjectGet(bucket_name string, object_name string) (*minio.Object, error)
	BucketCreate(name_bucket string) error
	BucketDelete(bucket_name string) error
	BucketExist(bucket_name string) (bool, error)
	EnsureBucketExist(bucket_names []string) error
}

type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}
