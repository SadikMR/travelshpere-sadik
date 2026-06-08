document.addEventListener("DOMContentLoaded", () => {
    const logoutBtn = document.getElementById("logout-btn");

    if (!logoutBtn) return;

    logoutBtn.addEventListener("click", async () => {
        const form = document.createElement("form");
        form.method = "POST";
        form.action = "/logout";

        document.body.appendChild(form);
        form.submit();
    });
});