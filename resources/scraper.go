// Package resources - This file is automatically generated.
// Do not edit this file manually.
// Check `/generator/resources.go` to change generated content
package resources

//ScraperPath -
var ScraperPath = "scraper.js"

//ScraperScript -
var ScraperScript = "optStrings={selectors:{feedLoading:\".tiktok-ui-loading-container\",modalArrowRight:\"div > div.video-card-container > img.arrow-right\",modalClose:\"div > div.video-card-container > img.control-icon.close\",modalPlayer:\"div.video-card-container > div.video-card-browse > video\",modalCaption:\"div.content-container > div.video-infos-container > h1\",modalSoundLink:\"div.content-container > div.video-infos-container > h2.music-info > a\",modalUploader:\".user-username\",videoPlayer:\"div.video-card-container > div > video\",videoCaption:\"div.content-container > div.video-infos-container > h1\",videoSoundLink:\"div.content-container > div.video-infos-container > h2 > a\",videoUploader:\".user-username\"},classes:{feedVideoItem:\"video-feed-item-wrapper\",modalCloseDisabled:\"disabled\",titleMessage:\"title\"},tags:{resultTag:\"video_urls\",resultParentTag:\"body\"},attributes:{src:\"src\"},tiktokMessages:[\"Couldn't find this account\",\"No videos yet\",\"Video currently unavailable\"]},currentState={preloadCount:0,finished:!1,limit:0},checkForErrors=function(){var e=document.getElementsByClassName(optStrings.classes.titleMessage);if(e&&e.length){var t=Array.from(e).find(e=>optStrings.tiktokMessages.includes(e.textContent)).textContent;if(t)return createVidUrlElement(\"ERR: \"+t),!0}return!1},createVidUrlElement=function(e){var t=document.createElement(optStrings.tags.resultTag);t.innerText=JSON.stringify(e),document.getElementsByTagName(optStrings.tags.resultParentTag)[0].appendChild(t),currentState.finished=!0},buldVidUrlArray=function(e){document.getElementsByClassName(optStrings.classes.feedVideoItem)[0].click();var t=[],r=window.setInterval(o=>{t.push(getCurrentModalVideo()),currentState.limit>0&&t.length>=currentState.limit&&(window.clearInterval(r),document.querySelector(optStrings.selectors.modalClose).click(),e(t));var n=document.querySelectorAll(optStrings.selectors.modalArrowRight)[0];!n||n.classList.contains(optStrings.classes.modalCloseDisabled)?(window.clearInterval(r),document.querySelector(optStrings.selectors.modalClose).click(),e(t)):n.click()},20)},getCurrentModalVideo=function(){var e=document.querySelector(optStrings.selectors.modalPlayer).getAttribute(optStrings.attributes.src),t=window.location.href,r=document.querySelector(optStrings.selectors.modalCaption).textContent,o=document.querySelector(optStrings.selectors.modalSoundLink),n=document.querySelector(optStrings.selectors.modalUploader).textContent,i=o?o.getAttribute(\"href\"):\"\";return{url:e,shareLink:t,caption:r,uploader:n,sound:{title:o?o.text:\"\",link:i}}},getCurrentVideo=function(){if(!checkForErrors()){var e=document.querySelector(optStrings.selectors.videoPlayer).getAttribute(optStrings.attributes.src),t=window.location.href,r=document.querySelector(optStrings.selectors.videoCaption).textContent,o=document.querySelector(optStrings.selectors.videoSoundLink),n=document.querySelector(optStrings.selectors.videoUploader).textContent,i=o?o.getAttribute(\"href\"):\"\";return{url:e,shareLink:t,caption:r,uploader:n,sound:{title:o?o.text:\"\",link:i}}}},scrollBottom=()=>window.scrollTo(0,document.body.scrollHeight),scrollWhileNew=function(e){var t={count:0},r=window.setInterval(o=>{scrollBottom();var n=t.count;if(t.count=document.getElementsByClassName(optStrings.classes.feedVideoItem).length,currentState.limit>0&&(currentState.preloadCount>=currentState.limit||t.count>=currentState.limit)&&(e(createVidUrlElement),window.clearInterval(r)),checkForErrors())window.clearInterval(r);else if(0!=t.count)if(n!==t.count)currentState.preloadCount=t.count;else{if(isLoading())return;window.clearInterval(r),e(createVidUrlElement)}},1e3)},isLoading=function(){var e=document.querySelector(optStrings.selectors.feedLoading);return e&&0!=e.getClientRects().length},bootstrapIteratingVideos=function(e){return currentState.limit=e,scrollWhileNew(buldVidUrlArray),\"bootstrapIteratingVideos\"},bootstrapGetCurrentVideo=function(){var e=getCurrentVideo();return createVidUrlElement(e),\"bootstrapGetCurrentVideo\"},init=()=>{const e=navigator.__proto__;return delete e.webdriver,navigator.__proto__=e,\"script initialized\"},init();"
