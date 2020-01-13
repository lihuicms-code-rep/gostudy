package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map:Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现
//Go语言中的map是引用类型，必须初始化才能使用(注意)

func basic71() {
	//使用
	scoreMap := make(map[string]int, 8)   //其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//判断某个键是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//遍历
	//使用delete()函数删除键值对
	delete(scoreMap, "lili")  //这里删除不存在的key不会引起崩溃

	//按照指定顺序遍历,做法就是对key排序然后遍历取出值
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap1 = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap1[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap1 {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap1[key])
	}

	//元素为map类型的切片(json数组)
	var mapSlice = make([]map[string]string, 3)
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "lllll"
	mapSlice[0]["age"] = "16"

	for index, v := range mapSlice {
		fmt.Printf("index = %d, v = %v\n", index, v)
	}

	//元素为切片的map
	var sliceMap = make(map[string][]string, 3)
	sliceMap["1"] = []string{"hello", "world"}
	fmt.Printf("sliceMap:%v\n", sliceMap)
}

func main() {
	basic71()
}