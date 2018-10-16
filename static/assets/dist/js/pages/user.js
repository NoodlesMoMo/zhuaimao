
function init_group_selection() {
    let roles_selection = $("#roles");
    if(roles_selection != null){
        roles_selection.select2({
            ajax: {
                type: "GET",
                url: "/role/all/",
                dataType: 'json',
                delay: 250,
                data: function (params) {
                    /*
                    return {
                        q: params.term,
                    };
                    */
                },
                processResults: function (data) {
                    console.log(data);

                    if(!data || data["code"] !== 0){
                        return null
                    }
                    roles = data["data"];
                    let items = new Array();
                    for(let i=0; i<roles.length; i++) {
                        let role = roles[i];
                        items[i] = {id: role["ID"], text: role["Name"]}
                    }

                    console.log(items);

                    return {
                        results: items
                    };
                },
                cache: true
            },
            escapeMarkup: function (markup) { return markup; },
            minimumInputLength: 1,
        });
    }
}

function add_user() {
    let user_name = $("#user_name").val();
    let email = $("#email").val();

    if(user_name==="" || email==="") {
        return
    }

    $.ajax({
        type: "PUT",
        url: "/user/user/",
        headers: {
            "Content-Type": "application/json; charset=utf-8"
        },
        data: JSON.stringify({
            user_name: user_name,
            email: email
        }),
        dataType: "json",
        success: function (data) {
            console.log("success")
        }
    })
}

function add_group() {
    let name = $("#name").val();
    let roles = $("#roles").val();
    console.log(roles)
    if(name === "") {
        toastr.error("invalid param.");
        return
    }

    $.ajax({
        type: "PUT",
        url: "/user/group/",
        headers: {"Content-Type": "application/json; charset=utf8"},
        data: JSON.stringify({
            name: name,
            roles: roles
        }),
        dataType: "json",
        success: function (data) {
            if(data["code"] === 0) {
                toastr.success("success")
            } else {
                toastr.error(data["msg"])
            }
        },
        error: function (data) {
            toastr.error("critical: " + data)
        }
    })
}

$(function () {
    init_group_selection()
});
