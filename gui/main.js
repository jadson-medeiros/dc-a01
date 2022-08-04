function loadItems() {
    let node_url = "http://127.0.0.1:3000/tb01"
    let data = getItems(node_url);
    let items = JSON.parse(data);
    let table = document.getElementById("table");
    items.forEach(element => {
        let row = getRow(element)
        table.appendChild(row);
    });
}

function addItem() {
    preventDefault()
    let go_url = "http://127.0.0.1:8000/tb01"
    let col_text = document.getElementById("col_text").value
    body = {
        "col_text": col_text,
    }

    let req = new XMLHttpRequest()
    req.open("POST", go_url, true)
    req.setRequestHeader("Content-type", "application/json")
    req.setRequestHeader("Access-Control-Allow-Origin", "*")

    req.send(JSON.stringify(body))

    req.onload = function() {
        console.log(this.responseText)
    }

    return req.responseText
}


function getItems(url){
    let req = new XMLHttpRequest()
    try {
        req.setRequestHeader("Access-Control-Allow-Origin", "*")
        req.open("GET", url, false)
        req.send()
        return req.responseText
    } catch (err) {
        console.log(err)
    }
}

function getRow(element) {
    row = document.createElement("tr");
    tdId = document.createElement("td");
    tdText = document.createElement("td");
    tdDate = document.createElement("td");

    tdId.innerHTML = element.id
    tdText.innerHTML = element.col_text
    tdDate.innerHTML = `${element.col_dt.substring(0,10)} ${element.col_dt.substring(11,19)}`

    row.appendChild(tdId);
    row.appendChild(tdText);
    row.appendChild(tdDate);

    return row;
}
