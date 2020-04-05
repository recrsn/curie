import client from "./client";

describe('/ping', () => {
    it('should return message = pong', async () => {
        const response = await client()
            .get('/ping');

        expect(response.status).toBe(200);
        expect(response.text).toBe("pong");
    });
});