// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9a\x5f\x31\x15\x5b\x22\x81\xa3\xd0\x6b\x61\x4f\xcd\x1e\x52\x68\x52\xd8\x40\x0f\xdb\x25\xc8\xf6\x78\x23\x6a\xcb\xae\xa4\x94\x06\xa1\xff\x5e\x24\xd9\x0b\x3d\xd9\xf2\x7b\xfa\xe6\x3d\x4f\x8c\xd8\xd3\x60\x2c\x21\x9f\xb4\xb1\x1c\x53\x82\xbb\x3b\xfc\x32\xf7\x84\x17\xb2\xe4\x74\xa0\x1e\xdb\x77\xbc\x21\x1b\xba\xeb\xa7\x1b\x85\xf7\x4f\xf8\xf8\xf4\x8c\x87\xfb\x87\x67\x05\x8b\xee\x7e\xe9\x0b\x61\x66\x00\x98\x69\x99\x5d\x40\x01\x8c\xcf\x9e\x03\xe3\xed\x7b\xa0\xfc\x12\x23\x06\x9a\x96\x51\x07\x42\x5e\x5d\xbe\x8c\x2c\xd2\xe2\x8c\x0d\x03\xf2\x8f\xbf\x39\xaa\xef\x2b\x31\x25\x90\x00\x31\xe2\xae\xd5\x9e\xf0\xf3\x1e\xcb\x73\xd3\xf3\xdd\x3f\xda\xa1\xef\x5e\x69\xd2\x1e\xf7\x78\x3a\x93\x0d\xea\xc1\x06\x72\x83\xee\x28\x16\xb4\xd3\xf6\x42\xb8\x7b\x69\x70\x67\xf5\x54\x30\xea\x51\x4f\xe4\x33\x9f\xb1\x18\x6f\x57\x7e\x4a\x2a\x1f\xae\x51\x7c\x4c\x7c\xbd\x93\x52\x53\x58\x64\x7b\xbc\x4d\x09\x12\xc0\xf0\x66\xbb\xd2\x59\x48\x8c\xc0\x72\x90\xd1\x58\xf2\x78\x3a\x9f\xce\xb9\x34\xb0\x61\x76\xf8\xd2\xac\xf9\xf2\xdc\x1a\x65\xcb\x1b\x81\xb1\xb6\x41\x72\x2e\x6b\xdf\xb4\xf3\xaf\x7a\x3c\x16\x51\x54\x8f\x04\xc6\xcc\x50\x1c\x1f\xf6\x68\xcd\x58\xee\xb0\x41\x9b\x51\x90\x73\x59\xce\x15\xea\xdc\x3d\xea\x65\x21\xdb\x8b\x72\x6c\xb0\x95\x90\xd5\xd9\xab\x63\xe8\xe7\xb7\xa0\x7e\x38\x13\x48\x94\x7d\xa8\xaf\xb3\xb1\x9b\xb1\xc6\x15\xfc\xa7\xe5\x52\xca\x6b\xb7\x6d\x4a\x1e\x3f\xbb\x52\xb2\xb2\xc8\xb9\xca\x3a\x06\x67\xec\x25\x7b\xd4\x21\x7b\x84\x94\xc5\x73\xf8\x6b\x82\xf8\x54\x48\xff\x6d\xbd\x96\xaa\x4b\x5f\x7f\x66\x4a\xf0\x2f\x00\x00\xff\xff\xb5\xb1\x2f\xf6\x87\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 647, mode: os.FileMode(420), modTime: time.Unix(1566131522, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5b\x6f\xdb\x36\x14\x7e\x96\x7e\xc5\xa9\x81\x06\x52\xe0\xc9\x5d\xdf\xa6\xc2\x0f\x45\x9b\x02\xd9\x25\x1d\x9a\x6e\x2f\x45\xd1\xd2\xd2\xa1\xcd\x56\x22\x55\x92\x76\x9a\x06\xf9\xef\xc3\xe1\x45\x17\xdb\xc9\xb6\x16\x09\x10\x58\x3a\x57\x9e\xef\x7c\x3c\xa4\x16\x0b\x78\xa1\xba\x6b\x2d\xd6\x1b\x0b\x4f\x9f\xfc\xfc\xcb\x4f\x9d\x46\x83\xd2\xc2\x2b\x56\xe1\x4a\xa9\xcf\x70\x2e\xab\x02\x9e\x37\x0d\x38\x23\x03\xa4\xd7\x3b\xac\x8b\x74\xb1\x80\xb7\x1b\x61\xc0\xa8\xad\xae\x10\x2a\x55\x23\x08\x03\x8d\xa8\x50\x1a\xac\x61\x2b\x6b\xd4\x60\x37\x08\xcf\x3b\x56\x6d\x10\x9e\x16\x4f\xa2\x16\xb8\xda\xca\x9a\x42\x08\xe9\x4c\x7e\x3f\x7f\x71\x76\x71\x79\x06\x5c\x34\x18\x65\x5a\x29\x0b\xb5\xd0\x58\x59\xa5\xaf\x41\x71\xb0\xa3\x7c\x56\x23\x16\x69\xda\xb1\xea\x33\x5b\x23\x34\x8a\xd5\x69\x2a\xda\x4e\x69\x0b\x59\x9a\xcc\x50\x56\xaa\x16\x72\xbd\xf8\x64\x94\x9c\xa5\xc9\x8c\xb7\x96\x7e\x34\xf2\x06\x2b\x3b\x4b\xd3\x64\xb6\x16\x76\xb3\x5d\x15\x95\x6a\x17\x3c\x14\x2c\x64\xb5\x5d\x31\xab\xf4\x02\xa5\x5d\x98\x6a\x83\x2d\x5b\x60\xbd\xc6\xff\xe4\x30\xfb\x1f\x41\xb9\xc0\xa6\x9e\xa5\x79\x4a\x30\x5c\x3a\x19\x68\x0c\x0d\x30\xc0\x24\xa0\xb4\x45\x50\xd8\x0d\xb3\x70\xc5\x8c\xab\x13\x6b\xe0\x5a\xb5\xc0\xa0\x52\x6d\xd7\x08\x02\xdb\xa0\x86\x80\x45\x91\xda\xeb\x0e\x63\x48\x63\xf5\xb6\xb2\x70\x93\x26\x17\xac\x45\x00\x20\x89\x90\x6b\x00\xf8\x48\xd0\x94\x33\xc9\x5a\x9c\xab\x56\x58\x6c\x3b\x7b\x3d\xfb\x98\x26\x67\xf5\x1a\x0d\x00\xbc\x7b\x7f\x4a\x8f\xbd\x25\xe1\x60\xa6\xa6\xaf\xa8\x0a\xe3\x4c\xdd\x63\x34\x75\xd5\xed\xd9\x9e\xcb\x1a\xbf\xa2\x21\x5b\xf7\x18\x6d\x85\x97\x4f\x8c\x6f\x1d\x2c\x3e\xe4\x21\x2a\x5e\xfe\x1d\xa0\x78\xc7\x43\x4c\xfc\x5f\x8f\xcc\x3d\xd8\xbc\xa5\x30\xfd\x9f\x2b\xb3\x70\xb2\xe0\x41\x69\xf6\x3c\xd8\x1a\xee\xcb\x61\xd9\x7a\xea\x70\x29\xbe\x8d\x52\x9c\x0a\x69\xc3\x63\x70\x30\xe2\xdb\x5e\x8a\x17\x1b\xa6\x0d\x46\xb3\xd3\x21\x47\xf0\xa8\xbc\x7e\xea\xf4\x97\x14\x5f\xb6\x7d\xa2\x95\x52\xcd\x34\xcd\xd6\xe9\xa7\x3e\x17\xa2\x69\xd8\xaa\xc1\xbb\x7c\x64\xd0\x4f\xbd\x5e\x77\x56\x28\xc9\x9a\xbb\xbc\x54\xd0\x4f\xbd\x5e\x22\x67\xdb\xc6\xde\xb9\xbe\xda\xeb\xf7\x8a\xea\x6a\x66\x31\xba\x1e\x29\xca\xe9\x3f\x1c\xf5\x3d\x6f\xdb\xad\xed\xab\x3b\xf4\x15\x51\x3f\x75\xfb\x9b\x35\xa2\xa6\xdd\x4d\x5b\x06\x86\x6e\x45\xb7\x5d\xaf\x3f\xc2\x70\xb7\xbf\x0e\x09\xee\xc4\xdf\xc1\x6f\xe7\x77\x84\xde\x81\x11\xff\xce\xea\xa9\xe1\x3d\x64\xde\x33\xdc\xe7\xf0\x1b\xe4\x3e\xf9\xd4\x4e\x23\xff\x70\x98\xfd\x0d\xf2\xc0\xdc\xc9\xb8\xd1\xc8\xef\x60\x6c\xe8\xcd\x3d\x44\x3d\x97\x3b\xd4\x06\xf7\x4d\x85\x17\xef\xa7\xff\xb2\x15\x1a\xeb\x3d\x5b\x1d\xc4\x47\xba\xe6\xc7\xd7\x61\xdb\xbc\xfc\x3b\xfa\xe6\x1d\x87\xc6\x85\x4a\x7b\x0e\xde\x53\x69\x18\xd6\xef\xde\x4f\x91\xbe\x7b\x56\xef\x5b\x1e\x19\xd5\xbe\xca\x0b\xbc\x72\xfd\xa8\x34\x32\x8b\xae\xc8\x50\x11\x05\xf7\x65\xb9\xa7\x1a\x4d\xa5\x45\x67\x95\x2e\x52\xbe\x95\x55\xf4\xcc\xb0\x86\x53\xb2\x28\x5e\xf6\x16\x79\x68\xf2\x4d\x9a\x48\x84\x72\x09\x27\xf4\x7a\x93\x26\x44\xad\xd2\xd3\x00\xeb\xe2\x2d\x5b\xcf\x49\x76\xdd\x61\xd9\xcb\x88\x8d\x69\xe2\x58\xdd\x0b\xe9\x85\x84\x1e\xb1\xd2\x0b\xfd\x0b\x89\x03\x0f\x4a\x27\x0e\x2f\x24\x8f\x3d\x2f\x49\x1e\x5f\xbc\x82\x87\xf8\x4e\xc1\x43\xfc\xdb\x34\x11\x1c\x34\x72\x5a\xb2\xd7\x3c\x73\xaf\x8f\x96\x20\x45\x43\xe5\x24\x12\x49\x0c\xcb\xbe\x7c\x8d\x3c\x77\xae\x1a\xed\x56\x4b\x90\x18\x90\xfd\x83\x69\xb3\x61\x4d\x38\xa2\xdd\x55\x05\xdd\x9d\x67\x74\xe4\x0b\x69\x51\xd3\x0d\x82\x9e\x14\x30\xf8\xf5\xf2\xf5\x05\x39\x3b\x7a\x55\x4c\xc2\x8a\x90\x27\xd7\xda\x9b\x50\x80\xe0\xac\x56\x9f\xb0\xb2\xe1\x27\x34\x65\x92\x34\x33\x31\x37\xb1\x36\x64\xca\x21\x5b\xc1\xbb\xf7\xab\x6b\x8b\x73\x40\xad\xe9\x9f\x3a\x76\x93\x26\xc6\xb5\xca\xfb\xde\x78\x80\x84\xf4\x97\xb3\x2c\x5c\xa9\x5c\x7f\x5e\xf3\x10\x39\xcf\x5d\x6b\xb2\xfc\x36\x4d\x02\xc3\x5c\xc8\x72\x09\x86\x71\xf4\x5c\x8c\xb6\x0e\x5c\xd2\x8e\xd0\x8c\x98\x89\x66\x0e\xbc\xb5\xc5\x19\xad\x85\x67\xb3\xb0\xf0\xc7\x5f\x4a\x78\xbc\x9b\xcd\xc1\x78\x0a\x90\xbb\x07\x9b\x2b\x0d\x1f\xe6\xe0\x3a\xa5\x99\x24\xa6\x7a\xe2\x53\x54\x5e\x93\x98\x8f\x08\x99\xe5\x69\x92\x18\x67\x7d\xe2\x56\x45\x66\x23\x8e\xf9\x73\x7e\x20\xda\x88\x93\x51\x15\x89\x39\xa2\x70\xaf\xf2\x3c\x1e\xd1\x33\x6a\x06\x8e\xf6\xa7\x6a\x39\x24\x8b\xe7\x28\xa9\xe3\xf1\x39\xa8\xa3\xc4\xa9\xfb\x63\xab\x8c\xea\x5e\xe2\xf4\xc3\xf9\xe4\x0c\x1a\x94\x19\xaf\x8b\x41\x9a\x3b\xab\x70\x64\x96\xc3\x02\xe3\x21\xea\x7b\xe2\xab\x18\x9f\xae\xa5\xab\x62\x72\xde\x0e\xa6\xb7\x69\x42\x3d\xe5\x75\xe1\x6e\x33\x8f\x96\xf0\xc4\xe1\x9f\x18\xee\x25\x4b\x38\x09\xca\x89\x75\xbc\xc9\x3c\x5a\xc2\x6c\xd6\x7b\x44\xa9\x77\x0a\x6f\xc1\xcf\x14\x61\xae\x2d\x81\x75\x1d\xca\x3a\x8b\x92\x39\x98\xb0\xff\xfc\x30\x1c\xf3\xcf\x4d\xcd\x87\xa4\x1f\x0e\xf4\x73\xd9\x5d\x50\x53\xf8\x69\x3d\x5a\xea\x99\x5f\x5a\x3f\x31\x27\xcc\xcc\x7d\xc8\x78\x41\x1e\x17\x10\x2e\xd3\x0f\x59\x82\xa8\xbf\x0e\x45\x84\x35\xb8\xc0\x41\x21\xea\xaf\x07\xfb\xa8\x88\x97\xfc\x51\x89\xe7\x71\xf9\x27\xee\xc9\x35\xd5\x95\x5d\x82\x8b\xe1\x21\x20\xa9\xef\x5b\xe9\xa4\xa1\x87\xe3\xcd\x43\xe2\x61\xdb\xdc\x4e\x66\x2b\x9d\x65\x45\x18\x71\x99\xc9\xc3\xa0\x1d\x46\x0d\x5c\x69\xd6\x19\x37\x23\x7d\xfd\x91\x36\x2d\xda\x8d\xaa\xe1\x4a\xd8\x0d\x68\xac\xd4\x8e\xbe\x5a\x15\xa0\x34\x5b\x8d\x20\x15\x74\x4c\x8a\xca\xd0\x57\x69\xeb\xc3\x0b\xb9\x0e\x23\xf5\x60\x92\x1d\xcc\x53\x1e\xcf\xdc\xfe\xb3\x65\x7f\xb2\xd6\xc8\x51\x03\x85\xcb\x72\x8f\x2e\x87\x9d\xc3\xdd\x2f\x26\xcb\x9f\xc1\x6e\xdc\xd6\x84\xfc\x97\x47\x3a\x1a\x2b\xf2\x0b\x0e\xcd\xdd\x51\x5b\xc2\x04\x06\x17\xc4\xef\x9b\x5b\xea\x57\xc0\x6e\xe2\x9e\xe5\x73\x67\x35\x00\xe8\x39\x7b\x80\x9f\x17\xff\x28\x7c\xe3\x8d\x78\x80\x9e\xdf\x39\x1e\x3c\x32\x7c\x40\xec\x7c\x35\x47\xa0\xc3\xb0\x63\xef\x43\xce\x17\x71\x00\x5c\xdc\x0b\x07\xd0\x45\xc5\x8f\x82\x37\x1d\x02\x07\xf0\x89\xfe\x83\xbb\xbf\x9c\x3e\x20\x82\xb1\xa8\x23\x18\x8a\x7e\x28\xdc\x87\x62\xac\x66\xc0\xd1\x15\xda\xdf\x33\x2c\x8c\x6f\x1a\xf9\xe4\x8d\xd6\x46\x63\xcb\x16\xbf\x09\x59\x67\x39\x2c\x97\xbd\xfe\x4f\xab\xdd\xd2\xe9\xfc\xb0\xc5\x59\x83\x6d\x36\x19\x1d\x36\xbd\x4d\xff\x09\x00\x00\xff\xff\xf1\x30\x41\xcf\x0b\x13\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 4875, mode: os.FileMode(420), modTime: time.Unix(1566979817, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
