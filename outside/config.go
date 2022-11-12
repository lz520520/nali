package outside

import (
	"github.com/zu1k/nali/internal/constant"
	"github.com/zu1k/nali/internal/db"
	"log"
	"os"
)

func SetConfigPath(path string) {
	constant.ConfigDirPath = path
	constant.DataDirPath = path

	prepareDir(constant.ConfigDirPath)
	prepareDir(constant.DataDirPath)

	_ = os.Chdir(constant.DataDirPath)
}

func UpdateDB() {
	db.UpdateDB([]string{}...)
}

func prepareDir(dir string) {
	stat, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal("can not create config dir:", dir)
		}
	} else {
		if !stat.IsDir() {
			log.Fatal("path already exists, but not a dir:", dir)
		}
	}
}
