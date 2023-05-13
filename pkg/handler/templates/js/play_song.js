
    var audioPlayer = document.getElementById("audioPlayer");
    var seekBar = document.getElementById("seekBar");
    var currentTime = document.getElementById("currentTime");
    var duration = document.getElementById("duration");
    function toggleAudio() {
    if (audioPlayer.paused) {
        audioPlayer.play();
    } else {
        audioPlayer.pause();
    }
    }

    function initializeTime() {
        duration.innerHTML = formatTime(audioPlayer.duration);
    }

    function updateTime() {
        currentTime.innerHTML = formatTime(audioPlayer.currentTime);
        seekBar.value = (audioPlayer.currentTime / audioPlayer.duration) * 100;
    }

    function seek() {
        var seekTime = (seekBar.value / 100) * audioPlayer.duration;
        audioPlayer.currentTime = seekTime;
    }

    function formatTime(time) {
        var minutes = Math.floor(time / 60);
        var seconds = Math.floor(time % 60);
        if (seconds < 10) {
            seconds = "0" + seconds;
        }
        return minutes + ":" + seconds;
    }
