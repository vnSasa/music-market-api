
const formUpdate = document.getElementById('update-form');
    formUpdate.addEventListener('submit', async (e) => {
        e.preventDefault();
        const formData = new FormData(formUpdate);
        const response = await fetch(formUpdate.action, {
            method: 'PUT',
            body: formData
        });
        if (response.ok) {
            window.location.href = '/';
        } else {
            console.log('Error update:', error);
        }
    });