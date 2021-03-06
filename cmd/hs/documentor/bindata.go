// Code generated by go-bindata.
// sources:
// demo.css
// demo.js
// demo.tmpl
// item.tmpl
// DO NOT EDIT!

package documentor

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

var _demoCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xcd\x6e\xdb\x30\x10\x84\xef\x7a\x8a\x85\x8b\x5e\x02\xd1\x3f\xb1\x9d\x36\xca\xa5\x40\xdb\x5b\xfb\x10\x14\xb9\x72\x09\x89\xbb\x02\xb9\x4a\xac\x16\x7e\xf7\x42\x3f\x96\xe1\xc4\x07\x07\x88\x6f\x92\x38\x33\x1f\x77\x28\x2e\xee\x40\xd8\x32\x14\x8e\x2c\x68\x78\xd1\x2d\x08\x43\x40\x09\x0e\x9f\x11\xe4\x0f\x82\xe1\x8a\x43\x84\x22\xb0\x87\x82\x83\xd7\x22\x18\x16\xd3\xd3\x7c\xc7\x70\xb7\x48\x92\x9c\x6d\x0b\xff\x12\x00\x80\x82\x49\x54\xa1\xbd\xab\xda\x0c\x66\xbf\x1a\xe3\xac\x86\xef\x4c\x91\x2b\x9c\xa5\xf0\x9b\x49\x1b\x4e\xc1\x33\x71\xac\xb5\xc1\xa7\xe4\x90\x24\x73\xdc\xa3\x69\x04\xad\x32\xde\x8e\x41\x3d\x3a\x83\x4f\xeb\x2f\x0f\x5b\xbd\x79\xea\xbf\x79\x1d\x76\x8e\x54\xce\x22\xec\x33\xb8\xdf\xd6\xfb\xc1\xaf\x8d\xb8\x67\x54\x1e\xa9\x49\x61\x5e\xeb\x80\x24\xd3\x5b\x60\x5f\xcb\xab\x54\xb3\x79\x2c\x96\x0f\x27\x33\x53\x9c\xa4\x29\xcc\x9d\xa0\x8f\xa3\x65\x84\x0a\xd7\x19\xac\x26\x62\x89\x6d\x0a\xf3\xb7\xdb\xdd\xe0\xa3\x5e\x9a\x41\xd3\xa5\x28\x8b\xd1\x9c\x29\x8f\x81\x15\x16\x92\xc1\x6a\xd9\x25\xc2\xab\x21\x46\xa9\xe0\x5e\x94\x45\xc3\x41\x77\x5b\xcc\xa0\x21\x8b\xa1\x72\x34\xd6\x26\xce\xe3\x39\x7f\x17\xb0\x1d\xd6\xc6\xe9\x2f\x34\x7b\x81\x7f\x95\xe1\x58\xfb\xa9\x84\x28\x96\x1b\xb9\xd0\xd3\x7a\x98\xea\xad\x77\x3d\xe1\xf2\xca\x51\xe9\x68\xa7\x4c\x13\x22\x87\x3e\xe4\x58\xe2\xfd\xcf\xf5\x8f\xcd\xd7\x2e\x40\xbd\x60\x5e\x3a\x51\x9a\x9c\x1f\x3b\x58\x45\xe8\xbd\x10\x05\x6b\x85\x64\xc1\x51\xe1\xc8\x09\xf6\x06\xcf\x7f\xdf\xa3\x8e\xef\x10\xf3\xf5\xda\x2b\x85\x87\x24\xf9\x56\x62\x5b\x04\xed\x31\xc2\xac\xd7\xcd\xfa\x2a\xba\x2b\x97\x76\xf7\xf1\xec\x78\x25\x68\x8a\xc3\xef\xdd\x51\x0e\x09\xc0\x76\xf9\xf9\x5c\x93\x57\xda\x94\xc3\x6a\x17\xdf\x17\x72\x62\x0c\x5b\xf9\x58\xc2\x78\x46\x37\x1f\x24\xde\x1c\xc1\x37\x25\xfc\x0f\x00\x00\xff\xff\x17\xfd\x8f\x2c\x71\x05\x00\x00")

