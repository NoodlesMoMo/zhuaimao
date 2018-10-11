
function add_role() {
    let role_name = $("#role_name").val();

    if(role_name === "") {
        toastr.error("role name is empty.");
        return
    }

    $.ajax({
        type: "PUT",
        url: "/role/",
        headers: {
            "Content-Type": "application/json; charset=utf-8"
        },
        data: JSON.stringify({
            role_name: role_name,
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

function append_role(role) {
    $("#role_table").append("<tr><td></td></tr>")
}
