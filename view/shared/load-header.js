document.addEventListener('DOMContentLoaded', function () {
    fetch('../view/shared/header.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('header').innerHTML = data;
        });
}); 