document.addEventListener("DOMContentLoaded", function () {
    var form = document.getElementById('chroma');
    var elms = document.getElementsByTagName("select");
    for (e of elms) {
        e.addEventListener('change', () => form.submit());
    }
});
