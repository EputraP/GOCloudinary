package images

import "context"

// interface ni penting buat misal pake third pary gitu
// soalnya third party gini rawan untuk digantikan
// yang penting nama method sama parameternya sama, isinya bisa beda
// tipe data returnnnya bisa beda gk papa untuk Build Third Party nya

type ImageInterface interface {
	Upload(ctx context.Context, file interface{}, path string) (string, error)
	Delete(ctx context.Context, path string) error
}

type Image struct {
	BucketName string
	APIKey     string
	APISecrete string
}

func NewImages(bucket, apiKey, apiSecrete string) Image {
	return Image{
		BucketName: bucket,
		APIKey:     apiKey,
		APISecrete: apiSecrete,
	}
}

func (i Image) BuildCloudinary() Cloudinary {
	return NewCloudinary(i.BucketName, i.APIKey, i.APISecrete)
}

func (i Image) BuildGCS() Image {
	return NewImages(i.BucketName, i.APIKey, i.APISecrete)
}

func (i Image) Upload(ctx context.Context, file interface{}, path string) (string, error) {
	panic("unimplement")
}
func (i Image) Delete(ctx context.Context, path string) error {
	panic("unimplement")
}
