// This is a momument to linux, the monolith that murdered minix
// GomMVC is gunning for all that JS running on the backend,
//  WHY WOULD YOU EVER MAKE IT SO SIMPLE TO PROGRAM A SERVER? DANGEROUS!
// everything is still better than PHP. reddit.com/r/lolPHP

// LoadCSRF pulls the json result of /api/csrf
function LoadCSRF(formDestination, setToID) {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      var csrfTag = document.getElementById(setToID);
      csrfTag.value = this.responseText;
    }
  };
  xhttp.open("POST", "/api/csrf", true);
  xhttp.setRequestHeader("Content-Type", "application/json");
  xhttp.send(JSON.stringify({
    "formDestination" : formDestination,
  }));
}

function submitRegistrationForm(csrfID){
  // pull the values from the form
  var email = document.getElementById('email');
  // make sure the email is valid format
  if (!((/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email.value)))) {
    alert("You have entered an invalid email address!");
    email.value = "try again!"
    return;
  }

  var username = document.getElementById('userName');
  if (username.value.length < 3 || username.value.length > 64) {
    alert("Username must be 3-64 characters")
    username.value = "more than 3 and less than 65 characters!"
    return;
  } else if (!(/^[A-Za-z0-9]+$/.test(username.value))) {
    alert("Username contains more than charaters and numbers!")
    username.value = "alphanumeric characters only!"
    return;
  }

  var firstname = document.getElementById('firstName');
  if (firstname.value.length < 1 || firstname.value.length > 100) {
    alert("First name required")
    firstname.value = "First name required!"
    return;
  }

  var lastname = document.getElementById('lastName');
  if (lastname.value.length < 1 || lastname.value.length > 100) {
    alert("Last name required!")
    lastname.value = "Last name required!"
    return;
  }

  // make sure these are >3 <64 characters
  var password = document.getElementById('password');
  // make sure it's >10 char, measure "entropy"?
  var passwordConfirm = document.getElementById('passwordConfirm');
  // confirm is same as password
  var csrf = document.getElementById(csrfID);
  // confirm the CSRF token loaded.
  var submitData = {
    "email" : email.value,
    "userName" : userName.value,
    "firstName" : firstName.value,
    "lastName" : lastName.value,
    "password" : password.value,
    "passwordConfirm" : passwordConfirm.value,
    "csrf" : JSON.parse(csrf.value),
  }
  console.log(JSON.stringify(submitData));
}
