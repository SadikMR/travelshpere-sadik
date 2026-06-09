(function () {
  "use strict";

  const container = document.getElementById("wishlist-rows");
  if (!container) return;

  // ── API helper ────────────────────────────────────────────
  async function api(url, method, body) {
    const opts = {
      method,
      headers: {
        "Content-Type": "application/json",
      },
    };

    if (body) {
      opts.body = JSON.stringify(body);
    }

    const resp = await fetch(url, opts);

    if (resp.status === 401) {
      window.location.href = "/login";
      return null;
    }

    if (!resp.ok) {
      const text = await resp.text();
      throw new Error(text || resp.statusText);
    }

    if (resp.status === 204) {
      return true;
    }

    return resp.json();
  }

  // ── Refresh rows partial ──────────────────────────────────
  async function refreshRows() {
    const resp = await fetch("/wishlist/rows");

    if (!resp.ok) return;

    container.innerHTML = await resp.text();
    bindEvents();
  }

  // ── Button states ─────────────────────────────────────────
  function setSaved(btn) {
    btn.textContent = "Saved";
    btn.classList.add("wl-btn-save--saved");
  }

  function setUnsaved(btn) {
    btn.textContent = "Save";
    btn.classList.remove("wl-btn-save--saved");
  }

  // ── Save wishlist entry ───────────────────────────────────
  function onSave(e) {
    const btn = e.currentTarget;
    const id = btn.dataset.id;

    const noteInput = container.querySelector(
      `.js-note-input[data-id="${id}"]`
    );

    const statusSelect = container.querySelector(
      `.js-status-select[data-id="${id}"]`
    );

    const note = noteInput ? noteInput.value : "";
    const status = statusSelect ? statusSelect.value : "Planned";

    api(`/api/wishlist/${id}`, "PUT", {
      note,
      status,
    })
      .then(() => {
        setSaved(btn);
      })
      .catch((err) => {
        console.error("Save failed:", err);
      });
  }

  // ── Delete wishlist entry ─────────────────────────────────
  function onDelete(e) {
    const btn = e.currentTarget;
    const id = btn.dataset.id;

    if (!confirm("Remove this entry from your wishlist?")) {
      return;
    }

    api(`/api/wishlist/${id}`, "DELETE")
      .then(() => {
        refreshRows();
      })
      .catch((err) => {
        console.error("Delete failed:", err);
      });
  }

  // ── Mark row dirty after edits ────────────────────────────
  function onRowChange(e) {
    const id = e.target.dataset.id;

    if (!id) return;

    const btn = container.querySelector(
      `.js-save-note[data-id="${id}"]`
    );

    if (btn) {
      setUnsaved(btn);
    }
  }

  // ── Bind events ───────────────────────────────────────────
  function bindEvents() {
    container.querySelectorAll(".js-save-note").forEach((btn) => {
      btn.removeEventListener("click", onSave);
      btn.addEventListener("click", onSave);
    });

    container.querySelectorAll(".js-delete-btn").forEach((btn) => {
      btn.removeEventListener("click", onDelete);
      btn.addEventListener("click", onDelete);
    });

    container
      .querySelectorAll(".js-note-input, .js-status-select")
      .forEach((el) => {
        el.removeEventListener("input", onRowChange);
        el.removeEventListener("change", onRowChange);

        el.addEventListener("input", onRowChange);
        el.addEventListener("change", onRowChange);
      });
  }

  bindEvents();
})();