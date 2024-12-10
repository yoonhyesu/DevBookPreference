document.addEventListener('DOMContentLoaded', function () {
    fetch('../shared/header.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('header').innerHTML = data;
        });
}); 