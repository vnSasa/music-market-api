function addToPlaylist(id) {
    if (confirm('Are you sure you want to add this song to playlist?')) {
        fetch(`/api_user/add_to_playlist/${id}`, {
            method: 'POST'
        })
    }
}