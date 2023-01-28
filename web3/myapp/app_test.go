package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler()) //새로운 서버를 만드는 거다. 서버에서는 새로운 핸들러를 집어넣어야 한다.
	//근데 실제 http서버가 아니고, test 하기위한 서버를 생성한것. 그리고 항상 닫아줘야함
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler()) //새로운 서버를 만드는 거다. 서버에서는 새로운 핸들러를 집어넣어야 한다.
	//근데 실제 http서버가 아니고, test 하기위한 서버를 생성한것. 그리고 항상 닫아줘야함
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User ID:89")
}
