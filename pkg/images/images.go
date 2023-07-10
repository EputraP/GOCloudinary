package images

type ImageInterface interface {
	Upload()
	Unlink()
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
