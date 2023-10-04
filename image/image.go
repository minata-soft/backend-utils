package image

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go"
)

type ImageMinio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Ctx             context.Context
	Client          *minio.Client
	Bucket          string
}

func NewImageStorage(img Image, options ...func(*ImageMinio)) Image {
	return img
}

func WithEndpoint(endpoint string) func(*ImageMinio) {
	return func(img *ImageMinio) {
		img.Endpoint = endpoint
	}
}

func WithAccessKeyID(accessKeyID string) func(*ImageMinio) {
	return func(img *ImageMinio) {
		img.AccessKeyID = accessKeyID
	}
}

func WithSecretAccessKey(secretAccessKey string) func(*ImageMinio) {
	return func(img *ImageMinio) {
		img.SecretAccessKey = secretAccessKey
	}
}

func WithUseSSL(useSSL bool) func(*ImageMinio) {
	return func(img *ImageMinio) {
		img.UseSSL = useSSL
	}
}

func WithBucket(bucket string) func(*ImageMinio) {
	return func(img *ImageMinio) {
		img.Bucket = bucket
	}
}

func (img *ImageMinio) Connect(config Config) error {
	client, err := minio.New(config.Endpoint, config.AccessKeyID, config.SecretAccessKey, config.UseSSL)
	if err != nil {
		return err
	}
	img.Client = client
	img.Ctx = context.Background()
	return err
}

func (img *ImageMinio) Init() error {
	err := img.Connect(Config{})
	if err != nil {
		return err
	}
	exists, err := img.BucketExist(img.Bucket)
	if err != nil {
		return err
	}
	if !exists {
		err = img.CreateBucket()
		if err != nil {
			return err
		}
	}
	return nil
}

func (img *ImageMinio) UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error {
	panic("unimplemented")
}

func (img *ImageMinio) CreateBucket() error {
	// implementation
	panic("unimplemented")
}

func (img *ImageMinio) DeleteBucket(bucket_name string) error {
	// implementation
	panic("unimplemented")
}

func (img *ImageMinio) BucketExist(bucket_name string) (bool, error) {
	// implementation
	panic("unimplemented")
}
