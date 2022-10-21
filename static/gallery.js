async function addAlbum() {
    let album_name = document.getElementById("album_name").value;
    const response = await fetch("http://localhost:9001/album", {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: `{
           "name": "${album_name}"
        }`,
    });

    response.json().then(data => {
        console.log(data);
        alert(data["Message"])
    });

}

async function viewAllAlbums() {

    const response = await fetch("http://localhost:9001/albums", {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    });
    const albumSpace = document.getElementById("album_space")
    var parentDIV = document.getElementsByClassName("grid-container-parent")[0];

    response.json().then(data => {
        console.log(data);
        data["albums"].forEach(item => {
            const albumName = item.split("/")[2]

            const folderButton = document.createElement('button')
            const label = document.createElement('label')
            const br = document.createElement('br')
            label.innerHTML = albumName
            label.style.position = "relative"
            label.style.right = "40%"
            folderButton.className = "fa fa-folder"
            folderButton.style.fontSize = "60px"
            folderButton.onclick = getAllImages(albumName)
            parentDIV.appendChild(folderButton)
            parentDIV.appendChild(br)
            parentDIV.appendChild(label)


        });
    });

}

async function getAllImages(name) {
    console.log(name)
    const response = await fetch("http://localhost:9001/album/" + name + "/images", {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    });
    const row = document.createElement("row")
    row.className = "row"
    const column = document.createElement("column")
    column.className = "column"
    row.appendChild(column)
    response.json().then(data => {
        console.log(data);
        if (data["images"]) {

            data["images"].forEach(item => {
                const img = document.createElement('img')
                img.src = item
                column.appendChild(img)
            });
        }
    });
}

function openForm() {
    document.getElementById("myForm").style.display = "block";
}

function closeForm() {
    document.getElementById("myForm").style.display = "none";
}