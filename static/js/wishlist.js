(function () {
  "use strict";

  const container = document.getElementById("wishlist-container");
  if (!container) return;

  // ── Helper: call API and return JSON ─────────────────────
  async function api(url, method, body) {
    const opts = {
      method,
      headers: { "Content-Type": "application/json" },
    };
    if (body) opts.body = JSON.stringify(body);

    const resp = await fetch(url, opts);

    if (resp.status === 401) {
      window.location.href = "/login";
      return null;
    }

    if (!resp.ok) {
      const text = await resp.text();
      throw new Error(text || resp.statusText);
    }

    // 204 No Content (delete)
    if (resp.status === 204) return true;

    return resp.json();
  }

  // ── Helper: reload only the rows section via SSR partial ─
  async function refreshRows() {
    const resp = await fetch("/wishlist/rows");
    if (resp.ok) {
      container.innerHTML = await resp.text();
      bindEvents();
    }
  }

  // ── Save note + status (AJAX, no page reload) ──────────
  function onSave(e) {
    const btn = e.target;
    const id = btn.dataset.id;
    const noteInput = container.querySelector(
      `.js-note-input[data-id="${id}"]`
    );
    const statusSelect = container.querySelector(
      `.js-status-select[data-id="${id}"]`
    );

    const note = noteInput ? noteInput.value : "";
    const status = statusSelect ? statusSelect.value : "want_to_visit";

    api(`/api/wishlist/${id}`, "PUT", { note, status })
      .then(() => refreshRows())
      .catch((err) => console.error("Save failed:", err));
  }

  // ── Delete (AJAX, row removed without page reload) ──────
  function onDelete(e) {
    const btn = e.target;
    const id = btn.dataset.id;

    if (!confirm("Remove this entry from your wishlist?")) return;

    api(`/api/wishlist/${id}`, "DELETE")
      .then(() => refreshRows())
      .catch((err) => console.error("Delete failed:", err));
  }

  // ── Bind all event listeners ────────────────────────────
  function bindEvents() {
    container.querySelectorAll(".js-save-note").forEach((el) => {
      el.addEventListener("click", onSave);
    });

    container.querySelectorAll(".js-delete-btn").forEach((el) => {
      el.addEventListener("click", onDelete);
    });
  }

  // ── Initial bind ────────────────────────────────────────
  bindEvents();
})();
