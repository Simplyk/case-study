"use client";

import { NextUIProvider, Pagination } from "@nextui-org/react";
import { useState, useEffect, useMemo } from "react";
import {
    Table,
    TableHeader,
    TableBody,
    TableColumn,
    TableRow,
    TableCell,
} from "@nextui-org/react";
import { TransactionWithDonation } from "@/models/donation";
import { getFormattedDate } from "@/utils/date";

const columns = [
    { key: "createdAtUtc", label: "Date" },
    { key: "firstName", label: "First Name" },
    { key: "lastName", label: "Last Name" },
    { key: "amount", label: "Amount" },
];

export default function DonationsTable() {
    const [donations, setDonations] = useState<TransactionWithDonation[]>([]);
    const [page, setPage] = useState(1);
    const [pages, setPages] = useState(1);
    const rowsPerPage = 10;

    const items = useMemo(() => {
        const start = (page - 1) * rowsPerPage;
        const end = start + rowsPerPage;
        return donations.slice(start, end);
    }, [page, donations]);

    useEffect(() => {
        TransactionWithDonation.getAll()
            .then((donations) => {
                setDonations(donations);
                setPages(Math.ceil(donations.length / rowsPerPage));
            })
            .catch((err) => {
                console.error(err);
            });
    }, []);

    return (
        <NextUIProvider>
            <Table
                aria-label="donations table"
                bottomContent={
                    <Pagination
                        isCompact
                        showControls
                        showShadow
                        color="primary"
                        page={page}
                        total={pages}
                        onChange={setPage}
                    />
                }
            >
                <TableHeader columns={columns}>
                    {(column) => (
                        <TableColumn key={column.key}>
                            {column.label}
                        </TableColumn>
                    )}
                </TableHeader>
                <TableBody items={items} emptyContent={"No rows to display."}>
                    {(item) => (
                        <TableRow key={item.id} className="text-black">
                            <TableCell>
                                {getFormattedDate(item.donation.createdAtUtc)}
                            </TableCell>
                            <TableCell>{item.donation.firstName}</TableCell>
                            <TableCell>{item.donation.lastName}</TableCell>
                            <TableCell>
                                {(
                                    item.donation.amount - item.refundedAmount
                                ).toLocaleString()}
                            </TableCell>
                        </TableRow>
                    )}
                </TableBody>
            </Table>
        </NextUIProvider>
    );
}
