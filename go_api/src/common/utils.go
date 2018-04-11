package common

import (
	"github.com/go-gomail/gomail"
	"github.com/satori/go.uuid"
	. "core"
	"encoding/json"
)

func GetIntArray(i interface{}) (array []int) {

	var original []interface{}
	var ok bool

	original, ok = i.([]interface{})

	if !ok {
		return []int{}
	}

	var result []int

	for _, val := range original {
		result = append(result, val.(int))
	}

	return result
}

func GetStringArray(i interface{}) (array []string) {

	var original []interface{}
	var ok bool

	original, ok = i.([]interface{})

	if !ok {
		return []string{}
	}

	var result []string

	for _, val := range original {
		result = append(result, val.(string))
	}

	return result
}

func GetStringValue(i interface{}) string {
	if i == nil{
		return ""
	}

	return i.(string)
}

func GetFloatValue(i interface{}) float64 {
	if i == nil{
		return 0
	}

	return i.(float64)
}

func GetBoolValue(i interface{}) bool {
	if i == nil{
		return false
	}

	return i.(bool)
}

func SetStructEntity(requireHandle RequireHandle, i interface{}){
	json.Unmarshal(requireHandle.Body, i)
}

func GetUUID() string {
	return uuid.NewV4().String()
}

const (
	USERNAME = "519955464@qq.com"
	PASSWORD = "dokkpjlccjlzbicb"
	NAME = "王佩剑"
	HOST = "smtp.qq.com"
	PORT = 465
)

func SendToMail(to, subject, body string, htmlMail bool) error {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", USERNAME, NAME)  // 发件人
	m.SetHeader("To",  // 收件人
		m.FormatAddress(to, ""),
	)
	m.SetHeader("Subject", subject)  // 主题

	if htmlMail {
		m.SetBody("text/html", body) // 正文
	} else {
		m.SetBody("text/plain", body) // 正文
	}

	d := gomail.NewDialer(HOST, PORT, USERNAME, PASSWORD)  // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}else{
		return err
	}
}
