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

  // ── Autocomplete ────────────────────────────────────────
  const suggestions = document.getElementById("search-suggestions");

  function showSuggestions(countries) {
    if (!suggestions) return;

    if (!countries.length || !searchInput.value.trim()) {
      suggestions.classList.add("hidden");
      suggestions.innerHTML = "";
      return;
    }

    // Show top 8 matches
    const items = countries.slice(0, 8);
    suggestions.innerHTML = items.map((c) => `
      <a href="/countries/${encodeURIComponent(c.name)}"
         class="flex items-center gap-3 px-4 py-2 hover:bg-gray-50 cursor-pointer transition-colors"
         style="text-decoration:none;color:inherit;">
        <img src="${esc(c.flag)}" alt="" class="w-8 h-5 object-cover rounded" onerror="this.style.display='none'" />
        <div>
          <div class="text-sm font-medium text-gray-900">${esc(c.name)}</div>
          <div class="text-xs text-gray-400">${esc(c.capital || "")}${c.region ? " · " + esc(c.region) : ""}</div>
        </div>
      </a>
    `).join("");
    suggestions.classList.remove("hidden");
  }

  function autocomplete() {
    const query = searchInput.value.trim();
    if (!query) {
      showSuggestions([]);
      return;
    }

    fetch("/api/countries?search=" + encodeURIComponent(query))
      .then((r) => r.ok ? r.json() : [])
      .then((countries) => showSuggestions(countries))
      .catch(() => showSuggestions([]));
  }

  // Hide suggestions on outside click
  document.addEventListener("click", (e) => {
    if (suggestions && !suggestions.contains(e.target) && e.target !== searchInput) {
      suggestions.classList.add("hidden");
    }
  });

  // Show suggestions on focus if there's text
  searchInput.addEventListener("focus", () => {
    if (searchInput.value.trim()) autocomplete();
  });

  // ── Events ───────────────────────────────────────────────
  const debouncedFilter = debounce(filter, 300);
  const debouncedAutocomplete = debounce(autocomplete, 200);

  searchInput.addEventListener("input", () => {
    debouncedFilter();
    debouncedAutocomplete();
  });
  regionSelect.addEventListener("change", filter);

  // ── Boot: render initial SSR data ────────────────────────
  render(allCountries);
})();