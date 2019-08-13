package controller

import (
	"fmt"
	// "bytes"
	"net/http"
	"testing"
	// "encoding/json"
	"github.com/stretchr/testify/assert"
	httpEntity "example_app/entity/http"
)

type UserServiceMock struct{}

func (service *UserServiceMock) GetUserByID(id int) httpEntity.UserDetailResponse {
	// var status string = "Status Test"
	// statusPointer := &status
	return httpEntity.UserDetailResponse{
		ID: uint(id),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
		Status: nil,
		Avatar:nil,
	}
}

func (service *UserServiceMock) GetAllUser(page int,count int) []httpEntity.UserResponse{
	users := []httpEntity.UserResponse{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID1),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID2),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	users = append(users, httpEntity.UserResponse{
		ID: uint(ID3),
		Name: "Test Name Mock",
		IDCardNumber: "IDCARDTEST12345",
		Address: "Address Test Street",
	})
	return users
}

func (service *UserServiceMock) UpdateUserByID(id int) bool{
	return true
}

func TestUserGetByID(t *testing.T) {
	assert := assert.New(t)
	c, r, resp := LoadRouterMock()

	var idTest uint = 1
	url := "/v1/users" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)

	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")
	// doing test result from router
	// res := objects.V1UserObjectResponse{}
	// err := json.Unmarshal([]byte(resp.Body.String()), &res)
	// assert.Equal(err, nil, "should have no error")
	// assert.Equal(idTest, res.ID, "It should be same ID")
	// assert.Equal("testing name controller mock", res.Name, "It should be same Name")
	// assert.Equal("testing_controller_mock@mail.com", res.Email, "It should be same Email")
}