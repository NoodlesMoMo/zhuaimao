
function initUserTable() {
    $('#user_table').bootstrapTable({
        url: '/user/',
        method: 'GET',                      //请求方式（*）
        //toolbar: '#toolbar',              //工具按钮用哪个容器
        striped: true,                      //是否显示行间隔色
        cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
        pagination: true,                   //是否显示分页（*）
        sortable: true,                     //是否启用排序
        sortOrder: "asc",                   //排序方式
        sidePagination: "server",           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber: 1,                      //初始化加载第一页，默认第一页,并记录
        pageSize: 5,                     //每页的记录行数（*）
        pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
        search: false,                      //是否显示表格搜索
        strictSearch: true,
        showColumns: true,                  //是否显示所有的列（选择显示的列）
        showRefresh: true,                  //是否显示刷新按钮
        minimumCountColumns: 2,             //最少允许的列数
        clickToSelect: true,                //是否启用点击选中行
        //height: 500,                      //行高，如果没有设置height属性，表格自动根据记录条数觉得表格高度
        uniqueId: "ID",                     //每一行的唯一标识，一般为主键列
        showToggle: true,                   //是否显示详细视图和列表视图的切换按钮
        cardView: false,                    //是否显示详细视图
        detailView: false,                  //是否显示父子表
        columns: [{
            field: 'id',
            title: 'ID'
        }, {
            field: 'name',
            title: 'Name'
        }, {
            field: 'email',
            title: 'Email'
        }, {
            field: 'role',
            title: 'Role'
        }]
    });
}

function add_user() {
    var user_name = $("#user_name").val();
    var email = $("#email").val();

    if(user_name === "" || email === ""){
        // TODO: add alert
        return
    }

    $.ajax({
        type: "PUT",
        url: "/user/",
        data: {
            user_name: user_name,
            email: email,
        },
        dataType: "json",
        success: function (data) {
            console.log(data)
        }
    })
}

$(function () {
    //modalInit();
    initUserTable();
});
