package image_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/minata-soft/backend-utils/image"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

type TestStruct struct {
	img image.Image
}

func setup(t *testing.T) *TestStruct {
	img := image.NewImageStorage(&image.ImageMinio{})
	require.NotNil(t, img)

	err := img.Connect(image.Config{
		Endpoint:        "127.0.0.1:9000",
		AccessKeyID:     "minio_access_key",
		SecretAccessKey: "minio_secret_key",
		UseSSL:          false,
	})
	require.Nil(t, err)

	return &TestStruct{img}
}

func TestMinioNewServer(t *testing.T) {
	img := image.NewImageStorage(&image.ImageMinio{})
	assert.NotNil(t, img)
}

func TestConnect(t *testing.T) {

	img := image.NewImageStorage(&image.ImageMinio{})
	assert.NotNil(t, img)

	testData := []struct {
		Name          string
		Config        image.Config
		ExpectedError bool
	}{
		{
			Name:          "empty conf",
			Config:        image.Config{},
			ExpectedError: true,
		},
		{
			Name: "only fake endoint",
			Config: image.Config{
				Endpoint: "fake",
			},
			ExpectedError: false,
		},
		{
			Name: "only fake AccesKeyID",
			Config: image.Config{
				AccessKeyID: "fake",
			},
			ExpectedError: true,
		},
		{
			Name: "only fake SecretAccessKey",
			Config: image.Config{
				SecretAccessKey: "fake",
			},
			ExpectedError: true,
		},
	}

	t.Run("configuration", func(t *testing.T) {

		for _, test := range testData {
			t.Run(test.Name, func(t *testing.T) {
				err := img.Connect(test.Config)
				t.Run("ExpectedError "+strconv.FormatBool(test.ExpectedError), func(t *testing.T) {
					if test.ExpectedError {
						require.NotNil(t, err)
					} else {
						require.Nil(t, err)
					}
				})
			})
		}

	})

}

func TestCreateBucket(t *testing.T) {
	var (
		test_utils         = setup(t)
		name_bucket string = strings.ToLower(faker.RandomString(6))
	)

	fmt.Printf("name_bucket %s\n", name_bucket)

	err := test_utils.img.BucketCreate(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

}

func TestBucketExist(t *testing.T) {
	var (
		test_utils         = setup(t)
		name_bucket string = strings.ToLower(faker.RandomString(6))
	)

	err := test_utils.img.BucketCreate(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

	b, err := test_utils.img.BucketExist(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)
	assert.True(t, b)
}

func TestDeleteBucket(t *testing.T) {
	var (
		test_utils         = setup(t)
		name_bucket string = strings.ToLower(faker.RandomString(6))
	)

	err := test_utils.img.BucketCreate(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

	b, err := test_utils.img.BucketExist(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)
	assert.True(t, b)

	err = test_utils.img.BucketDelete(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

}

func TestUploadImage(t *testing.T) {
	var (
		test_utils         = setup(t)
		name_bucket string = strings.ToLower(faker.RandomString(6))
	)

	err := test_utils.img.BucketCreate(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

	b, err := test_utils.img.BucketExist(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)
	assert.True(t, b)

	err = test_utils.img.BucketDelete(name_bucket)
	assert.Nil(t, err, image.ErrBucketCreate)

}
