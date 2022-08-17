let headers = new Headers();
headers.append('Content-Type', 'application/json');
headers.append('Accept', 'application/json');
headers.append('Origin', 'http://localhost:8080');
var apiUrl = 'http://localhost:8080';
fetch(apiUrl, {
  mode: 'no-cors',
  method: 'GET',
  headers: headers
}).then((response) => {
    if (response.ok) {
      return response.json();
    } else {
      throw new Error("NETWORK RESPONSE ERROR");
    }
  })
  .then(data => {
    console.log(data);
    getJoke(data)
  })
  .catch((error) => console.error("FETCH ERROR:", error));





//   <!DOCTYPE html>
// <html>

// <body>

//   <h1>Dad Jokes</h1>
//   <button id="btn" onclick="get()">Click me</button>

//   <p id="output">Joke: </p>
//   <div v-html="link.FUNCTIONVARIABLE"></div>

//   <script type="text/javascript">
//     function get()
//     {
//       var apiUrl = 'http://localhost:8080';
//       let headers = new Headers();
//       let output = document.getElementById('output')
//       headers.append('Content-Type', 'application/json');
//       headers.append('Accept', 'application/json');
//       headers.append('Origin', 'http://localhost:8080');

//       fetch(apiUrl, {
//         mode: 'no-cors',
//         method: 'GET',
//         headers: headers
//       }).then(response => response.json())
//         .then(json => console.log(json))
//         .catch(error => console.log(error.message))
//     }
//     window.addEventListener('load', event =>
//     {
//       get();
//     });
//     let btn = document.getElementById("btn");
//     btn.addEventListener('click', event =>
//     {
//       get();
//       document.getElementById('output').innerHTML = get()
//     });
//   </script>
// </body>

// </html>