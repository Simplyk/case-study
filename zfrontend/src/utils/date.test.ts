import { expect } from "@jest/globals";

const { getFormattedDate } = require("./date");

describe("Test date utils", () => {
    it("should be UTC in test", () => {
        expect(new Date().getTimezoneOffset()).toBe(0);
    });

    test("getFormattedDate returns correct date in UTC", () => {
        const timestamp = 1600000000000;
        const expected = "September 13 at 12:26 p.m.";
        const actual = getFormattedDate(timestamp);
        expect(actual).toBe(expected);
    });
});
