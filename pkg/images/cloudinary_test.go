package images

import (
	"GOCloudinary/utils/config"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "./../../.env"
var cloud Cloudinary

func init() {
	config.LoadConfig(path)
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRETE")

	cloud = NewCloudinary(cloudName, apiKey, apiSecret)
}

func TestBuildCloudinary(t *testing.T) {
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	// apiSecret := os.Getenv("CLOUDINARY_API_SECRETE")

	cloudinary := NewCloudinary(cloudName, apiKey, "ZCFn1Ub_ESBMzH4zRx9kENfT9pc")

	require.Equal(t, "daylf7rsj", cloudName)
	require.Nil(t, cloudinary.isError)
}

func TestUpload(t *testing.T) {
	_, err := cloud.Upload(context.Background(), "https://upload.wikimedia.org/wikipedia/id/3/36/Naruto_Uzumaki.png", "ngetesAja")

	require.Nil(t, err)
}

func TestDelete(t *testing.T) {

	// url := "test1"

	url := "testImage/ngetesAjaa"
	// url := "v1689082144/testImage/ngetesAja/2698ed4f-deed-4d15-b5c1-9cde4f2347a1.png"

	err := cloud.Delete(context.Background(), url)

	require.Nil(t, err)
}
