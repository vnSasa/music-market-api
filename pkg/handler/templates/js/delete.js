function deleteArtist(id) {
    if (confirm('Are you sure you want to delete this artist?')) {
        fetch(`/api_admin/delete_artist/${id}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                window.location.href = '/api_admin/artist';
            } else {
                console.log('Error delete:', response.status);
            }
        })
        .catch(error => console.log('Error delete:', error));
    }
}

function deleteSong(id) {
    if (confirm('Are you sure you want to delete this song?')) {
        fetch(`/api_admin/delete_song/${id}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                window.location.href = '/api_admin/song';
            } else {
                console.log('Error delete:', response.status);
            }
        })
        .catch(error => console.log('Error delete:', error));
    }
}