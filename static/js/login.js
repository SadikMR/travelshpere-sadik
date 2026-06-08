document.addEventListener("DOMContentLoaded", () => {

    const form = document.getElementById("loginForm");

    form.addEventListener("submit", (event) => {

        const username =
            document
                .getElementById("username")
                .value
                .trim();

        if (!username) {

            event.preventDefault();

            alert("Username is required");
        }
    });

});