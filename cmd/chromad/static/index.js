document.addEventListener("DOMContentLoaded", function () {
    var style = document.createElement('style');
    var ref = document.querySelector('script');
    ref.parentNode.insertBefore(style, ref);

    var form = document.getElementById('chroma');
    var textArea = form.elements["text"];
    var styleSelect = form.elements["style"];
    var languageSelect = form.elements["language"];
    var csrfToken = form.elements["gorilla.csrf.Token"].value;
    var output = document.getElementById("output");
    var htmlCheckbox = document.getElementById("html");

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
    }

    function getFormJSON() {
        return {
            "language": languageSelect.value,
            "style": styleSelect.value,
            "text": textArea.value,
            "classes": htmlCheckbox.checked,
        }
    }

    function update(event) {
        fetch("api/render", {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            headers: {
                'X-CSRF-Token': csrfToken,
                'Content-Type': 'application/json',
            },
            redirect: 'follow',
            referrer: 'no-referrer',
            body: JSON.stringify(getFormJSON()),
        }).then(data => {
            data.json().then(
                value => {
                    if (value.language) {
                        languageSelect.value = value.language;
                    }
                    style.innerHTML = "#output { " + value.background + "}";
                    if (htmlCheckbox.checked) {
                        output.innerText = value.html;
                    } else {
                        output.innerHTML = value.html;
                    }
                }
            );
        }).catch(reason => {
            console.log(reason);
        });

        event.preventDefault();
    }

    var eventHandler = (event) => update(event);
    var debouncedEventHandler = debounce(eventHandler, 250);

    languageSelect.addEventListener('change', eventHandler);
    styleSelect.addEventListener('change', eventHandler);
    htmlCheckbox.addEventListener('change', eventHandler);

    textArea.addEventListener('input', debouncedEventHandler);
    textArea.addEventListener('change', debouncedEventHandler);
});
