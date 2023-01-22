package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetSpecNameFromUrl(url string) string {
	splitUrl := strings.Split(url, "/")
	specName := strings.SplitN(splitUrl[len(splitUrl)-1], "_", 2)[1]
	specName = strings.Replace(specName, ".yaml", "", 1)
	specName = strings.Replace(specName, ".yml", "", 1)
	return specName
}

func SaveFile(file *zip.File, output string, specName string, subFolder string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(rc)
	if err != nil {
		return err
	}
	fileSubPath := strings.Split(file.Name, "/")
	err = os.MkdirAll(fmt.Sprintf("%s/%s/%s", output, specName, strings.Replace(strings.Join(fileSubPath[:len(fileSubPath)-1], "/"), subFolder, "/", 1)), 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s/%s", output, specName, strings.Replace(file.Name, subFolder, "/", 1)), b, 0644)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
