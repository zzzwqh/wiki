package main

import "io/ioutil"

//func main() {
//	// 内部封装了 OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, perm)
//	ioutil.WriteFile("writeTest.txt", []byte("This is www.itsky.tech"), 0666)
//
//}

func main() {
	ioutil.WriteFile("fileWrite.txt", []byte{'d', 'a'}, 0666)
}
