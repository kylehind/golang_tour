// I DID NOT GET WHAT I WAS DOING HERE - COME BACK
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read([]byte) (int, error) {

}

func main() {
	reader.Validate(MyReader{})
}
