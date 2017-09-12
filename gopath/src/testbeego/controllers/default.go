package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"

	bt, err := json.Marshal(c.Data)

	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(bt))

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(c.Data)

	c.Ctx.WriteString(string(bt))
}
