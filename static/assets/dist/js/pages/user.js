
function add_user() {
    let user_name = $("#user_name").val();
    let email = $("#email").val();

    if(user_name==="" || email==="") {
        return
    }

    $.ajax({
        type: "PUT",
        url: "/user/",
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

