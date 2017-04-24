// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tAuthorize struct {}
var Authorize tAuthorize


func (_ tAuthorize) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Authorize.Index", args).URL
}

func (_ tAuthorize) EditAuthorize(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Authorize.EditAuthorize", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Main(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Main", args).URL
}


type tFocus struct {}
var Focus tFocus


func (_ tFocus) Index(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Index", args).URL
}

func (_ tFocus) Add(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Add", args).URL
}

func (_ tFocus) Edit(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Edit", args).URL
}


type tCategory struct {}
var Category tCategory


func (_ tCategory) Index(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Index", args).URL
}

func (_ tCategory) Add(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Add", args).URL
}

func (_ tCategory) Edit(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Edit", args).URL
}

func (_ tCategory) Delete(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Delete", args).URL
}


type tFocusCate struct {}
var FocusCate tFocusCate


func (_ tFocusCate) Index(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Index", args).URL
}

func (_ tFocusCate) Add(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Add", args).URL
}

func (_ tFocusCate) Edit(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Edit", args).URL
}

func (_ tFocusCate) Delete(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Delete", args).URL
}


type tContent struct {}
var Content tContent


func (_ tContent) Index(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Index", args).URL
}

func (_ tContent) Left(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Left", args).URL
}

func (_ tContent) List(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.List", args).URL
}

func (_ tContent) Keywords(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Keywords", args).URL
}

func (_ tContent) Delete(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Delete", args).URL
}

func (_ tContent) CateNameSearch(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Content.CateNameSearch", args).URL
}

func (_ tContent) Push(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Push", args).URL
}

func (_ tContent) Remove(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Remove", args).URL
}

func (_ tContent) Comment(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Comment", args).URL
}

func (_ tContent) Relationlist(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Relationlist", args).URL
}

func (_ tContent) AddContent(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.AddContent", args).URL
}

func (_ tContent) Add(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Add", args).URL
}

func (_ tContent) Edit(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Edit", args).URL
}


type tExtend struct {}
var Extend tExtend


func (_ tExtend) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Extend.Index", args).URL
}


type tCopyfrom struct {}
var Copyfrom tCopyfrom


func (_ tCopyfrom) Index(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Index", args).URL
}

func (_ tCopyfrom) Add(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Add", args).URL
}

func (_ tCopyfrom) Edit(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Edit", args).URL
}

func (_ tCopyfrom) Delete(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Delete", args).URL
}


type tModule struct {}
var Module tModule


func (_ tModule) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Module.Index", args).URL
}


type tAnnounce struct {}
var Announce tAnnounce


func (_ tAnnounce) Index(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Index", args).URL
}

func (_ tAnnounce) Add(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Add", args).URL
}

func (_ tAnnounce) Edit(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Edit", args).URL
}

func (_ tAnnounce) Delete(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Delete", args).URL
}


type tComplaints struct {}
var Complaints tComplaints


func (_ tComplaints) Index(
		complaints interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "complaints", complaints)
	return revel.MainRouter.Reverse("Complaints.Index", args).URL
}

func (_ tComplaints) Delete(
		complaints interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "complaints", complaints)
	return revel.MainRouter.Reverse("Complaints.Delete", args).URL
}


type tPanel struct {}
var Panel tPanel


func (_ tPanel) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Panel.Index", args).URL
}


type tPlugin struct {}
var Plugin tPlugin


func (_ tPlugin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Plugin.Index", args).URL
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Index", args).URL
}

func (_ tCaptcha) GetCaptchaId(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.GetCaptchaId", args).URL
}


type tKindeditor struct {}
var Kindeditor tKindeditor


func (_ tKindeditor) Manager(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.Manager", args).URL
}

func (_ tKindeditor) TitleImage(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.TitleImage", args).URL
}

func (_ tKindeditor) Upload(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.Upload", args).URL
}

func (_ tKindeditor) AnnounceImage(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.AnnounceImage", args).URL
}


type tTest struct {}
var Test tTest


func (_ tTest) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Test.Index", args).URL
}


