package image

import "mime/multipart"

type MockImage struct {
	Connected bool
	InitError error
	Bucket    string
	UploadErr error
}

// BucketCreate implements Image.
func (*MockImage) BucketCreate(name_bucket string) error {
	return nil
}

// BucketDelete implements Image.
func (*MockImage) BucketDelete(bucket_name string) error {
	return nil
}

// BucketExist implements Image.
func (*MockImage) BucketExist(bucket_name string) (bool, error) {
	return true, nil
}

// Connect implements Image.
func (*MockImage) Connect(Config) (Image, error) {
	return &MockImage{Connected: true}, nil
}

// EnsureBucketExist implements Image.
func (*MockImage) EnsureBucketExist(bucket_names []string) error {
	return nil
}

// UploadImage implements Image.
func (*MockImage) UploadImage(objectName string, file *multipart.FileHeader, bucket_name string) error {
	return nil
}
