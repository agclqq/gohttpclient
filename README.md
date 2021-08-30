# gohttpclient
[中文版](README_ZH.MD)

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
	url:="https://httpbin.org/"
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
	resPost,err:=client.Post(url+"post",header,strings.NewReader(postData))
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
	fmt.Printf("post request:\n  code of resPost:%d\n  header of resPost:%v\n  body of resGet:%s\n", resPost.StatusCode, resPost.Header,bodyPost)

	//------postForm------
	resPostForm,err:=client.PostForm(url+"post",header,formData)
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
	fmt.Printf("postForm request:\n  code of resPostForm:%d\n  header of resPostForm:%v\n  body of resPostForm:%s\n", resPostForm.StatusCode, resPostForm.Header,bodyPostForm)

	//------put------
	putHeader:=header
	putHeader["Content-Type"] = "application/json"
	resPut,err:=client.Put(url+"put",putHeader,strings.NewReader(postData))
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
	fmt.Printf("put request:\n  code of resPut:%d\n  header of resPut:%v\n  body of resPut:%s\n", resPut.StatusCode, resPut.Header,bodyPut)

	//------putForm------
	resPutForm,err:=client.PutForm(url+"put",header,formData)
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
	fmt.Printf("putForm request:\n  code of resPutForm:%d\n  header of resPutForm:%v\n  body of resPutForm:%s\n", resPutForm.StatusCode, resPutForm.Header,bodyPutForm)

	//------delete------
	resDelete,err:=client.Delete(url+"delete",header)
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
	fmt.Printf("delete request:\n  code of resDelete:%d\n  header of resDelete:%v\n  body of resDelete:%s\n", resDelete.StatusCode, resDelete.Header,bodyDelete)
}
```