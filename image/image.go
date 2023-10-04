package image

import (
	"context"
	"errors"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return err
	}
	img.Client = client
	img.Ctx = context.Background()
	return err
}

func (m *ImageMinio) UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error {
	var (
		info        minio.UploadInfo
		buffer, err = file.Open()
		fileBuffer  = buffer
		contentType = file.Header["Content-Type"][0]
		fileSize    = file.Size
	)
	if err != nil {
		return errors.New("Fail to open file :" + err.Error())
	}

	// Upload the zip file with PutObject
	if info, err = m.Client.PutObject(m.Ctx, bucket_name, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType}); err != nil {
		return errors.New("Fail to upload file :" + err.Error())
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return nil
}

func (img *ImageMinio) BucketCreate(bucket_name string) error {
	err := img.Client.MakeBucket(img.Ctx, bucket_name, minio.MakeBucketOptions{})
	if err != nil {
		return errors.New(ErrBucketCreate + " err: " + err.Error())
	}

	return err
}

func (img *ImageMinio) BucketDelete(bucket_name string) error {
	err := img.Client.RemoveBucket(img.Ctx, bucket_name)
	if err != nil {
		return errors.New(ErrBucketDelete + " err: " + err.Error())
	}

	return err
}

func (img *ImageMinio) BucketExist(bucket_name string) (bool, error) {
	found, err := img.Client.BucketExists(context.Background(), bucket_name)
	if err != nil {
		return false, errors.New(ErrBucketCheck + " err: " + err.Error())
	}

	return found, err
}
