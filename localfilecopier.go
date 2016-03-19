package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Starting printing ")

	listFiles("/data/test/source", "/data/test/target", "")

	fmt.Println("Done")
}

func listFiles(folderName, targetBaseFolder, currentFolderPath string) {

	currFolder, _ := ioutil.ReadDir(folderName)

	for _, y := range currFolder {
		updatedPath := currentFolderPath + "/" + y.Name()
		if y.IsDir() {
			subFolder := folderName + "/" + y.Name()
			fmt.Println("--> ", updatedPath)
			currentTarget := targetBaseFolder + "/" + updatedPath
			os.Mkdir(currentTarget, os.ModePerm)
			listFiles(subFolder, targetBaseFolder, updatedPath)

		} else {
			filePath := targetBaseFolder + "/" + updatedPath
			reader, err := os.Open(filePath)
			if err != nil {
				fmt.Errorf("Error")
			}
			target := filePath + ".gz"

			writer, err := os.Create(target)
			defer writer.Close()

			archiver := gzip.NewWriter(writer)
			archiver.Name = target
			defer archiver.Close()

			fmt.Println("Writing ", filePath, target)
			_, err = io.Copy(archiver, reader)
			// fmt.Println(x, y.Name())
		}

	}

}
