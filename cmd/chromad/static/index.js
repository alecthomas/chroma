document.addEventListener("DOMContentLoaded", function () {
    var style = document.createElement('style');
    var ref = document.querySelector('script');
    ref.parentNode.insertBefore(style, ref);

    var form = document.getElementById('chroma');
    var textArea = form.elements["text"];
    var csrfToken = form.elements["gorilla.csrf.Token"].value;
    var elms = document.getElementsByTagName("select");
    var output = document.getElementById("output");

    function debounce(func, wait, immediate) {
        var timeout;
        return function () {
            var context = this, args = arguments;
            var later = function () {
                timeout = null;
                if (!immediate) func.apply(context, args);
            };
            var callNow = immediate && !timeout;
            clearTimeout(timeout);
            timeout = setTimeout(later, wait);
            if (callNow) func.apply(context, args);
        };
    };

    function getFormJSON() {
        return {
            "language": document.getElementById("language").value,
            "style": document.getElementById("style").value,
            "text": document.getElementById("text").value,
        }
    }

    function update(event) {
        fetch("api/render", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'cors', // no-cors, cors, *same-origin
            cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            credentials: 'same-origin', // include, *same-origin, omit
            headers: {
                'X-CSRF-Token': csrfToken,
                'Content-Type': 'application/json',
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: 'follow', // manual, *follow, error
            referrer: 'no-referrer', // no-referrer, *client
            body: JSON.stringify(getFormJSON()),

        }).then(data => {
            data.json().then(
                value => {
                    style.innerHTML = "#output { " + value.background + "}";
                    output.innerHTML = value.html;
                }
            );
        }).catch(reason => {
            console.log(reason);
        })

        event.preventDefault();
    }

    var eventHandler = (event) => update(event);
    for (e of elms) {
        e.addEventListener('change', eventHandler);
    }
    form.addEventListener('submit', eventHandler);

    var debouncedEventHandler = debounce(eventHandler, 250);
    textArea.addEventListener('input', debouncedEventHandler);
    textArea.addEventListener('change', debouncedEventHandler)
});
