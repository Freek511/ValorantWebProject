const login = document.getElementById("login");
const signup = document.getElementById("signup")

login.addEventListener('submit', (e) =>  {
  e.preventDefault();
  console.dir(e);
  console.log(e.target.children[1].value)
  console.log(e.target.children[2].value)
  data = {
    name: e.target.children[1].value,
    password: e.target.children[2].value
  }
  fetch('127.0.0.1/check_user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(data)
  }).then(res => res.json()).then(res => {
    console.log(res);
    if (res.status === 200) {
      // window.location.href = "localhost:8080/";
      alert("Logged in");
    }
    else {
      alert("Wrong password or username");
    }
  })
})
