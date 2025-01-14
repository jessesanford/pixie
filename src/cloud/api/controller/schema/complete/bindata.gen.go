// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\x0a\x02\x31\x10\x45\xfb\x39\xc5\x4f\xa7\x57\x48\x67\x23\x58\x28\x88\xa5\x58\x0c\xeb\xec\x1a\xd8\x24\x4b\x66\x14\x17\xd9\xbb\x0b\x81\x88\x62\x37\xbc\x79\xbc\x2f\x4f\x93\x74\x85\xcd\x93\xe0\x78\x97\x32\xe3\x45\x00\x17\x0b\x3d\x77\xa6\xab\x76\x1d\x38\x8a\xc7\xc9\x4a\x48\x83\x5b\x7b\x6c\x9a\xb1\x4b\x7d\x76\xb4\x10\xd5\xc4\x0f\xae\xa9\x60\x12\xd5\xe3\xdc\x3e\xee\xf2\x6f\x57\xf1\x21\x45\x43\x4e\x9f\x11\x02\xba\x1b\xa7\x41\xc6\x3c\x7c\x43\x0b\x51\xd4\x38\x4e\x7b\xf5\xd8\x8e\x99\xcd\xd1\xf2\x0e\x00\x00\xff\xff\xc2\xab\x64\xad\xc7\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 199, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xdf\x6f\xe3\x36\xf2\x7f\xf7\x5f\x31\xdb\x7d\x68\x02\x04\x8b\xe2\x8b\x6f\x8b\x83\x9f\x4e\xb5\xd5\xae\x2e\x89\xe3\x8b\x9d\xdd\x2b\x0e\x8b\x05\x2d\x8e\x2d\x22\x12\xa9\x92\x54\x12\xdf\x61\xff\xf7\xc3\xf0\x87\x24\xda\x4a\xf7\xb6\xb8\x37\x8b\x3f\x3e\xf3\x99\xe1\xcc\x70\x86\xc6\x17\x8b\x92\x83\x3d\xb6\x08\x7f\xef\x50\x1f\xe1\xdf\x33\x80\xce\xa0\x9e\xc3\x83\x41\x5d\xc8\xbd\x7a\x33\x03\x50\xfa\x30\x87\x3b\x7d\x88\xdf\xb4\x62\x83\xd6\x0a\x79\x30\x7e\x65\xfc\x8a\xb3\x99\xb5\x5a\xec\x3a\x8b\x61\x7e\xf8\x0e\x78\x34\x68\xe6\xf0\xcf\x5e\xcc\x27\x9a\x28\xeb\xce\x58\xd4\x17\x82\xcf\xa1\x58\xbe\xb9\x9c\xc3\xc2\x8f\x44\xc9\x61\xc1\xcf\xc7\x15\x6b\xf0\x42\xb2\x06\xe7\xb0\xb1\x5a\xc8\xc3\xeb\x8b\x49\xcc\x78\x66\x2c\x69\xa1\xa4\xc4\xd2\x0a\x25\xcf\x65\x0e\x73\x03\xa0\xc8\xb4\x15\x7b\x56\xda\x0b\x16\x7e\x6c\x8f\x2d\xce\x21\x1b\x7d\x39\x88\x9b\x22\x0e\xd1\x46\xd6\x59\x55\xaa\xa6\xad\xd1\xe2\x85\x90\x6d\x67\x23\xed\x2b\x28\x3b\x6d\x94\x5e\x2b\x33\x87\x42\xda\x2b\x60\x4e\xe4\x1c\xb2\xd1\x9e\xcc\x8d\x11\xf8\x55\x64\xfe\x50\x2c\x23\xc6\x65\xba\xf8\x1e\x4d\x57\x9f\x89\xfd\x45\x60\xcd\x4f\x65\xef\x69\x30\x68\x30\x5a\x9b\x4b\x2b\xec\xf1\x5a\x48\x7e\x35\x03\x00\xd0\xf8\x7b\x27\x34\xf2\x4c\x1f\x68\x31\x19\x74\x7a\xf9\xa7\x57\xe8\x25\xcb\x37\xdd\xe1\x80\x86\x14\xfa\x34\x9b\x01\xbc\x85\x4d\xa9\x45\x6b\x9b\x83\x06\x94\xbc\x55\x42\x5a\x73\x05\x1a\xf7\xa8\xc1\x2a\xe0\xaa\x34\x20\x24\x94\xb5\xea\x38\x6b\xc5\xbb\x56\x2b\xab\x66\x00\xb5\x78\xc2\x0f\x02\x9f\x89\xce\x4d\xf8\x7d\x8b\x96\x71\x66\x99\x3f\xe4\xb8\x62\xa1\xa4\x45\x69\xcd\xe8\x8c\x6f\x4e\xa6\x68\xb9\x71\x3c\x08\xce\x33\x4a\xc1\xfc\xec\x04\xd4\x26\x99\x78\xe3\x75\x5a\x62\x5b\xab\x23\x3c\xe2\xd1\xcc\x00\xb8\xfb\x6a\x50\xda\x6b\x3c\x92\x80\xe5\x78\xc0\xe3\x27\x6b\x46\xf0\xe9\x52\x8f\x9e\xad\x8b\x08\xcd\x5a\x11\x30\xb3\x75\xd1\x83\xf9\xd1\x11\x4a\x98\x9c\x7d\x99\xcd\xc6\x51\x7f\xdb\x59\x46\x27\xe1\x02\x7f\xa1\x91\x59\x0c\xde\x9f\x44\x13\xfc\x95\x63\xab\xb1\x64\x16\xf9\x85\x46\x66\xc8\x41\xbf\x0b\x0b\x0c\x30\x8d\x20\xd5\x33\x94\x0e\x80\xc3\x93\x60\xd0\xbe\x04\x8d\xbe\xbb\x9c\x01\x3c\xb4\x9c\x59\xfc\x20\xfe\x25\x5c\x5c\xed\xc5\xe1\x22\x38\x0a\xf9\x49\xb1\x7c\x73\x05\x4f\xa3\xc9\x39\xe4\x5c\x58\xb6\xab\x93\x2d\x13\x21\xee\x29\x27\x26\x3a\xb3\x18\xc0\x12\xc9\xef\x96\xaf\x18\xf8\x67\xa5\x6a\x64\x72\x80\xf3\xb6\x1a\x6c\x16\x01\xfc\xf7\xf4\x4e\xaf\xe0\x38\x15\x5e\x98\x3e\x43\x46\x65\x92\x4c\x79\x79\x9e\x39\x37\x68\xd3\x64\x79\xc1\x46\x79\x74\x8c\x32\xca\xa7\x97\x53\x19\xb6\x90\x4f\xc2\xd3\xb9\xc0\x86\x89\xba\xcf\x92\x14\xf3\xda\xd8\xd5\x38\x73\x5e\x41\xcd\x4e\x86\x2e\xe3\x05\x40\x30\xa9\x7e\x6b\xd4\x8d\x30\x46\x28\x69\x2e\x28\xd5\xf7\x07\xd8\xa5\x93\x29\xe1\xd1\xc4\x00\xee\xcf\xd0\x43\xdf\xe9\x43\x6f\x39\xa5\x0f\x3d\xaa\x1a\xc6\x07\xc4\xd1\x62\x42\xeb\xaf\xa6\x2f\xb3\x99\x73\xeb\x08\xef\xdc\x3a\x9c\xd7\x0c\x20\xb9\x2f\x66\x00\xa9\x69\x66\x00\xad\x28\x6d\xa7\x93\x35\x4a\x1f\x56\x27\xdb\x02\xbd\x61\x40\x98\xac\x6d\xb5\x7a\x42\x3e\xf2\x89\xc8\x25\x90\xfb\x1a\x15\x49\x7a\x79\x18\x56\x9b\x09\x9c\xb1\xb3\x38\x30\x26\x59\x7d\xb4\xa2\x34\x77\xad\x55\x94\xd5\x47\xee\x18\x05\x8d\x37\x0f\x1e\xe2\xb6\x5b\xd5\xe9\x0d\xa2\x7c\x6d\x9f\xbb\x2a\x5e\x71\xba\x69\x80\xe9\x5d\xff\x15\xe7\x9e\xa8\x8f\xb1\x13\x63\x3d\x52\x34\x0e\xb6\x0a\x69\x26\xb3\xb7\x66\x0e\xbf\xd4\x8a\x59\x9f\x3f\x4d\x39\xac\x8a\x78\x49\xd0\xff\x4f\x60\x51\x76\xcd\xc4\x6d\xb9\xb1\xcc\xa2\x13\x90\xe5\x9b\xcf\x0f\xab\xeb\xd5\xdd\xc7\x55\xf8\x5a\xe7\xab\x65\xb1\xfa\x35\x7c\xdd\x3f\xac\x56\xc3\xd7\x2f\x59\x71\x93\x2f\xc3\xc7\x36\xbf\xbf\x2d\x56\xd9\x36\x5f\x4e\x4a\x1a\xca\x00\x2f\x28\xdb\x8e\x04\xbd\x85\x4c\x02\x72\x61\x43\x05\x01\xaa\xa4\xd2\x02\xc4\x1e\x98\x8b\x4d\xa8\x98\x81\x46\x71\xb1\x17\xc8\xc1\x56\x08\xfe\xb0\x2c\xbe\x58\xd8\x1d\x41\x48\x83\x9a\x8e\x0a\x94\x06\x4e\x19\x8f\x7e\x97\x15\xd3\xac\xa4\x34\xff\xce\x09\xd9\x56\x82\xae\xe3\xb2\xee\x38\x1a\xba\x44\xdc\x06\xe9\xf0\x1e\xf1\xb8\x53\x4c\x73\x60\x92\x43\xcb\x8c\x07\x50\x4d\xc3\x24\x77\xdb\x89\x71\xbe\x2c\xb6\x9e\x2e\x18\xac\xb1\x1c\xf8\xca\xfa\x38\x4d\xba\xac\x94\x41\x09\x4c\x26\x15\x0d\x98\xbe\x90\x78\x17\x69\x71\x41\x77\x94\x01\x57\x20\xbc\x75\xa4\x92\x2d\xb6\x62\x16\x84\x05\x53\xa9\xae\xe6\xd0\xa8\x27\x74\x8b\x48\xd4\xf7\x26\xd4\x62\x54\x75\xd0\xa0\x24\xc3\x30\x0a\xc9\x56\x0b\x3a\x5d\xcb\x76\x51\x8b\x4d\x7e\x93\x2f\xb6\x7f\xe0\x0f\x54\x0e\x05\x77\xb8\x4e\xdc\xe1\xfa\xf3\xfa\x6e\x19\x7e\x6d\x3e\x2c\xe2\xaf\xc5\x7d\xb1\xde\x86\x8f\x55\x76\x9b\x6f\xd6\xd9\x22\x1f\xc2\x62\xb2\x7e\x72\xf8\x8f\x42\xf2\xd7\xca\xb7\x93\x44\x13\xdc\x99\xca\x15\x57\x62\xf6\xa3\x0d\xb3\x65\x85\xbc\x90\x1c\x5f\x5c\x79\x57\x48\xfb\x89\x6a\x1e\x72\xea\x29\x70\xe7\xed\x3d\xbb\x2d\xdb\x9d\x90\x22\x3f\x21\xff\xe2\xf8\x02\x6a\xef\xac\x69\xd9\xce\x9b\xdf\x56\x68\xc6\x87\xe7\xeb\x87\xbd\xd2\x64\x5b\xcb\x76\x8e\x85\x2b\x86\x1d\xd0\xc7\x0a\x6d\x85\x3a\x38\x0b\x79\x14\x1b\x6d\xa6\x7d\x60\xe9\xf0\x09\xdf\x0b\x7c\x16\x75\x0d\x0d\x7b\xf4\x47\x1b\xfc\x0f\xf0\x05\xcb\xce\x65\x25\x92\x33\x7c\x65\x7b\x4b\x49\x8a\xc0\x87\x74\x04\x63\x7e\x7f\x50\xbf\x4e\x9d\x8f\xaf\xbf\x47\x66\xd8\x2b\xdd\x30\x4b\x85\x91\x0f\x38\x22\xdb\x47\x9f\x09\xa5\xf8\x73\x25\xca\xca\x79\xfb\x0e\x51\x42\xcb\xb4\x41\x4e\x61\x79\xee\xc3\xaa\x77\x74\xef\xe4\x6c\xb7\xb1\xaa\x85\x56\x19\xe1\xf8\x92\x7e\xbd\xcc\x62\x5c\xf1\x27\x06\x3d\xe5\x40\xbc\x18\x3c\xb1\x5a\xf0\xab\x91\x7d\xa2\x01\xdf\xb9\x8b\x2e\xef\xc7\xc7\xc6\x7a\x0b\x59\x5d\x27\x47\x4a\xc7\x82\xac\xac\x46\xa7\x4f\x24\x4d\x38\xe3\x4d\x62\xdd\xc4\x7f\x06\xa3\x52\x49\xcd\x84\x44\x4d\xde\xd6\xf9\x0b\xe4\xf4\xde\x9c\x4e\xda\xc1\x6f\x87\x65\x0d\x1a\xc3\x0e\xc9\x50\x2c\x63\x4f\x6f\x8c\xeb\xbf\x98\xfc\x09\xa5\x3f\xc0\x89\x7d\xae\x7e\xda\x8a\x06\x13\x89\x54\x41\x9d\x0c\x46\xc0\xb5\xe2\x7f\x4a\x81\xce\x7c\xa3\x06\x00\x65\xb4\x98\xeb\x79\x53\xf3\xf9\xc6\x00\x49\x35\x9a\x8d\x6a\xd2\x70\xcc\x62\xa1\xb8\x1e\xb1\xf5\x0e\xcc\x71\xcf\xc8\xa5\x9d\x59\x29\x37\x4b\x65\xab\xe0\x31\x8f\x52\x3d\x4b\x3a\xd5\xc5\x26\xb9\x8c\x68\x5f\x58\x6f\xa0\x42\x56\xdb\xea\x48\x5b\x2b\x64\xda\xee\x90\x59\x1f\xf6\x1a\x4b\x14\x4f\xc8\xe9\x0a\xd1\x78\xe8\x6a\xa6\x41\x48\x8b\x9a\xaa\x20\x77\x8f\xd8\xca\xbb\x79\xe8\x17\x08\x4e\xa3\x69\x95\xe4\xc4\xc0\x2a\xd7\x9d\xa2\xb1\x26\x90\x78\x9f\x67\x37\xdb\xf7\xbf\x9d\x93\xe8\xe4\x88\x86\xcb\x0c\x03\x62\xe9\x7b\x7d\xba\x17\x15\xac\xc5\x8b\x40\x58\x50\xbf\xe9\x18\x08\x03\x54\x96\x09\x1e\x23\x68\xd0\xe1\x0a\x76\x2e\xa0\xe5\xf7\x16\x7e\xef\x50\x1f\x5d\xc4\x90\xf3\x1b\xd5\x60\x38\xa1\x70\x3b\x69\x34\xd8\xec\x6a\x34\xf0\x7e\xbb\x5d\x7f\x6f\xe0\xc7\x1f\x7e\x08\x07\xdd\xdb\x6f\x9a\xbc\x4b\x68\x07\xe5\xba\x61\x61\x06\xae\x41\x8f\x5f\xef\xd7\x8b\xa8\x01\xa5\xc4\x9d\x46\xf6\x68\xde\x39\x80\x4a\xb5\xe8\x13\x0e\xb3\xfd\x95\x18\x15\x77\xb8\x25\x11\xdd\xb1\xf2\x91\x2e\x60\x21\xd1\xa9\xac\xd1\x74\x0d\xa5\x07\x08\x8c\x3c\x93\xc0\x73\x59\x6c\x16\x77\xab\x55\xbe\xd8\xba\xca\xe5\xd4\xce\x54\xcf\xd3\xd9\x3c\x57\x28\x4f\x0d\x2d\xfc\x48\xab\x55\x89\xc6\x50\x76\x88\xcb\xa3\x0d\xd6\xcb\x6c\xeb\xcb\x23\x8f\xeb\xfb\x42\x5f\x07\x44\xcd\xbd\xd9\x69\x48\x2a\x0b\x86\xa2\x95\xc9\x23\x28\x97\xd7\xf6\x9d\xf6\x17\x86\x77\x63\x87\x8f\x06\xd8\x4e\x75\xde\x04\xcf\x21\x01\x0a\x3b\xf6\x4d\xa5\x4f\xa9\x9c\xeb\x18\xb8\x3c\x33\x03\x56\x1f\x83\xff\x79\x01\x9e\xd2\x9e\x89\x1a\x7b\xaf\xa1\x9e\x58\x48\x60\xb0\x63\x3c\x31\xa0\x53\x32\x8f\xb5\x5f\x4c\x14\xe3\x5e\xd7\x45\x5f\xcb\x8c\xb1\x95\x56\xdd\xa1\xca\x5d\x7f\x30\xd5\x5f\x8c\xdb\xf4\xb4\xbe\x8d\x49\x24\x09\xeb\x98\xac\xde\x47\x1f\x4e\xf2\x4e\xda\x84\x27\xcd\x77\x3f\xfb\x01\xb5\x11\x27\x79\xc7\x4b\x78\x7d\xe6\xb4\x85\x6a\x35\x5a\x7b\x5c\x4c\x4f\x9e\x3f\x21\xc5\xdc\xa6\x55\xbd\xae\x99\xc4\x3e\xa5\xba\x62\xa5\xff\xf2\x29\x4e\x76\xcd\x4a\x71\xf4\xef\x69\x61\xa0\x90\xc6\xea\x8e\xda\x00\xe4\xe3\xc9\x13\x23\xa6\x4f\x7e\xde\x9c\x6d\xc6\xb9\x46\x93\x24\x63\xab\x1e\x71\xe2\xea\x18\x5a\x66\xb7\xf5\xac\xbd\x14\x6e\xee\x46\xc8\xc7\x64\xef\x5b\xb8\xff\xca\x63\x97\x43\x3f\x7d\xe3\xfa\x4a\x47\x79\xd6\xb8\x7c\xa3\x98\xf8\xa0\x15\xee\x41\x2f\x73\x7e\xc6\xc2\x9d\xe6\x4b\x1d\x57\x8f\x19\x3c\x09\xf3\xb7\xcd\xdd\xea\xcf\x90\x48\x1f\xe0\xbe\x49\x53\xa0\xbc\x10\x59\xa6\xf1\xf2\x4d\xc2\x5f\xd1\xff\xe4\x69\x30\x38\x66\xaa\x7a\xdf\x15\x8c\x5e\x85\x1d\x0c\x40\xd2\xb2\xb9\xcf\x9b\x62\xf5\xf0\x8f\xcf\xd9\xed\xf2\xa7\xff\x8f\x43\xcb\xec\xfe\x63\xb1\x4a\xc7\x16\x77\xab\x6d\x56\xac\xf2\xfb\xcf\x9b\x7c\xfb\xf9\xb7\xec\xf6\x66\x33\x3d\x35\x81\x97\x2e\xd8\xe6\xb7\xeb\x1b\x4a\x3f\x1e\xa4\x0f\x81\xe1\xc9\xda\xff\x0d\xa0\x13\xdf\x35\x15\xfb\xbf\x1f\x7f\x4a\x74\x4c\x7b\xfd\x6f\xc9\x5e\xd3\x2f\x05\xa3\x37\x22\x7f\xe2\xe7\xcf\x2a\xe7\x1b\x47\x4f\x41\x3e\xe8\x5e\x79\x48\x99\x7d\xf9\x4f\x00\x00\x00\xff\xff\x1c\x98\x27\x45\xee\x18\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 6382, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
