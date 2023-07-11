package images

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Cloudinary struct {
	Cloud   *cloudinary.Cloudinary
	isError error
}

type Option struct {
}

// menit 53.56
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

func (c Cloudinary) Upload(ctx context.Context, file interface{}, path string) error {

	filename := uuid.NewString()

	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "testImage/" + path + "/" + filename,
		Eager:    "q_10",
	})

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil

}

func (c Cloudinary) Delete(ctx context.Context, path string) error {
	res, err := c.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: path,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", res)
	return nil
}
