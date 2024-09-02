package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// area 网段， 10.0.1.231 网段是 10.0
var area string

// vlan 根据 vlan id 设置
var vlan uint

// ip1 子网段，10.0.1.231 网段是 1
var ip1 uint

// ip2 地址，10.0.1.231 地址是 231
var ip2 uint

// num 次数，ping 命令执行 ICMP 连接的次数
var num uint

func main() {
	// 绑定命令行参数并赋值
	flag.StringVar(&area, "area", "10.0", "指定网段")
	flag.UintVar(&vlan, "vlan", 9, "指定 Vlan")
	flag.UintVar(&ip1, "ip1", 0, "子网段（10.0.2.* 则该值为 2）")
	flag.UintVar(&ip2, "ip2", 1, "地址")
	flag.UintVar(&num, "num", 4, "Ping 命令次数")

	// 解析命令行参数
	flag.Parse()
	// 拼接命令行参数
	//commandStr := fmt.Sprintf("ping  %v.%v.%v -c %v", area, ip1, ip2, num)
	commandStr := fmt.Sprintf("/usr/bin/ping -I vlan%v %v.%v.%v -c %v", vlan, area, ip1, ip2, num)
	fmt.Println("执行命令 => ", commandStr)

	// 执行命令
	cmd := exec.Command("sh", "-c", commandStr)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command Exec Error:", err)
	}
	fmt.Println("=====================")
	fmt.Println()
	fmt.Println("执行 ping 命令的输出结果：")
	fmt.Println(string(res))
	fmt.Println("=====================")
	// 根据正则匹配丢包率
	r := regexp.MustCompile(`(\d+)% packet loss`)
	result := r.FindStringSubmatch(string(res))
	fmt.Println("丢包率", result[1]+"%")

	var delayTimeSlice []string
	// 根据正则获取到每一行的 time 延时，并添加到 delayTimeSlice 切片
	re := regexp.MustCompile(`icmp_seq=\d+ ttl=\d+ time=(\d+\.\d+) ms`)
	lines := strings.Split(string(res), "\n")
	for _, line := range lines {
		if re.MatchString(line) {
			match := re.FindStringSubmatch(line)
			delayTimeSlice = append(delayTimeSlice, match[1])
		}
	}
	// Output:
	// [0.236 0.192 0.248 0.495]

	// 将 delayTimeSlice 切片中的数据转成 float64 类型，并相加
	var delayTimeFloatSlice []float64
	for _, str := range delayTimeSlice {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}
		delayTimeFloatSlice = append(delayTimeFloatSlice, f)
	}

	sum := 0.0
	for _, f := range delayTimeFloatSlice {
		sum += f
	}

	// 计算平均时长，根据 ping 命令获取的时间总和，除以执行 ping 命令的次数
	avgTime := sum / float64(num)
	fmt.Printf("平均时长为：%.3f \n", avgTime)
	if avgTime > 1000.0 {
		fmt.Println("当前 ping 命令执行 ICMP 连接时，平均时长超过 1000 毫秒。")
	} else {
		fmt.Println("当前 ping 命令执行 ICMP 连接时，平均时长小于 1000 毫秒。")
	}
}
