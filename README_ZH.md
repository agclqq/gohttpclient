# gohttpclient
This is a GO version of the restful API client that currently supports the following methods (including the WithContext method)：
* Head
* Get
* Post
* PostForm
* Put
* PutForm
* Delete
## install
```shell
go get github.com/agclqq/gohttpclient
```


## useage
```go
package main

import (
	"fmt"
	"github.com/agclqq/gohttpclient"
	"io"
	"strings"
)
func main()  {
	url:="http://127.0.0.1:8080/test/?t1=a"
	header:=make(map[string]string)
	header["x-a"]="abc0"

	formData:=make(map[string][]string)
	formData["a"]=[]string{"A"}
	formData["b"]=[]string{"B"}
	formData["c"]=[]string{"C1","C2"}

	postData:="{\"ping\":\"pong\"}"

	client:=httpclient.New()

	//------head------
	resHead,err:=client.Head(url,header)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("head request:\n  code of resGet:%d\n  header of resGet:%v\n", resHead.StatusCode, resHead.Header)
	//------get------
	resGet,err:=client.Get(url,header)
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyGet,err:=io.ReadAll(resGet.Body)
	resGet.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("get request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyGet)

	//------post------
	resPost,err:=client.Post(url,header,strings.NewReader(postData))
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyPost,err:=io.ReadAll(resPost.Body)
	resPost.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("post request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyPost)

	//------postForm------
	resPostForm,err:=client.PostForm(url,header,formData)
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyPostForm,err:=io.ReadAll(resPostForm.Body)
	resPostForm.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("postForm request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyPostForm)

	//------put------
	putHeader:=header
	putHeader["Content-Type"] = "application/json"
	resPut,err:=client.Put(url,putHeader,strings.NewReader(postData))
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyPut,err:=io.ReadAll(resPut.Body)
	resPut.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("put request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyPut)

	//------putForm------
	resPutForm,err:=client.PutForm(url,header,formData)
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyPutForm,err:=io.ReadAll(resPutForm.Body)
	resPutForm.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("putForm request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyPutForm)

	//------delete------
	resDelete,err:=client.Delete(url,header)
	if err!=nil {
		fmt.Println(err)
		return
	}
	bodyDelete,err:=io.ReadAll(resDelete.Body)
	resDelete.Body.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("delete request:\n  code of resGet:%d\n  header of resGet:%v\n  body of resGet:%s\n", resGet.StatusCode, resGet.Header,bodyDelete)

}
```