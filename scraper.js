optStrings = {
    selectors: {
        feedVideoItem: 'video-feed-item-wrapper',
        modalArrowLeft: 'div.video-card-modal > div > img.arrow-right',
        modalClose: '.video-card-modal > div > div.close',
        modalPlayer: 'div > div > main > div.video-card-modal > div > div.video-card-big > div.video-card-container > div > div > video',
        modalShareInput: '.copy-link-container > input',
        videoPlayer: 'div.video-card-container > div > div > video',
        videoShareInput: 'div.content-container.border > div.copy-link-container > input',
    },
    classes: {
        modalCloseDisabled: 'disabled',
    },
    tags: {
        resultTag: 'video_urls',
        resultParentTag: 'body',
    },
    attributes: {
        src: "src",
    },
};

createVidUrlElement = function(outputObj) {
    var urlSetElement = document.createElement(optStrings.tags.resultTag);
    urlSetElement.innerText = JSON.stringify(outputObj);
    document.getElementsByTagName(optStrings.tags.resultParentTag)[0].appendChild(urlSetElement);
}

buldVidUrlArray = function(finishCallback) {
    var feedItem = document.getElementsByClassName(optStrings.selectors.feedVideoItem)[0];
    feedItem.click();

    var videoArray = [];
    var intervalID = window.setInterval(x => {
        videoArray.push(getCurrentModalVideo());

        var arrowRight = document.querySelectorAll(optStrings.selectors.modalArrowLeft)[0];
        if (arrowRight.classList.contains(optStrings.classes.modalCloseDisabled)) {
            window.clearInterval(intervalID);
            document.querySelector(optStrings.selectors.modalClose).click();
            finishCallback(videoArray);
        } else {
            arrowRight.click();
        }
    }, 500);
};

getCurrentModalVideo = function() {
    var modalPlayer = document.querySelector(optStrings.selectors.modalPlayer);
    var vidUrl = modalPlayer.getAttribute(optStrings.attributes.src);
    var shareLink = document.querySelector(optStrings.selectors.modalShareInput).value;

    return {
        url: vidUrl,
        shareLink: shareLink
    };
}

getCurrentVideo = function() {
    var player = document.querySelector(optStrings.selectors.videoPlayer);
    var vidUrl = player.getAttribute(optStrings.attributes.src);
    var shareLink = document.querySelector(optStrings.selectors.videoShareInput).value;

    return {
        url: vidUrl,
        shareLink: shareLink
    };
}

scrollWhileNew = function(finishCallback) {
    var state = { count: 0 };
    var intervalID = window.setInterval(x => {
        var oldCount = state.count;
        state.count = document.getElementsByClassName(optStrings.selectors.feedVideoItem).length;
        if (oldCount !== state.count) {
            window.scrollTo(0, document.body.scrollHeight);
        } else {
            window.clearInterval(intervalID);
            finishCallback();
        }
    }, 1000);
};

bootstrapIteratingVideos = function() {
    var intervalID = window.setInterval(() => {
        window.scrollTo(0, document.body.scrollHeight);
        if (document.getElementsByClassName(optStrings.selectors.feedVideoItem).length > 0) {
            window.setTimeout(() => buldVidUrlArray(createVidUrlElement), 100);
            window.clearInterval(intervalID);
        }
    }, 500);
    return 'bootstrapIteratingVideos';
};

bootstrapGetCurrentVideo = function() {
    var video = getCurrentVideo();
    createVidUrlElement(video);
    return 'bootstrapGetCurrentVideo';
}

init = () => {
    const newProto = navigator.__proto__;
    delete newProto.webdriver;
    navigator.__proto__ = newProto;
};

init();
'script initialized'