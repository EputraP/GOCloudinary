package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "./../../.env"

func TestLoadConfig(t *testing.T) {
	err := LoadConfig(path)
	cloud_name := os.Getenv("CLOUDINARY_NAME")
	require.Nil(t, err)
	require.Equal(t, "daylf7rsj", cloud_name)
}
