(function () {
  "use strict";

  const data = window.__COUNTRY_DETAIL__;
  if (!data) return;

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
  const feedback = document.getElementById("wishlist-feedback");

  function showFeedback(msg, isError) {
    if (!feedback) return;
    feedback.textContent = msg;
    feedback.style.color = isError ? "#e53e3e" : "#38a169";
    setTimeout(() => { feedback.textContent = ""; }, 3000);
  }

  if (wishBtn) {
    wishBtn.addEventListener("click", async function () {
      const isWishlisted = this.dataset.wishlisted === "true";
      const countryName = this.dataset.country;

      try {
        if (!isWishlisted) {
          // Add to wishlist
          const resp = await fetch("/api/wishlist", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ country_name: countryName, note: "" }),
          });

          if (resp.status === 401) {
            window.location.href = "/login";
            return;
          }

          if (!resp.ok) {
            const text = await resp.text();
            showFeedback(text || "Failed to add", true);
            return;
          }

          const created = await resp.json();
          this.dataset.wishlisted = "true";
          this.dataset.wishlistId = created.id;
          this.textContent = "✓ Added to Wishlist";
          this.classList.add("bg-gray-800", "text-white");
          showFeedback("Added to wishlist!", false);
        } else {
          // Remove from wishlist
          const wid = this.dataset.wishlistId;
          if (wid) {
            const resp = await fetch("/api/wishlist/" + wid, {
              method: "DELETE",
            });
            if (!resp.ok && resp.status !== 204) {
              showFeedback("Failed to remove", true);
              return;
            }
          }

          this.dataset.wishlisted = "false";
          delete this.dataset.wishlistId;
          this.textContent = "+ Add to Wishlist";
          this.classList.remove("bg-gray-800", "text-white");
          showFeedback("Removed from wishlist", false);
        }
      } catch (err) {
        showFeedback("Network error", true);
      }
    });
  }
})();
