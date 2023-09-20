package config

import "os"

type CDNConfig struct {
	UploadPath string
}

func LoadCDNConfig() CDNConfig {

	os.MkdirAll(os.Getenv("UPLOAD_PATH"), os.ModePerm)

	return CDNConfig{
		UploadPath: os.Getenv("UPLOAD_PATH"),
	}
}
