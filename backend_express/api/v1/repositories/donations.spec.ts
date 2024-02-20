import { getDonations } from "./donations";

describe("getDonations", () => {
  it("returns the last 100 donations", () => {
    const donations = getDonations();
    expect(donations).toHaveLength(100);
  });
})