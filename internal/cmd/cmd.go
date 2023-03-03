package cmd

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"gongzhaoweishop/api/backend"
	"gongzhaoweishop/internal/consts"
	"gongzhaoweishop/internal/controller"
	"gongzhaoweishop/internal/dao"
	"gongzhaoweishop/internal/model/entity"
	"gongzhaoweishop/internal/service"
	"gongzhaoweishop/utility"
	"gongzhaoweishop/utility/response"
	"strconv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				CacheMode:        2,
				ServerName:       "shop",
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  LoginFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/user/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true,
			}
			//todo 抽取方法
			err = gfAdminToken.Start()
			if err != nil {
				return err
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//gtoken中间件绑定
				//err := gfAdminToken.Middleware(ctx, group)
				//if err != nil {
				//	panic(err)
				//}
				group.Bind(
					//controller.Hello,        //示例
					controller.Rotation,     // 轮播图
					controller.Position,     // 手工位
					controller.Admin.Create, // 管理员
					controller.Admin.Update, // 管理员
					controller.Admin.Delete, // 管理员
					controller.Admin.List,   // 管理员
					controller.Login,        // 登录
					//controller.Data,         // 数据大屏相关
					controller.Role, // 角色
					controller.File,
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) //for jwt
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
						"/hello":              controller.Hello,
					})
					//group.Middleware(service.Middleware().GTokenSetCtx, ) //for gtoken
					//todo 优化代码 返回的数据格式和之前的一致
					//group.ALL("/backend/admin/info", func(r *ghttp.Request) {
					//	r.Response.WriteJson(gfAdminToken.GetTokenData(r).Data)
					//})
				})
			})
			s.Run()
			return nil
		},
	}
)

func LoginFunc(r *ghttp.Request) (string, interface{}) {

	name := r.Get("name").String()
	password := r.Get("password").String()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或者密码错误"))
		r.ExitAll()
	}
	//return name, "1"

	ctx := context.TODO()

	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		return "", nil
	}
	//gutil.Dump("加密后密码：", utility.EncryptPassword(name, adminInfo.UserSalt))
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		return "", nil
	}
	//return "admin:" + gconv.String(adminInfo.Id), "1"
	//因为我们是前后台一体的项目，前台项目的user和后台项目的admin的id一定有重合，所以要加前缀区分
	//为什么用冒号分隔？因为商业项目要把token保存到redis中，:分隔 数据可视化优化
	//唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// todo 迁移到合适的位置
// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		g.Dump("adminId:", adminId)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    10 * 24 * 60 * 60, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump(respData)
	g.Dump(respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	//g.Dump("adminfo:", adminInfo)
	fmt.Println(adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	//if adminInfo.DeletedAt != nil {
	//	response.AuthBlack(r)
	//	return
	//}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
