optStrings = {
    selectors: {
        feedVideoItem: 'video-feed-item-wrapper',
        modalArrowLeft: 'div.video-card-modal > div > img.arrow-right',
        modalClose: '.video-card-modal > div > div.close',
        modalPlayer: 'div > div > main > div.video-card-modal > div > div.video-card-big > div.video-card-container > div > div > video',
        modalShareInput: '.copy-link-container > input',
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

createVidUrlElement = function(videoArray) {
    var urlSetElement = document.createElement(optStrings.tags.resultTag);
    urlSetElement.innerText = JSON.stringify(videoArray);
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
};

init = () => {
    const newProto = navigator.__proto__;
    delete newProto.webdriver;
    navigator.__proto__ = newProto;
    bootstrapIteratingVideos();
};

init();
'script initialized'