package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/foo", fooHandle)
	http.HandleFunc("/api/foo/post", fooPostHandle)

	fmt.Println("服务已经启动成功，端口监听在8082...")
	http.ListenAndServe(":8082", nil)
}

type Foo struct {
	Bar string `json:"bar"`
}

func fooPostHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var foo Foo
	err := json.NewDecoder(r.Body).Decode(&foo)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err.Error())
		return
	}

	// buf, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// r.Body.Close()

	// var foo Foo
	// err = json.Unmarshal(buf, &foo)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Println(err.Error())
	// 	return
	// }

	fmt.Println("Foo:", foo)
	foo.Bar = foo.Bar + "--ok"

	json.NewEncoder(w).Encode(foo)

	// buf, _ = json.Marshal(foo)

	// w.WriteHeader(200)
	// w.Write(buf)
}

func fooHandle(w http.ResponseWriter, r *http.Request) {
	// TODO: 处理Foo的请求及响应

	fmt.Printf("RequestURI:%s,Method:%s", r.RequestURI, r.Method)

	m := map[string]string{
		"foo": "bar",
	}
	buf, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