func demoCssBytes() ([]byte, error) {
	return bindataRead(
		_demoCss,
		"demo.css",
	)
}

func demoCss() (*asset, error) {
	bytes, err := demoCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "demo.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _demoJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x57\xdb\x6e\xdb\x38\x13\xbe\xae\x9f\x62\x7e\xe3\x47\x25\x61\x63\xb9\xbd\x75\x63\x05\xdd\x74\x81\x14\xdb\xec\x02\xdb\xee\x55\x50\x14\xb4\x38\xb6\x08\x4b\xa4\x40\x52\x49\x0c\x45\xef\xd2\x67\xe9\x93\x2d\x48\x9d\xa8\x83\x37\xe9\xea\xc2\x91\xc8\xe1\xcc\x37\x87\x6f\x38\xb9\x27\x12\x3e\x6a\xcc\x60\x0b\x7f\x21\x89\x75\x18\x4b\x24\x1a\xaf\x53\xa2\x94\x5f\x2e\x00\x24\x72\x8a\x72\x03\xfb\x82\xc7\x9a\x09\xee\x07\x60\x96\xcd\x86\x2e\x24\x87\x4b\xca\xee\x21\x36\xe2\x7f\x90\x0c\xb7\x4b\xa6\x31\x5b\x46\x56\xc2\x3c\x97\x2a\x27\xdc\xd9\xf7\x8e\x78\xf2\xa2\x52\x27\x4c\x85\xb9\x14\xb9\x0a\xcd\x81\xdf\xf1\x54\x5d\xae\x8d\xe8\xf9\x93\x33\x67\x60\xbb\x05\xcf\x83\xa7\x27\x18\xed\x85\x14\x55\xdc\xec\x5e\x99\x9f\x0d\x78\x66\x7d\x65\xd6\xbd\x0a\x28\xe1\x07\x94\xa2\x50\xe9\xe9\x33\xea\x8f\x9c\xa3\xbc\xf9\x72\xfb\x69\x5b\x96\xdf\xbe\x25\x3a\x4b\x37\xb3\x0a\xab\x2a\x7a\x0e\xa4\x17\x67\xd4\xfb\x0f\xea\xe3\x8c\x36\x70\x5f\xbf\x9e\x6c\x9a\x1f\x15\xa6\xc8\x0f\x3a\x81\x08\xde\x18\x97\x22\xe3\xd3\x8c\x96\x11\xc6\xcb\x35\x65\xf7\xd1\xbb\x05\x40\xb5\xa8\x82\x77\x8b\x85\xc9\xf7\x2d\xf2\xe2\xe7\xf3\x6d\x4e\x5a\x24\xb0\x3d\x83\x30\x23\xb9\xdf\x1d\xc3\xb4\x3d\x08\xf0\xaa\xad\x15\x5b\x69\x4d\xf2\xb6\x25\xa6\xe1\x11\x4f\x95\x5d\xb0\x5f\xe6\xa5\x5a\xd7\xc8\xab\x60\xd1\x99\xd5\x92\xb0\xf4\x53\xed\xff\xc0\xb8\xdd\x68\x22\x33\x94\x9e\x93\x1b\xe0\x33\xb6\x2e\x80\x3d\x06\x65\x97\x4c\xeb\xa1\x7a\x1f\x6b\x76\x8f\xb0\x05\xf6\x68\x52\xe2\xda\x5e\xc1\xdb\x4e\x78\xbe\xfc\xcb\xee\xfc\x15\x78\xc4\xbe\xad\x32\xe4\x85\x07\x9b\x1f\xdf\xbd\x9c\x48\xe4\xba\x5e\xa8\xe0\x68\x62\x60\xc3\x67\xa2\x10\x95\x7d\x9d\xd5\x59\x1b\x84\xa1\x31\xe7\x37\xf6\xc7\xac\x33\x2a\x1d\xd6\x95\x16\x75\xd5\x97\xe9\x0c\x49\xd5\xb2\x36\xa9\x5c\x73\x73\xc2\xc4\x06\x4c\x2d\xa3\x71\xb1\x2f\x8f\x78\x5a\x46\x2a\x27\x31\xee\x88\x6c\xca\x0e\x76\x24\x3e\x3e\xa3\x31\x97\x22\xcb\xb5\x83\xd7\x3c\x57\xe0\x32\xbc\x50\x28\xaf\x13\xc1\x62\x13\xc9\xf9\xf5\xcd\x8f\xef\x13\xfa\xed\x52\xc6\x8f\x8c\x1f\x56\x71\x21\x95\x90\x5e\xf4\xd4\xc0\x72\x62\xe1\x62\x73\x3e\x82\x19\x96\x5c\x0b\xae\xa5\x48\x53\x94\x67\xf9\x72\x30\x04\x67\x9a\x91\xf4\xb3\x26\x1a\xcf\x37\xca\xd2\x96\x03\xd1\x48\x0d\x0b\xd4\x06\xee\x6c\xf8\x27\x2c\xbe\xa8\x4b\x6e\x03\x77\xa3\x8d\xaf\xd5\xd7\xca\x62\xbc\x58\x00\x24\x84\xd3\x14\xff\xee\xa2\xe1\xd8\xed\x43\x64\x10\x58\x08\x6c\x0f\xce\x72\xdd\x69\x82\x06\x58\x4f\xb3\x94\x28\xfd\xde\x05\xd9\x92\x48\x19\xcf\xc2\x21\xfe\xbb\xb3\x3b\x6d\xa7\x5a\xc1\xdb\xaf\x8b\x09\x2b\x27\x46\x1a\x72\xaa\x94\xc5\xe8\x07\x9d\x3c\xc7\x47\x6d\xb6\xcf\xe1\x07\xaf\xef\x2f\xb5\x80\xdb\x0c\x8c\xc8\xdb\xce\xc3\x56\xaa\x96\xc8\x45\xde\xd8\x31\x4f\x6b\x07\x1a\xaa\xdf\x0d\xd4\x58\x1f\x5a\x51\x17\x01\x78\x6d\xdd\x7b\x35\x4f\x01\x53\x85\x0e\x22\x47\xef\xb7\x70\xcf\x38\xf5\xa7\x9e\xf7\xbd\xf3\x02\x66\x1b\xe7\xa0\xd5\xd4\xed\xd2\x78\xd6\x03\xe9\xe4\xaa\x60\x10\x8b\xde\xfa\x16\x0a\x4e\x71\xcf\x38\xd2\x49\x3c\x1c\x8c\xed\xab\x85\x33\xab\xca\xde\x51\xff\xab\x2b\x67\x88\xef\x7c\x25\xe4\x85\x4a\xfc\xa6\xcc\x5b\x3d\x0e\xd2\x11\x86\xf9\x00\xf5\x2e\x8e\x43\xec\xa4\xd4\xd8\x69\x35\xf5\xfa\x6b\xd2\x57\x75\x11\x4e\xb5\x0f\xf2\xe9\xc4\x74\xf1\xf3\x4e\x75\xa4\xad\x1b\x6f\xe0\xa8\x40\x6d\x1b\x83\x3f\xe1\xff\x59\x0b\x6d\xcf\x7f\x60\x9c\x8a\x87\x50\xc5\xa6\x05\x7d\x11\xfe\x9b\x0b\x2a\xe2\x22\x43\xae\xc3\x9d\xa0\xa7\x66\xe3\x06\xd9\x21\xd1\x41\xd7\x19\x62\x91\xe5\x82\x23\xd7\x1f\x18\xbd\x15\x05\xd7\x7e\x77\xc5\x19\x5e\x29\x4c\xf7\x0d\xaf\x9b\xd5\xa1\x52\xc1\x8f\x78\xca\x25\x2a\x73\xcf\xf7\x45\xe9\x5c\x93\x46\x43\x38\xee\x3f\xfe\x67\x2d\x19\x3f\x84\x7b\x29\xb2\xeb\x84\xc8\x6b\x41\xd1\xc7\xf0\x21\x61\x71\x12\xb4\x19\xe9\xdb\xd7\xec\x94\x31\x3f\x66\xcc\xe6\xe0\xd9\x51\xa3\x99\x26\x9c\xc1\xea\xaa\x95\xb0\x7d\xdf\xce\x40\x36\x59\x76\xf2\xa8\xd3\x36\x9a\x44\x9c\x9a\xb0\x8b\xfd\x67\xb5\x8e\x60\x33\xd0\xe7\xdc\x29\xaf\xda\x40\x8d\xae\x3d\x0f\x1f\x31\x2e\x34\xd2\xd5\x0b\x87\x44\x6f\xbd\x63\x7c\xbd\x23\x2a\x81\x55\x0c\x1e\xfc\xe2\x7a\x65\x07\xbd\x97\x58\x55\x9a\x8a\x42\x7b\x13\x39\xf3\xf8\x3a\x41\x10\x85\xce\x0b\x0d\x62\x0f\xe6\x2b\x16\x59\x46\x38\x85\x07\x51\xa4\x14\x76\x08\x94\xa9\x3c\x25\x27\xa4\x90\xa0\xc4\x60\x62\x6e\x78\xd5\x8f\x96\xaa\xbe\xa1\x6b\x96\x19\x9e\x65\xc2\xd4\x9a\x1f\x84\x7b\x21\x33\xa2\x7d\xef\xe6\x66\x93\x65\x9e\xad\xdf\x17\x8d\x39\xfd\x9d\xbc\x1c\x4e\x18\xd3\xf9\xc4\x98\x5c\x46\xa5\xf9\xd3\xfe\x7b\x01\xff\x87\x44\xc1\x6a\x0f\xe1\xba\xdc\xb3\x14\x39\xb1\x7b\xc3\x71\x65\x27\xd7\xce\x2c\x55\x0f\x49\xff\x3e\x30\xd8\xd9\xe0\xc3\x9f\xb7\x61\x5d\xd8\x06\xff\xe5\x68\x80\xa8\x6b\xab\x2e\xac\x75\x64\x48\xd0\x11\xef\x80\xfa\xb7\x14\xcd\xeb\xaf\xa7\x8f\xd4\xf7\x62\xc1\x35\x61\x1c\xa5\x17\x2c\x82\x77\xff\x04\x00\x00\xff\xff\x28\xdf\x9a\xb3\xa3\x0d\x00\x00")

