package filepathhelper

import (
	"errors"
	"github.com/apaxa-io/iohelper/ioutilhelper"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func rename(path string, info os.FileInfo, err error) error {
	error := os.Rename(path, path+"1")
	return error
}

//TODO check sort = false
func TestWalk(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("Got error while creating temp dir: %v", errTempDir)
	}
	errChdir := os.Chdir(nameDir)
	if errChdir != nil {
		t.Errorf("Got error while changing dir: %v", errChdir)
	}
	dirs := []string{"dir4", "dir3", "dir2", "dir1"}
	for i, v := range dirs {
		errMkdir := os.Mkdir(v, 0777)
		if errMkdir != nil {
			t.Errorf("Got error while making dir №%v: %v", i, errMkdir)
		}
	}
	//call Walk
	err := Walk(nameDir, rename, true, true)
	if err != nil {
		t.Errorf("Walk error: %v", err)
	}
	//check dirs
	sNew := []string{"dir11", "dir21", "dir31", "dir41"}
	s, err1 := ioutilhelper.ReadDirNames(nameDir+"1", true)
	if !reflect.DeepEqual(s, sNew) {
		t.Errorf("Check dir names:%v\nerror: %v", s, err1)
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir + "1")
	if errRemove != nil {
		t.Errorf("Got error while removing dir: %v", errRemove)
	}
	err = Walk("", rename, true, true)
	if err == nil {
		t.Errorf("Error expected but got nil")
	}
}

func TestWalk2(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("Got error while creating temp dir: %v", errTempDir)
	}
	errChdir := os.Chdir(nameDir)
	if errChdir != nil {
		t.Errorf("Got error while changing dir: %v", errChdir)
	}
	dirs := []string{"dir4", "dir3", "dir2", "dir1"}
	for i, v := range dirs {
		errMkdir := os.Mkdir(v, 0777)
		if errMkdir != nil {
			t.Errorf("Got error while making dir №%v: %v", i, errMkdir)
		}
	}
	//call Walk
	err := Walk(nameDir, rename, false, true)
	if err != nil {
		t.Errorf("Walk error: %v", err)
	}
	//check dirs
	sNew := []string{"dir11", "dir21", "dir31", "dir41"}
	s, err1 := ioutilhelper.ReadDirNames(nameDir, true)
	if !reflect.DeepEqual(s, sNew) {
		t.Errorf("Check dir names:%v\nerror: %v", s, err1)
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir)
	if errRemove != nil {
		t.Errorf("Got error while removing dir: %v", errRemove)
	}
}

func rename3(path string, info os.FileInfo, err error) error {
	return errors.New("Func for return error.")
}

func TestWalk3(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("Got error while creating temp dir: %v", errTempDir)
	}
	//call Walk
	err := Walk(nameDir, rename3, true, true)
	if err == nil {
		t.Errorf("Error expected but got nil")
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir)
	if errRemove != nil {
		t.Errorf("Got error while removing dir: %v", errRemove)
	}
}
