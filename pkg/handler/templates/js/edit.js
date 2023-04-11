function editArtist(event, name, birth, about) {
    event.preventDefault();
    const url = event.target.href;
    const form = document.createElement('form');
    form.action = url;
    form.method = 'GET';

    const nameInput = document.createElement('input');
    nameInput.type = 'text';
    nameInput.name = 'name_artist';
    nameInput.id = 'name_artist';
    nameInput.className = 'form-control';
    nameInput.value = name;

    const birthInput = document.createElement('input');
    birthInput.type = 'text';
    birthInput.name = 'date_of_birth';
    birthInput.id = 'date_of_birth';
    birthInput.className = 'form-control';
    birthInput.value = birth;

    const aboutInput = document.createElement('textarea');
    aboutInput.name = 'about_artist';
    aboutInput.id = 'about_artist';
    aboutInput.className = 'form-control';
    aboutInput.rows = '5';
    aboutInput.textContent = about;

    form.appendChild(nameInput);
    form.appendChild(birthInput);
    form.appendChild(aboutInput);

    document.body.appendChild(form);
    form.submit();
}