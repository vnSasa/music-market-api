function addToPlaylist(id) {
    Swal.fire({
      title: 'Are you sure?',
      text: 'Do you want to add this song to your playlist?',
      icon: 'question',
      showCancelButton: true,
      confirmButtonText: 'Yes, add it!',
      cancelButtonText: 'No, cancel'
    }).then((result) => {
      if (result.isConfirmed) {
        fetch(`/api_user/add_to_playlist/${id}`, {
          method: 'POST'
        })
        .then(response => {
          if (response.ok) {
            Swal.fire(
              'Success!',
              'The song has been added to your playlist.',
              'success'
            ).then(() => {
              window.location.href = '/api_user/get_song';
            });
          } else {
            Swal.fire(
              'Error!',
              'You are trying to add a song that you have previously added.',
              'error'
            ).then(() => {
              window.location.href = '/api_user/get_song';
            });
          }
        })
      .catch(error => console.log('Error:', error));
    }
  });
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