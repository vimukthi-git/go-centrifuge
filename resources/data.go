// Code generated by go-bindata.
// sources:
// ../build/configs/default_config.yaml
// ../build/configs/testing_config.yaml
// DO NOT EDIT!

package resources

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

var _goCentrifugeBuildConfigsDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x59\x73\xdb\xba\x15\x7e\xd7\xaf\x38\x23\xbf\xb4\x33\xa5\xcc\x7d\xd1\xcc\x9d\x8e\xbc\x25\xf7\xc6\x71\xbd\xc8\xf1\x8d\x9f\x02\x02\x87\x24\x62\x0a\x60\x00\x50\x4b\x7e\x7d\x07\x20\xe5\xd8\x71\xec\xdb\x4e\xa7\x7e\x31\x05\xe0\x7c\x38\xcb\x77\x3e\x00\x07\x70\x82\x15\xe9\x5b\x03\x0c\xd7\xd8\xca\x6e\x85\xc2\x80\x41\x6d\x04\x1a\x20\x35\xe1\x42\x1b\x50\x5c\x3c\x60\xb9\x9b\x50\x14\x46\xf1\xaa\xaf\xf1\x02\xcd\x46\xaa\x87\x39\xa8\x5e\x6b\x4e\x44\xc3\xdb\x76\xe2\xc0\xb8\x40\x30\x0d\x02\x1b\x71\xc5\xb0\x52\x83\x69\x88\x81\xe3\x47\x04\x58\x11\x2e\x8c\xc5\x9f\xec\x97\xcc\x27\x00\x07\x70\x2e\x29\x69\x9d\x0b\x5c\xd4\x40\xa5\x30\x8a\x50\x03\x84\x31\x85\x5a\xa3\x06\x81\xc8\xc0\x48\x28\x11\x34\x1a\xd8\x70\xd3\x00\x8a\x35\xac\x89\xe2\xa4\x6c\x51\xcf\x26\xb0\xb7\xb7\x90\x00\x9c\xcd\x21\x8a\x22\xf7\x8d\xa6\x41\x85\xfd\x6a\x8c\xe0\x77\x36\x87\x3c\xca\x87\xb9\x52\x4a\xa3\x8d\x22\xdd\x25\xa2\xd2\x83\xad\x07\xd3\x43\xde\xc5\x87\x41\x98\xcd\xfc\x99\x3f\x0b\x0e\x0d\xed\x0e\xa3\x3c\xf4\xc3\x43\xde\x55\xfa\xf0\x6a\xb5\xbc\xda\x96\x9b\x87\xfe\xfe\xf3\xe7\x93\xaa\xff\xbe\x2c\xb7\xa7\x8b\x6b\x5c\x5e\x1c\x9f\xcb\xef\xbb\x5d\x92\xe4\xeb\x2b\x51\x7f\x5a\x5f\x7e\xfc\x7a\xfe\xf9\x61\xfa\x17\xa0\xd1\x1e\xf4\x53\x95\x9e\x5e\xa4\xab\x87\x6f\x77\xf8\xf5\xee\xc3\x5d\xf8\xed\xb2\x0f\xd2\x3f\x3b\xf6\x2e\x7a\xf8\x43\x06\xcb\x68\xd5\x90\xe6\xf2\x28\xb9\xc1\x44\x04\x03\xe8\x3e\x55\x8b\x7d\xa6\x86\x00\x6c\xf8\x28\x0c\x37\xbb\x33\x42\x8d\x54\xbb\x39\x4c\xa7\x3f\xcd\x5c\x63\xcd\xb5\x79\x36\x45\x04\x6d\xa4\xba\xc6\x4e\x6a\xfe\x93\x55\x47\x76\x96\x26\xff\x2a\x5b\x5e\x13\xc3\xa5\x70\x73\xae\x78\x1f\x09\x17\xbf\xa4\xd2\x58\xe3\x09\x3c\x65\xcc\xe0\xe0\x01\x5c\xf4\x2b\x54\x9c\xc2\xef\x27\x20\x2b\xc7\x9e\x27\x3c\xf9\x61\x39\x14\x32\x09\x46\xab\xa3\x7d\xb5\xa0\xe5\xda\x58\x4b\x21\x19\xbe\x24\x5a\xa7\xe4\x9a\xbb\x09\xe9\xb0\x9f\x38\xb0\x77\xef\x2f\xab\x1f\x25\xb3\x30\x4c\x66\xa1\xef\xcf\xe2\xf0\x67\x06\x04\xe1\x49\xf4\x41\xca\xbb\x73\xce\xe9\xd5\xa7\xcd\xb2\x59\x1e\x7d\x4e\xb7\x1f\xe8\xa5\x3c\xaf\xd2\xeb\xab\xcf\x7f\x9c\x75\x9b\x2a\x50\x59\xb2\x39\xdf\x86\xf7\xd7\x51\x77\xcc\x82\xe9\xaf\xe0\xf3\x74\x16\x06\xfe\x6b\xf0\x57\xf7\x1f\x17\xf9\xbb\xcb\xf7\x6a\x7d\x7a\x7f\x54\x6c\xd8\x83\xbc\xa5\x8b\xc5\xea\xf8\xfe\x7d\x57\xe0\x6e\x77\x1f\xdf\x9c\xe6\xf5\x99\x8a\x9a\xe5\xc5\x9f\xd3\x31\x47\xa7\x23\xdb\xf7\x59\xb4\x29\xf6\xe0\x7a\xec\xe7\x57\xfa\x21\x1e\x8d\xcf\x89\x4d\x0f\x30\xec\x5a\xb9\x43\x06\x37\x2b\xa2\x0c\x1c\x8f\x34\xd3\x50\x49\xe5\x12\x5a\xf3\x35\x8a\x67\xa9\x7c\x49\x45\x78\x95\x8b\xfe\xb6\xf0\x59\x58\xc4\x49\x16\x60\x16\xe5\x71\x98\x16\x19\x49\xd3\x32\x23\x45\x41\xfc\x82\xb1\x94\x66\x11\x8b\x92\x94\xbd\xc1\x5a\x7f\x5b\xa4\xa9\x4f\xfd\xa8\x60\x51\x10\xc4\x49\x44\x2a\x9f\x25\x39\x4d\xd2\x34\xcd\xc2\x88\x15\x34\xac\x48\xc6\x52\xa4\x6f\xf0\xdb\xdf\x66\x55\x9e\xc4\xac\x22\x45\xee\x07\x21\xcb\x2a\x92\x24\x34\xf7\xa3\xb2\x24\x61\x98\xfa\x25\x65\x88\x71\x99\x20\x7b\xab\x13\xfc\x2d\x2b\xfd\x24\x0f\x16\x45\x14\xe6\x69\x1a\xe7\x49\x12\x85\xf9\x82\x9d\x94\xfe\x69\x98\x04\x41\x1e\xa7\xb1\x5f\x15\x98\x9c\xb8\x9e\x29\x51\x09\xd2\x36\xc8\xeb\xc6\x8c\xa4\x3b\x38\x38\x18\x2b\xf0\x41\xae\x89\x80\xb3\xc5\xd5\xf8\xdb\x83\x3b\xab\x76\x5c\x54\xbd\x22\xb0\x93\x3d\xd4\x56\xa6\x05\xa0\x52\x52\x59\x3a\x2d\x1b\xae\x41\xe1\xb7\xde\x56\x8e\x6b\x10\xd2\x80\xee\xbb\x4e\x2a\x83\x0c\x4a\xa4\xa4\xd7\x68\x2d\x95\xeb\x16\xbb\x44\xf5\x42\x58\xa9\x75\x42\xaa\x0d\x31\xb6\x65\x7a\x3b\x34\x83\xeb\x5e\x0c\xe3\x9e\x37\x8e\xfd\x46\x14\x6d\xf8\x1a\x67\xd3\x7f\x8c\x4e\x01\x6c\x6c\xc7\x19\x09\x4c\xfe\xd3\x59\x10\x68\x9d\x88\x77\x44\x71\xb3\x1b\x36\x72\x28\x0f\x2e\x1e\xac\xe7\xc3\xcf\x2f\xe3\x02\xcf\xa3\x0d\xe1\xe2\xb7\x61\xda\xf3\xac\xb7\xbf\x45\x7e\xe4\xc7\xe0\x79\x1b\xa2\xba\xf1\x9f\x57\x12\xa5\x38\x2a\x48\xd2\xdc\xf7\x7d\x1f\x3c\x4f\x48\x8f\x08\xca\x51\x18\xaf\x6c\x25\x7d\xd0\xc3\x98\x46\xb5\x46\xaf\xb5\x49\x05\xcf\x5b\x91\xad\xd7\xd9\xa6\x86\x30\xb1\x46\x5a\x90\x4e\x37\xd2\x8c\x83\x6e\x6c\xc5\xc5\xb3\x9f\xd6\x67\x42\x0d\x5f\x23\x78\x9e\x25\xb3\x4d\x91\xac\xaa\x97\x99\x00\xcf\x63\xa5\x47\xe5\xaa\xb3\xeb\xa5\x00\xad\x99\x0d\x89\xd0\x06\x3d\xcd\xbf\x23\xc4\x7e\x91\x82\xe7\x7d\xd5\x52\xa8\x8e\x7a\x8d\xd4\x46\x03\x69\xdb\x27\x63\x5c\x18\x54\x15\xa1\x68\xc7\xbf\x3c\x2f\xf7\xcb\x64\xfe\xaa\xf2\x47\x36\x7c\x64\xb6\xf7\x04\x0e\x8e\x18\x09\x77\x58\xde\xd8\x71\xa3\xc1\xe5\x44\x41\xa5\xe4\x0a\x7a\x61\x54\xaf\x2d\x25\xa4\xe2\x35\x17\x73\x98\xcd\xa6\xaf\xd6\xd3\x36\xf9\x8b\x5a\x7e\xf1\xbc\x5e\x68\x52\xa1\x87\xdb\x4e\x6a\xfc\x02\x55\x4b\xea\x9f\x08\xfc\xdf\x29\x7b\xf8\x3f\x2a\xfb\xb3\x5e\xfa\x8f\xb5\x3d\xf0\xe3\x59\x90\xc4\xb3\x20\x9f\x25\x2f\x4e\xf7\xbd\xf8\x5e\xea\x94\x13\xbc\xed\xcf\xee\x2f\xfa\xe0\xdd\x76\xad\x77\x47\xcb\x1b\xb5\xd4\xc5\xda\x1c\xa5\xa5\xf9\xb8\x10\xef\xcf\xe4\xf9\xd7\xf2\xe1\xfb\x31\x99\xfe\x02\x3e\x99\x05\x79\x32\x0b\xa3\xec\xd5\x0d\x8e\xdf\xd1\x0d\x5f\x7e\x95\x1f\xee\xde\x57\x47\x24\xce\xc3\xdb\x4b\x43\xf0\x76\x7b\x71\xbe\x61\xf9\xf7\x52\x1c\x05\x37\xd9\x06\x17\xf7\xb7\xdb\xfb\xb7\xd5\xdd\x89\xc6\xab\xda\x1e\xfe\x1f\xc4\xfd\x0d\x6d\xcf\x93\x32\x0a\xab\x8c\x44\x55\xec\xc7\x79\x50\x05\x61\x14\xc5\x7e\x1c\xa4\x99\x4f\x73\x5a\xa2\x9f\x55\x19\xcb\x0a\xfa\xa6\xb6\x27\x31\xc1\x28\x8b\x2a\xbf\x48\x2b\x52\x85\xac\x4c\xcb\x9c\xc4\x69\x16\x64\xd4\x2f\x8b\x1c\x69\x45\xfc\x2c\x61\xec\x4d\x6d\x8f\xe3\xb8\x4a\xe3\x02\x23\x3f\x8b\xe3\x10\xb3\x94\xd2\x2a\x8b\xb2\x38\x4d\x31\x09\xab\x20\xf5\x8b\xb2\xc8\xc3\xd4\x7f\x5b\xdb\xfd\x38\xc8\xb0\x8c\xb2\x22\x0e\x82\x34\x8e\xd2\x3c\xf6\x83\x93\x34\x4d\x8b\x3c\xa6\xa7\x27\x59\x5a\xc4\x8b\x23\x7a\x54\x06\xd3\xc9\xe4\x00\x2c\xd9\x3c\x23\x9d\xae\xd8\xb4\x55\xbc\xee\x95\xc3\xd2\x93\x2e\xec\x86\xfb\xee\x92\xaf\x50\xf6\x06\x36\x0d\x0a\x90\x1d\x8a\xf1\xda\x3b\x36\xb1\x23\xb7\x13\xa6\x09\xec\x87\x47\x93\x39\x4c\x23\x5f\xbb\x9d\xae\x7a\xec\xf1\xa7\x2d\x5c\x09\x89\xde\x09\xda\x28\x29\x64\xaf\x6d\xbf\x50\xd4\x9a\x8b\x7a\xf2\xcd\x1a\x0c\x0e\x0c\x97\x76\xed\xaa\x2d\xfa\x55\x89\xca\x76\x9c\x25\x0c\x2a\x7d\x48\xa5\xd0\xb6\x89\xc7\xee\xdb\xd8\x5b\x53\xe9\x54\x4a\x52\x62\x05\x84\x18\x7b\x68\x28\xd3\x77\x13\xb0\xf6\x77\x83\xe1\x1c\x42\x87\x7e\xa6\x10\x35\xf4\x1d\x1c\x5f\xde\x02\xdd\xd1\x16\xf5\x10\xea\xb0\x81\x3d\x80\x36\x84\xbb\xbb\xbe\xf5\x17\xd7\x28\x8c\x0d\x75\x98\xbe\x23\xdc\x45\xfb\xf1\x66\x0e\x81\x0d\xf4\x91\xf2\xba\x43\xca\x2b\x4e\x9f\x07\x3d\xd9\x53\x7e\x08\xed\x06\x5b\xb4\x64\xde\x34\x9c\x36\x8f\xed\x00\x84\x52\xd9\x0b\x27\x71\xf6\x34\x1c\x95\x49\xda\x24\x8c\x92\xc2\x80\x0f\xb2\x47\x7b\x6d\xe4\x6a\xdc\x04\x2a\xde\xe2\x04\xf6\x6f\x9b\xc5\x00\x73\x41\x56\x38\x87\xa9\x7d\xcf\x4c\x1f\x5f\x30\x4e\x7f\x47\xe0\xc7\x7d\x69\x6b\x0f\xaa\x41\x43\xff\xb6\x41\x77\x4e\x73\x85\xb0\xd1\x20\x15\xf0\x8e\x8e\xcf\x1a\xfb\x8a\xb1\x9f\x94\x18\xeb\xb6\x4b\xc9\xdf\x6d\x76\x25\xc3\xdb\xeb\xf3\x39\x34\xc6\x74\xf3\xc3\x43\x77\x30\xd8\xd3\x64\x5e\x24\x71\xb2\x2f\xa6\x7b\x76\xd5\xc4\xc6\xc2\xa9\x75\xb7\x26\xfa\xd2\x7e\xce\x21\xf0\xf7\x7f\x2f\x16\xb7\x7c\xc5\xcd\xb0\xf8\xdc\x7e\xce\x21\xce\x82\x30\xca\xf3\x67\x24\x35\xd2\x55\x6b\xa0\x96\xf8\x11\x99\x51\x44\x68\xf2\x78\xea\xd8\x18\x18\x1b\x9e\x69\x04\xdc\xc1\x0c\x44\xb0\x31\x14\x30\x8a\xd7\x35\x2a\x64\x03\xa5\x0d\x6e\xcd\xbe\xd0\x03\xad\x53\xdf\xf2\xfa\xb5\x8d\x15\x12\x06\x52\xb4\x3b\xdb\x2e\x7b\xb2\xef\xdf\xaa\x7b\x97\x7e\x40\x5f\x23\x61\xcf\xe1\x83\x64\x44\xbf\xb0\x95\x78\xea\x7b\x27\x65\x0b\x2b\xb2\x05\x85\x46\xf1\xe1\x64\xd1\x28\x18\x90\x67\xcb\xe4\x1a\xd5\x04\xec\xc2\xeb\x61\xdd\x1c\xc2\x31\xa7\xbf\x86\x74\xc7\xfb\x9a\xb4\x0e\x77\x37\x34\x00\xb1\x0e\xd2\x5e\x29\xf7\x4e\x7a\x62\xd1\x10\x0d\x25\xa2\x7d\x48\x19\xa4\xc6\xa5\x69\x0f\x60\xf7\xb3\x7a\x16\x8e\x11\x9c\x70\xed\xd8\xe2\x10\xb5\x5c\xbd\x60\x9b\x06\x26\x9f\xde\x02\xc1\x6c\x9d\x47\xa4\xe3\xf6\x99\xbc\xbd\x94\xb2\x5d\x50\x2b\x0b\xa7\xc2\x22\xb1\x39\x18\xd5\xa3\xed\x35\x22\x76\xc0\xb0\xec\xeb\x7a\x94\x24\xdb\x02\x4e\x00\x6a\x09\x76\x93\x89\x9b\x1d\x5a\xad\xeb\x94\xac\x5c\x79\x1e\x4d\x26\x30\x8c\xce\xa1\x22\xad\xc6\xc9\xbf\x03\x00\x00\xff\xff\x6c\x54\x71\x54\x6d\x10\x00\x00")

