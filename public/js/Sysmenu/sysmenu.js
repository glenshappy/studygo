// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

/**
 * 菜单管理
 */

//添加和修改操作
//
function update(menu_name,action) {
	if(action=="add"){
		 art.dialog.open('/EditAuthorize/', {
                id : 'sysmenu_add',
                title : menu_name,
                width : 700,
                height : 500,
                lock : true,
                ok : function() {
                        var iframe = this.iframe.contentWindow;

                        var pars = {};

                        var cn_name = iframe.$('#cn_name').val();
                        var en_name = iframe.$('#en_name').val();
                        pars.cn_name = cn_name;
                        pars.en_name = en_name;
                        pars.remark = iframe.$('#remark').val();
                        $.ajax({
                                type : "POST",
                                url : "/Authorize/Add/",
                                data : pars,
                                success : function(tmp) {
                                        if (tmp.status == 1) {
                                                window.location.reload();
                                                notice_tips("添加菜单成功!");
                                        } else {
                                                notice_tips(tmp.message);
                                        }
                                }
                        });
                },
                cancel : true
        });

        }
}

/**
 * 用户详细
 */
function user_info(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/User/userinfo/' + uid + '/', {
		id : 'user_info',
		title : '个人信息',
		width : 700,
		height : 500,
		lock : true,
		ok : function() {},
		cancel : true
	});
}

//锁定
function lock() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要锁定用户吗?', function() {
		$.post("/User/lock/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("锁定成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了锁定用户操作!");
	});
}

//解锁
function unlock() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要解锁用户吗?', function() {
		$.post("/User/unlock/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("解锁成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了解锁用户操作!");
	});
}

function show_move() {

	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	if ($("#show_move").css("display") == "none") {
		$("#show_move").fadeIn("slow");
	}else{
		$("#show_move").fadeOut("slow");
	}
}

//移动
function move() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	var groupid = $('#groupid').val();

	art.dialog.confirm('确定要移动用户吗?', function() {
		$.post("/User/move/",{'ids':ids,'groupid':groupid,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("移动成功!");
			} else {
				notice_tips(tmp.message);
			}
		});
	}, function() {
		notice_tips("取消了移动用户操作!");
	});
}

/**
 * 删除用户
 */
function confirm_delete() {

	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要删除用户吗?', function() {
		$.post("/User/delete/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("删除成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了删除用户操作!");
	});
}

//搜索
function Search() {
	url = '/User/';

	search = "";

	search += 'islock:' + $('#islock').val() + '|';
	search += 'type:' + $('#type').val() + '|';
	search += 'keyword:' + $('#keyword').val() + '|';
	search += 'start_time:' + $('#start_time').val() + '|';
	search += 'end_time:' + $('#end_time').val() + '';

	$.base64.encode(search)

	url += $.base64.encode(search)+'/1/';

	redirect(url);
};

//重置
function Reset() {
	redirect("/Authorize/");
}
