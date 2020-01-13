package main

import (
	"fmt"
	"math"
	"strings"
	"unsafe"
)

//基本数据类型

//整型
func basic21() {
    //无符号
    var a uint8       //0~2^8-1, 0~255
    var b uint16      //0~2^16-1, 0~65535
    var c uint32      //0~2^32-1, 0~42 9496 7295
    var d uint64      //0~2^64-1, 0~1844 6744 0737 0955 1615

    //有符号
    var e int8        //-128~127
    var f int16       //-32768~32767
    var g int32       //-2147483648~2147483647
    var h int64       //-9223372036854775808~9223372036854775807

    //特殊整型,涉及到二进制传输,读写文件的结构描述时,不要使用uint和int
    var i uint        //uint和int取决于所在系统是32还是64
    var j int
    var k uintptr     //无符号整型,存放指针

    fmt.Println(a, b, c, d, e, f, g, h)
    fmt.Println("size of  i", unsafe.Sizeof(i), " size of j ", unsafe.Sizeof(j), " size of k", unsafe.Sizeof(k))
}

//数字字面量
func basic22() {
	//十进制
	var a int32 = 32
	fmt.Printf("十进制a=%d \n", a)
	fmt.Printf("二进制a=%b \n", a)
	fmt.Printf("八进制a=%o \n", a)
	fmt.Printf("十六进制a=%x \n", a)

	//八进制
	var b int32 = 0765
	fmt.Printf("十进制b=%d \n", b)
	fmt.Printf("二进制b=%b \n", b)
	fmt.Printf("八进制b=%o \n", b)
	fmt.Printf("十六进制b=%x \n", b)

	//十六进制
	var c int32 = 0xfe
	fmt.Printf("十进制c=%d \n", c)
	fmt.Printf("二进制c=%b \n", c)
	fmt.Printf("八进制c=%o \n", c)
	fmt.Printf("十六进制c=%x \n", c)
}

//浮点型
func basic23() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("float32 max %f\n", math.MaxFloat32)
	fmt.Printf("float64 max %f\n", math.MaxFloat64)
}

//复数
func basic24() {
	var c1 complex64 = 1 + 2i   //complex64的实部和虚部都是32位
	var c2 complex128 = 2 + 3i  //complex128的实部和虚部都是64位
	fmt.Println(c1, c2)
}

//布尔值

//字符串,go中字符串内部实现使用UTF-8编码
func basic25() {
	//打印window下一个文件路径:E:\GamePro\texas
	//转义字符的使用
	str := "\"E:\\GamePro\\texas\""
	fmt.Println(str)

	//如果要定义一个多行字符串,使用反引号``
	//里面的转义字符无效
	str1 := `我是第一行\n
             我是第二行\n
             我是第三行
            `
	fmt.Println(str1)

	//常用操作
	fmt.Println(len(str1))         //求长度
	str2 := fmt.Sprintf("%s:%d", str, 2)  //拼接
	fmt.Println(str2)

	str3 := "github.com/lihuicms-code-rep/go.study"
	strArr := strings.Split(str3, ".")  //分割
	fmt.Println(strArr)

	fmt.Println(strings.Contains(str3, "rep"))  //是否包含

	fmt.Println(strings.HasPrefix(str3, "github"))  //前缀判断
	fmt.Println(strings.HasSuffix(str3, "study"))   //后缀判断
	fmt.Println(strings.Index(str3, ".com"))        //子串出现的位置
	fmt.Println(strings.LastIndex(str3, "."))
	fmt.Println(strings.Join(strArr, "|"))            //join

}


//byte和rune类型
//字符串的元素叫做字符,Go中的字符有两类: uint8,也就是byte类型,代表一个ASCII字符,另一个是rune类型(int32),代表一个UTF-8字符
func basic26() {
  //遍历字符串
  s := "hello李辉"
  for i := 0; i < len(s); i++ {           //这样是按照byte输出
  	fmt.Printf("%v, %c", s[i], s[i])
  }
  fmt.Println()

  for _, r := range s {                   //在utf-8编码中一个中文汉字由3~4个字节来组成,所以用rune来读出是比较好的
  	fmt.Printf("%c ", r)           //这样是按照rune输出,汉语言文字使用这个比较方便
  }
  fmt.Println()

  //修改字符串,需要先将其转换成[]byte或者[]rune,改完后再转为string
  s1 := "abc"
  byteS1 := []byte(s1)      //字符串的底层数据结构是byte数组,可以转换
  byteS1[0] = 'd'
  s1 = string(byteS1)
  fmt.Println(s1)

  s2 := "红苹果"
  runeS2 := []rune(s2)
  runeS2[0] = '黄'
  s2 = string(runeS2)
  fmt.Println(s2)

}


//类型转换,Go语言中只有强类型转换,没有隐式转换

func main() {
	basic21()
	basic22()
	basic23()
	basic24()
	basic25()
	basic26()
}
