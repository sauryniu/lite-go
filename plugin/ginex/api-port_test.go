package ginex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type defaultAPI struct {
	Name string `validate:"max=6,min=1"`
}

func (m defaultAPI) Call() (interface{}, error) {
	return m.Name, nil
}

func Test_apiPort_Listen(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "endpoint"
		name := "default"
		req, _ := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("/%s/%s", endpoint, name),
			strings.NewReader(`{"name":"ok"}`),
		)
		req.Header.Add("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		self := make(apiPort, 0)

		mockAPIFactory := api.NewMockIFactory(ctrl)
		option := NewPostOption(
			mockAPIFactory,
			validator.New(),
		)
		self = append(self, option)
		self = append(
			self,
			NewServerHTTPOption(req, resp),
		)

		mockAPIFactory.EXPECT().Build(endpoint, name).Return(&defaultAPI{})

		self.Listen()

		assert.Equal(
			t,
			resp.Body.String(),
			`{"data":"ok","err":0}`,
		)
	})

	t.Run("verify", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "endpoint"
		name := "default"
		req, _ := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("/%s/%s", endpoint, name),
			strings.NewReader(`{"name":""}`),
		)
		req.Header.Add("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		self := make(apiPort, 0)

		mockAPIFactory := api.NewMockIFactory(ctrl)
		option := NewPostOption(
			mockAPIFactory,
			validator.New(),
		)
		self = append(self, option)
		self = append(
			self,
			NewServerHTTPOption(req, resp),
		)

		mockAPIFactory.EXPECT().Build(endpoint, name).Return(&defaultAPI{})

		self.Listen()

		assert.Equal(
			t,
			resp.Body.String(),
			fmt.Sprintf(`{"data":"","err":%d}`, errorex.VerifyCode),
		)
	})
}
