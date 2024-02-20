export function getFormattedDate(timestamp: number) {
    const date = new Date(timestamp);
    const formattedDate = date.toLocaleDateString("en-CA", {
        day: "numeric",
        month: "long",
        hour: "numeric",
        minute: "numeric",
    });
    return formattedDate;
}
