package api

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	url := "https://api.github.com/users/aierui/orgs"
	apiObj := NewGetApi(url)
	apiObj.AddProtocol("http")
	initRes := apiObj.InitRequest()
	if initRes == false {
		fmt.Println("initRes return false")
	}
	rtn := apiObj.GetResponse()

	fmt.Printf("rtn %v\n", rtn)

}
