GoDS
===========

Simple lib for text synchronization between server and client based on [Differential Synchronization](https://neil.fraser.name/writing/sync/) (DS) method by N. Fraser.
This is basic implementation of this methods and readme document describe basic information about method. For better understanding of method read link above.

#### Install
Installation process is standard as any other package installation. 
Type the following command in your terminal:

	go get github.com/lehovec/gods

After it the package is ready to use.

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/lehovec/gods"
```
No more is needed.

#### Usage:
GoDS is lib for client server communication but DS method is same for both sides.
Client or server side of lib is created by:
```go
gods.New()
```
Every instance of gods synchronize one document. For every document you must create new instance of gods struct.

DS method creates sync cycles. Every initialisation of GoDS can communicate as server or as client at the same time by creating synchronization cycles. Every cycle synchronize document between client and server.
You can create cycles by:
```go
diffSync := gods.New()
diffSync.AddClientConnection(<id>)
diffSync.AddServerConnection(<id>)
```
<id> param identifies synchronization cycle. Id can be any comparable type (see [Go comparison operators](https://golang.org/ref/spec#Comparison_operators))

There are two main method of GoDS lib. Fist is **_GetPatch_**. This method generate patches for send to client or server. Method receive id of connection cycle.
```go
diffSync := gods.New()
// connection creation
patches, err := diffSync.GetPatch(<id>)
```

Second method is **_ApplyPatch_**. Method receive generate patches and apply them to document. This method then return patched document.
```go
diffSync := gods.New()
// connection creation
document, err := diffSync.ApplyPatch(patches, <id>)
```

All documents is save to memory ( to variable ) by default. If you want to use another storage, you can control saving and obtaining of documents by impelmenting **_Storage_** interface.
```go
type MyStorage struct {
}
func (storage *MyStorage) Get(docType int, id interface{}) (string, error) {
	// Your storage code
}
func (storage *MyStorage) Set(docType int, id interface{}, text string) error {
	// Your storage code
}
myStorage := new(MyStorage)
diffSync := gods.New()
diffSync.SetStorage(myStorage)
```
Methods Set and Get save three types of documents defined by first param docType. There are three constants defined in GoDS which represents document types
* DOCTYPE_DOCUMENT
* DOCTYPE_SHADOW
* DOCTYPE_BACKUP
Second param is id of connection cycle. Every connection cycle has own unique DOCTYPE_SHADOW, and DOCTYPE_BACKUP. DOCTYPE_DOCUMENT is common for entire GoDS instance.
Get function returns document content.
Set function last param is saved document content.

#### Examples
All examples is available in examples dir.
There is only one example but more examples will be added in future

#### Notes
Lib is not tested and I am beginner at Go. Please criticize my code or my english and creates issues. Thanks.
Test will be added later.

#### Support
If you do have a contribution for the package feel free to put up a Pull Request or open Issue.
