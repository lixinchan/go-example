package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	Path "path/filepath"
	"strings"
	"sync"
)

func IsValidDir(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsValidFile(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsRegular()
}

func IsValidDirOrFile(name string) bool {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fileInfo.IsDir() || fileInfo.Mode().IsRegular()
}

func FilePutContents(filename, content string) (int, error) {
	return WriteBytesToFile(filename, []byte(content))
}

func FileReadContents(filename string) (contents string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		contents = string(data)
	}
	return
}

func FileRemove(filename string) (err error) {
	if IsValidFile(filename) {
		err = os.Remove(filename)
	}
	return
}

func WriteBytesToFile(filename string, bytes []byte) (int, error) {
	os.MkdirAll(path.Dir(filename), os.ModePerm)
	fw, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(bytes)
}

func IsSameAbsPath(filename1, filename2 string) bool {
	var (
		abs1, abs2 string
		err        error
	)
	if abs1, err = Path.Abs(filename1); err != nil {
		return false
	}
	if abs2, err = Path.Abs(filename2); err != nil {
		return false
	}
	return abs1 == abs2
}

func ListFiles(dir, suffix string) []string {
	var res = make([]string, 0)
	valid := IsValidDir(dir)
	if !valid {
		return res
	}
	if fileInfos, err := ioutil.ReadDir(dir); err == nil {
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				file := fileInfo.Name()
				if strings.HasSuffix(file, suffix) {
					res = append(res, file)
				}
			}
		}
	}
	return res
}

func ListAllFile(dir string) []string {
	var res = make([]string, 0)
	valid := IsValidDir(dir)
	if !valid {
		return res
	}
	if fileInfos, err := ioutil.ReadDir(dir); err == nil {
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				file := fileInfo.Name()
				absPath := dir + "/" + file
				res = append(res, absPath)
			}
		}
	}
	return res
}

func ForEachLine(filename string, lineHandler func(line string) error) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		err = lineHandler(line)
		if err != nil {
			break
		}
	}
	return err
}

func ForEachLineConcurrent(filename string, lineHandler func(line string) error, concurrent int) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	var waitGroup sync.WaitGroup
	jobs := make(chan bool, concurrent)
	started := make(chan bool)
	for scanner.Scan() {
		waitGroup.Add(1)
		line := scanner.Text()
		jobs <- true
		go func() {
			started <- true
			defer func() {
				waitGroup.Done()
				<-jobs

			}()
			_ = lineHandler(line)
		}()
		<-started
	}
	waitGroup.Wait()
	return err
}
