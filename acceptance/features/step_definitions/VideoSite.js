const {Given, When, Then} = require("cucumber");
const openUrl = require("../support/action/openUrl");
const waitForSelector = require("../support/action/waitForSelector");
const checkUrlContains = require("../support/check/checkUrlContains");
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
    await this.page.waitForTimeout(4000);
});

Given(/^that User is on Video Site Project's HomePage$/, async function () {
    await openUrl.call(this, "/");
});

When(/^User clicks "([^"]*)" video$/, async function (videoName) {
    const videoContainer = ".video-container";
    await this.page.$$eval(
        videoContainer,
        async (items, videoName) => {
            const n = items.find(item => item.querySelector("#title").textContent.includes(videoName));
            await n.querySelector(".cover-image").click();
        },
        videoName
    );
});

Then(/^User should see watch url correctly$/, async function () {
    await checkUrlContains.call(this, false, "/watch?id=2");
    await this.page.waitForTimeout(4000);
});

When(/^User hovers "([^"]*)" video$/, async function (name) {
    const videoContainer = await this.page.$$('.video-container');
    for (let image of videoContainer) {
        let title = await image.$("#title");
        const nameTextContent = await this.page.evaluate(videoName => videoName.textContent, title);
        if(name === nameTextContent) {
            await image.hover();
            this.condition = true;
            assert.strictEqual(name, nameTextContent);
        }
    }
});

Then(/^User should see hovered image$/, async function () {
    assert.strictEqual(this.condition, true);
    await this.page.waitForTimeout(7000);
});
