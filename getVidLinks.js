createVidUrlElement = function(videoSet) {
    var videoArray = Object.entries(videoSet).map(x => { 
        return {
            shareLink: x[1].shareLink,
            url: x[0],
        };
    });

    var urlSetElement = document.createElement("video_urls");
    urlSetElement.innerText = JSON.stringify(videoArray);
    document.getElementsByTagName("body")[0].appendChild(urlSetElement);
}

buldVidUrlSet = function(finishCallback) {
    var feedItem = document.getElementsByClassName("video-feed-item-wrapper")[0];
    feedItem.click();

    var videoSet = {};
    var intervalID = window.setInterval(x => {
        var players = document.getElementsByClassName("video-player");
        for (var i = 0; i < players.length; i++) {
            var vidUrl = players[i].getAttribute("src");
            if(!videoSet[vidUrl]) {
                var shareLink = document.querySelector(".copy-link-container > input").value;
                videoSet[vidUrl] = {
                    shareLink: shareLink
                };
            }
        }
        var arrowRight = document.querySelectorAll("div.video-card-modal > div > img.arrow-right")[0];
        if (arrowRight.classList.contains("disabled")) {
            window.clearInterval(intervalID);
            document.querySelector(".video-card-modal > div > div.close").click();
            finishCallback(videoSet);
        } else {
            arrowRight.click();
        }
    }, 500);
};

scrollWhileNew = function(finishCallback) {
    var state = { count: 0 };
    var intervalID = window.setInterval(x => {
        var oldCount = state.count;
        state.count = document.getElementsByClassName("video-feed-item").length;
        if (oldCount !== state.count) {
            window.scrollTo(0, document.body.scrollHeight);
        } else {
            window.clearInterval(intervalID);
            finishCallback();
        }
    }, 1000);
};

init = () => {
    const newProto = navigator.__proto__;
    delete newProto.webdriver;
    navigator.__proto__ = newProto;

    window.setTimeout(x => {
        window.scrollTo(0, document.body.scrollHeight);
        window.setTimeout(x => buldVidUrlSet(createVidUrlElement), 2000);
    }, 1000)
};

init();
'script initialized'