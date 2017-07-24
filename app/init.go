// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package app

import "github.com/revel/revel"
import (
	"time"
	"fmt"
)
func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	//a 需要格式化的时间戳，b为格式成的格式
	revel.TemplateFuncs["format_time"] = func(a,b interface{}) string {
		tmpTimestamp,err:=a.(int64);
		if err!=false {
			fmt.Println("error %v",err);
		}
		tmpFormat,err1:=b.(string)
		if err1!=false {
			fmt.Println("error %v",err1);
		}
		tm:=time.Unix(tmpTimestamp,0);
		return tm.Format(tmpFormat);
	};
	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB())
	// revel.OnAppStart(FillCache())
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Content-Type", "text/html; charset=utf-8")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
