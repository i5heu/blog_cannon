function GoSocialAjaxSend(data){
      $("#LoadingIndicator").show();

    var xhr = new XMLHttpRequest();
    var url = "/gosocial";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var json = JSON.parse(xhr.responseText);
            $("#LoadingIndicator").hide();
            console.log(json.Status);
            GetComments();
            return true
            }
      };

        console.log(data);

        xhr.send(data);

  }


function GetComments(){
$("#LoadingIndicator").show();

var data = JSON.stringify({ "APP":"GetComments","Slug":window.location.pathname});

var xhr = new XMLHttpRequest();
var url = "/gosocial";
xhr.open("POST", url, true);
xhr.setRequestHeader("Content-type", "application/json");
xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
        var json = JSON.parse(xhr.responseText)
        $(json.Comments).each(function(index, item) {
          var foo ="<div class='GoSocial_Comment' id=GoSocial_GetCommentID_'" + item.Id + "' data-gosocialid='"+item.Id+"'>"
          //foo += "<div class='GoSocial_Votes'><span class='GoSocial_Upvotes'>"+ item.Upvotes +"</span>↑-↓<span class='GoSocial_Downvotes'>" + item.Downvotes+ "</span></div>"
          // THE VOTING IS UNDER DEV
          foo += "<div class='GoSocial_Title'>"+ item.Title +"</div>"
          foo += "<div class='GoSocial_Name'>by "+ item.Name +"</div>"
          foo += "<div class='GoSocial_Text'>"+ item.Text +"</div>"
          foo += "</div>"
          $('#GoSocial_Comments').append(foo)
        })
        return true
        }
  };

    console.log(data);

    xhr.send(data);
}

var SEND = false
$( document ).ready(function() {


GetComments();

$('#GoSocial_SubmitForm').on( 'click', '#GoSocial_SubmitForm_Send', function () {

    if(SEND == true){
      alert("ALREDY SENDED A COMMENT - PLS REALOAD THE PAGE");
      return;
    }

    SEND = true
    var data = JSON.stringify({ "APP":"WriteComment","Title":$("#GoSocial_SubmitForm_Title").val(),"Name":$("#GoSocial_SubmitForm_Name").val(),"Text":$("#GoSocial_SubmitForm_Text").val(),"Slug":window.location.pathname});
    $.when(GoSocialAjaxSend(data)).done(function() {
      $('#GoSocial_SubmitForm_Send').html("OK");
    });
  });




}); // DO NOT REMOVE DOC RDY
