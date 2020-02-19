//this is where the CSRF code will be loaded and called.

function LoadCSRF() {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
     document.getElementById("csrfTag").innerHTML = this.responseText;
    }
  };
  xhttp.open("GET", "/api/csrf", true);
  xhttp.send();
}
