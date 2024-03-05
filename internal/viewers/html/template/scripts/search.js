$(document).ready(function () {
    $("#searchButton").on("click", function () {
        const filter = $("#searchInput").val();
        searchTable(filter);
    });

    $("#searchInput").on("keyup", function (event) {
        if (event.key === "Enter") {
            const filter = $(this).val();
            searchTable(filter);
        }
    });
});

function searchTable(filter) {
    let tableEmpty = true;

    $("#tableBody").find("tr").each(function () {
        const cell = $(this).find("td:eq(0)");
        const cellText = cell.text().toLowerCase();

        if (cellText.indexOf(filter.toLowerCase()) !== -1) {
            $(this).show();
            tableEmpty = false;
        } else {
            $(this).hide();
        }
    });

    if (tableEmpty) {
        $("#tableEmpty").show();
    } else {
        $("#tableEmpty").hide();
    }
}
