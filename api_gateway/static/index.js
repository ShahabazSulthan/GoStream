window.addEventListener('load', function() {
    fetch('/video/all') 
    .then(async response => {
        let res = await response.json();
        console.log(res);

        res.Video.forEach(video => { 
            console.log(video);
            const videoElement = document.createElement('video');
            videoElement.controls = true;
            
            videoElement.autoplay = true;       
            
            if (Hls.isSupported()) {
                const hls = new Hls();
                hls.loadSource('/stream/' + video.VideoId + '/playlist.m3u8');
                hls.attachMedia(videoElement);
            } else if (videoElement.canPlayType('application/vnd.apple.mpegurl')) {
                videoElement.src = '/stream/' + video.VideoId + '/playlist.m3u8';
            } else {
                console.error('HLS is not supported on this browser.');
            }
            
            document.getElementById('videoContainer').appendChild(videoElement);
        });
    })
    .catch(error => {
        console.error('Error fetching video data:', error);
    });
});