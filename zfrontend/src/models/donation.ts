import { BasePathV1 } from "@/requests/base";

export class Donation {
    constructor(
        public id: string,
        public firstName: string,
        public lastName: string,
        public createdAtUtc: number,
        public amount: number,
        public thankYouComment: string,
        public isAnonymous: boolean,
        public companyName: string,
        public __typename: string
    ) {}
}

// TransactionWithDonation class est utilisé comme modèle et pour récupérer les donations
export class TransactionWithDonation {
    constructor(
        public id: string,
        public type: string,
        public refundedAmount: number,
        public donation: Donation,
        public __typename: string
    ) {}

    static async getAll(): Promise<TransactionWithDonation[]> {
        const headers: Headers = new Headers();
        headers.set("Content-Type", "application/json");
        headers.set("Accept", "application/json");

        const request: Request = new Request(BasePathV1 + "/donations", {
            method: "GET",
            headers,
        });

        const response = await fetch(request);
        if (!response.ok) {
            throw new Error(response.statusText);
        }

        const json = await response.json();
        if (json.errors) {
            throw new Error(json.errors[0].message);
        }

        return json;
    }
}
