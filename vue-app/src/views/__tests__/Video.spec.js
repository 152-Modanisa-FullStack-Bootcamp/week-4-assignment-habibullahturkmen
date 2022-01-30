import {shallowMount} from '@vue/test-utils';
import Video from "@/components/Video";

describe("Video.vue",() => {

    it("Video component should exist", () => {
        const wrapper = shallowMount(Video, {
            propsData: {
                video: {
                    "id": 1,
                    "videoAddress": "https://www.youtube.com/watch?v=FXpIoQ_rT_c",
                    "coverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/1-cover.webp",
                    "hoverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/1-hover.webp",
                    "title": "Vue.js Course for Beginners [2021 Tutorial]",
                    "viewCount": 254,
                    "publishDateInMonth": 4,
                    "ownerImage": "https://yt3.ggpht.com/ytc/AKedOLTtJvQ1Vfew91vemeLaLdhjOwGx3tTBLlreK_QUyA=s68-c-k-c0x00ffffff-no-rj",
                    "ownerName": "freeCodeCamp.org",
                    "description": "Learn Vue 3 by in this full course. Vue.js is an open-source model–view–view model front end JavaScript framework for building user interfaces and single-page applications."
                }
            }

        });
        expect(wrapper.exists()).toBeTruthy();
    });

    it("When user clicks on video image, he/she should be able to navigate to watch page", async () => {
        const navigateToVideoSpy = jest.spyOn(Video.methods, 'navigateToVideo');
        let video = {
            "id": 21,
            "videoAddress": "https://www.youtube.com/watch?v=qZXt1Aom3Cs",
            "coverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/2-cover.webp",
            "hoverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/2-hover.webp",
            "title": "Vue JS Crash Course",
            "viewCount": 623,
            "publishDateInMonth": 10,
            "ownerImage": "https://yt3.ggpht.com/ytc/AKedOLSxHOOxxa9Af8Bfb2XMop3lm4tor9bViWiC-d5aaw=s68-c-k-c0x00ffffff-no-rj",
            "ownerName": "Traversy Media",
            "description": "Learn the fundamentals of Vue JS (v3) in this project-based crash course",
            "favorite": true
        }

        let routerPushMock = jest.fn();

        const wrapper = shallowMount(Video, {
            propsData: {
                video: video
            },
            mocks: {
                $router: {
                    push: routerPushMock
                }
            }
        });

        const videoImage = wrapper.find("#videoID");
        await videoImage.trigger("click");

        expect(navigateToVideoSpy).toBeCalled();
        expect(routerPushMock).toHaveBeenCalledWith({"path": "watch", "query": {"id": video.id}});
    });

    it("When user hovers on video image, he/she should see the hoverImage/GIF", () => {
        let video = {
            "id": 21,
            "videoAddress": "https://www.youtube.com/watch?v=qZXt1Aom3Cs",
            "coverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/2-cover.webp",
            "hoverImage": "https://raw.githubusercontent.com/modanisa/bootcamp-video-db/main/video-images/2-hover.webp",
            "title": "Vue JS Crash Course",
            "viewCount": 623,
            "publishDateInMonth": 10,
            "ownerImage": "https://yt3.ggpht.com/ytc/AKedOLSxHOOxxa9Af8Bfb2XMop3lm4tor9bViWiC-d5aaw=s68-c-k-c0x00ffffff-no-rj",
            "ownerName": "Traversy Media",
            "description": "Learn the fundamentals of Vue JS (v3) in this project-based crash course",
            "favorite": true
        }

        const wrapper = shallowMount(Video, {
            propsData: {
                video: video
            },
            data() {
                return {
                    hover: true
                }
            }
        });

        const hoverImage = wrapper.find("img").attributes().src;
        expect(hoverImage).toEqual(video.hoverImage);
    });

});
