package gobatis

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

type TUser struct {
	Id       int64      `field:"id"`
	Name     string     `field:"name"`
	Password NullString `field:"pwd"`
	Email    NullString `field:"email"`
	CrtTm    NullTime   `field:"crtTm"`
}

func TestGoBatis(t *testing.T) {
	ConfInit("")
	if nil == conf {
		log.Println("db config == nil")
		return
	}

	gb := NewGoBatis("datasource1")

	//result := make(map[string]interface{})
	//result := make([]interface{}, 0)
	//var result interface{}
	//result := make([]TUser, 0)
	var result *TUser
	err := gb.Select("userMapper.findById", map[string]interface{}{
		"id": 2,
	})(&result)

	fmt.Println("result:", result, "err:", err)

	u := &TUser{
		Name: "wenj1991",
		Password: NullString{
			String: "654321",
			Valid:  true,
		},
	}

	id, a, err := gb.Insert("userMapper.saveUser", u)
	fmt.Println("id:", id, "affected:", a, "err:", err)

	uu := &TUser{
		Id:   1,
		Name: "wenj1993",
		Password: NullString{
			String: "654321",
			Valid:  true,
		},
	}

	// test set
	affected, err := gb.Update("userMapper.updateByCond", uu)
	fmt.Println("updateByCond:", affected, err)

	param := &TUser{
		Name: "wenj1993",
	}

	// test where
	res := make([]*TUser, 0)
	err = gb.Select("userMapper.queryStructsByCond", param)(&res)
	fmt.Println("queryStructsByCond", res, err)

	// test trim
	res2 := make([]*TUser, 0)
	err = gb.Select("userMapper.queryStructsByCond2", param)(&res2)
	fmt.Println("queryStructsByCond", res2, err)

	affected, err = gb.Delete("userMapper.deleteById", map[string]interface{}{
		"id": 3,
	})
	fmt.Println("delete affected:", affected, "err:", err)
}
