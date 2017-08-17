$( document ).ready(function() {
  console.log("START");

  if (!!$.cookie('pwd')) {
   // have cookie
   $("#ApiContainer").html(`<h1>Create New Article</h1><input id="title">title</input><br><input id="category">category</input><br><input id="tags">tags</input><br><textarea id="text">text</textarea><br><button id="articlesend">SEND</button>`);
  } else {
   // no cookie
   PwdInput();
  }

  $('#ApiContainer').on( 'click', '#PwSend', function () {
    console.log("PwdSend");
    PWSave();
    location.reload();
  });

  $('#ApiContainer').on( 'click', '#articlesend', function () {
    console.log("ArticleSend");
    ArticleSend();
  });

}); // DO NOT REMOVE DOC RDY

expireAt = new Date;
expireAt.setMonth(expireAt.getMonth() + 1);


function PwdInput(){
  $("#ApiContainer").html(`<input type="password" id="pw">PW</input><button id="PwSend">Send</button>`);
}

function PWSave(){
  var foo = $('#pw').val();
  pwdhash = CryptoJS.SHA256(foo).toString(CryptoJS.enc.Hex);

  document.cookie = "pwd=" + pwdhash + ";expires=" + expireAt.toGMTString() + ";path=/"
}

function ArticleSend(){
  data = '{"PWD":"'+  $.cookie("pwd") + `", "Title":"` + $('#title').val() + `", "Category":"` + $('#category').val() + `", "Tags":"` + $('#tags').val() + `", "Text":"` + $('#text').val() + '"}';

  $.ajax({
              type:"POST",
              url: "/api",
              data:data,
              success: function (response){
                    $(".flexparent").append('<div class="SearchFlexChild" id="TextSearch"></div>')
                    var json = $.parseJSON(response);
                      $("#ApiContainer").html(json.Status);
                    console.log("ArticleSend");

                  }
        });


}
