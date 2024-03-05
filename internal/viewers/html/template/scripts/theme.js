$(document).ready(function () {
    setTheme(getPreferredTheme());
    showActiveTheme(getPreferredTheme());

    window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", function () {
        const storedTheme = getStoredTheme();
        if (storedTheme !== "light" && storedTheme !== "dark") {
            setTheme(getPreferredTheme());
        }
    });

    $(".dropdown-item").click(function () {
        const selectedTheme = $(this).data("theme-value");
        setTheme(selectedTheme);
        setStoredTheme(selectedTheme);
        showActiveTheme(selectedTheme);
    });
});

const getStoredTheme = () => localStorage.getItem("theme");

const setStoredTheme = (theme) => localStorage.setItem("theme", theme);

function getPreferredTheme() {
    const storedTheme = getStoredTheme();
    if (storedTheme) {
        return storedTheme;
    }
    return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

function setTheme(theme) {
    if (theme === "auto") {
        document.documentElement.setAttribute("data-bs-theme",
            (window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light"));
    } else {
        document.documentElement.setAttribute("data-bs-theme", theme);
    }
}

function showActiveTheme(theme) {
    $(".theme-item").css("font-weight", "normal");
    $(`#theme-${theme}`).css("font-weight", "bold");

    $("[data-bs-theme-value]").each(function () {
        $(this).attr("aria-pressed", "false");
    });
}
