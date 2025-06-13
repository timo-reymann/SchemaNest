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

var _postgres01_initialize_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\xc8\x2a\xce\xcf\x8b\x2f\x4e\xce\x48\xcd\x4d\x54\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\x22\xa4\x32\xbe\x2c\xb5\xa8\x38\x33\x3f\x4f\x01\xae\x05\x10\x00\x00\xff\xff\xb9\x0b\xc5\xe0\x5c\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgres/01_initialize_schema.down.sql", size: 92, mode: os.FileMode(420), modTime: time.Unix(1745597948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres01_initialize_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xd1\x4e\xc3\x20\x14\x86\xef\x79\x8a\xff\x72\x4d\xfa\x06\x5e\x21\x3b\x1a\x22\x52\xa5\x34\x71\x57\x4d\xd3\x62\xc6\xcc\x60\xb6\xe8\xf3\x1b\xed\xba\xac\x9d\x53\x6e\xf9\xe0\xff\xcf\x77\x84\x21\x6e\x09\x96\xdf\x2a\xc2\x6e\x88\xa1\x1e\xda\xad\xdb\x37\x6c\xc5\x00\xc0\x77\x98\x4e\x49\x46\x72\x85\x27\x23\x1f\xb9\xd9\xe0\x81\x36\xf9\x11\x71\x21\xf9\x57\xef\x7a\x58\x7a\xb1\xd0\x85\x85\xae\x94\x42\xa5\xe5\x73\x45\x2c\xbb\x61\xec\x5a\x4a\xfd\xe9\xfa\xc1\xc7\x70\x99\xf6\x57\xe2\xf1\x51\xbd\x6f\x76\xb1\x07\xa4\xb6\x74\x4f\xe6\x14\xbc\x80\x7c\xf8\x1f\x3a\x34\xa9\xdd\x5e\x83\xda\x18\x92\x0b\x69\x6a\xf5\x33\x24\xb0\x80\x3a\x37\xb4\xbd\x3f\x24\x1f\xc3\x04\x8d\x17\xe7\xd3\xfa\xee\x22\x02\x86\xee\xc8\x90\x16\x54\x9e\xa3\x58\xf9\x2e\x43\xa1\xb1\x26\x45\x96\x20\x78\x29\xf8\x9a\xc6\x2f\x45\xa1\x4b\x6b\xb8\xd4\xf6\x37\x97\xf5\xc7\xfb\xdb\xc9\xeb\x54\x7a\xdc\x05\x56\x33\x75\xf9\x5c\x52\x3e\xd7\x91\x2f\xaa\x67\xdf\x9b\xfc\x0a\x00\x00\xff\xff\x0a\xcd\xc9\xff\x2f\x02\x00\x00")

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

	info := bindataFileInfo{name: "postgres/01_initialize_schema.up.sql", size: 559, mode: os.FileMode(420), modTime: time.Unix(1749483609, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite301_initialize_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\xc8\x2a\xce\xcf\x8b\x2f\x4e\xce\x48\xcd\x4d\x8c\x2f\x4b\x2d\x2a\xce\xcc\xcf\xb3\xe6\x22\xa4\xd2\x9a\x0b\x10\x00\x00\xff\xff\xdc\x18\x48\xc0\x4c\x00\x00\x00")

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

	info := bindataFileInfo{name: "sqlite3/01_initialize_schema.down.sql", size: 76, mode: os.FileMode(420), modTime: time.Unix(1745597967, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite301_initialize_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\x41\x52\xc3\x30\x0c\xbc\xfb\x15\x3a\xb6\x33\xf9\x01\x67\x7e\xd0\x03\xb7\x8c\x71\xb6\x54\x4d\x22\xa7\xb2\xc2\x4c\x7f\xcf\x04\xea\x62\xa7\x53\x0a\xe8\x96\x78\xa5\xd5\xae\x36\x28\xbc\x81\xcc\xbf\x0e\xa0\x63\x8a\xd2\xa6\x70\xc0\xe8\xdd\xc6\x11\x11\x71\x47\xb9\x58\x0c\x6f\x50\x97\xbf\x43\x94\x64\xea\x59\xac\xec\x6b\xa7\xfe\x8a\x58\x6a\x52\x1e\xbd\x9e\xa9\xc7\x99\xfc\x6c\x91\x25\x28\x46\x88\x35\x97\xf9\x10\xe3\x3d\x43\x69\xf7\xfc\xb2\x23\x89\x46\x32\x0f\xc3\x23\x96\xf9\xd4\xb7\xdf\xbd\x15\xe3\x2c\x7c\x9a\xe1\xb6\x4f\xce\xdd\x13\xd7\xbe\x43\x13\x47\xb9\x15\x59\x08\xfd\xf5\x2e\x97\x61\x7f\x54\x9e\xbb\x46\x7f\x8c\x7a\x4b\xba\x02\xb1\x3c\x06\x4d\xde\xc2\xe1\x1e\x28\x44\x31\x88\xe5\xe5\x3e\xcd\x26\x5a\x81\x3a\xa4\xa0\x3c\x19\x47\xc9\xa0\xaf\x87\x52\x2e\x77\xff\x76\xa8\x2d\x7f\xee\x6b\xbb\x14\x7b\x28\x24\x20\x55\x31\x2c\x21\x51\xa8\xc3\x00\x03\x05\x9f\x82\xef\x70\x95\xf6\x13\xe9\x12\x95\x7c\xef\x3a\x23\xb4\xa9\x8e\xd0\xd4\x76\x37\xb5\xb1\xcd\xca\x84\xed\x92\xb0\x8f\x00\x00\x00\xff\xff\xbe\xa7\x45\x30\x3e\x03\x00\x00")

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

	info := bindataFileInfo{name: "sqlite3/01_initialize_schema.up.sql", size: 830, mode: os.FileMode(420), modTime: time.Unix(1749483609, 0)}
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
