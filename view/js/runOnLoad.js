// This is a momument to linux, the monolith that murdered minix
// GomMVC is gunning for all that JS running on the backend,
//  WHY WOULD YOU EVER MAKE IT SO SIMPLE TO PROGRAM A SERVER? DANGEROUS!
// everything is still better than PHP. reddit.com/r/lolPHP

// LoadCSRF pulls the json result of /api/csrf
function LoadCSRF(inputID,formDestination) {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
     csrfHiddenTag = document.getElementById(inputID)
     csrfHiddenTag.value = this.responseText
    }
  };
  xhttp.open("POST", "/api/csrf", true);
  xhttp.setRequestHeader("Content-Type", "application/json");
  xhttp.send(JSON.stringify({
    "formDestination" : formDestination,
  }));
}

function noJSWarn(){
  document.getElementById("noJavascript").remove()
}
