package images

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (c Cloudinary) Upload(ctx context.Context, file interface{}, path string) (string, error) {

	filename := uuid.NewString()

	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "testImage/" + path + "/" + filename,
		// for handle image transformation
		Eager: "q_10",
	})

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", res)

	url := res.SecureURL

	return url, nil

}

func (c Cloudinary) Delete(ctx context.Context, path string) error {

	res, err := c.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		// path formula
		// the file must not contains file format, ex: (wrong) image.png => valid image
		// if the image in the home => <file_name>
		// if the image inside a folder => <folder_name>/<file-name>
		// if you type a folder inside the url path, but the folder is not existing in the cloudinary, cloudinary will create the folder automaticly
		// ex => testiamge/test/image.png => assume test folder is not existing in cloudinary, so cloudinary will create folder test inside testimage folder and the image.png will placed inside test folder
		PublicID: path,
	})
	if err != nil {
		return err
	}
	if strings.Contains(res.Result, "not found") {
		return errors.New("Image not found")
	}
	fmt.Printf("%+v\n", res)
	return nil
}
