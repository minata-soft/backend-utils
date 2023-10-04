package image_test

import (
	"strconv"
	"testing"

	"github.com/minata-soft/backend-utils/image"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
