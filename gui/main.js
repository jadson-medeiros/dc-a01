function loadItems() {
    var node_url = "http://localhost:3000/tb01"
    var raw = `<table>
                    <thead>
                        <th id="id">ID</th>
                        <th id="col_texto">Text</th>
                        <th id="col_dt">Date</th>
                    </thead>
                </table>`

    var table = document.getElementsByTagName("table")[0]
    table.innerHTML = raw

    fetch(node_url)
        .then(res => res.json())
        .then(res => res.forEach((row, i) => {
            if ((i % 2) === 0) {
                table.innerHTML = table.innerHTML +
                    `<tr class="even">
                        <td>${row.id}</td>
                        <td>${row.col_texto}</td>
                        <td>${row.col_dt}</td>
                    </tr>`
            } else {
                table.innerHTML = table.innerHTML +
                    `<tr class="odd">
                        <td>${row.id}</td>
                        <td>${row.col_texto}</td>
                        <td>${row.col_dt}</td>
                    </tr>`
            }
        }))
}

function addItem() {
    var go_url = "http://localhost:8000/tb01"
    var col_text = document.getElementById("col_text").value

    fetch(go_url, {
        method: "POST",
        body: JSON.stringify({col_texto: col_text})
    }).then(() => loadItems())
}