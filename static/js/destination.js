(function () {
  "use strict";

  const data = window.__COUNTRY_DETAIL__;
  if (!data) return;

  // ── Format population ──────────────────────────────────────
  const popEl = document.getElementById("js-population");
  if (popEl) {
    const n = parseInt(popEl.dataset.pop, 10);
    if (!isNaN(n)) {
      let formatted;
      if (n >= 1_000_000_000) formatted = (n / 1_000_000_000).toFixed(1) + "B";
      else if (n >= 1_000_000) formatted = (n / 1_000_000).toFixed(1) + "M";
      else if (n >= 1_000) formatted = (n / 1_000).toFixed(1) + "K";
      else formatted = String(n);
      popEl.textContent = formatted;
    }
  }

  // ── Render attractions ─────────────────────────────────────
  const listEl = document.getElementById("attractions-list");
  if (listEl && data.attractions && data.attractions.length > 0) {
    listEl.innerHTML = data.attractions
      .map(
        (a) => `
      <div style="display:flex;justify-content:space-between;align-items:center;padding:8px 12px;background:#f9fafb;border-radius:8px;">
        <div>
          <div style="font-weight:600;font-size:14px;color:#111;">${a.name}</div>
          <div style="font-size:12px;color:#888;">${a.category}</div>
        </div>
        <div style="text-align:right;">
          <div style="font-size:13px;font-weight:500;color:#111;">⭐ ${a.rate}</div>
          <div style="font-size:11px;color:#999;">${(a.distance / 1000).toFixed(1)} km</div>
        </div>
      </div>`
      )
      .join("");
  } else if (listEl) {
    listEl.innerHTML =
      '<div style="color:#aaa;font-size:14px;background:#f9fafb;border-radius:8px;padding:16px;">No attractions found nearby.</div>';
  }

  // ── Wishlist toggle ────────────────────────────────────────
  const wishBtn = document.getElementById("wishlist-btn");
  if (wishBtn) {
    wishBtn.addEventListener("click", function () {
      const current = this.dataset.wishlisted === "true";
      const next = !current;
      this.dataset.wishlisted = String(next);
      this.textContent = next ? "✓ Added to Wishlist" : "+ Add to Wishlist";
      if (next) {
        this.classList.add("bg-gray-800", "text-white");
      } else {
        this.classList.remove("bg-gray-800", "text-white");
      }
    });
  }
})();
