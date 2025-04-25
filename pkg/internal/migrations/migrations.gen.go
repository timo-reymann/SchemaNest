// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../migrations/postgres/01_initialize_schema.down.sql
// ../../../migrations/postgres/01_initialize_schema.up.sql
// ../../../migrations/sqlite3/01_initialize_schema.down.sql
// ../../../migrations/sqlite3/01_initialize_schema.up.sql
package migrations

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

var _postgres01_initialize_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\xc8\x2a\xce\xcf\x8b\x2f\x4e\xce\x48\xcd\x4d\xb4\xe6\x22\xa4\x22\xbe\x2c\xb5\xa8\x38\x33\x3f\xcf\x1a\x10\x00\x00\xff\xff\xbc\x27\x60\x32\x4b\x00\x00\x00")

func postgres01_initialize_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres01_initialize_schemaDownSql,
		"postgres/01_initialize_schema.down.sql",
	)
}

func postgres01_initialize_schemaDownSql() (*asset, error) {
	bytes, err := postgres01_initialize_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/01_initialize_schema.down.sql", size: 75, mode: os.FileMode(420), modTime: time.Unix(1745573910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres01_initialize_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xd1\x4e\x83\x30\x18\x85\xef\xfb\x14\xe7\x12\x12\xde\xc0\xab\xda\xfd\x9a\xc6\x5a\xb4\xfc\x24\xee\x8a\x10\xa8\x59\x67\x56\x14\xd0\xe7\x37\xba\xb1\x0c\xe6\x1c\xd7\x1f\xff\x39\xe7\xab\x72\x24\x99\xc0\xf2\xd6\x10\xb6\x43\x17\xab\xa1\xd9\xf8\x5d\x2d\x12\x01\x00\xa1\xc5\xf4\x15\xe4\xb4\x34\x78\x72\xfa\x51\xba\x35\x1e\x68\x9d\x1d\x10\x1f\xc7\xf0\x1a\x7c\x0f\xa6\x17\x86\xcd\x19\xb6\x34\x06\xa5\xd5\xcf\x25\x89\xf4\x46\x88\x4b\x29\xd5\x97\xef\x87\xd0\xc5\xf3\xb4\xff\x12\x0f\x3f\x55\xbb\x7a\xdb\xf5\x80\xb6\x4c\xf7\xe4\x8e\xc1\x0b\x28\xc4\xeb\xd0\x7b\x3d\x36\x9b\x4b\x50\xd3\xc5\xd1\xc7\x71\x6a\xf5\x3b\x12\x58\x40\xa7\xa3\x42\x7b\x76\x09\x8e\xee\xc8\x91\x55\x54\x9c\xa2\x48\x42\x9b\x22\xb7\x58\x91\x21\x26\x28\x59\x28\xb9\xa2\xfd\x49\x95\xdb\x82\x9d\xd4\x96\xff\x52\x56\x7d\x7e\xbc\x1d\xf5\x4d\xdd\xf6\xca\x91\xcc\x0c\x65\x73\x17\xd9\x7c\x75\xb6\xa8\x9e\xfe\x3c\xd8\x77\x00\x00\x00\xff\xff\x56\xd0\x54\x3a\x16\x02\x00\x00")

func postgres01_initialize_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres01_initialize_schemaUpSql,
		"postgres/01_initialize_schema.up.sql",
	)
}

func postgres01_initialize_schemaUpSql() (*asset, error) {
	bytes, err := postgres01_initialize_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/01_initialize_schema.up.sql", size: 534, mode: os.FileMode(420), modTime: time.Unix(1745576250, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite301_initialize_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\xc8\x2a\xce\xcf\x8b\x2f\x4e\xce\x48\xcd\x4d\xb4\xe6\x22\xa4\x22\xbe\x2c\xb5\xa8\x38\x33\x3f\xcf\x1a\x10\x00\x00\xff\xff\xbc\x27\x60\x32\x4b\x00\x00\x00")

