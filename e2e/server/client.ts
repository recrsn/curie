import * as process from "process"
import * as request from "superagent";

const {E2E_SERVER_ENDPOINT} = process.env;
const endpoint = E2E_SERVER_ENDPOINT || 'http://localhost:8080';

class Client {
    private readonly baseUrl: string;

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    get(path: string) {
        return request.get(this.url(path));
    }

    private url(path: string) {
        return `${this.baseUrl}${path}`;
    }
}

export default function () {
    return new Client(endpoint);
}