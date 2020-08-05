package isgit

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


func Path(dir string) (bool, error){
	rd, err := GetRootDir(dir)
	return rd != "", err
}

func WD() (bool, error){
	p, err := os.Getwd()
	if err != nil {
		return false, err
	}
	return Path(p)
}

func GetRootDir(path string) (string, error) {
	return findDirReverse(path, containsGit)
}

func GetRootDirWD() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return GetRootDir(p)
}


func containsGit(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, f := range files {
		if f.Name() == ".git" {
			return true, nil
		}
	}
	return false, nil
}

func findDirReverse(path string, e func(string)(bool, error)) (string, error){
	r, err := e(path)
	if err != nil {
		return "", err
	}

	if r {
		return path, nil
	}

	np, err := filepath.Abs(filepath.Join(path,".."))
	if err != nil {
		return "", err
	}
	if np == path {
		return "", nil
	}
	return findDirReverse(np, e)
}
