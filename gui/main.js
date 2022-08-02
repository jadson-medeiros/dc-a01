function addItem() {
  let data = document.getElementById("text-input").value
  let url = "http://127.0.0.1:8000/tb01"

  fetch(url, {
      method: "POST",
      body: JSON.stringify({text: data}),
      headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then(() => loadItems())
}

function loadItems() {

  let raw = `<table>
              <thead>
              <th id="id">ID</th>
              <th id="text">Text</th>
              <th id="timestamp">Date</th>
              </thead>
          </table>`

  let table = document.getElementsByTagName("table")[0]
  table.innerHTML = raw

  fetch("http://127.0.0.1:3000/tb01")
      .then(res => res.json())
      .then(res => res.forEach((row, i) => {
          if ((i % 2) === 0) {
              table.innerHTML = table.innerHTML +
                  `<tr class="even">
                      <td>${row.id}</td>
                      <td>${row.text}</td>
                      <td>${row.date}</td>
                  </tr>`
          } else {
              table.innerHTML = table.innerHTML +
                  `<tr class="odd">
                      <td>${row.id}</td>
                      <td>${row.text}</td>
                      <td>${row.date}</td>
                  </tr>`
          }
      }))
}