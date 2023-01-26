// Code generated for package server by go-bindata DO NOT EDIT. (@generated)
// sources:
// frontend/public/404.css
// frontend/public/404.in.html
// frontend/public/bootstrap.min.css
// frontend/public/tty-share.in.html
// frontend/public/tty-share.js
package server

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	return data, nil
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

//go:embed frontend/public/404.css
var __404Css string

func _404CssBytes() ([]byte, error) {
	return bindataRead(
		[]byte(__404Css),
		"404.css",
	)
}

func _404Css() (*asset, error) {
	bytes, err := _404CssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "404.css", size: 3393, mode: os.FileMode(420), modTime: time.Unix(1668774616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

//go:embed frontend/public/404.in.html
var __404InHtml string

func _404InHtmlBytes() ([]byte, error) {
	return bindataRead(
		[]byte(__404InHtml),
		"404.in.html",
	)
}

func _404InHtml() (*asset, error) {
	bytes, err := _404InHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "404.in.html", size: 615, mode: os.FileMode(420), modTime: time.Unix(1668774616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

//go:embed frontend/public/bootstrap.min.css
var _bootstrapMinCss string

func bootstrapMinCssBytes() ([]byte, error) {
	return bindataRead(
		[]byte(_bootstrapMinCss),
		"bootstrap.min.css",
	)
}

func bootstrapMinCss() (*asset, error) {
	bytes, err := bootstrapMinCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.min.css", size: 140930, mode: os.FileMode(420), modTime: time.Unix(1668774616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

//go:embed frontend/public/tty-share.in.html
var _ttyShareInHtml string

func ttyShareInHtmlBytes() ([]byte, error) {
	return bindataRead(
		[]byte(_ttyShareInHtml),
		"tty-share.in.html",
	)
}

func ttyShareInHtml() (*asset, error) {
	bytes, err := ttyShareInHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tty-share.in.html", size: 532, mode: os.FileMode(420), modTime: time.Unix(1668774616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

//go:embed frontend/public/tty-share.js
var _ttyShareJs string

func ttyShareJsBytes() ([]byte, error) {
	return bindataRead(
		[]byte(_ttyShareJs),
		"tty-share.js",
	)
}

func ttyShareJs() (*asset, error) {
	bytes, err := ttyShareJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tty-share.js", size: 284790, mode: os.FileMode(420), modTime: time.Unix(1668776956, 0)}
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
	"404.css":           _404Css,
	"404.in.html":       _404InHtml,
	"bootstrap.min.css": bootstrapMinCss,
	"tty-share.in.html": ttyShareInHtml,
	"tty-share.js":      ttyShareJs,
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
	"404.css":           &bintree{_404Css, map[string]*bintree{}},
	"404.in.html":       &bintree{_404InHtml, map[string]*bintree{}},
	"bootstrap.min.css": &bintree{bootstrapMinCss, map[string]*bintree{}},
	"tty-share.in.html": &bintree{ttyShareInHtml, map[string]*bintree{}},
	"tty-share.js":      &bintree{ttyShareJs, map[string]*bintree{}},
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
