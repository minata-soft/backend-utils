package image

import "mime/multipart"

type MockImage struct {
	Connected bool
	InitError error
	Bucket    string
	UploadErr error
}

func (m *MockImage) Connect() error {
	m.Connected = true
	return nil
}

func (m *MockImage) Init() error {
	return m.InitError
}

func (m *MockImage) UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error {
	return m.UploadErr
}

func (m *MockImage) CreateBucket() error {
	return nil
}

func (m *MockImage) DeleteBucket(bucket_name string) error {
	return nil
}

func (m *MockImage) BucketExist(bucket_name string) (bool, error) {
	return m.Bucket == bucket_name, nil
}
