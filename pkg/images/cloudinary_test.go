package images

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildCloudinary(t *testing.T) {
	cloudName := "daylf7rsj"
	apiKey := os.Getenv("843337241332197")
	apiSecret := os.Getenv("ZCFn1Ub_ESBMzH4zRx9kENfT9pc")

	cloudinary := NewCloudinary(cloudName, apiKey, apiSecret)

	require.Nil(t, cloudinary.isError)
}
