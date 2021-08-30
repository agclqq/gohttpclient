package httpclient

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var httpClient=New()
var uri ="http://127.0.0.1:8080/test/?paramA=123"

func TestClientHead(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	res,err:=httpClient.Head(uri,header)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	fmt.Printf("Test_client_Get:%v\n",res.Header)
}
func TestClientGet(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	res,err:=httpClient.Get(uri,header)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",string(rs))
}

func TestClientPost(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	header["Content-Type"] = "application/json"
	body:=strings.NewReader("{\"ping\":\"pong\"}")
	res,err:=httpClient.Post(uri,header,body)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",rs)
}

func TestClientPostForm(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	data:=make(map[string][]string)
	data["test"]=[]string{"t1","t2"}
	res,err:=httpClient.PostForm(uri,header,data)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",rs)
}

func TestClientPut(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	header["Content-Type"] = "application/json"
	body:=strings.NewReader("{\"ping\":\"pong\"}")
	res,err:=httpClient.Put(uri,header,body)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",rs)
}

func TestClientPutForm(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	data:=make(map[string][]string)
	data["test"]=[]string{"t1","t2"}
	res,err:=httpClient.PutForm(uri,header,data)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",rs)
}

func TestClientDelete(t *testing.T) {
	header:=make(map[string]string)
	header["x-a"]="a"
	res,err:=httpClient.Delete(uri,header)
	if err!=nil{
		t.Errorf("Test_client_Delete has error, got:%v",err)
	}
	rs,_:=io.ReadAll(res.Body)
	fmt.Printf("Test_client_Get:%s\n",rs)
}

