package services

import (
	"mime/multipart"

	"github.com/kumin/BityDating/entities"
)

func ConvertMultipartToFile(multipartFile *multipart.FileHeader) (*entities.File, error) {
	if multipartFile == nil {
		return nil, nil
	}

	file := &entities.File{
		Name:        multipartFile.Filename,
		ContentType: multipartFile.Header.Get("Content-Type"),
		Size:        multipartFile.Size,
	}
	buffer, err := multipartFile.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Close()
	file.Buffer = buffer

	return file, nil
}
