package common

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

type FileList struct {
	Name       string `json:"name"`
	UpdateTime int64  `json:"updateTime"`
	Type       string `json:"type"`
	IsDir      bool   `json:"isDir"`
	URL        string `json:"url,omitempty"`
}

type FileLists []FileList

func (s FileLists) Len() int {
	return len(s)
}

func (s FileLists) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s FileLists) Less(i, j int) bool {
	return s[i].UpdateTime < s[j].UpdateTime
}

func CreateDir(outPath string) {
	if _, err := os.Stat(outPath); err != nil {
		os.MkdirAll(outPath, 0777)
	}
}

// 检查文件或目录是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files FileLists, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		file := FileList{
			Name:       fi.Name(),
			UpdateTime: fi.ModTime().Unix(),
			Type:       GetFileExt(fi.Name()),
			IsDir:      false,
			URL:        fi.Name(),
		}

		if fi.IsDir() {
			file.IsDir = true
			file.Type = "dir"
			file.UpdateTime = file.UpdateTime - 30*365*24*60*60 //减去30年，目的让文件夹排前面
		}

		files = append(files, file)
	}
	sort.Sort(files)

	return files, nil
}

func GetFileExt(filename string) string {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".ppt", ".pptx":
		ext = "ppt"
	case ".doc", ".docx":
		ext = "doc"
	case ".xls", ".xlsx":
		ext = "xls"
	case ".pdf":
		ext = "pdf"
	case ".html", ".htm":
		ext = "htm"
	case ".txt", ".log":
		ext = "txt"
	case ".swf":
		ext = "flash"
	case ".zip", ".rar", ".7z":
		ext = "zip"
	case ".mp3":
		ext = "mp3"
	case ".mp4":
		ext = "mp4"
	default:
		ext = "file"
	}

	return ext
}
