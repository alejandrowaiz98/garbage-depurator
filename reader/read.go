package reader

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var allErrors []string

func (r *Reader) BuildExcel() error {

	err := filepath.WalkDir("Files", r.walkDirFunc)

	if err != nil {
		xerr := fmt.Errorf("List of errors reading files: \n%v", strings.Join(allErrors, ",\n"))
		return xerr
	}

	return nil

}

func (r *Reader) walkDirFunc(path string, fileInfo fs.DirEntry, err error) error {

	if !fileInfo.IsDir() {

		//os.Setenv("excel_filename", "/Files/"+fileInfo.Name())
		os.Setenv("excel_filename", path)

		if err != nil {
			obtainedErr := fmt.Sprintf("[ Path: %v | Step: WalkDir entry point | Error: %v ]", path, err)
			allErrors = append(allErrors, obtainedErr)
			return err
		}

		err = retry(3, 1, func() error { return r.e.BuildFinalExcel() })

		if err != nil {
			obtainedErr := fmt.Sprintf("[ Path: %v | Step: WalkDir entry point | Error: %v ]", path, err)
			allErrors = append(allErrors, obtainedErr)
			return err
		}

		return nil

	}

	return nil

}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
