package s3

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	MaxFileSize    = 20 << 20
	MaxFileCount   = 3
	StorageDir     = "storage"
	FileExtensions = []string{".jpg", ".jpeg", ".png", ".doc", ".docx", ".xls", ".xlsx", ".csv", ".pdf"}
)

type FileUploadInfo struct {
	FileName string
	FilePath string
}

type FileUploadForm struct {
	req          *http.Request
	exts         []string
	maxFileCount int
	maxFileSize  int64
	storageDir   string
}

func NewFileUpload(req *http.Request) *FileUploadForm {
	return &FileUploadForm{
		req:          req,
		exts:         FileExtensions,
		maxFileSize:  int64(MaxFileSize),
		storageDir:   StorageDir,
		maxFileCount: MaxFileCount,
	}
}

func (f *FileUploadForm) SetExtensions(exts []string) {
	if len(exts) > 0 {
		f.exts = exts
	}
}

func (f *FileUploadForm) SetMaxFileSize(fileSize int64) {
	if fileSize > 0 {
		f.maxFileSize = fileSize
	}
}

func (f *FileUploadForm) SetStorageDir(dir string) {
	if len(dir) > 0 {
		f.storageDir = fmt.Sprintf("%s/%s", f.storageDir, dir)
	}
}

func (f *FileUploadForm) createStorageDir() error {
	if _, err := os.Stat(f.storageDir); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(f.storageDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (f *FileUploadForm) isValidFileExtension(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	for _, validExt := range f.exts {
		if ext == validExt {
			return true
		}
	}
	return false
}

func (f *FileUploadForm) Parse(key string) ([]*multipart.FileHeader, error) {
	err := f.req.ParseMultipartForm(f.maxFileSize)
	if err != nil {
		return nil, err
	}

	files, ok := f.req.MultipartForm.File[key]
	if !ok {
		return nil, fmt.Errorf("the key name %s does not exists", key)
	}
	if len(files) > f.maxFileCount {
		return nil, fmt.Errorf("number of files exceeds the limit (%d)", f.maxFileCount)
	}
	return files, nil
}

func (f *FileUploadForm) Handle(key string) ([]*FileUploadInfo, error) {
	files, err := f.Parse(key)
	if err != nil {
		return nil, err
	}

	if err = f.createStorageDir(); err != nil {
		return nil, err
	}
	var results []*FileUploadInfo
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			_ = file.Close()
			return nil, err
		}

		if !f.isValidFileExtension(fileHeader.Filename) {
			return nil, fmt.Errorf("invalid file extension for file %s", fileHeader.Filename)
		}

		fPath := filepath.Join(f.storageDir, fileHeader.Filename)
		tpf, err := os.Create(fPath)
		if err != nil {
			_ = tpf.Close()
			return nil, fmt.Errorf("error creating output file: %v", err)
		}

		_, err = io.Copy(tpf, file)
		if err != nil {
			_ = file.Close()
			_ = tpf.Close()
			return nil, fmt.Errorf("error copying file: %v", err)
		}

		results = append(results, &FileUploadInfo{
			FileName: fileHeader.Filename,
			FilePath: fPath,
		})
		_ = file.Close()
		_ = tpf.Close()
	}
	return results, nil
}
