package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type BizParam struct {
	ID_NO     string `json:"id_no"`
	ProductID int    `json:"product_id" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

func main() {
	jsonStr := `{
                "id_no": "1001607",
                "phone": "1001607",
                "product_id":2
            }`

	ret := &BizParam{}
	err := jsoniter.Unmarshal([]byte(jsonStr), ret)
	fmt.Printf("err %v\n", err)
	fmt.Println(ret)

	jsonStr2 := `{
                "id_no": "1001607",
                "phone": "1001607"
            }`

	ret2 := &BizParam{}
	err2 := jsoniter.Unmarshal([]byte(jsonStr2), ret2)
	fmt.Printf("err %v\n", err2)
	fmt.Println(ret2)

	m1 := make(map[string]interface{}, 0)

	//m1["product_id"] = 1
	m1["id_no"] = "test"
	m1["phone"] = "1234567"
	m1["product_id"] = 1 // reset

	rtn3 := &BizParam{}
	err3 := convertInterfaceToStruct(m1, rtn3)
	fmt.Printf("err3: %v\n", err3)
	fmt.Printf("%#v\n", rtn3)

}

func convertInterfaceToStruct(p interface{}, v interface{}) error {
	s, err := jsoniter.Marshal(p)
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(s, v)
	if err != nil {
		return err
	}
	return nil
}
