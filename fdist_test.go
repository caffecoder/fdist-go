package fdist

import (
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	var fd = NewFileDistribution("/tmp/storage")

	exp := "/tmp/storage"
	val := fd.GetPath()

	if val != exp {
		t.Errorf("Expected %s, got %s.", exp, val)
	}
}

func TestCase1(t *testing.T) {
	var fd = NewFileDistribution("/tmp/storage")
	fd.SetExtension("tmp")
	fd.SetExtension(".dat")
	fd.HexPath(102423)

	exp := "/tmp/storage/01/90/17.dat"
	val := fd.GetPath()

	if val != exp {
		t.Errorf("Expected %s, got %s.", exp, val)
	}
}

func TestCase2(t *testing.T) {
	var fd = NewFileDistribution("/tmp/storage")
	fd.SetExtension("dat")
	fd.HexPath(256)

	exp := "/tmp/storage/01/00.dat"
	val := fd.GetPath()

	if val != exp {
		t.Errorf("Expected %s, got %s.", exp, val)
	}
}

// in most cases this is wrong way
func TestCase3(t *testing.T) {
	var fd = NewFileDistribution("/tmp/storage")
	fd.SetExtension("")
	fd.HexPath(256)

	exp := "/tmp/storage/01/00"
	val := fd.GetPath()

	if val != exp {
		t.Errorf("Expected %s, got %s.", exp, val)
	}
}

func TestCase4(t *testing.T) {
	var fd = NewFileDistribution("/tmp/storage")
	fd.HexPath(1)

	exp := "/tmp/storage/01.dat"
	val := fd.GetPath()

	if val != exp {
		t.Errorf("Expected %s, got %s.", exp, val)
	}
}

func TestCase5(t *testing.T) {
	srcFile := "/tmp/storage/test.txt"
	dstDir := "/tmp/storage"

	var fd = NewFileDistribution(dstDir)
	fd.SetExtension(".dat")
	fd.HexPath(256)

	os.MkdirAll(dstDir, os.ModeDir|os.ModePerm)
	os.Create(srcFile)

	_, err1 := os.Stat(srcFile)

	if os.IsNotExist(err1) {
		t.Error(err1)
	}

	fd.RenameFrom(srcFile)

	_, err2 := os.Stat("/tmp/storage/01/00.dat")

	if os.IsNotExist(err2) {
		t.Error(err2)
	}

	os.RemoveAll(dstDir)
}

func TestCase6(t *testing.T) {
	srcFile1 := "/tmp/storage/test1.txt"
	srcFile2 := "/tmp/storage/test2.txt"
	dstDir := "/tmp/storage"

	var fd = NewFileDistribution(dstDir)
	fd.SetExtension(".dat")

	os.MkdirAll(dstDir, os.ModeDir|os.ModePerm)
	os.Create(srcFile1)
	os.Create(srcFile2)

	fd.HexPath(1)
	fd.RenameFrom(srcFile1)

	fd.HexPath(256)
	fd.RenameFrom(srcFile2)

	_, err1 := os.Stat("/tmp/storage/01.dat")
	_, err2 := os.Stat("/tmp/storage/01/00.dat")

	if os.IsNotExist(err1) {
		t.Error(err1)
	}

	if os.IsNotExist(err2) {
		t.Error(err2)
	}

	os.RemoveAll(dstDir)
}
