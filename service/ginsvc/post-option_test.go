package ginsvc

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewPostOption(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode)
	app := gin.New()

	relativePath := "/mobile/api"
	mockApi := contract.NewMockIApi(ctrl)
	NewPostOption(relativePath, func(_ *gin.Context) (contract.IApi, error) {
		return mockApi, nil
	})(app)

	mockApi.EXPECT().Call().Return("ok", nil)

	req := httptest.NewRequest(
		"POST",
		relativePath,
		strings.NewReader(``),
	)
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)

	res, err := ioutil.ReadAll(
		resp.Result().Body,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		string(res),
		`{"data":"ok","err":0}`,
	)
}
