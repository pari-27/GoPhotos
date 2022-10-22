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
    });

    location.reload();
}

async function addImage() {
    let album_name = document.getElementById("batchSelect").value;

    const formData = new FormData();
    let ins = document.getElementById('files').files.length;
    for (let x = 0; x < ins; x++) {
        formData.append("multiplePhotos", document.getElementById('files').files[x]);

    }
    formData.append("albumName", album_name)
    console.log(formData)
    const response = await fetch("http://localhost:9001/upload/image", {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
        },
        body: formData,
    });

    response.json().then(data => {
        console.log(data);
    });


}

async function deleteAlbum() {
    let album_name = document.getElementById("album_name").value;
    const response = await fetch("http://localhost:9001/album/" + album_name, {
        method: 'DELETE',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    });

    response.text().then(data => {
        console.log(data);
        // alert(data)
    });
    location.reload();

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
    const dropdown = document.getElementById("batchSelect")
    let defaultOption = document.createElement('option');
    defaultOption.text = 'Choose Album';

    dropdown.add(defaultOption);
    dropdown.selectedIndex = 0;

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
            folderButton.className = "fa fa-light fa-folder"
            folderButton.style.fontSize = "60px"
            folderButton.onclick = function () {
                getAllImages(albumName)
            };
            folderButton.onblur = function () {
                imageSpace.style.display = "none"
                removeChildElements(albumName)
            };
            albumSpace.appendChild(folderButton)
            albumSpace.appendChild(br)
            albumSpace.appendChild(label)


            const option = document.createElement('option');
            option.text = albumName;
            option.value = albumName;
            dropdown.add(option);


        });
    });

}

function removeChildElements(name) {
    const parentDIV = document.getElementById(name);
    parentDIV.innerHTML = ''
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


                const img = document.createElement('img')
                img.src = item
                img.style.width = "20%"

                img.style.padding = "1%"
                img.onclick = function () {
                    window.open(this.src)
                };
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

function openForm1() {
    document.getElementById("myForm1").style.display = "block";
}

function closeForm1() {
    document.getElementById("myForm1").style.display = "none";
}