package myCasbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func CasBin() {
	e, err := casbin.NewEnforcer("myCasbin/model.conf", "myCasbin/policy.csv")
	//if err != nil {
	//	fmt.Println("================")
	//	fmt.Println("%s", err)
	//}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	added, err := e.AddPolicy("alice", "data1", "read")
	fmt.Println(added, err)
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		fmt.Println("===============")
		fmt.Println("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("casBin ok")
	} else {
		// deny the request, show an error
		fmt.Println("casBin not ok")
	}
}
func CasBinByGorm() {
	a, _ := gormadapter.NewAdapter("mysql", "root:celeste0922@tcp(127.0.0.1:3306)/casbin", true)
	e, _ := casbin.NewEnforcer("myCasbin/model.conf", a)
	e.AddFunction("my_func", KeyMatchFunc) //注册自定义方法
	sub := "alice"                         // the user that wants to access a resource.
	obj := "data1"                         // the resource that is going to be accessed.
	act := "read"                          // the operation that the user performs on the resource.
	//added, err := e.AddPolicy("alice", "data4", "read") //添加p规则到数据库
	//fmt.Println(added, err)
	//filteredPolicy := e.GetFilteredPolicy(0, "alice") //查询第0项为“alice"的规则
	//fmt.Println(filteredPolicy)
	//updated, err := e.UpdatePolicy([]string{"alice", "data4", "read"}, []string{"alice", "data2", "read"})//更新
	//fmt.Println(updated, err)
	//added, err := e.AddGroupingPolicy("alice", "admin")//添加g规则
	//fmt.Println(added, err)
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// handle err
		fmt.Println("===============")
		fmt.Println("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("casBin ok")
	} else {
		// deny the request, show an error
		fmt.Println("casBin not ok")
	}
}

// 自定义方法
func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
