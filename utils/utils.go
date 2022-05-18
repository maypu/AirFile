package utils

import (
	"errors"
	"fmt"
	"github.com/OrlovEvgeny/go-mcache"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"
)

var MCache = mcache.New()

func InitConfig() {
	//configs.yaml
	corrPath, err := os.Getwd() //获取项目的执行路径
	if err != nil {
		fmt.Println(err)
	}
	config := viper.New()
	config.AddConfigPath(corrPath)               //设置读取的文件路径
	config.SetConfigName("configs")          //设置读取的文件名
	config.SetConfigType("yaml")             //设置文件的类型
	err = config.ReadInConfig()					//尝试进行配置读取
	for _,value := range config.AllKeys() {
		err = MCache.Set("config_"+value, fmt.Sprintf("%v", config.Get(value)), mcache.TTL_FOREVER)
		if err != nil {
			panic(err)
		}
	}
}

func GetConfig(key string) string {
	if data, ok := MCache.Get("config_"+key); ok {
		return fmt.Sprintf("%v", data)
	}
	return ""
}

func Random(n int) string {
	//letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//letters := []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func GetLocalIP() []string {
	var ipStr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces error:", err.Error())
		return ipStr
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv6
					/*if ipnet.IP.To16() != nil {
					    fmt.Println(ipnet.IP.String())
					    ipStr = append(ipStr, ipnet.IP.String())
					}*/
					//获取IPv4
					if ipnet.IP.To4() != nil {
						//fmt.Println(ipnet.IP.String())
						ipStr = append(ipStr, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ipStr
}

//获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func ModelStructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Name == "Model" {
			continue
		}
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

//func StructTimeFormat(obj interface{}) interface{}  {
//	obj1 := reflect.TypeOf(obj)
//	obj2 := reflect.ValueOf(obj)
//
//	data := obj
//	for i := 0; i < obj1.NumField(); i++ {
//		_, fieldType := obj2.Field(i).Interface().(time.Time)
//		if  fieldType == true {
//			fmt.Println(obj2.Field(i).Interface().Format("2006-01-02 15:04:05"))
//		}
//	}
//	return data
//}