func goCentrifugeBuildConfigsDefault_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsDefault_configYaml,
		"go-centrifuge/build/configs/default_config.yaml",
	)
}

func goCentrifugeBuildConfigsDefault_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsDefault_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/default_config.yaml", size: 4205, mode: os.FileMode(420), modTime: time.Unix(1544092466, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goCentrifugeBuildConfigsTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xc9\x6e\xe3\x48\x0c\x86\xef\x7a\x0a\x81\x97\x5c\xbc\xd4\xbe\xbd\xc1\x20\x98\xd3\x0c\x90\x33\xab\xc8\x8a\x05\xdb\xb2\x5a\x4b\x12\x23\xc8\xbb\x37\xe4\x38\x9d\x6b\x1a\xba\x90\x04\x7f\xfe\xa4\xea\xe3\xf9\xc0\x23\x2f\xe7\xd4\xb4\x2d\x96\x72\x59\xfa\x79\x5a\xe3\xb6\x3d\x63\xd7\xa7\xf6\x16\xb6\xed\x91\xaf\xa9\x7d\x78\x07\x24\x1a\x79\x9a\x20\x41\x88\x59\x60\x70\x36\xe8\x62\x8c\x31\x58\x2a\x79\x99\x8d\xd3\x2c\x48\x17\x6b\x91\xa5\x91\x0a\x2d\x6c\xa0\x8c\xd7\x61\xbe\x40\x7a\x87\xd2\x0d\x07\x1e\x21\x01\xf2\xb4\x95\x2a\x6c\xcb\x3c\xae\x0d\xb7\xf2\xcc\x6f\x33\x24\x28\xde\xc7\x1a\xb4\x8f\xe4\xbd\xa0\xa8\x4a\x2d\x92\x88\x0c\x86\xaa\x25\x59\x14\x48\x25\x54\x85\x22\x2b\x94\x46\x48\xed\x05\x69\xa7\x45\xd5\xa1\x88\x12\xf0\xcf\xbc\x01\x47\x3c\x4f\xab\x6d\xf7\x02\x09\xb4\x2b\xd2\x05\xf6\x3a\xd7\x18\x44\x65\x6f\xb3\xf0\xca\xd7\x10\x05\x7a\x89\x04\x1f\x1b\x38\x52\x85\x04\xd3\x6d\x61\xb8\xa5\xdf\x43\xe8\x78\xe2\x1e\x92\x56\x1b\xe8\x21\x29\xa7\xa4\x31\x1b\x18\x20\xc9\x0d\x8c\x90\xc2\x06\x26\x3c\xad\x07\x10\xcb\xcc\xd2\xb1\x2e\x31\xc8\x68\x0c\x49\x2e\xa8\x72\xc8\xca\xb3\x61\xc7\x22\xdb\x5c\xb3\xd1\x99\x85\xf6\x0e\x2d\x85\x10\x62\x45\xe7\x23\xaa\x20\x95\x5a\x17\x39\x63\x59\x7f\x45\x91\x2a\xe4\x20\xad\xb5\x36\xa3\x64\x24\x5f\x90\xa3\x70\x82\x43\x30\x0a\x6b\xc1\xa0\xad\x23\xe1\x8c\xb5\x99\x22\x5a\x6f\x55\x46\x57\x4b\x11\x51\x71\x5d\x27\x75\x04\x09\x8c\x65\xe1\x04\xba\x2d\x29\xe4\xad\xd1\x39\x6c\xa3\x52\x75\x6b\x4c\x50\xd1\xc4\x48\xda\x13\x6c\xe0\x85\xc7\xa9\xbb\xac\x47\x7e\x3c\xdc\x1f\x7e\xc0\x69\x7a\xbd\x8c\x94\xda\x87\xaf\xd2\x9d\x81\xd4\xfe\x14\x81\xa6\xe9\x88\xfb\xb9\x9b\xaf\xff\x50\x6a\x41\xbc\x09\xf9\xfd\x41\xd3\xfc\x5a\x78\xe1\x15\xba\x7e\x39\x3f\x5d\xc6\x23\x8f\x53\x6a\x55\xd3\xb6\xaf\xb7\xe4\x09\xbb\xf9\xff\xee\xcc\xff\xfe\x97\x5a\xd9\x34\x47\xbe\xde\x08\x9d\xba\xe7\xbe\xeb\x9f\x3f\x61\x1d\x96\x7c\xea\xca\xe3\x4a\xe9\x6e\xb7\xdf\xed\xf6\x79\xe9\x4e\xb4\x1f\x79\xba\x2c\x63\xe1\x69\x7f\xef\x7e\xe4\xeb\x6e\x58\xf2\x6e\xe0\xf3\xa7\x6e\xec\x5e\x70\xe6\x9f\x09\x8f\xab\xf8\x26\xe4\xf9\x80\xcb\x7c\xf8\xa1\xf7\xbd\xfb\x2f\x8d\xbf\x54\x5f\xae\xbf\x03\x00\x00\xff\xff\xb0\x1c\xaf\x3f\xaa\x03\x00\x00")

func goCentrifugeBuildConfigsTesting_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsTesting_configYaml,
		"go-centrifuge/build/configs/testing_config.yaml",
	)
}

func goCentrifugeBuildConfigsTesting_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsTesting_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/testing_config.yaml", size: 938, mode: os.FileMode(420), modTime: time.Unix(1540471826, 0)}
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
	"go-centrifuge/build/configs/default_config.yaml": goCentrifugeBuildConfigsDefault_configYaml,
	"go-centrifuge/build/configs/testing_config.yaml": goCentrifugeBuildConfigsTesting_configYaml,
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
	"go-centrifuge": &bintree{nil, map[string]*bintree{
		"build": &bintree{nil, map[string]*bintree{
			"configs": &bintree{nil, map[string]*bintree{
				"default_config.yaml": &bintree{goCentrifugeBuildConfigsDefault_configYaml, map[string]*bintree{}},
				"testing_config.yaml": &bintree{goCentrifugeBuildConfigsTesting_configYaml, map[string]*bintree{}},
			}},
		}},
	}},
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
