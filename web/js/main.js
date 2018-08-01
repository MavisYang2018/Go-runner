

$(function (){

	$.ajax({
		type: 'GET',
		url: 'http://127.0.0.1:8080/counts',
		
		error: function(){
			console.log('error!!!')
		},
		
		
	});
});
