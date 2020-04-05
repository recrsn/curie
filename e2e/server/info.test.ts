import client from "./client";
import {promisify} from 'util';
import {exec} from 'child_process';

describe('/info', () => {
    it('should return version info', async () => {
        const {stdout:version } = await promisify(exec)("git describe --tags --always --first-parent");

        const response = await client()
            .get('/info');

        expect(response.status).toBe(200);
        expect(response.body).toEqual({version: version.trim()});
    });
});