type tPublic struct {}
var Public tPublic


func (_ tPublic) Map(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Public.Map", args).URL
}

func (_ tPublic) CreateHtml(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.CreateHtml", args).URL
}

func (_ tPublic) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Search", args).URL
}

func (_ tPublic) Message(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Message", args).URL
}


type tAjax struct {}
var Ajax tAjax


func (_ tAjax) GetCaptcha(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetCaptcha", args).URL
}

func (_ tAjax) Pos(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Ajax.Pos", args).URL
}

func (_ tAjax) GetPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.GetPanel", args).URL
}

func (_ tAjax) DelPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.DelPanel", args).URL
}

func (_ tAjax) AddPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.AddPanel", args).URL
}

func (_ tAjax) GetMessage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetMessage", args).URL
}

func (_ tAjax) ScreenLock(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.ScreenLock", args).URL
}

func (_ tAjax) ScreenUnlock(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Ajax.ScreenUnlock", args).URL
}


type tMenu struct {}
var Menu tMenu


func (_ tMenu) Index(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Index", args).URL
}

func (_ tMenu) Add(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Add", args).URL
}

func (_ tMenu) Edit(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Edit", args).URL
}

func (_ tMenu) Delete(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Delete", args).URL
}


type tSetting struct {}
var Setting tSetting


func (_ tSetting) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Setting.Index", args).URL
}


type tRole struct {}
var Role tRole


func (_ tRole) Index(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Index", args).URL
}

func (_ tRole) Member(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Member", args).URL
}

func (_ tRole) Add(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Add", args).URL
}

func (_ tRole) Edit(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Edit", args).URL
}

func (_ tRole) SetStatus(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.SetStatus", args).URL
}

func (_ tRole) Delete(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Delete", args).URL
}


type tTask struct {}
var Task tTask


func (_ tTask) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Task.Index", args).URL
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Index", args).URL
}

func (_ tAdmin) Add(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Add", args).URL
}

func (_ tAdmin) Edit(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Edit", args).URL
}

func (_ tAdmin) Delete(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Delete", args).URL
}


type tLogs struct {}
var Logs tLogs


func (_ tLogs) Index(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.Index", args).URL
}

func (_ tLogs) DelAll(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.DelAll", args).URL
}


type tStyle struct {}
var Style tStyle


func (_ tStyle) Index(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Index", args).URL
}

func (_ tStyle) File(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.File", args).URL
}

func (_ tStyle) Import(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Import", args).URL
}

func (_ tStyle) Setstatus(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Setstatus", args).URL
}

func (_ tStyle) Edit(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Edit", args).URL
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Index", args).URL
}

func (_ tUser) Add(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Add", args).URL
}

func (_ tUser) Edit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Edit", args).URL
}

func (_ tUser) Delete(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Delete", args).URL
}

func (_ tUser) Lock(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Lock", args).URL
}

func (_ tUser) Unlock(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Unlock", args).URL
}

func (_ tUser) Move(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Move", args).URL
}

func (_ tUser) UserInfo(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.UserInfo", args).URL
}

func (_ tUser) Login(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Login", args).URL
}

func (_ tUser) Logout(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Logout", args).URL
}

func (_ tUser) EditInfo(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditInfo", args).URL
}

func (_ tUser) AdminPanel(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.AdminPanel", args).URL
}

func (_ tUser) EditPwd(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditPwd", args).URL
}

func (_ tUser) Left(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("User.Left", args).URL
}


type tGroup struct {}
var Group tGroup


func (_ tGroup) Index(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Index", args).URL
}

func (_ tGroup) Add(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Add", args).URL
}

func (_ tGroup) Edit(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Edit", args).URL
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}


type tPprof struct {}
var Pprof tPprof


func (_ tPprof) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Profile", args).URL
}

func (_ tPprof) Symbol(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Symbol", args).URL
}

func (_ tPprof) Cmdline(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Cmdline", args).URL
}

func (_ tPprof) Trace(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Trace", args).URL
}

func (_ tPprof) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Index", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


