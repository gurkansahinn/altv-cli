package pkg

import (
	"fmt"
	"runtime"
)

type BinaryFile struct {
	RootPath string `json:"rootPath"`
	CdnUrl   string `json:"cdnUrl"`
}

func GetBinaryFiles(branch string) []BinaryFile {
	var linuxBinaryFiles = []BinaryFile{
		{
			RootPath: "data/vehmodels.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/vehmodels.bin", branch),
		},
		{
			RootPath: "data/vehmods.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/vehmods.bin", branch),
		},
		{
			RootPath: "data/clothes.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/clothes.bin", branch),
		},
		{
			RootPath: "modules/libjs-module.so",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/js-module/%s/x64_linux/modules/js-module/libjs-module.so", branch),
		},
		{
			RootPath: "libnode.so.102",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/js-module/%s/x64_linux/modules/js-module/libnode.so.102", branch),
		},
		{
			RootPath: "start.sh",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/others/start.sh", branch),
		},
		{
			RootPath: "altv-server",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/server/%s/x64_linux/altv-server", branch),
		},
	}

	var windowsBinaryFiles = []BinaryFile{
		{
			RootPath: "data/vehmodels.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/vehmodels.bin", branch),
		},
		{
			RootPath: "data/vehmods.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/vehmods.bin", branch),
		},
		{
			RootPath: "data/clothes.bin",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/data/%s/data/clothes.bin", branch),
		},
		{
			RootPath: "modules/js-module.dll",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/js-module/%s/x64_win32/modules/js-module/js-module.dll", branch),
		},
		{
			RootPath: "libnode.dll",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/js-module/%s/x64_win32/modules/js-module/libnode.dll", branch),
		},
		{
			RootPath: "altv-server.exe",
			CdnUrl:   fmt.Sprintf("https://cdn.altv.mp/server/%s/x64_win32/altv-server.exe", branch),
		},
	}

	os := runtime.GOOS
	switch os {
	case "windows":
		return windowsBinaryFiles
	case "darwin", "linux":
		return linuxBinaryFiles
	default:
		return linuxBinaryFiles
	}
}
