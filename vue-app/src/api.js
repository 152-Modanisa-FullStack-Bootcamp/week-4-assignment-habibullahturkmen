import axios from 'axios'
import adapter from "axios/lib/adapters/http";

axios.defaults.adapter = adapter;

export class API {
    useProxy = false
    constructor(url, useProxy) {
        this.useProxy = useProxy
        if (url === undefined || url === "") {
            url = process.env.VUE_APP_BASE_API_URL;
        }
        if (url.endsWith("/")) {
            url = url.substr(0, url.length - 1)
        }
        this.url = url
    }

    withPath(path) {
        if (this.useProxy) {
            return path
        }

        if (!path.startsWith("/")) {
            path = "/" + path
        }
        return `${this.url}${path}`
    }

    async getVideos() {
        return axios.get(this.withPath('/api/v1/videos')).then(r => r.data)
    }

}

export default new API(process.env.VUE_APP_BASE_API_URL, true);
