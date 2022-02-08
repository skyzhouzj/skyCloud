package signature

import (
	"net/url"
	"testing"
	"time"
)

const (
	key    = "admin"
	secret = "12878dd962115106db6d"
	ttl    = time.Minute * 10
)

func TestSignature_Generate(t *testing.T) {
	path := "/p/c/Captcha"
	method := "GET"

	params := url.Values{}
	//params.Add("a", "a1")
	//params.Add("d", "d1")
	//params.Add("c", "c1 c2")

	authorization, date, err := New(key, secret, ttl).Generate(path, method, params)
	t.Log("authorization:", authorization)
	t.Log("date:", date)
	t.Log("err:", err)
}

func TestSignature_Verify(t *testing.T) {

	authorization := "blog y7a326f3aWvIxdeNIgRo0P7FSDnCNSsN8gJi/4y+cZo="
	date := "2021-04-06 16:15:26"

	path := "/p/c/Captcha"
	method := "GET"
	params := url.Values{}
	//params.Add("a", "a1")
	//params.Add("d", "d1")
	//params.Add("c", "c1 c2*")

	ok, err := New(key, secret, ttl).Verify(authorization, date, path, method, params)
	t.Log(ok)
	t.Log(err)
}
