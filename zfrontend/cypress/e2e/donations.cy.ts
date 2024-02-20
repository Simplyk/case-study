describe("Test Donation Table Page", () => {
    beforeEach(() => {
        cy.visit("http://localhost:3000/");
    });

    context("Table Elements", () => {
        specify("initial state", () => {
            cy.intercept("GET", "http://localhost:1323/api/v1/donations").as(
                "noDonations"
            );

            cy.get("table").should("be.visible");

            cy.get("thead").should("be.visible");
            cy.get("thead").should("not.be.empty");
            cy.get("thead").should("have.length", 1);
            cy.get("thead").should("contain", "Date");
            cy.get("thead").should("contain", "First Name");
            cy.get("thead").should("contain", "Last Name");
            cy.get("thead").should("contain", "Amount");

            cy.get("tbody").should("be.visible");
            cy.get("tbody").should("not.be.empty");
            cy.get("tbody").should("have.length", 1);

            cy.get("tr").should("have.length", 3);
            cy.get("tr").should("contain", "No rows to display.");
        });
    });

    context("Table Data", () => {
        specify("after loading", () => {
            cy.intercept("GET", "http://localhost:1323/api/v1/donations", {
                fixture: "100-donations.json",
            }).as("TransactionWithDonation.getAll");

            cy.wait("@TransactionWithDonation.getAll").then((interception) => {
                assert.isNotNull(interception.response?.body);
                cy.log(interception.response?.body);

                const body = interception.response?.body;
                cy.get("tbody").should("be.visible");
                cy.get("tbody").should("not.be.empty");
                cy.get("tr").should(
                    "have.length",
                    Math.min(body.length, 10) + 2
                );
            });
        });
    });
});