func demoJsBytes() ([]byte, error) {
	return bindataRead(
		_demoJs,
		"demo.js",
	)
}

func demoJs() (*asset, error) {
	bytes, err := demoJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "demo.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _demoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x31\x6f\x1c\x21\x10\x85\xfb\xfd\x15\x13\x1a\x57\x3b\xf8\xce\x89\x72\x8a\xd8\x6d\xec\x58\x49\x9a\xa4\x70\x8a\x94\x2c\x8c\x75\x9c\x60\x39\xc1\xf8\xe2\x13\xda\xff\x1e\xed\x71\x9b\x44\x72\xbb\x15\x9a\x61\x78\xdf\x43\x3c\xd4\xbb\x87\xef\xf7\x4f\xbf\x7e\x7c\x86\x3d\x07\xdf\x37\x6a\x59\x48\xdb\xbe\x01\x00\x50\x81\x58\x83\xd9\xeb\x94\x89\x3b\xf1\xf3\xe9\xb1\xdd\x09\x90\xd7\x4d\x76\xec\xa9\xff\x12\x39\xef\xc9\x7b\xb0\x14\x22\xb4\x50\x0a\x7e\x65\x0a\xf8\x40\xd9\x4c\x93\x92\x75\xaa\x9e\xc8\x26\xb9\x23\x43\x4e\xa6\x13\x7b\xe6\x63\xfe\x24\xa5\xb1\xe3\x21\xa3\xf1\xf1\xc5\x3e\x7b\x9d\x08\x4d\x0c\x52\x1f\xf4\xab\xf4\x6e\xc8\x32\x91\x36\x2c\x6f\x71\xf3\x1e\x3f\xd6\x02\x83\x1b\xf1\x90\x45\xaf\x64\xd5\x5b\x53\xbc\xb5\x31\xac\x08\x18\xf4\x40\xbe\x35\x31\x91\xfc\x80\x3b\xdc\xde\xc9\x21\xc5\xdf\x99\xd2\x8a\x8c\x97\xd1\x52\xca\x33\x03\x0f\x59\x6e\x70\x87\x77\xff\xf5\xda\xf5\x40\x21\x06\x1a\x79\x86\x6c\x71\xb3\xc5\xdb\xa5\xf1\x96\x70\x45\xf0\xd9\x53\x5f\x0a\xde\xe7\x3c\x27\xa1\xd6\x8d\x92\x35\x60\x6a\x88\xf6\xdc\x37\xca\xba\x13\x38\xdb\x09\x13\x47\xd6\x6e\xa4\x34\x2b\x59\x77\xea\x1b\xf5\x57\xef\xa4\x13\x38\xa6\x00\x1d\x94\xc2\x14\x8e\x5e\x33\x81\x98\x5b\x02\x2e\x79\x9b\xa6\xcb\xd0\xb3\xf3\x34\xea\x40\xd0\xc1\x4d\x29\xf8\x78\x2d\xa7\xe9\xa6\x69\xfe\xf9\x5b\x6e\xcf\xe7\x23\x75\x82\xe9\x95\xeb\x4b\x89\xd9\xed\xb7\x6a\x76\x19\x95\x57\x9b\xf2\xf2\x3b\xfe\x04\x00\x00\xff\xff\x27\x72\x5c\x2e\x34\x03\x00\x00")

