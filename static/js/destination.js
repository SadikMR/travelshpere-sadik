(function () {
  "use strict";

  const data = window.__COUNTRY_DETAIL__;
  if (!data) return;

  // ── Format population ─────────────────────────────────────
  const popEl = document.getElementById("js-population");
  if (popEl) {
    const n = parseInt(popEl.dataset.pop, 10);
    if (!isNaN(n)) {
      if (n >= 1_000_000_000) popEl.textContent = (n / 1_000_000_000).toFixed(1) + "B";
      else if (n >= 1_000_000) popEl.textContent = (n / 1_000_000).toFixed(1) + "M";
      else if (n >= 1_000)     popEl.textContent = (n / 1_000).toFixed(1) + "K";
      else                     popEl.textContent = String(n);
    }
  }

  // ── Render attractions ────────────────────────────────────
  const listEl = document.getElementById("attractions-list");
  if (listEl) {
    if (data.attractions && data.attractions.length > 0) {
      listEl.innerHTML = data.attractions.map((a) => {
        // category is comma-separated tags like "mountain_peaks,geological_formations"
        const tags = (a.category || "")
          .split(",")
          .map((t) => t.trim().replace(/_/g, " "))
          .filter(Boolean)
          .join(", ");

        return `
          <div class="dest-attraction-row">
            <span class="dest-attraction-name" title="${esc(a.name)}">${esc(a.name)}</span>
            <span class="dest-attraction-tags">${esc(tags)}</span>
          </div>`;
      }).join("");
    } else {
      listEl.innerHTML = '<div class="dest-empty">No attractions found nearby.</div>';
    }
  }

  // ── Wishlist toggle ───────────────────────────────────────
  const wishBtn  = document.getElementById("wishlist-btn");
  const feedback = document.getElementById("wishlist-feedback");

  function showFeedback(msg, isError) {
    if (!feedback) return;
    feedback.textContent = msg;
    feedback.style.color = isError ? "#c0392b" : "#27ae60";
    setTimeout(() => { feedback.textContent = ""; }, 3000);
  }

  if (wishBtn) {
    wishBtn.addEventListener("click", async function () {
      const isWishlisted = this.dataset.wishlisted === "true";
      const countryName  = this.dataset.country;

      try {
        if (!isWishlisted) {
          const resp = await fetch("/api/wishlist", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ country_name: countryName, note: "" }),
          });

          if (resp.status === 401) { window.location.href = "/login"; return; }
          if (!resp.ok) { showFeedback((await resp.text()) || "Failed to add", true); return; }

          const created = await resp.json();
          this.dataset.wishlisted = "true";
          this.dataset.wishlistId = created.id;
          this.textContent = "✓ Added to Wishlist";
          this.classList.add("added");
          showFeedback("Added to wishlist!", false);
        } else {
          const wid = this.dataset.wishlistId;
          if (wid) {
            const resp = await fetch("/api/wishlist/" + wid, { method: "DELETE" });
            if (!resp.ok && resp.status !== 204) { showFeedback("Failed to remove", true); return; }
          }
          this.dataset.wishlisted = "false";
          delete this.dataset.wishlistId;
          this.textContent = "+ Add to Wishlist";
          this.classList.remove("added");
          showFeedback("Removed from wishlist", false);
        }
      } catch {
        showFeedback("Network error", true);
      }
    });
  }

  // ── Helpers ───────────────────────────────────────────────
  function esc(str) {
    return String(str ?? "")
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;");
  }
})();