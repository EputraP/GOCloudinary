package images

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloudinary struct {
	Cloud   *cloudinary.Cloudinary
	isError error
}

type Option struct {
}

// menit 44.50
func NewCloudinary(cloudName, apiKey, apiSecrete string) Cloudinary {
	c, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecrete)
	if err != nil {
		return Cloudinary{}
	}
	return Cloudinary{
		Cloud:   c,
		isError: err,
	}
}

func (c Cloudinary) Upload(ctx context.Context, file interface{}) error {
	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "test1",
	})

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil

}