func sqlite301_initialize_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite301_initialize_schemaDownSql,
		"sqlite3/01_initialize_schema.down.sql",
	)
}

func sqlite301_initialize_schemaDownSql() (*asset, error) {
	bytes, err := sqlite301_initialize_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/01_initialize_schema.down.sql", size: 75, mode: os.FileMode(420), modTime: time.Unix(1745573910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite301_initialize_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xc1\x6e\xb3\x30\x10\x84\xef\x7e\x8a\x3d\x26\x12\x6f\xf0\x9f\xff\x37\xc8\xa1\x37\xb4\x35\x43\xb3\x01\xd6\x64\xbd\x54\xca\xdb\x57\x54\x21\xc5\x44\x69\xda\xfa\x06\x9a\xf5\xec\x7c\x9e\x68\x60\x07\x39\xbf\xf6\xa0\x53\x4e\x5a\xe7\x78\xc4\xc0\x61\x17\x88\x88\xa4\xa1\xe5\x88\x3a\xde\x60\x61\xf9\x8e\x49\xb3\x1b\x8b\xfa\x7a\xae\x1e\xbb\x9b\x62\x3e\xa3\xc9\xc0\x76\xa1\x0e\x17\xe2\xc9\x93\x68\x34\x0c\x50\xaf\xae\xf7\x43\x5d\x5a\x81\xd1\xe1\xff\xcb\x81\x34\x39\xe9\xd4\xf7\xcf\x5c\xa6\x73\x57\x7f\xcd\x16\x8e\x93\xca\x79\x42\xd8\xff\x0b\xe1\x51\xb8\xfa\x1d\x96\x25\xe9\x7d\xc8\x55\xd0\x1f\xef\x72\xbd\xec\x97\xc9\x97\xa9\x81\x4f\xc9\xee\x4d\x37\x22\xd1\xe7\xa2\x91\x3d\x1e\x1f\x89\x62\x52\x87\xfa\xb2\xdc\x27\x6c\xa2\x8d\x68\x9d\x4a\x9a\x3f\x83\xa8\xd7\x3f\xdb\x92\x8a\xa1\x85\x41\x23\x72\xd1\xb6\xb5\x24\x29\x35\xe8\xe1\xa0\xc8\x39\x72\x83\x5b\x82\xef\x4c\xe7\x46\x2c\xcf\x5a\x56\x81\x76\x05\xeb\xaa\xa4\x5a\x95\xfc\xaa\x0d\x84\xfd\x5c\xa4\x8f\x00\x00\x00\xff\xff\x89\x87\x38\xac\x25\x03\x00\x00")

func sqlite301_initialize_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite301_initialize_schemaUpSql,
		"sqlite3/01_initialize_schema.up.sql",
	)
}

func sqlite301_initialize_schemaUpSql() (*asset, error) {
	bytes, err := sqlite301_initialize_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/01_initialize_schema.up.sql", size: 805, mode: os.FileMode(420), modTime: time.Unix(1745575789, 0)}
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
	"postgres/01_initialize_schema.down.sql": postgres01_initialize_schemaDownSql,
	"postgres/01_initialize_schema.up.sql":   postgres01_initialize_schemaUpSql,
	"sqlite3/01_initialize_schema.down.sql":  sqlite301_initialize_schemaDownSql,
	"sqlite3/01_initialize_schema.up.sql":    sqlite301_initialize_schemaUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"postgres": &bintree{nil, map[string]*bintree{
		"01_initialize_schema.down.sql": &bintree{postgres01_initialize_schemaDownSql, map[string]*bintree{}},
		"01_initialize_schema.up.sql":   &bintree{postgres01_initialize_schemaUpSql, map[string]*bintree{}},
	}},
	"sqlite3": &bintree{nil, map[string]*bintree{
		"01_initialize_schema.down.sql": &bintree{sqlite301_initialize_schemaDownSql, map[string]*bintree{}},
		"01_initialize_schema.up.sql":   &bintree{sqlite301_initialize_schemaUpSql, map[string]*bintree{}},
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
