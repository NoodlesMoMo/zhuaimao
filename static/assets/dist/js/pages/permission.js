
function init_methods_selection() {
    $("#perm_methods").select2();
}

function add_permission() {
    let perm_name = $("#perm_name").val();
    let perm_slug = $("#perm_slug").val();
    let perm_methods = $("#perm_methods").val();
    let perm_path = $("#perm_path").val();

    if(perm_name === "" || perm_slug === "" || perm_path==="" || perm_methods==="") {
        toastr.error("invalid param");
        return
    }

    $.ajax({
        type: "PUT",
        url: "/permission/",
        headers: {
            "Content-Type": "application/json; charset=utf-8"
        },
        data: JSON.stringify({
            name: perm_name,
            slug: perm_slug,
            methods: perm_methods,
            path: perm_path
        }),
        dataType: "json",
        success: function (data) {
            if(data["code"] !== 0) {
                toastr.error(data["msg"])
            } else {
                toastr.success("success!")
            }
        },
        error: function(data) {
            if(data) {
                toastr.error(data);
            }
        }
    })
}

$(function () {
    init_methods_selection()
});

