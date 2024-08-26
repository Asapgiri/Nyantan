function watchInfo(id) {
    var pub = document.getElementById('translation-'+id)
    if (pub.classList.contains('already-in-load') ||
        pub.classList.contains('loaded')) return
    pub.classList.add('already-in-load')

    pub.innerText = "Loading..."

    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            pub.outerHTML = this.responseText
        }
    }
    xmlHttp.open("GET", '/api/translations/'+id, true);
    xmlHttp.send();
}
