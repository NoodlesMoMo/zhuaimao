
function add_user() {
    console.log("xxxxxxxxx");

    let user_name = $("#user_name").val();
    let email = $("#email").val();

    if(user_name==="" || email==="") {
        return
    }

    console.log(user_name);
    console.log(email);

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

$(function () {
    /*
    $.ajaxSetup({
        contentType : 'application/json',
        processData : false
    });
    $.ajaxPrefilter( function( options, originalOptions, jqXHR ) {
        if (options.data){
            options.data=JSON.stringify(options.data);
        }
    });
    */
});

