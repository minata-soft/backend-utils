package image

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	backend_utils "github.com/minata-soft/backend-utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ImageMinio struct {
	Ctx    context.Context
	Client *minio.Client
}

func NewImageStorage(img Image) Image {
	return img
}

func (img *ImageMinio) Connect(config Config) (Image, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return img, err
	}
	img.Client = client
	img.Ctx = context.Background()
	return img, err
}

func (m *ImageMinio) GetImage(objectName string, bucket_name string) ([]byte, error) {
	obj, err := m.Client.GetObject(m.Ctx, bucket_name, objectName, minio.GetObjectOptions{})
	if err != nil {

	}
	defer obj.Close()

	// Read the object's content
	data, err := io.ReadAll(obj)
	if err != nil {
		fmt.Println(err)
		backend_utils.Debug.Error("error while reading the object: %v", err)
		return []byte{}, err
	}

	return data, nil

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

func (img *ImageMinio) EnsureBucketExist(bucket_names []string) error {
	var (
		err   error
		exist bool
	)

	for _, v := range bucket_names {
		if exist, err = img.BucketExist(v); err != nil {
			return errors.New("an error occurs when checking the bucket " + err.Error())
		}
		if !exist {
			if err = img.Client.MakeBucket(img.Ctx, v, minio.MakeBucketOptions{}); err != nil {
				return errors.New("Fail to create bucket " + err.Error())
			}
		}

	}
	return nil
}
