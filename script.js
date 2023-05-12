$('#search-button').on('click', function(){
    $.ajax({
        url: 'localhost:8080',
        type: 'get',
        datatype: 'json',
        data: {
            's': $('#search-input').val()
        },
        success: function(result){
            if(result.Response == "True"){
                console.log("hai");
            }
            else{
                $('product-list').html('<h1>'+ result.Error +'<h1>');
            }
        }
    })
})