import * as Base64 from "./base64.js";

document.addEventListener("DOMContentLoaded", function () {
  var style = document.createElement('style');
  var ref = document.querySelector('script');
  ref.parentNode.insertBefore(style, ref);

  var form = document.getElementById('chroma');
  var textArea = form.elements["text"];
  var styleSelect = form.elements["style"];
  var languageSelect = form.elements["language"];
  var copyButton = form.elements["copy"];
  var csrfToken = form.elements["gorilla.csrf.Token"].value;
  var output = document.getElementById("output");
  var htmlCheckbox = document.getElementById("html");

  (document.querySelectorAll('.notification .delete') || []).forEach((el) => {
    const notification = el.parentNode;
    el.addEventListener('click', () => {
      notification.parentNode.removeChild(notification);
    });
  });

  // https://stackoverflow.com/a/37697925/7980
  function handleTab(e) {
    var after, before, end, lastNewLine, changeLength, re, replace, selection, start, val;
    if ((e.charCode === 9 || e.keyCode === 9) && !e.altKey && !e.ctrlKey && !e.metaKey) {
      e.preventDefault();
      start = this.selectionStart;
      end = this.selectionEnd;
      val = this.value;
      before = val.substring(0, start);
      after = val.substring(end);
      replace = true;
      if (start !== end) {
        selection = val.substring(start, end);
        if (~selection.indexOf('\n')) {
          replace = false;
          changeLength = 0;
          lastNewLine = before.lastIndexOf('\n');
          if (!~lastNewLine) {
            selection = before + selection;
            changeLength = before.length;
            before = '';
          } else {
            selection = before.substring(lastNewLine) + selection;
            changeLength = before.length - lastNewLine;
            before = before.substring(0, lastNewLine);
          }
          if (e.shiftKey) {
            re = /(\n|^)(\t|[ ]{1,8})/g;
            if (selection.match(re)) {
              start--;
              changeLength--;
            }
            selection = selection.replace(re, '$1');
          } else {
            selection = selection.replace(/(\n|^)/g, '$1\t');
            start++;
            changeLength++;
          }
          this.value = before + selection + after;
          this.selectionStart = start;
          this.selectionEnd = start + selection.length - changeLength;
        }
      }
      if (replace && !e.shiftKey) {
        this.value = before + '\t' + after;
        this.selectionStart = this.selectionEnd = start + 1;
      }
    }
    debouncedEventHandler(e);
  }

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

  function share(event) {
    let data = JSON.stringify(getFormJSON())
    data = Base64.encodeURI(data);
    location.hash = "#" + data;
    try {
      navigator.clipboard.writeText(location.href);
    } catch (e) {
      console.log(e);
    }
    event.preventDefault();
  }

  if (location.hash) {
    let json = Base64.decode(location.hash.substring(1))
    json = JSON.parse(json);
    textArea.value = json.text;
    languageSelect.value = json.language;
    styleSelect.value = json.style;
    htmlCheckbox.checked = json.classes;
    update(new Event('change'));
  }

  var eventHandler = (event) => update(event);
  var debouncedEventHandler = debounce(eventHandler, 250);

  languageSelect.addEventListener('change', eventHandler);
  styleSelect.addEventListener('change', eventHandler);
  htmlCheckbox.addEventListener('change', eventHandler);
  copyButton.addEventListener('click', share);

  textArea.addEventListener('keydown', handleTab);
  textArea.addEventListener('change', debouncedEventHandler);
});
