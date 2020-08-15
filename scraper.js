optStrings = {
    selectors: {
        feedLoading: '.tiktok-ui-loading-container',
        modalArrowRight: 'div > div.video-card-container > img.arrow-right',
        modalClose: 'div > div.video-card-container > img.control-icon.close',
        modalPlayer: 'div.video-card-container > div.video-card-browse > video',
        modalCaption: 'div.content-container > div.video-infos-container > h1',
        modalSoundLink: 'div.content-container > div.video-infos-container > h2.music-info > a',
        modalUploader: '.user-username',
        videoPlayer: 'div.video-card-container > div > video',
        videoCaption: 'div.content-container > div.video-infos-container > h1',
        videoSoundLink: 'div.content-container > div.video-infos-container > h2 > a',
        videoUploader: '.user-username',
    },
    classes: {
        feedVideoItem: 'video-feed-item-wrapper',
        modalCloseDisabled: 'disabled',
        titleMessage: 'title',
    },
    tags: {
        resultTag: 'video_urls',
        resultParentTag: 'body',
    },
    attributes: {
        src: "src",
    },
    tiktokMessages: [
        "Couldn't find this account",
        "No videos yet",
        "Video currently unavailable",
    ],
};

currentState = {
    preloadCount: 0,
    finished: false,
    limit: 0
};

checkForErrors = function () {
    var titles = document.getElementsByClassName(optStrings.classes.titleMessage);
    //debugger;
    if (titles && titles.length) {
        var error = Array.from(titles).find(x => optStrings.tiktokMessages.includes(x.textContent)).textContent;
        if (error) {
            createVidUrlElement("ERR: " + error);
            return true;
        }
    }
    return false;
};

createVidUrlElement = function (outputObj) {
    var urlSetElement = document.createElement(optStrings.tags.resultTag);
    urlSetElement.innerText = JSON.stringify(outputObj);
    document.getElementsByTagName(optStrings.tags.resultParentTag)[0].appendChild(urlSetElement);
    currentState.finished = true;
};

buldVidUrlArray = function (finishCallback) {
    var feedItem = document.getElementsByClassName(optStrings.classes.feedVideoItem)[0];
    feedItem.click();

    var videoArray = [];
    var intervalID = window.setInterval(x => {
        videoArray.push(getCurrentModalVideo());
        if (currentState.limit > 0) {
            if (videoArray.length >= currentState.limit) {
                window.clearInterval(intervalID);
                document.querySelector(optStrings.selectors.modalClose).click();
                finishCallback(videoArray);
            }
        }
        var arrowRight = document.querySelectorAll(optStrings.selectors.modalArrowRight)[0];
        if (!arrowRight || arrowRight.classList.contains(optStrings.classes.modalCloseDisabled)) {
            window.clearInterval(intervalID);
            document.querySelector(optStrings.selectors.modalClose).click();
            finishCallback(videoArray);
        } else {
            arrowRight.click();
        }
    }, 20);
};

getCurrentModalVideo = function () {
    var modalPlayer = document.querySelector(optStrings.selectors.modalPlayer);
    var vidUrl = modalPlayer.getAttribute(optStrings.attributes.src);
    var shareLink = window.location.href;
    var caption = document.querySelector(optStrings.selectors.modalCaption).textContent;
    var soundLink = document.querySelector(optStrings.selectors.modalSoundLink);
    var uploader = document.querySelector(optStrings.selectors.modalUploader).textContent;
    var soundHref = soundLink ? soundLink.getAttribute("href") : '';
    var soundText = soundLink ? soundLink.text : '';

    return {
        url: vidUrl,
        shareLink: shareLink,
        caption: caption,
        uploader: uploader,
        sound: {
            title: soundText,
            link: soundHref,
        },
    };
};

getCurrentVideo = function () {
    //debugger;
    if (checkForErrors()) return;
    var player = document.querySelector(optStrings.selectors.videoPlayer);
    var vidUrl = player.getAttribute(optStrings.attributes.src);
    var shareLink = window.location.href;
    var caption = document.querySelector(optStrings.selectors.videoCaption).textContent;
    var soundLink = document.querySelector(optStrings.selectors.videoSoundLink);
    var uploader = document.querySelector(optStrings.selectors.videoUploader).textContent;
    var soundHref = soundLink ? soundLink.getAttribute("href") : '';
    var soundText = soundLink ? soundLink.text : '';

    return {
        url: vidUrl,
        shareLink: shareLink,
        caption: caption,
        uploader: uploader,
        sound: {
            title: soundText,
            link: soundHref,
        },
    };
};

scrollBottom = () => window.scrollTo(0, document.body.scrollHeight);

scrollWhileNew = function (finishCallback) {
    var state = {
        count: 0
    };
    var intervalID = window.setInterval(x => {
        scrollBottom();
        var oldCount = state.count;
        state.count = document.getElementsByClassName(optStrings.classes.feedVideoItem).length;
        if (currentState.limit > 0) {
            if (currentState.preloadCount >= currentState.limit || state.count >= currentState.limit) {
                finishCallback(createVidUrlElement);
                window.clearInterval(intervalID);
            }
        }
        if (checkForErrors()) {
            window.clearInterval(intervalID);
            return;
        } else if (state.count == 0) {
            return;
        }
        if (oldCount !== state.count) {
            currentState.preloadCount = state.count;
        } else {
            if (isLoading()) {
                return;
            }
            window.clearInterval(intervalID);
            finishCallback(createVidUrlElement);
        }
    }, 1000);
};

isLoading = function () {
    var loadingElement = document.querySelector(optStrings.selectors.feedLoading);
    return loadingElement && loadingElement.getClientRects().length != 0;
}

bootstrapIteratingVideos = function (limit) {
    currentState.limit = limit;
    scrollWhileNew(buldVidUrlArray);
    return 'bootstrapIteratingVideos';
};

bootstrapGetCurrentVideo = function () {
    var video = getCurrentVideo();
    createVidUrlElement(video);
    return 'bootstrapGetCurrentVideo';
};

init = () => {
    const newProto = navigator.__proto__;
    delete newProto.webdriver;
    navigator.__proto__ = newProto;
    return 'script initialized';
};

init();