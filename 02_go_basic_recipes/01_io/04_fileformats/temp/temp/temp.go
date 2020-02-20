package temp

import (
	"io/ioutil"
	"os"
)


// Let's explore some basic ways to work with temporary
// files
func CreateTempFile(dirLocation string) (string, *os.File, error) {

	// --> Create TempDir
	// If you are in need of a temp place to store files with
	// the same name i.e template-1-10.html a temp directory
	// is the solution.

	// Check if tempDir location exists

	tDir, err := ioutil.TempDir(dirLocation, "tmp")
	if err != nil {
		return "", nil, err
	}

	// Use defer to schedule cleanup for the directory at exit
	// defer os.RemoveAll(t)

	// ---> Create TempFile
	// Directory must exist, so check for errors
	tf, err := ioutil.TempFile(tDir, "tmp")
	if err != nil {
		return tDir, nil, err
	}

	// Here we return the name of the temp dir for cleanup, the filepointer
	// for usage of the temp file and an empty error
	return tDir, tf, nil
}
