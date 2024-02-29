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

func BuildNewPathWithDate(file fs.DirEntry, srcPath string, dstPath string, dateOption string) (string, error) {
	oldFilePath := srcPath + "/" + file.Name()
	fileDate, err := GetFileDate(oldFilePath)
	if err != nil {
		log.Printf("Error while getting the date file of %v. Error: %v\n", file.Name(), err)
		return "", err
	}
	defaultDateConfig := GetDateConfig(fileDate)
	dateConfig, err := SetDateConfig(defaultDateConfig, types.DateOption(dateOption))
	if err != nil {
		log.Println(err)
		return "", err
	}
	newPathWithDate := GetNewPathWithDate(dstPath, dateConfig)
	return newPathWithDate, nil
}

func MoveFilesDateOption(option map[string]types.Paths, dateOption string) {
	for key, paths := range option {
		srcDir := ReadFilesFromSrcDir(paths.SrcDir)
		for _, file := range srcDir {
			if !FileNameContainsKey(file.Name(), key) {
				continue
			}

			newPathWithDate, err := BuildNewPathWithDate(file, paths.SrcDir, paths.DstDir, dateOption)
			if err != nil {
				log.Println(err)
				continue
			}
			err = MkdirNewFolders(file, newPathWithDate, key)
			if err != nil {
				log.Println(err)
				continue
			}
			MoveFile(newPathWithDate, paths.SrcDir, file.Name())
		}

	}
}

func MoveFilesWithDate(configuration types.Configuration) {
	MoveFilesDateOption(configuration.WithDate.WithDateYearly, types.Yearly)
	MoveFilesDateOption(configuration.WithDate.WithDateMonthly, types.Monthly)
	MoveFilesDateOption(configuration.WithDate.WithDateDaily, types.Daily)
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

func GetDatePaths(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := CreateFolders(path)
		if err != nil {
			return fmt.Errorf("failed to create folder(s): %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error while checking folder existence: %v", err)
	}
	return nil
}

func CreateFolders(path string) error {
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
