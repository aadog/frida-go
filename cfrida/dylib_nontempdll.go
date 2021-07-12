// +build !tempdll
// +build !memorydll

package cfrida

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
