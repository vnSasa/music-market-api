function addToPlaylist(id) {
    if (confirm('Are you sure you want to add this song to playlist?')) {
        fetch(`/api_user/add_to_playlist/${id}`, {
            method: 'POST'
        })
        .then(response => {
            if (response.ok) {
                confirm('The song has been added to your playlist')
                window.location.href = '/api_user/user_playlist';
            } else {
                confirm('You are trying to add a song that you have previously added')
                window.location.href = '/api_user/user_playlist';
            }
        })
        .catch(error => console.log('Error delete:', error));
    }
}

function deleteSongFromPlaylist(id) {
    if (confirm('Are you sure you want to delete this song?')) {
        fetch(`/api_user/delete_from_playlist/${id}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                window.location.href = '/api_user/user_playlist';
            } else {
                console.log('Error delete:', response.status);
            }
        })
        .catch(error => console.log('Error delete:', error));
    }
}