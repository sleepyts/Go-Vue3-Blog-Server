package vo

import (
	"Go-Vue3-Blog-Server/utils/redis_util"
	"fmt"
	"strconv"
)

type Verify struct {
	Key       string `json:"key"`
	Var1      string `json:"var1"`
	Var2      string `json:"var2"`
	VerifyVar string `json:"verifyVar,omitempty"`
}

func (v *Verify) Verify() string {
	var realVerify Verify
	err := redis_util.GetObject(v.Key, &realVerify)
	if err != nil {
		return "验证码已过期,请重新获取"
	}
	var1Int, err1 := strconv.Atoi(realVerify.Var1)
	if err1 != nil {
		fmt.Println("Var1 转换失败:", err1)
		return "验证码格式错误"
	}

	var2Int, err2 := strconv.Atoi(realVerify.Var2)
	if err2 != nil {
		fmt.Println("Var2 转换失败:", err2)
		return "验证码格式错误"
	}

	if v.VerifyVar == strconv.Itoa(var1Int*var2Int) {
		return "ok"
	}
	return "验证码错误"
}
