package pkg

import (
	"altv-cli/util"
	"log"
	"os"
)

func DownloadBinaryFiles(branch string) {
	directories := [2]string{"data", "modules"}

	for _, name := range directories {
		exists := util.ExistsDirectory(name)

		if !exists {
			err := os.Mkdir(name, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, binaryFile := range GetBinaryFiles(branch) {
		util.DownloadFile(binaryFile.RootPath, binaryFile.CdnUrl)
	}
	os.Exit(1)
}
