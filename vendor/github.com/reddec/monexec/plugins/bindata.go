// Code generated by go-bindata.
// sources:
// ../ui/dist/index.html
// ../ui/dist/main.js
// DO NOT EDIT!

package plugins

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8f\xb1\x8e\xc2\x30\x0c\x86\x5f\xc5\x97\xfd\x2e\xeb\x0d\x8e\x97\x3b\xd8\x10\x0c\x65\x60\x0c\x89\x45\x53\xb5\x69\x15\x9b\xa8\x7d\x7b\x54\x5a\x31\xd9\xdf\x27\xd9\xbf\x7e\xfc\xfa\x3f\xff\x35\xb7\xcb\x01\x5a\x1d\x7a\x02\x5c\x07\xf4\x3e\x3f\x1c\xe7\x15\xd9\x47\x02\x1c\x58\x3d\x84\xd6\x17\x61\x75\xd7\xe6\xf8\xfd\x4b\x80\x9a\xb4\x67\x3a\x8d\x99\x67\x0e\x68\x37\x04\xb4\xfb\xcd\x7d\x8c\x0b\x01\xc6\x54\x21\x45\x27\xcf\x89\x4b\x4d\x32\x16\x21\xb4\x31\x55\x02\x94\x50\xd2\xa4\xa0\xcb\xc4\xce\x28\xcf\x6a\x3b\x5f\xfd\x66\x0d\x48\x09\xce\x0c\x3e\xe5\x9f\x4e\x0c\xa1\xdd\x3c\xa1\xdd\x1f\x7f\x78\x5f\xd6\xe4\x77\x87\x57\x00\x00\x00\xff\xff\x65\x5a\x0d\x38\xd4\x00\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 212, mode: os.FileMode(420), modTime: time.Unix(1530190485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7b\xeb\x92\xdb\xb6\x92\xf0\xab\x68\x58\x3e\x2c\xa0\xd2\xc3\x68\x7c\x92\xaa\x7c\x54\x90\xf9\x62\xe7\xea\x24\x76\x62\x27\x71\x12\xad\x6a\x0e\x45\x82\x12\x3c\x14\x20\x83\xe0\x5c\xa2\xe1\xfe\xde\xa7\xd8\xda\x67\xd9\x47\xd9\x27\xd9\x6a\x80\x20\x41\x49\xe3\x4b\x4e\xd5\xf9\xb5\xa9\xd4\x40\x04\x1a\x8d\x46\xa3\x1b\x7d\x41\xfb\xa4\x6c\x64\x6e\x84\x92\xc4\xd0\xdd\x55\xa6\x27\x9c\xed\xda\x99\xef\x9c\x48\xa2\xe8\x4e\x94\x84\xcf\xd5\x82\x6a\x6e\x1a\x2d\x27\xf8\x3b\xe1\x37\x5b\xa5\x4d\x3d\xc3\x29\x9a\x61\x17\xdb\x89\x54\x41\x95\x9e\x9c\x41\x37\x98\xee\xda\x76\xd6\x4d\x32\x38\x29\xcf\xaa\x8a\x68\x3f\x17\x34\x0c\xbf\x25\x05\x9d\x54\xec\x64\x3a\xf4\xb5\x32\xd9\x30\x03\x32\xc9\x19\x07\x99\x14\x6c\x20\x15\x38\x28\xba\x93\x89\xc2\x9f\xf4\xee\xee\xd9\xf2\x15\xcf\x4d\x52\xf0\x52\x48\xfe\xa3\x56\x5b\xae\xcd\xad\x05\xdb\x71\xd9\x6c\xb8\xce\x96\x15\x4f\x4f\xa6\xb0\xe2\x26\x55\x2d\x6d\x41\x26\x9a\x85\x5b\x8f\x1a\xe9\x66\x17\xd1\x09\x33\xb7\x5b\xae\xca\xc9\x8b\xdb\xcd\x52\x55\x71\xec\xda\xc4\xa8\x17\x46\x0b\xb9\xfa\x39\x5b\xc5\xf1\x7d\x2b\x1e\xc2\xc2\xee\x2a\xab\x1a\x9e\x46\x3f\xa8\xa2\xa9\x78\xd4\x52\xb8\x6f\x72\x74\x71\xc1\xeb\x0e\xcc\x4f\x3b\x99\x3a\x72\xcd\x68\xfb\xf6\x50\xce\x62\x1e\xc7\xc4\x30\xdc\x00\x85\x4f\x62\xee\x4f\xc8\xcc\x44\x49\x3e\xc2\xd1\x48\xd9\xa5\x22\xe6\xf7\x64\xe2\x18\xff\x4f\x86\x95\x86\x49\x78\x96\x8a\x75\xc4\xe5\x9a\x67\x86\x13\xd9\x54\x15\x45\x74\x32\xd1\x44\xdd\x47\xba\x82\xa8\xe0\x65\xd6\x54\x26\xda\xe7\xb8\xdb\x85\x69\x29\x3c\xb4\x04\xd5\x96\x2f\x03\x93\x0d\x2d\x95\x26\x56\x8c\x26\x42\x4e\x0c\x95\x49\x41\x14\x68\xe8\xb7\xcb\xe9\xae\x17\x22\xbe\x68\x93\xa5\x90\x85\xa5\x0b\x34\xa5\x5e\xbe\x14\xf2\x48\xb2\x43\x69\xde\xdb\xed\x79\x0f\x31\x60\x4d\x3a\xda\xdb\xf4\xc8\x60\x2f\xc1\x48\x17\x87\x28\x8b\x80\x53\xe0\xb8\x9c\xda\x3b\x92\x0e\xb0\x63\xd1\x56\x2b\xa3\x70\x93\xc9\x3a\xab\x9f\x5d\x4b\xcf\x2c\xa7\x05\x38\x01\x71\x6c\x59\x14\x81\x24\x32\xa9\xd9\x94\xb6\x64\x3e\x92\x71\x89\x72\x59\xf3\x09\xf2\x2c\x37\xd1\xa0\x96\x8a\xd0\x5d\xdb\x7f\x69\xb7\xbc\xe7\xa3\x44\x3e\x72\x6a\xe6\x72\xc1\xf8\x5c\x2e\x7a\x15\x1c\x66\x08\x37\x83\x27\xd9\x76\xcb\x65\xf1\x78\x2d\xaa\x82\x18\x3a\x00\xd4\x7e\x7d\x9e\x08\x59\x73\x6d\x1e\xf1\x52\x69\x4e\x0c\xc8\x00\x2a\x47\x36\x9b\x64\x9b\x69\x2e\xcd\x53\x55\xf0\x44\xf3\x8d\xba\xe2\x87\xf8\xb2\x3d\x12\xd9\x74\x26\x3f\x35\x49\xc5\xe5\xca\xac\x67\xf2\x03\x76\x66\xe9\x8d\x63\xfc\x8b\x8c\x0e\xe6\x56\xb8\x4a\xb7\x87\x42\xe5\xcd\x86\x4b\x2f\x9e\x5f\x56\x1c\xbf\x46\x4b\x35\xf7\x83\xff\xcc\x6f\x2c\x99\x23\xf8\x92\xdc\x07\xfe\x58\x6d\x2c\xf6\x28\x0a\xc0\x0b\xcf\x19\x93\x64\x45\xf1\xe5\x15\x97\xe6\x7b\x51\x1b\x2e\xb9\x26\x1c\x24\x9c\x9c\x05\xc0\xab\x01\xd8\x71\xe6\x2d\xf0\x9b\x01\xbe\x36\xb7\x15\x4f\x6a\x6e\x7a\x25\xc3\x81\x16\xd5\x90\x53\xab\xaa\x6b\xb6\xd3\x8d\x94\x42\xae\xf0\xce\x35\x3a\x93\xb5\x40\x2c\x75\x3a\x5f\xc0\x52\x35\xb2\x48\xad\x96\x58\x4c\xf5\x9a\x73\xe3\xbe\xb3\xdc\x88\x2b\xfe\xbc\xa9\x38\xde\xd0\xb0\xd5\x6a\x23\x6a\xde\x8d\x15\x85\x3d\xd3\xb5\xa8\x93\x00\x63\xb2\x6d\xea\x35\x31\x14\xec\x40\xb7\xea\xdd\x1d\x09\x3f\xed\xdd\xcd\x5f\x37\xbc\x36\x9f\x4b\xb1\xc9\x70\xe2\x57\x3a\xdb\x70\x07\x65\x09\xf2\x53\xec\x07\xb3\x3f\x25\xbf\x31\x4e\xa5\xf1\x93\x52\x4a\x5b\xa4\x02\xc9\xeb\x2f\xba\x13\x0b\x39\xec\x83\xee\x72\x25\x6b\x33\x31\xac\x22\x91\xed\x8e\xe8\xac\x3f\xbc\x35\xcf\x8a\x3d\xc1\x86\x75\x30\x9b\x99\xc4\xb6\xad\xc5\x1a\x70\x63\xce\x17\x9e\xc0\x71\x2f\x6e\x6d\x8f\x84\x4e\x2d\x2c\x99\xff\xf8\xff\x97\xfc\xb6\xc4\xad\xd6\x93\x07\x3b\xde\x4e\x1e\xec\x4c\xfb\x8f\x83\x19\x79\x5d\x5b\x84\x9d\xd4\xe3\x46\x71\xf3\xa4\x63\x77\xcf\xc6\xb3\x99\xdf\xdd\xb5\x90\x85\xba\x4e\xb6\x5c\x97\x4a\x6f\x32\x99\xf3\x44\xaa\x6b\x42\x67\x15\x37\x78\xb5\xed\x9f\x52\xa7\x4e\xa8\x65\x33\x7e\x7a\x3a\xf3\x7c\x92\x07\xa0\x73\xbe\x98\x49\xbc\xa2\x56\x3a\xdb\xc4\xb1\xf9\x8c\xf5\x5f\x09\x97\x45\x1c\xcb\xa4\x50\x92\x13\x8a\x57\x14\x97\x85\x90\x2b\x0f\xe5\xbe\x92\xda\x64\xda\x20\x9c\xfd\x41\xfa\x01\x9c\xd1\x6d\xe5\x9c\xc8\xa4\xd9\x16\x68\x46\xf6\x44\x87\x9d\x4c\x69\xda\x4f\xb9\xbb\x3b\xd8\x49\xbd\xad\x44\xce\x09\x87\x33\xda\x8a\x72\x24\x67\xf4\x6d\x42\x46\x67\xbc\xaa\xf9\xc4\x4f\x0b\xc5\x06\xf9\x66\xd8\x5b\x0e\xc6\xf1\xcf\x20\xff\xf6\x21\x0b\x5e\x71\xc3\x9d\x6c\xd2\xd9\xbe\xa4\xb0\x5d\xdb\xb6\x10\xc2\xa0\xfc\x7a\x55\xce\x3c\xbd\xec\xa0\xc7\x6e\xd7\x90\x08\x26\x11\x4d\x4a\x51\x19\xae\x89\x61\x9f\x99\x38\x3e\x3d\x63\x8c\xa1\xb0\x15\xfc\xe6\x59\x49\x38\xa5\xc9\x2b\x25\xa4\x03\x6d\x61\xa5\x55\xb3\x7d\xd6\x18\xad\x6a\x2f\x46\xca\x7e\xb1\x9d\xe6\x9b\x4c\xd8\xcb\x61\x0a\x68\x71\x96\x59\x7e\x89\x37\x43\xdb\xc2\x75\x26\x4c\x4a\x28\xfb\x8c\xac\x93\x4e\xfb\xef\xee\x86\xdf\xec\x47\xd7\x26\x9a\xd7\xaa\xba\x42\x21\xe8\xc7\x12\xb3\xe6\x92\xe0\xdc\xdd\x00\x8f\x37\x47\x4b\x03\x28\x1a\x78\x90\x5b\xc7\x05\x77\xeb\xf3\xc4\xa8\x4b\x2e\x47\x2e\x26\xfa\x70\x0a\x04\xd4\xce\xd1\x74\x10\x27\x8c\xc9\xce\x31\x99\xf5\x84\x14\x4c\xc4\xf1\x6e\x2e\x16\x69\xdd\x76\x3a\x92\x33\x4d\x34\xd9\xb5\xc0\x93\xdc\xdc\x50\x18\x60\x29\x64\x68\xfb\x09\x4f\xf2\x46\xa3\x7d\x62\x86\xe2\x87\xda\x6c\x95\xe4\xd2\x40\x4e\x67\x3c\x59\x56\x2a\xbf\xb4\x50\xf6\x57\x7d\xee\x7f\x24\xa5\xd2\x5f\x66\xf9\x9a\x58\x93\xc7\x3e\xdb\x21\x49\xca\x3a\x4f\x64\x9d\x8c\xf8\x0e\x26\x51\x8e\x25\x26\x29\xc8\x19\x12\xe1\x70\xa0\x01\xee\x78\xd3\xd2\xb4\xeb\x75\x20\x59\x92\x13\x0a\xd9\x3c\x4b\xc4\x79\x24\xa2\x34\xda\x44\x0b\xc2\x93\x8d\x6a\xa4\x21\x88\x21\x93\xf9\x5a\x69\xfc\xd5\x53\x9c\x68\xa5\x0c\x5a\x04\xb2\x43\x6e\x77\xf8\x58\xd6\xaf\x17\xec\x03\xfd\xf1\xcc\xaa\x4e\xe8\x11\x79\x47\x20\x8e\x23\xdf\x1d\x78\x86\xf6\x6c\x5b\xf4\x25\xf1\x20\xdc\x27\x8a\xe1\x4e\x11\x77\xee\x70\x06\x3c\xb1\xfe\x1c\x18\xda\x82\x1f\xca\x33\x93\xaf\xe1\x21\xf0\x84\x6b\xad\x34\x8e\x59\xba\x1d\xdb\x4f\x18\xe3\xfd\xdd\xe0\x9d\x35\xd2\x77\xc1\x94\xc2\xc9\xb4\x45\x8d\x75\xc7\x1f\x4e\xc3\x45\xc3\x39\x07\x44\xc0\xc9\x34\x14\x8f\xdd\xbc\x1b\x5a\xa4\xa6\x6d\x07\xb3\xba\xec\x6d\x5a\xc1\x6b\xa3\xd5\x2d\x53\xee\x3a\x2a\x85\xe6\x24\xea\x3a\xa3\xee\x8e\xaa\xb9\xf1\xe3\x17\xa5\xce\x56\xd6\xa8\x14\xe4\xe4\xec\x84\x31\x7f\x8f\xf5\x03\xf6\x80\xbb\xbe\xda\x64\x06\xc3\xa7\x61\xe1\xab\x91\x63\x68\x4e\x98\x39\xe7\x8c\xf1\xd4\x6e\xef\xee\xce\x1c\x73\xd0\xef\xee\x8e\x1d\xce\x80\xf3\x22\x54\x28\x63\x7d\x66\xbb\xfa\x3a\x93\x45\xc5\x75\x1d\xc7\xe3\xef\xb9\x59\x24\xb5\xbd\x4d\x9d\x1f\xdf\xfb\xdb\x8a\x4d\x67\xea\x53\xe9\x6f\x3d\x85\x4e\xd8\xce\xc5\x73\x72\xae\x16\x33\x9d\x5c\x5c\xe0\xcd\xe1\x4c\x7d\xf0\xe5\xa2\x34\xe7\xc6\xae\x45\x8d\xfe\xf0\x68\xf4\x8c\x06\x2c\xb8\x0d\x7c\xe9\x81\x4b\xc3\xf8\x8d\xbf\x25\x7b\x82\x8f\x45\x20\x60\x92\x0b\x74\x12\x18\x77\x2d\xea\xdc\xd6\x9a\x0a\xc6\xc1\x69\x06\xe3\xb6\xb9\xbb\x33\x80\x17\xac\xd2\x9c\x71\xd7\xde\xdd\x79\xdd\xc1\xaf\x61\xe9\xcb\x11\x27\xf7\x99\xe6\x5d\x81\xb0\x8f\xcd\x17\x74\x88\x07\xac\x53\xc4\x29\xec\x72\x34\xcd\x55\x18\x3d\x20\x4e\xc3\x64\x70\x71\xcf\xfe\xdd\x59\x4b\x67\xd8\x0c\x1a\xb6\x80\x4b\x4f\x7b\x09\xbd\x40\x05\xb7\x57\x1a\x06\x75\xce\xf0\x21\xed\x17\xa8\xd7\xbd\xcb\xd5\xf7\xe0\x59\x3c\x0b\x3b\x97\xd6\x51\x77\xbc\xa3\xe3\x21\x25\x8f\x76\x67\xa5\xe1\xda\x8f\x1c\xa0\x0f\xdd\xd3\x57\x41\x58\x35\x9c\x25\xe0\x7d\x0e\x0a\x1d\x17\x2f\x5a\xc2\x85\x72\x0e\xa8\x10\x65\xc9\x75\x4d\xcc\x5c\x2c\x80\xcf\xc5\x82\xc6\x31\x91\x73\xb1\x60\x0a\xfd\x80\x99\xc2\x20\x36\x50\x20\x7f\xa3\x53\xe8\xb5\x4d\x73\xbc\x05\x1b\x94\x86\x50\xd7\xfc\x30\x4a\x84\x17\x7b\x17\x1d\x1e\x83\xf2\x2a\xeb\x97\x73\xca\x6f\x41\x22\xd8\xe5\xeb\x4c\xae\x78\x91\x4a\xe8\x6e\xa0\x34\xdc\xe2\x56\xf3\x2b\xa1\x9a\x3a\xe5\xed\x3e\xba\x64\x7b\x74\x39\x87\xdd\xb9\x3f\xef\x85\x9e\x06\x0c\x7f\x86\x0c\x77\xee\x08\x86\xb1\x9d\xa6\x52\x74\x60\x45\x69\x08\x25\x01\xec\xf3\x4e\x91\x46\xc4\xcd\xf7\x68\x0d\xac\x8d\x01\x7e\x77\x67\x95\xab\xc5\x23\xbb\x66\xbb\xee\x1e\x4c\x97\x36\x4d\x72\x0b\xb8\x85\xf4\x02\x94\x4c\x2f\xa1\xe6\x26\x7d\x0a\xc3\x41\xa4\x0a\x50\x52\xd3\x57\x70\x61\x8d\x56\xfa\x1c\xfc\x41\xa7\x57\xad\x8d\x4f\x1e\xb3\x5d\xc9\x4d\xbe\x7e\x81\x5b\x24\x9d\x2f\x8a\x7f\x4b\xb1\x4a\x4d\xeb\x44\x68\xc5\x0d\xe9\x55\xca\x82\x93\x28\x49\x3e\x14\xb2\x36\xa8\x54\x1f\x46\x1f\x70\x99\xab\x82\xff\xf2\xfc\xdb\xc7\xde\x10\x12\x93\x3c\xcd\x36\xa8\x78\x1b\x6e\xd6\xaa\x48\xa3\x15\x37\x51\x4b\x7b\x8b\xf5\xd1\xf4\x23\xeb\x36\x21\x73\x9b\xfa\x3c\x08\x93\xda\xd4\x24\xaf\x6a\x54\xd1\x01\x7a\xe7\x6f\x7d\xd2\x03\x1a\xef\x6c\xb6\xb4\xa5\xce\xc0\x11\xce\x3e\xb3\x5b\x50\x55\x67\xe8\x3a\x32\x30\x90\xa7\x2d\x38\x37\xd8\xef\xb2\x47\x04\xdd\x7e\xf9\x68\xbf\x36\xa6\xa1\xe1\xb5\x38\x5a\xde\x39\x0d\x30\xb0\xa3\x6e\xb6\x5c\x5f\x89\x5a\xe9\xe3\x0c\xe1\xfb\x0c\xd9\xaa\x7a\xc4\x11\xb4\xad\x0f\xa7\xd3\x80\x2b\xa3\xd5\xc3\x73\x9a\xf9\x4d\x5e\x67\x5a\x12\x0f\x0e\xfe\x07\xc6\xd0\x03\x53\xcc\x01\x53\x1c\x29\xd6\x09\x40\xa6\xa8\xed\x7b\xf0\xe4\x7d\x58\xf2\x66\x09\x79\x47\x86\x7c\xf4\x2f\x66\x48\xe0\xf6\xbe\x40\xbe\x54\x3c\xd3\x3f\x8b\x0d\x57\x8d\x71\x77\x92\x71\x1f\x81\x5a\xff\x1c\x18\xa9\x4e\x4f\x76\x39\xa1\x3b\x22\x31\xf8\x2d\xc4\x55\x44\x69\x92\x57\x59\x5d\xe3\x32\x2c\xca\x85\xce\x2b\x3e\x59\xae\x4e\x57\x9a\xdf\x4e\x96\x95\x90\x97\x42\xae\x26\xf5\x15\xaf\x0c\x3f\xfd\x68\xb9\xc9\xd5\xc3\xa8\x85\x8d\xc3\x5b\xe3\xe5\x65\x73\x51\x2e\xf0\x8f\xe3\x9c\xc8\x91\x69\xfa\xf2\x9f\x21\x80\xcb\xbf\xbe\xf0\x57\x7f\x79\x61\xcd\x8b\xbf\xbe\xec\x8f\x6f\x5c\x76\xd9\x18\xa3\x24\xae\x6c\xf8\x8d\x79\xac\xa4\x41\x27\x30\xca\xd7\x3c\x47\x36\x27\x49\x12\x81\x0c\xa9\x92\x22\xe7\x4b\xb3\xcf\x06\x90\x49\x21\xea\x6c\x59\xf1\x82\x9d\x4c\xdf\x83\xba\x6f\x43\xea\x82\x18\xaa\x0b\x34\x51\xdd\xda\x77\x25\x1a\xc1\x23\x40\x6b\x19\xe5\x95\xc8\x2f\x23\x50\xf4\x5d\x88\x7f\x0b\xb9\xb0\x1a\x61\x0c\xa9\xff\xfe\xcd\xd4\xdb\x1b\xf4\x3d\xc8\xcf\xb4\xf9\xd7\xd2\xff\x79\x40\x3f\x28\xd0\x90\x31\x0c\xd0\xf0\x36\x4b\x5e\x2a\x7d\xf9\x85\xd0\xa1\xc8\x0c\x82\x0a\x8a\x35\x24\xfa\xef\xff\xca\x4c\x3a\x89\x28\x68\xd6\x90\x8c\xc2\x06\x17\x2a\x95\x34\xa7\x2e\x89\x05\x91\x30\x59\x25\xf2\x28\x1c\x2a\xb3\x8d\xa8\x6e\x23\x88\x36\x4a\xaa\x7a\x9b\xe5\xdc\x0f\xe7\xaa\x52\x3a\x82\x68\xa5\xb3\x5b\xdf\xb7\x54\xba\xe0\xfa\xd4\x9e\x6c\x54\x89\xd5\xda\xe0\xe8\xa4\x50\xc6\xf0\x62\x72\xb6\xbd\xf1\x80\xdb\xac\xc0\xf8\xab\x83\x9c\x26\x1f\xf3\xcd\xfe\x50\xc5\x4b\x13\x41\xf4\xc9\xe1\x24\x8d\x88\xc7\x43\x9b\x4c\xaf\x84\x3c\x86\xae\x1b\xe9\xb0\x9d\x1e\xce\xf1\xd8\xdc\xd0\xc1\xe9\x80\x20\x0a\x24\x36\x1a\x24\x6d\x61\xeb\x23\x06\xc7\xf8\x38\xce\x4e\x18\x23\x87\x27\x81\x6e\xa6\x4e\x8a\xcc\x64\x18\x0a\xdf\xab\x51\x5f\x1c\x9c\x29\x94\x50\x0c\xd8\xbe\x57\xab\xaf\x44\xc5\xdf\x74\xae\x95\x5a\x75\xc7\x5a\x91\x28\x8b\x28\x64\xac\x21\x05\x46\x45\x6b\xcd\x4b\x56\xb2\x68\x6c\xc0\x3d\x6a\x14\xd2\x0f\xa2\x0f\x2b\xb5\x8a\xfe\x4f\x16\xfe\x82\x2c\x80\x20\x19\xe8\x63\x22\x51\xa0\x48\x1c\x1e\x22\x8a\x44\xe6\x44\xa2\xc0\x98\xd2\x83\x97\x08\xfe\x2e\xc7\xe4\x64\xca\x9d\xea\x1b\x64\xea\x87\xa3\x32\x05\x2b\xc6\x13\x89\xce\xc0\x9a\x75\x89\x8a\x03\xa9\x32\xda\x0a\x15\xfe\x5a\x77\x17\xc5\xca\xc9\x53\xf4\x6f\x72\x12\xfc\x17\x51\x28\x2d\x58\x11\x51\x28\x58\x43\xd6\xc8\xc8\x12\xa2\x6b\xa5\x8b\xd3\xa5\xe6\xd9\x65\x04\x91\x6d\x4f\xb3\xaa\x7a\x0b\x2f\x95\xe3\xa5\xfd\x2a\x5d\x53\x40\x39\x70\x96\x98\xc4\x45\xe5\x18\x4b\x3b\xb6\xd0\x38\x5e\x21\xdb\xfc\xa6\x02\x75\x5b\x51\x38\x3e\x61\x8d\x13\xfa\xcd\xe3\x8c\xc2\xcd\x58\xbf\x81\x99\xbf\xec\xbd\x1d\x81\x62\xbc\x43\x9e\x70\x69\xb4\xe0\x35\xe9\xcf\xea\x4b\x79\x25\xb4\x92\x18\xec\x20\xf7\xe6\x0b\x10\x6c\x3a\x13\x9f\x2a\x1f\x3c\x89\x0f\xd8\x19\xd5\x18\x7d\xfe\x40\xa6\xf0\x92\x70\x50\x20\xfa\x77\xc4\xe0\x20\xd0\x38\x47\xb4\x0f\x69\x0d\x9b\xce\xcc\xa7\xda\xa3\x31\x0e\x8d\x59\x24\x39\xa1\xb3\x91\xd1\xe1\xf2\xea\x6d\x06\x67\x36\xce\xc1\xe8\x51\x0e\x46\xcf\xd5\x22\x41\x9d\x70\xc1\x99\x3f\x02\x9b\x93\x3b\xe0\xe9\xee\x5d\x99\x31\xc4\xe6\x47\xf8\xd1\xbd\x13\xd4\xcc\xf3\x63\x86\x1c\x3a\xc7\x3f\x09\x2e\x5f\xd3\x94\xf4\x3c\xab\x29\xd8\x81\x9c\x74\x3f\x7a\x5a\x69\x6b\x63\x55\x31\x6c\xa8\xe7\xb6\x4d\x7b\xce\x7c\x3f\xf3\xcb\xb7\x7b\xa6\x37\x23\x36\x7f\x18\x9c\xfe\x4b\xff\x24\xe6\x48\xdc\x7f\x9f\x36\xc3\x0b\xb0\x95\x42\xfb\xe2\x39\x9f\x2e\x40\x39\x19\x73\xdf\x67\xf8\xcd\xb3\x7c\x7d\xd1\x75\xfa\x4f\x9b\x9f\x61\x12\xd4\xb0\xe0\x23\xa4\xe7\xc6\x65\xb6\xfa\x04\x44\x90\x99\x70\xeb\xa7\xae\x69\x61\x3f\x42\x31\x56\xa0\xfd\x3c\x21\x8d\x56\xfd\xf3\x51\x9f\x4a\xea\xf3\x90\xf3\x17\x0b\xe8\x1f\xae\x56\xdc\x74\x0f\x9a\x8f\x6e\xbf\x2d\x48\x34\x12\xa2\xce\x2c\xd0\xbb\xbb\x83\x34\x53\xf0\x08\x66\x12\x51\xb0\xa3\x13\xc1\x8c\xbd\xa7\xc4\x7a\xe7\xc9\x08\x74\x67\xad\x48\x6a\x47\xda\x44\xe9\x4c\xae\xf8\x51\x08\xcd\x8b\x36\x59\x56\xcd\xf1\x51\x1c\x68\x13\x1f\x00\xec\x81\x2c\xb3\xfc\x72\xa5\x55\x23\x8b\x53\x07\x5d\x89\x0d\xef\x56\x74\xae\xfb\xdb\x26\xb8\xc5\x5d\x7c\xf3\x76\xe4\x9d\xd5\x6b\x13\x17\x1e\xec\x4d\x28\x44\xbd\xad\xb2\xdb\x54\xc8\x4a\x48\x7e\x6a\x93\xe7\xb3\x6b\x51\x98\x75\xfa\x90\x6f\x66\x6b\x8e\xf3\xed\xcf\xce\x98\xea\xac\x10\x4d\x9d\x9e\xf1\xcd\xcc\x59\xac\x74\x9a\x3c\xe4\x9b\x36\xd1\xcd\xfe\x4e\x1d\x9a\xb3\xe9\xf4\x6f\xb3\xce\x56\xa6\x9f\x6c\x6f\xf6\x10\x7d\xbc\xbd\x69\x93\xce\x4d\xbd\x77\x33\xa9\x54\x92\xef\x4d\xfc\x7b\x8f\x2a\x75\x56\x3d\xab\xd7\xce\xaa\x7b\xc2\x3e\xde\xde\xcc\x9c\x43\x21\xfe\xe4\xe9\xd9\x27\xf7\xae\x94\xae\xd5\x15\xd7\xe1\x7a\xc8\xda\x99\xe3\xe0\xf5\x5a\x98\x7e\xed\xb0\x6b\x1f\xf5\x26\x13\xd2\xbd\x70\x1c\x67\x71\x59\xf1\x9b\x19\xfe\x39\x2d\x84\xe6\x56\x82\x53\xad\xae\xdb\x44\x3b\xe4\xfb\xbb\xef\x76\x6b\x0f\xa0\xc8\xf4\xe5\xfe\x2e\x47\xfe\x07\x1e\x48\x9b\x70\x79\x75\xff\x19\xa0\xf4\x9f\x66\x95\x58\xc9\x14\xdd\x8f\x59\xe0\x50\xa5\xbd\x3b\xd5\xe2\xd5\x89\x0a\xf2\x1e\x9b\xc8\x55\xd5\x6c\xe4\x6c\xa9\x6e\x4e\xeb\x75\x56\xa8\xeb\x8e\x5c\x7e\x3b\x99\x4e\xa6\x93\xbf\x1f\x3b\xf3\x91\x44\x84\x94\x3c\x57\x4b\x65\xd4\xa4\xe6\x5a\x94\xfe\x7c\x07\xd7\xed\x6c\x7b\x33\xa9\x55\x25\x0a\x7f\xc6\x67\x53\x64\xfd\x52\x19\xa3\x36\xa7\xdb\x4c\xf2\xea\x9d\x35\x62\x16\xf8\x62\x29\xba\x5b\xb3\xd0\x05\x1b\xf5\x38\xfc\xae\xcb\x13\xfe\xf1\x1e\xe1\x3d\x0b\xfd\x66\x3b\xa2\x10\x7d\xb8\xf1\xf1\xa8\x5d\x6b\xa4\x0b\x3e\x61\xb1\xb7\x8f\xfe\xad\x34\x1d\x5f\x6e\x7d\x7e\xe3\x61\x3d\x41\x15\x46\x03\x27\x4b\x21\x85\xe1\x6d\xf0\x2c\x7f\x7c\xd2\xee\xe3\xe9\xdf\x76\x6a\x9b\xe5\xc2\xdc\xa6\xd3\xb6\x8d\x40\x10\x03\xa3\x42\x02\xda\xda\x07\x3e\xf7\xb2\xd0\xa5\xa9\x7d\x1e\x9d\xcd\x17\x07\xcf\x41\xe3\x8a\xa0\x03\x07\x10\xd6\xb0\x85\x25\x5c\xc1\x05\xdc\xc2\x0d\x5c\xc2\x53\x78\x05\xcf\xe0\x39\x5c\xc3\x63\x78\x01\x3f\xc0\x4b\x78\x04\xaf\xe1\x4f\x78\x00\x5f\xc3\x37\xf0\x04\xbe\x83\x9f\xe0\x57\xf8\x8d\x8d\x9c\x51\xf8\x7d\xf8\x7e\xac\x36\x9b\x4c\x16\xf0\xc7\xd0\xf5\xb9\x5e\xd5\xdd\x33\x31\xba\x89\xc6\x30\xfb\x8c\xdc\x8f\x3f\xe7\x36\x74\x3e\x8f\xfe\xe7\x3f\xfe\x33\x4a\xf7\xbb\x81\x1b\x46\xf6\x3b\xbb\x24\xd5\x87\x67\xfc\xff\xd1\xc4\xa8\xaf\xc4\x0d\x2f\xc8\x43\x0a\x32\x84\x7d\x61\xd4\xf6\x1e\xc0\x20\xd8\x37\xc1\x4b\x24\xda\x4c\x9b\x82\xf3\x15\x03\x3f\x0f\xa9\xd7\xf3\x2f\xd3\xaf\x6c\x66\x5a\x1b\xa6\x0c\x41\x7f\xd5\x30\x6d\x88\xf3\x9f\xfa\x42\xa5\x37\xa2\xfb\x31\x40\xf7\x6d\xfa\xbd\x45\x97\x1b\x56\x5b\x74\x99\x61\xb9\x43\x07\x95\x39\x88\x1e\xe3\xf8\x73\x32\xc5\xb1\xc6\x1c\x84\x11\x71\xfc\x85\x1b\x2b\x83\xb1\xc0\xd9\x8a\xe3\x5f\xec\xf8\xbd\x31\x63\xff\x53\x18\xe7\x46\x8d\xfd\x7b\x1b\x41\x56\x24\x5a\x9f\x59\x2f\xbf\x21\xbf\x39\x17\xbf\x03\x89\x28\xac\x06\x14\x99\x43\xb1\x1e\x8d\x6f\x87\xf1\xe5\xf0\xf3\x0a\x7f\x6e\x35\x06\x8d\xe4\xc2\xba\x0d\xdb\xec\x48\x2e\x6b\x53\xa4\x11\xdc\x22\x3e\x44\x75\xc3\x1a\xf2\x3b\x85\x4b\xff\xfd\x94\x35\xe4\x0f\x0a\xaf\x8e\x86\x24\x95\x89\xe3\xca\x11\xf4\xec\x28\x40\x63\xe2\xb8\x71\x00\xcf\x47\x14\x97\x26\x8e\x4b\x37\x70\xed\x06\xfa\xa1\xc7\xc3\x0e\x5e\x0c\x54\xc3\x0f\x08\xa6\x9d\x7c\xda\xf0\xfb\x25\x6b\x88\x31\x14\x1e\x1d\x70\xf3\x75\x30\xed\xcf\x60\xda\xa4\xcb\xb7\xda\xe9\x0f\x58\x43\xb8\xa1\xf0\x35\x02\xd4\x11\x85\x6f\x0e\xf0\x3c\x09\xf0\x7c\x67\xc1\x8c\xda\x8e\x90\xfc\xc4\x1a\x82\xa1\xc7\xaf\x1e\xc9\x86\x64\x3e\xb8\xc5\x60\xd8\xf6\x28\x88\x3a\x83\x12\x41\x84\xd6\xc4\xf7\x8e\x2d\x4b\x04\x91\x56\xd7\x76\xec\x6a\x0f\xc7\x32\x0c\x38\x74\x73\x90\x5e\xdc\x86\xe3\xbd\x65\xde\x87\x7a\x11\x42\xb9\x6c\x6d\x67\x8c\xf7\x21\x5f\x87\x90\xe8\xdf\xdd\x07\xf8\x24\x04\x74\x1e\xe4\x3e\xc8\xe3\x11\xae\xc0\x74\x1d\xe6\x48\xc3\xec\xae\x13\xcf\xb7\x45\x56\x7d\x50\x6b\x12\xe4\xa8\x7b\x3c\x0e\x22\x5c\xe5\x22\xdc\xcc\x45\xb8\x36\xd0\x5d\x61\x93\x21\xfc\xaa\x87\x5f\xbb\xa1\xad\x6b\x96\xb0\xc5\xe6\x0a\x96\xd8\x5c\xc0\x15\x36\xb7\xae\xb9\x71\xcd\xa5\x6b\x9e\xba\xe6\x15\x42\x76\xba\xb0\x21\xcb\x1e\xed\x33\xec\xef\x54\x20\xec\x7f\x8e\xeb\x74\x1a\xd0\x47\x53\x20\xc8\xb5\x5b\xff\xb1\x6b\x5e\xc0\x63\x6c\x7e\x80\x17\xd8\xbc\x74\xcd\x23\xd7\xf9\xda\x35\x7f\xc2\x6b\x6c\x1e\xb8\xe6\x6b\xd7\x7c\xe3\xc6\x9e\xb8\xe6\x3b\x78\x82\xcd\x4f\xae\xf9\x15\x9e\xd8\x10\x93\x83\xa0\x3b\x6d\x30\x3e\x77\xf7\xad\xa0\x18\x9e\x8b\xae\xa4\x85\xf8\xbb\x57\x50\x6a\x75\xb4\xe3\xaf\xa6\xae\x44\xc5\xa5\x51\x7e\xc3\xe9\xbf\x31\x11\xda\x2a\xc4\x52\xba\x20\xff\x37\x0a\xb9\x5d\xc1\x5d\xc1\x6e\x85\xcc\xaf\xe0\xaf\x63\xbf\x42\x78\x22\xe1\x22\xbf\x23\x86\xdf\x87\x45\x3a\x03\x88\xb8\x6e\xdc\x3a\xbf\x87\xe0\x7f\x20\xf8\x1f\x03\xf8\xd8\x38\xe2\xac\xa7\x6e\xd6\x1f\x14\xc4\x9e\x0d\x38\xaf\xcc\x79\x65\x12\xc7\x9d\x94\x90\xca\x30\x34\x0a\x9e\xc2\xee\x70\x9f\x51\x9a\xe2\x59\x93\xaa\xdb\x4a\xe5\xaa\x42\x02\x7c\x9d\xdd\x38\x6f\xcc\x79\x13\xe0\x6b\x0c\xfb\x22\xc0\x17\x0a\x05\x4d\x51\x4c\x48\xd3\xa1\x6c\x0e\x50\x06\xe6\xe6\xbc\xc4\xff\x07\xb4\xa5\x61\xbf\x04\x68\x3b\x99\xba\xa6\x34\x45\x11\x23\x65\x87\xb3\xf4\x38\x07\x66\x19\x7b\x3c\xde\x61\x10\xc7\x1d\x86\xfd\x6e\xe4\xe1\x4b\xc7\x43\xbc\x7d\x07\x6c\xdc\x62\x43\x77\x62\x7f\xca\x3d\x5e\x02\x62\x7a\xe0\x30\xf1\x11\x26\x5b\x0f\x44\x64\x88\xe9\x0d\xce\x06\xa2\xf9\xc9\xa1\x91\x66\x3f\xb9\x6f\x25\xda\x4a\x57\x41\x7a\x1d\x2d\x48\xaf\x96\x05\xe9\x35\xb1\x20\xb4\x6d\xdb\x2e\x13\x70\xf0\xc0\x3f\xae\xa9\x70\xc5\x20\xb6\x0a\x6c\x54\x72\xe5\x2a\x35\x5d\x75\x58\xf0\xc0\xc7\x6a\xee\x99\x40\x86\xc1\xf0\xd5\x31\x29\x85\xcc\xaa\xea\xd6\x0d\x73\x62\xdf\x53\x1f\xf2\xbf\xd3\x76\x76\x08\xdb\x57\xe5\x71\x42\xdd\x5b\x64\x57\x9b\xf3\xe6\x62\x84\xb1\xb3\x7a\xac\x7e\xfc\x6c\x28\x1e\xb7\x25\x29\x21\x17\x8e\x14\x32\xb4\xd4\x66\x40\x4c\xa6\x57\xbc\xaf\xb3\x18\xaa\x0f\x50\x1a\x5d\x97\x2b\x74\xf3\x90\x60\xfa\x7a\xb7\x67\x7b\x3e\x36\xa5\xad\x26\x8f\x86\x62\x7a\xb8\xa6\x30\xee\x78\xec\xaa\xa1\x5f\xb3\x47\x83\x57\xf8\xe7\x1b\xdf\xf0\xb6\x07\x2e\x4f\xa5\xb2\xa2\x7b\xbd\x3b\xf2\x38\xb4\x4d\xd5\xbd\xc9\xc9\x07\x61\x0d\x91\x2f\x6e\xf0\x82\xdb\x82\x62\x92\x5f\x4f\x5e\x93\x1d\x4a\x4b\xea\xe2\x08\xb0\x65\x48\x69\x57\x8e\x04\x28\xa7\xa9\x6c\x47\xce\xa2\x1a\x33\xad\xa7\x49\xf5\x9c\x73\x84\x8d\x16\x6f\x67\x3e\x2b\x58\xc7\x31\x91\xdd\xef\xde\x41\xa5\xa0\x5c\x71\x91\xf4\x2a\xa1\x7c\x1a\x8a\x8c\x53\x6e\x5f\x1f\x4b\xb8\x76\x98\xdf\x92\x53\x45\x76\x3c\x39\x96\x53\x7d\xb7\x34\x2a\x2b\x83\xcd\xbe\x43\x8e\xd4\xc6\x02\xc1\x41\x39\x33\xe6\xea\x08\x1d\xbd\x74\xa7\xfa\x7b\xac\xee\x73\xa0\x35\x9b\xce\xea\x81\xfe\x3a\xc8\x81\xe6\xec\x09\x11\xa0\xa0\xa6\x33\x3d\xaf\x17\xe7\xf8\xc7\x5e\xad\xb9\xcd\x81\xd6\x6e\x8f\x39\x05\x3b\xe0\x72\xa0\xb5\xcd\x81\x06\xff\x0e\x02\xa4\xcf\x85\xd6\x03\xe1\xb5\x23\xbc\x7e\x5b\x2e\xd4\xe5\x40\xe1\x88\xb0\x7d\x13\x96\x13\xee\xf2\x54\xc1\x26\x55\x4e\x3c\x53\x15\xc0\x3d\x79\xf7\x9c\xa9\x17\x92\xb9\x3c\x92\x22\x75\x83\xc7\x92\xa4\xdf\xbd\x29\x49\xda\xa9\x41\x9d\x1e\x2b\x60\xa9\xa3\x7b\xcb\x76\x0e\x4b\x73\x7c\x25\xb2\xbd\xce\x92\x4d\xb6\xc5\xce\x77\x2f\x8b\x31\xf7\x97\x08\xf5\x6b\x51\xea\x4b\x37\xc2\x6b\x7b\xaf\xd0\xc7\x5e\xbe\xf7\x65\x76\xef\x4d\x0e\xf8\xea\xb4\xa0\x22\x6f\xe8\x0d\xaa\xee\xde\x39\x8f\x60\xeb\xfe\x05\x5e\x33\xdd\x0e\x53\x03\xb9\xb9\x49\x79\x7f\x17\xdb\x6a\xd4\xae\xb4\x36\xfd\x13\x70\xb7\xe9\xd7\x60\xf7\x97\x7e\xd3\xfd\x5b\xac\xa8\x3b\x9e\x08\xec\xee\xd2\x08\x27\x45\xfd\xbf\x73\xda\x92\x50\xdd\x05\x85\x2e\x7c\x2d\xd1\xeb\xeb\x6a\x98\xc3\x3b\x69\xf0\xbc\xbb\x41\xec\x17\xdd\x85\xce\x6c\xbf\xbd\xb3\x98\x35\x50\xa1\x96\x74\x37\x98\xa4\x3b\xce\x24\x7a\x36\xe6\x86\x71\xe8\xc9\x13\xb6\x52\x59\xa1\xe5\x0f\x28\xa2\x71\xbc\x25\x78\xbd\xdc\xdd\xf9\x05\x11\xc9\x50\x34\x28\x86\x12\xf0\x03\xdb\xdf\x97\x60\x1b\x0a\xc2\xd5\x66\x1f\x37\xf0\xff\x8c\x09\x73\xc3\x7b\xc5\x99\xc7\xeb\x32\x0f\x4b\x32\x8f\x54\x63\x0e\x85\x98\x9a\x7c\x37\xb2\x84\x33\x34\x2e\xdf\x91\x9d\x23\x21\xbd\xff\x21\x21\xd0\x3e\xea\x0c\xce\xae\x6d\x69\xbb\xa0\xb3\xff\x0d\x00\x00\xff\xff\x21\xd8\x47\x3d\x00\x3a\x00\x00")

func mainJsBytes() ([]byte, error) {
	return bindataRead(
		_mainJs,
		"main.js",
	)
}

func mainJs() (*asset, error) {
	bytes, err := mainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.js", size: 14848, mode: os.FileMode(420), modTime: time.Unix(1530193393, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
	"main.js": mainJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"main.js": &bintree{mainJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
