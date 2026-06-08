(function () {
  "use strict";

  const grid         = document.getElementById("country-results");
  const searchInput  = document.getElementById("search-input");
  const regionSelect = document.getElementById("region-select");
  const resultCount  = document.getElementById("result-count");

  // ── Initial data from SSR ────────────────────────────────
  let allCountries = window.__COUNTRIES__ || [];

  // ── Populate region dropdown ─────────────────────────────
  const regions = [...new Set(allCountries.map((c) => c.region).filter(Boolean))].sort();
  regions.forEach((r) => {
    const opt = document.createElement("option");
    opt.value = r;
    opt.textContent = r;
    regionSelect.appendChild(opt);
  });

  // ── Helpers ──────────────────────────────────────────────
  function formatPop(n) {
    if (n >= 1_000_000_000) return (n / 1_000_000_000).toFixed(1) + "B";
    if (n >= 1_000_000)     return (n / 1_000_000).toFixed(1) + "M";
    if (n >= 1_000)         return (n / 1_000).toFixed(1) + "K";
    return String(n);
  }

  function esc(str) {
    return String(str ?? "")
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;");
  }

  // ── Card builder ─────────────────────────────────────────
  function buildCard(c) {
    const capital  = esc(c.capital  || "—");
    const currency = esc(c.currency || "—");
    const langs    = esc((c.languages || []).join(", ") || "—");
    const pop      = c.population ? formatPop(c.population) : "—";
    const slug     = encodeURIComponent(c.name);

    return `
      <a href="/countries/${slug}" class="ts-card-link" style="text-decoration:none;color:inherit;display:block;">
      <div class="ts-card">
        <img
          class="ts-card__flag"
          src="${esc(c.flag)}"
          alt="Flag of ${esc(c.name)}"
          loading="lazy"
          onerror="this.style.background='#e5e2da';this.removeAttribute('src')"
        />
        <div class="ts-card__body">
          <div class="ts-card__name" title="${esc(c.name)}">${esc(c.name)}</div>
          <div class="ts-card__meta">
            <div><b>Capital:</b> ${capital}</div>
            <div><b>Population:</b> ${pop}</div>
            <div><b>Currency:</b> ${currency}</div>
            <div><b>Languages:</b> ${langs}</div>
          </div>
        </div>
      </div>
      </a>`;
  }

  // ── Render into #country-results only ────────────────────
  function render(list) {
    grid.innerHTML = list.length
      ? list.map(buildCard).join("")
      : `<div class="ts-empty">No countries match your search.</div>`;

    resultCount.innerHTML = `Showing <strong>${list.length}</strong> countries`;
  }

  // ── AJAX filter via GET /api/countries ────────────────────
  function filter() {
    const query  = searchInput.value.trim();
    const region = regionSelect.value;

    const params = new URLSearchParams();
    if (query)  params.set("search", query);
    if (region) params.set("region", region);

    // Show loading state
    grid.innerHTML = `<div class="ts-empty" style="color:#aaa;">Loading...</div>`;

    fetch("/api/countries?" + params.toString())
      .then((resp) => {
        if (!resp.ok) throw new Error("API error");
        return resp.json();
      })
      .then((countries) => {
        render(countries);
      })
      .catch((err) => {
        console.error("Country search failed:", err);
        grid.innerHTML = `<div class="ts-empty" style="color:#e53e3e;">Failed to load countries. Please try again.</div>`;
      });
  }

  // ── Debounce ─────────────────────────────────────────────
  function debounce(fn, ms) {
    let t;
    return (...args) => { clearTimeout(t); t = setTimeout(() => fn(...args), ms); };
  }

  // ── Events ───────────────────────────────────────────────
  searchInput.addEventListener("input", debounce(filter, 300));
  regionSelect.addEventListener("change", filter);

  // ── Boot: render initial SSR data ────────────────────────
  render(allCountries);
})();