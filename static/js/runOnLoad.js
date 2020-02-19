// This is a momument to linux, the monolith that murdered minix
// GomMVC is gunning for all that JS running on the backend,
//  WHY WOULD YOU EVER MAKE IT SO SIMPLE TO PROGRAM A SERVER? DANGEROUS!
// everything is still better than PHP. reddit.com/r/lolPHP

// LoadCSRF pulls the json result of /api/csrf
function LoadCSRF() {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
     document.getElementById("csrfTag").innerHTML = this.responseText;
    }
  };
  xhttp.open("POST", "/api/csrf", true);
  xhttp.send();
}
