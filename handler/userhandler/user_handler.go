package userhandler

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/wayuncle/module-a-user/model/usermodel"
	Usertype "github.com/wayuncle/module-a-user/type/usertype"
)

// User 用户
type User struct {
	Version string
}

// Index
// @Description: 首页
// @Author: Zhenwei Huo
// @Date: 2022-03-31 15:58:42
// @Param r *ghttp.Request
func (u *User) Index(r *ghttp.Request)  {
	usermodel.Index(r)
}

// Create
// @Description: 新增用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 16:35:38
// @Param r *ghttp.Request
func (u *User) Create(user *Usertype.AddReq,r *ghttp.Request) {
	/*var user *Usertype.AddReq
	if err := r.Parse(&user); err != nil {
		fmt.Println("err",err)
	}*/
	usermodel.Create(user)
}

// Update
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 16:36:15
// @Param r *ghttp.Request
func (u *User) Update(r *ghttp.Request) {
	var user *Usertype.UpdateReq
	if err := r.Parse(&user); err != nil {
		fmt.Println("err",err)
	}
	usermodel.Save(user)
}

// QueryById
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-04-01 11:00:09
// @Param r *ghttp.Request
func (u *User) QueryById(r *ghttp.Request) {
	id := r.GetInt("id")
	usermodel.Query(id)
}

// Delete
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-04-01 11:00:46
// @Param r *ghttp.Request
func (u *User) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	usermodel.Delete(id)
}
