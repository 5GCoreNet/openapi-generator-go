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

// SaveFile saves a file to the specified path and parse relative imports
func SaveFile(file *zip.File, output string, specName string, subFolders []string, rootPkg string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(rc)
	if err != nil {
		return err
	}
	fileSubPath := strings.Split(file.Name, "/")
	if strings.Contains(string(b), "./openapi") {
		b = []byte(strings.Replace(string(b), "./openapi", fmt.Sprintf("%s/%s", rootPkg, specName), -1))
	}
	if strings.Contains(string(b), "./go") {
		b = []byte(strings.Replace(string(b), "./go", fmt.Sprintf("%s/%s", rootPkg, specName), -1))
	}
	folderName := fmt.Sprintf("%s/%s/%s", output, specName, strings.Join(fileSubPath[:len(fileSubPath)-1], "/"))
	for _, subFolder := range subFolders {
		folderName = strings.Replace(folderName, subFolder, "", 1)
	}
	err = os.MkdirAll(folderName, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	fileName := fmt.Sprintf("%s/%s/%s", output, specName, file.Name)
	for _, subFolder := range subFolders {
		fileName = strings.Replace(fileName, subFolder, "", 1)
	}
	err = os.WriteFile(fileName, b, 0644)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		return err
	}
	return nil
}
