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


    const parentDIV = document.getElementById("gallery_space");
    const albumSpace = document.createElement("div")

    albumSpace.className = "grid-container grid-container-parent"

    parentDIV.appendChild(albumSpace)


    response.json().then(data => {
        console.log(data);
        data["albums"].forEach(item => {
            const albumName = item.split("/")[2]
            const imageSpace = document.createElement("div")
            albumSpace.id = "album_space"
            imageSpace.className = "image_space"
            imageSpace.style.display = "none"
            imageSpace.id = albumName
            parentDIV.appendChild(imageSpace)
            const folderButton = document.createElement('button')
            const label = document.createElement('label')
            const br = document.createElement('br')
            label.innerHTML = albumName
            label.style.position = "relative"
            label.style.right = "40%"
            folderButton.className = "fa fa-folder"
            folderButton.style.fontSize = "60px"
            folderButton.onclick = function () {
                getAllImages(albumName)
            };
            albumSpace.appendChild(folderButton)
            albumSpace.appendChild(br)
            albumSpace.appendChild(label)


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

    const parentDIV = document.getElementById(name);
    parentDIV.style.display = "block"
    const row = document.createElement("row")
    row.className = "row"
    const column = document.createElement("column")
    column.className = "column"
    parentDIV.appendChild(row)
    row.appendChild(column)
    response.json().then(data => {

        if (data["images"]) {
            data["images"].forEach(item => {
                item = item.replace("static/", "")
                console.log("hellooo" + item);
                const img = document.createElement('img')
                img.src = item
                img.style.width = "20%"
                img.style.padding = "1%"
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