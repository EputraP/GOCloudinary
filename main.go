package main

import (
	"GOCloudinary/pkg/images"
	"GOCloudinary/utils/config"
	"context"
	"os"
)

type CloudSvc interface {
	Upload(ctx context.Context, file interface{}, path string) (string, error)
	Delete(ctx context.Context, path string) error
}

type Services struct {
	cloud CloudSvc
}

func main() {
	config.LoadConfig(".env")
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRETE")

	image := images.NewImages(cloudName, apiKey, apiSecret)

	svc := Services{
		cloud: image.BuildGCS(),
	}

	svc.cloud.Upload(context.Background(), "", "")
}
