package movefiles

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lucasportella/go-move-files/types"
)

func GetNewPathWithDate(path string, dateConfig types.DateValues) string {
	newPathWithDate := path
	year := dateConfig.GetYear()
	if year != "" {
		newPathWithDate += "/" + year
	}
	month := dateConfig.GetMonth()
	if month != "" {
		newPathWithDate += "/" + month
	}
	day := dateConfig.GetDay()
	if day != "" {
		newPathWithDate += "/" + day
	}
	return newPathWithDate
}

func BuildNewFilePath(file fs.DirEntry) {

}

func MoveFilesWithDate(configuration types.Configuration) {
	datePaths := configuration.WithDate
	// for key, paths := range datePaths.WithDateDaily {
	// 	srcPath := paths.SrcDir
	// 	dstpath := paths.DstDir
	// }

	for _, paths := range datePaths.WithDateMonthly {
		srcPath := paths.SrcDir
		// dstpath := paths.DstDir

		srcDir := ReadFilesFromSrcDir(srcPath)

		for _, file := range srcDir {
			newFilePathWithoutDate := srcPath + "/" + file.Name()
			fileDate, err := GetFileDate(newFilePathWithoutDate)
			if err != nil {
				log.Printf("Error while getting the date file of %v. Error: %v\n", file.Name(), err)
				continue
			}
			defaultDateConfig := GetDateConfig(fileDate)
			dateConfig, err := SetDateConfig(defaultDateConfig, "monthly")
			if err != nil {
				log.Println(err)
				continue
			}
			newPathWithDate := GetNewPathWithDate(srcPath, dateConfig)
			finalFilePath := newPathWithDate + "/" + file.Name()
			fmt.Println(finalFilePath)
		}

	}

	// for key, paths := range datePaths.WithDateYearly {
	// 	srcPath := paths.SrcDir
	// 	dstpath := paths.DstDir
	// }

	// srcDir := ReadFilesFromSrcDir(srcPath)

	// for _, dateOption := range dateOptions {

	// }

	// for _, file := range dateOption {
	// 	//openFile in the src dir
	// 	oldFilePath := srcPath + "/" + file.Name()
	// 	fileInfo, err := GetFileDate(oldFilePath)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fmt.Println(fileInfo)
	// 	_, err = OpenFile(oldFilePath)
	// 	if err != nil {
	// 		log.Printf("Error while opening the file using the old path. Path: %v", oldFilePath)
	// 	}

	// }
}

// GetFileDate gets the date of the file and returns a string in the format YYYY-MM-DD
func GetFileDate(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("error while getting data from the following file: %v. Error: %v", path, err)
	}
	formattedTime := fileInfo.ModTime().Format(time.RFC3339)
	relevantTime := strings.Split(formattedTime, "T")[0]
	return relevantTime, nil
}

// verify if correct folder exists -- done
// if exists create it, if not, do nothing
// update newFilePath
// moveFile
func GetDatePaths(key string, config string, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := CreateDateFolders(path)
		if err != nil {
			return fmt.Errorf("failed to create folder(s): %v\n", err)
		}
	} else if err != nil {
		return fmt.Errorf("error while checking folder existence: %v\n", err)
	}
	return nil
}

func CreateDateFolders(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfFolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreatePathWithDate(path string, dateConfig string) string {
	return path + dateConfig
}

func GetDateConfig(date string) types.DateValues {
	dates := strings.Split(date, "-")
	return types.DateValues{
		Year:  dates[0],
		Month: dates[1],
		Day:   dates[2],
	}
}

func SetDateConfig(dateValues types.DateValues, dateOption types.DateOption) (types.DateValues, error) {
	switch dateOption {
	case types.DateOption(types.Yearly):
		return types.DateValues{
			Year:  dateValues.Year,
			Month: "",
			Day:   "",
		}, nil
	case types.DateOption(types.Monthly):
		return types.DateValues{
			Year:  dateValues.Year,
			Month: dateValues.Month,
			Day:   "",
		}, nil
	case types.DateOption(types.Daily):
		return dateValues, nil
	}
	return dateValues, fmt.Errorf("error: could not set date config with dateValues of %v\n and dateOption of %v", dateValues, dateOption)
}
