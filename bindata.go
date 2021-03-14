// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// resources/index.html
package main

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

var _resourcesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x4d\x6f\xd3\x40\x10\x3d\x37\xbf\x62\xb4\x07\xe4\x34\x89\xb7\x5f\x02\x89\xda\x41\x80\x2a\x81\xc4\x09\x90\x80\xa3\xbb\x3b\xa9\x57\xf5\x7a\x57\xde\x49\x6a\x37\xca\x7f\x47\xfb\x91\xc4\x01\x4a\x4e\x33\x2f\xf3\xde\xec\xcc\x3c\xb9\xa8\x49\x37\xcb\x49\x51\x63\x25\x97\x93\xb3\x82\x14\x35\xb8\xdc\x6e\xf3\xef\x3e\xd8\xed\x0a\x1e\x91\x09\x00\x40\xd1\xa8\xf6\x11\x3a\x6c\x4a\xe6\x68\x68\xd0\xd5\x88\xc4\xa0\xee\x70\x55\xb2\x9a\xc8\xba\xb7\x9c\xeb\xaa\x17\xb2\xcd\xef\x8d\x21\x47\x5d\x65\x7d\x22\x8c\xe6\x07\x80\x5f\xe7\xd7\xf9\x1b\x2e\x9c\x3b\x62\xb9\x56\x6d\x2e\x9c\x63\xa9\x91\x46\xaa\xa0\xad\x34\x96\x6c\xa3\xf0\xc9\x9a\x8e\x18\x08\xd3\x12\xb6\x54\xb2\x27\x25\xa9\x2e\x25\x6e\x94\xc0\x45\x48\xe6\xa0\x5a\x45\xaa\x6a\x16\x4e\x54\x0d\x96\x97\xcc\x0f\xe3\x44\xa7\x2c\x45\x49\xff\xdb\x54\x1d\xe8\x1e\x4a\xb8\xc8\x2f\x6e\x4f\xd1\xe1\x5f\x68\x57\x49\xb5\x76\x50\xc2\xd5\xf8\x9f\xd5\xba\x15\xa4\x4c\x0b\xcf\xc6\xe8\xf7\x94\xa5\x07\xd4\xa8\x1e\x6a\x9a\x43\x3f\x87\x61\x0a\xdb\x43\xf9\x5e\xcc\xaa\x1e\x9b\x6f\xea\x19\xa1\x84\x2c\x29\x9f\xc3\xd5\x14\x38\xbc\xbe\x19\xc9\xef\xeb\xfd\xdc\x9f\x82\xa6\x27\x44\x75\xe0\x10\xba\x4d\xe1\xfc\x44\xe3\x76\x72\x36\xa6\x36\xb8\x0a\x24\xdd\xc3\x22\x0d\x31\x85\x19\x64\xc7\x27\xf0\xbf\x49\x64\x6c\xe0\x0c\xb0\x80\x6c\xd4\xdc\x97\xfe\x9f\x1d\x56\x1a\x7a\xce\xa0\x87\xf3\xe3\xa4\xe3\x1a\xbf\x60\xdf\x62\x06\xc3\x0b\x25\x87\x6d\xa7\x80\xc3\xcd\x78\xed\xbb\xc9\x59\xc1\xf7\x17\x2d\x78\x34\x6c\x71\x6f\xe4\x90\x3c\x53\x5f\x9e\xf8\xb6\xbe\x4c\xb8\x5d\x7e\x6c\x94\x78\x04\xa5\xab\x07\x04\x32\xe1\x6e\x39\x7c\xc5\xc6\x54\x12\x6c\x02\x3b\x74\x54\x75\x94\x17\xdc\x26\x9e\x54\x9b\xa3\x75\x8a\xc8\x56\xb2\x64\x56\x09\x06\xae\x13\x25\xe3\x21\x8c\x6e\x64\xdb\x6d\xfe\xc3\x47\xbb\x9d\xed\x59\x72\x43\x40\xe3\x1e\x23\xcc\x93\x36\x0f\xe2\x2f\x58\xd4\x2a\x01\x25\x48\x23\xd6\x1a\x5b\xca\x1f\x90\xee\x1a\xf4\xe1\x87\xe1\xb3\xcc\x42\xff\xe9\x71\x2f\x56\x89\xbc\x92\xf2\x6e\x83\x2d\x7d\x51\x8e\xb0\xc5\x2e\x63\xc2\x8f\xcc\xe6\x07\xb3\x66\xf8\xa7\x27\x93\x7b\x3d\x3d\x39\xd8\x87\x7b\x17\x63\x6e\x56\x2b\x87\xf4\xf3\x18\xfe\x9a\x9e\xf0\x7d\xb5\xeb\xfc\x4b\xc3\x1e\xde\xe9\xbe\x64\x30\xf3\x66\x98\x01\x7b\xa5\x87\x98\x0d\x21\x8b\x17\x0d\x48\x0c\x47\x67\xf5\x4e\x1a\x1f\x36\x5e\xb4\xe0\xf1\xc3\xf4\x3b\x00\x00\xff\xff\x39\xa3\x23\x5f\xa0\x04\x00\x00")

func resourcesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesIndexHtml,
		"resources/index.html",
	)
}

func resourcesIndexHtml() (*asset, error) {
	bytes, err := resourcesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/index.html", size: 1184, mode: os.FileMode(420), modTime: time.Unix(1523352264, 0)}
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
	"resources/index.html": resourcesIndexHtml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{resourcesIndexHtml, map[string]*bintree{}},
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
