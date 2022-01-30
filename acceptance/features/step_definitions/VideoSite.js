const {Given, When, Then} = require("cucumber");
const openUrl = require("../support/action/openUrl");
const waitForSelector = require("../support/action/waitForSelector");
const assert = require("assert");

Given(/^that User goes to Video Site Project's HomePage$/, async function () {
    await openUrl.call(this, "/");
});


When(/^page is loaded$/, async function () {
    await waitForSelector.call(this, ".video-list-container");
});

Then(/^User can see some of videos' title like$/, async function (nameArray) {
    const videoContainer = ".video-container";
    for (let [videoName] of nameArray.rawTable) {
        let result = await this.page.$$eval(
            videoContainer,
            async (items, videoName) => {
                const video = items.find(item => item.querySelector("#title").textContent.includes(videoName));
                return !!video;
            },
            videoName
        );
        assert.strictEqual(result, true);
    }
});
