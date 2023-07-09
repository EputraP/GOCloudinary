package cloud

import "github.com/cloudinary/cloudinary-go/v2"

type Cloudinary struct {
	cloudinary.Cloudinary
	isError error
}

//menit 36.25
func NewCloudinary(cloudName, apiKey, apiSecrete string) Cloudinary {
	c, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecrete)
	if err != nil {
		return Cloudinary{}
	}
	return Cloudinary{
		CloudName:  cloudName,
		APIKey:     apiKey,
		APISecrete: apiSecrete,
	}
}