func demoTmplBytes() ([]byte, error) {
	return bindataRead(
		_demoTmpl,
		"demo.tmpl",
	)
}

func demoTmpl() (*asset, error) {
	bytes, err := demoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "demo.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _itemTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x0a\x02\x31\x10\x44\xfb\x7c\xc5\x70\x08\xd7\x1c\xf7\x01\x01\x2b\x6d\xc4\xd2\x52\x2c\x42\xb2\x4a\xd0\x04\x31\x69\xce\x75\xff\x5d\xb2\x77\x8a\x85\xdb\x0d\xf3\x78\x3b\xcc\x81\xce\x31\x13\xba\x58\x29\x75\x22\x86\x0d\x00\x04\x2a\xde\xa2\x67\x1e\xb7\x54\x3c\x5e\x38\xb8\x1c\x6b\x7c\x92\x48\x3f\x28\xe1\x53\x98\x81\x4d\x0a\xff\xfa\xe6\x2b\x16\x47\x0d\xed\x98\x1f\x2e\x5f\x08\xab\xd6\xc0\xae\x31\xee\x1a\x22\xf2\x25\x94\xba\xd2\xa4\x5e\xc5\xc6\x3d\x4d\xcd\xa8\x36\x0b\xe6\x4a\xe9\x7e\x73\xf5\xb3\x77\x96\x89\xc8\xf0\xf3\x86\x72\x58\xa4\x27\x23\x66\xc9\xef\x00\x00\x00\xff\xff\xef\x4c\xe8\xcf\xe9\x00\x00\x00")

func itemTmplBytes() ([]byte, error) {
	return bindataRead(
		_itemTmpl,
		"item.tmpl",
	)
}

func itemTmpl() (*asset, error) {
	bytes, err := itemTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "item.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"demo.css":  demoCss,
	"demo.js":   demoJs,
	"demo.tmpl": demoTmpl,
	"item.tmpl": itemTmpl,
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
	"demo.css":  &bintree{demoCss, map[string]*bintree{}},
	"demo.js":   &bintree{demoJs, map[string]*bintree{}},
	"demo.tmpl": &bintree{demoTmpl, map[string]*bintree{}},
	"item.tmpl": &bintree{itemTmpl, map[string]*bintree{}},
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
