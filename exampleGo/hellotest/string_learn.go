package main //每一个go文件都必须有一个package package后面是名称 
/*用来练习go的string
*/
import(//一次性导入用的包名 用圆括号包裹 每个包名占一行“
"fmt"
"strings"
)

func main(){//主方法 程序的入口方法
one_str := "TEST ONE ,I don't know why"
//输出一行
fmt.Println(one_str)
//转换为小写字符
lower_str := strings.ToLower(one_str)
fmt.Println(lower_str)
//是否包含某些字符
iscontains := strings.Contains(lower_str,"now")
fmt.Println(iscontains)
//获取字符串的某一段文字
substr := lower_str[3:8]
fmt.Println("print char 4-8:"+substr)
fmt.Println("lower_str_first_five:"+lower_str[0:5])
fmt.Println("lower_str_first_six:"+lower_str[0:6])
//分割字符串
words := strings.Split(lower_str," ")
fmt.Printf("%v \n“",words)
//使用独立的函数分割字符串
fields := strings.Fields(lower_str)
fmt.Printf("%v \n",fields)
}
