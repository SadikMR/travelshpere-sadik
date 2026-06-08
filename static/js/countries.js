(function () {
  "use strict";

  const countries = window.__COUNTRIES__ || [];
  const grid        = document.getElementById("country-grid");
  const searchInput = document.getElementById("search-input");
  const regionSelect = document.getElementById("region-select");
  const resultCount  = document.getElementById("result-count");

  // ── Populate region dropdown ──────────────────────────────
  const regions = [...new Set(countries.map((c) => c.region).filter(Boolean))].sort();
  regions.forEach((r) => {
    const opt = document.createElement("option");
    opt.value = r;
    opt.textContent = r;
    regionSelect.appendChild(opt);
  });

  // ── Helpers ───────────────────────────────────────────────
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

  // ── Card builder ──────────────────────────────────────────
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

  // ── Render ────────────────────────────────────────────────
  function render(list) {
    grid.innerHTML = list.length
      ? list.map(buildCard).join("")
      : `<div class="ts-empty">No countries match your search.</div>`;

    resultCount.innerHTML = list.length === countries.length
      ? `Showing all <strong>${countries.length}</strong> countries`
      : `Showing <strong>${list.length}</strong> of ${countries.length} countries`;
  }

  // ── Filter ────────────────────────────────────────────────
  function filter() {
    const query  = searchInput.value.trim().toLowerCase();
    const region = regionSelect.value;

    const result = countries.filter((c) => {
      const inRegion = !region || c.region === region;
      const inQuery  = !query  ||
        c.name.toLowerCase().includes(query) ||
        (c.capital && c.capital.toLowerCase().includes(query));
      return inRegion && inQuery;
    });

    render(result);
  }

  // ── Debounce ──────────────────────────────────────────────
  function debounce(fn, ms) {
    let t;
    return (...args) => { clearTimeout(t); t = setTimeout(() => fn(...args), ms); };
  }

  // ── Events ────────────────────────────────────────────────
  searchInput.addEventListener("input", debounce(filter, 250));
  regionSelect.addEventListener("change", filter);

  // ── Boot ──────────────────────────────────────────────────
  render(countries);
})();