(function () {
  "use strict";

  const grid         = document.getElementById("country-results");
  const searchInput  = document.getElementById("search-input");
  const regionSelect = document.getElementById("region-select");
  const resultCount  = document.getElementById("result-count");
  const suggestions  = document.getElementById("search-suggestions");
  const paginationEl = document.getElementById("pagination-controls");

  let allCountries = window.__COUNTRIES__ || [];

  // ── Pagination state ────────────────────────────────
  const PER_PAGE = 12;
  let currentPage = 1;
  let currentList = allCountries;

  // ── Autocomplete state ───────────────────────────────
  let acTimer    = null;
  let activeIdx  = -1;
  let acItems    = [];   // last suggestion list

  // ── Region dropdown ──────────────────────────────────
  const regions = [...new Set(allCountries.map((c) => c.region).filter(Boolean))].sort();
  regions.forEach((r) => {
    const opt = document.createElement("option");
    opt.value = r;
    opt.textContent = r;
    regionSelect.appendChild(opt);
  });

  // ── Helpers ──────────────────────────────────────────
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

  function debounce(fn, ms) {
    let t;
    return (...args) => { clearTimeout(t); t = setTimeout(() => fn(...args), ms); };
  }

  // ── Card builder ─────────────────────────────────────
  function buildCard(c) {
    const capital  = esc(c.capital   || "—");
    const currency = esc(c.currency  || "—");
    const langs    = esc((c.languages || []).join(", ") || "—");
    const pop      = c.population ? formatPop(c.population) : "—";
    const slug     = encodeURIComponent(c.name);

    return `
      <a href="/countries/${slug}" class="ts-card-link">
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

  // ── Pagination helpers ──────────────────────────────
  function totalPages(list) {
    return Math.max(1, Math.ceil(list.length / PER_PAGE));
  }

  function getPageSlice(list, page) {
    const start = (page - 1) * PER_PAGE;
    return list.slice(start, start + PER_PAGE);
  }

  function renderPagination(list) {
    if (!paginationEl) return;
    const total = totalPages(list);
    if (total <= 1) {
      paginationEl.innerHTML = "";
      return;
    }

    let html = "";

    // Prev button
    html += `<button class="ts-page-btn${currentPage === 1 ? " disabled" : ""}" data-page="${currentPage - 1}" ${currentPage === 1 ? "disabled" : ""}>
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"/></svg>
      Prev
    </button>`;

    // Page numbers
    const pages = buildPageNumbers(currentPage, total);
    pages.forEach((p) => {
      if (p === "…") {
        html += `<span class="ts-page-ellipsis">…</span>`;
      } else {
        html += `<button class="ts-page-num${p === currentPage ? " active" : ""}" data-page="${p}">${p}</button>`;
      }
    });

    // Next button
    html += `<button class="ts-page-btn${currentPage === total ? " disabled" : ""}" data-page="${currentPage + 1}" ${currentPage === total ? "disabled" : ""}>
      Next
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"/></svg>
    </button>`;

    paginationEl.innerHTML = html;

    // Bind click handlers
    paginationEl.querySelectorAll("[data-page]").forEach((btn) => {
      btn.addEventListener("click", () => {
        const p = parseInt(btn.dataset.page, 10);
        if (p >= 1 && p <= total && p !== currentPage) {
          currentPage = p;
          renderPage(currentList);
          // Smooth scroll to top of grid
          grid.scrollIntoView({ behavior: "smooth", block: "start" });
        }
      });
    });
  }

  // Build a smart page number array like [1, 2, "…", 5, 6, 7, "…", 10]
  function buildPageNumbers(current, total) {
    if (total <= 7) {
      return Array.from({ length: total }, (_, i) => i + 1);
    }
    const pages = new Set([1, 2, total - 1, total]);
    for (let i = current - 1; i <= current + 1; i++) {
      if (i >= 1 && i <= total) pages.add(i);
    }
    const sorted = [...pages].sort((a, b) => a - b);
    const result = [];
    for (let i = 0; i < sorted.length; i++) {
      if (i > 0 && sorted[i] - sorted[i - 1] > 1) result.push("…");
      result.push(sorted[i]);
    }
    return result;
  }

  // ── Render grid + pagination ────────────────────────
  function renderPage(list) {
    const pageItems = getPageSlice(list, currentPage);
    grid.innerHTML = pageItems.length
      ? pageItems.map(buildCard).join("")
      : `<div class="ts-empty">No countries match your search.</div>`;

    const start = (currentPage - 1) * PER_PAGE + 1;
    const end = Math.min(currentPage * PER_PAGE, list.length);
    if (list.length > 0) {
      resultCount.innerHTML = `Showing <strong>${start}–${end}</strong> of <strong>${list.length}</strong> countries`;
    } else {
      resultCount.innerHTML = `Showing <strong>0</strong> countries`;
    }

    renderPagination(list);
  }

  function render(list) {
    currentList = list;
    currentPage = 1;
    renderPage(list);
  }

  // ── AJAX grid filter ──────────────────────────────────
  function filter() {
    const query  = searchInput.value.trim();
    const region = regionSelect.value;
    const params = new URLSearchParams();
    if (query)  params.set("search", query);
    if (region) params.set("region", region);

    grid.innerHTML = `<div class="ts-empty" style="color:#aaa;">Loading…</div>`;
    if (paginationEl) paginationEl.innerHTML = "";

    fetch("/api/countries?" + params.toString())
      .then((r) => { if (!r.ok) throw new Error("API error"); return r.json(); })
      .then(render)
      .catch(() => {
        grid.innerHTML = `<div class="ts-empty" style="color:#c0392b;">Failed to load. Please try again.</div>`;
      });
  }

  const debouncedFilter = debounce(filter, 300);

  // ── Autocomplete ─────────────────────────────────────
  function hideSuggestions() {
    suggestions.classList.add("hidden");
    suggestions.innerHTML = "";
    activeIdx = -1;
    acItems = [];
  }

  function highlightItems() {
    const lis = suggestions.querySelectorAll("li:not(.dd-empty)");
    lis.forEach((li, i) => {
      li.classList.toggle("active", i === activeIdx);
      if (i === activeIdx) li.scrollIntoView({ block: "nearest" });
    });
  }

  function renderSuggestions(countries) {
    suggestions.innerHTML = "";
    activeIdx = -1;
    acItems = countries;

    if (!countries.length) {
      const li = document.createElement("li");
      li.className = "dd-empty";
      li.textContent = "No results found";
      suggestions.appendChild(li);
      suggestions.classList.remove("hidden");
      return;
    }

    countries.forEach((c, idx) => {
      const li = document.createElement("li");
      li.setAttribute("role", "option");
      li.setAttribute("data-idx", String(idx));

      const img = document.createElement("img");
      img.src = c.flag || "";
      img.alt = "";
      img.onerror = function () { this.style.display = "none"; };

      const name = document.createElement("span");
      name.className = "dd-name";
      name.textContent = c.name;

      const cap = document.createElement("span");
      cap.className = "dd-capital";
      if (c.capital) cap.textContent = c.capital + (c.region ? " · " + c.region : "");

      li.appendChild(img);
      li.appendChild(name);
      li.appendChild(cap);

      // mousedown before blur so input keeps focus
      li.addEventListener("mousedown", (e) => {
        e.preventDefault();
        selectSuggestion(c);
      });

      suggestions.appendChild(li);
    });

    suggestions.classList.remove("hidden");
  }

  function selectSuggestion(c) {
    searchInput.value = c.name;
    hideSuggestions();
    window.location.href = "/countries/" + encodeURIComponent(c.name);
  }

  function fetchSuggestions() {
    const q = searchInput.value.trim();
    if (!q) { hideSuggestions(); return; }

    fetch("/api/countries/search?q=" + encodeURIComponent(q), {
      headers: { "X-Requested-With": "XMLHttpRequest" }
    })
      .then((r) => r.ok ? r.json() : [])
      .then((data) => renderSuggestions(Array.isArray(data) ? data.slice(0, 8) : []))
      .catch(() => hideSuggestions());
  }

  const debouncedAC = debounce(fetchSuggestions, 200);

  // ── Keyboard navigation ───────────────────────────────
  searchInput.addEventListener("keydown", (e) => {
    const lis = suggestions.querySelectorAll("li:not(.dd-empty)");
    const open = !suggestions.classList.contains("hidden");

    if (e.key === "ArrowDown") {
      e.preventDefault();
      if (!open && searchInput.value.trim()) { fetchSuggestions(); return; }
      activeIdx = Math.min(activeIdx + 1, lis.length - 1);
      highlightItems();
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      activeIdx = Math.max(activeIdx - 1, 0);
      highlightItems();
    } else if (e.key === "Enter") {
      e.preventDefault();
      if (open && activeIdx >= 0 && acItems[activeIdx]) {
        selectSuggestion(acItems[activeIdx]);
      }
    } else if (e.key === "Escape") {
      hideSuggestions();
      searchInput.blur();
    }
  });

  // ── Input events ──────────────────────────────────────
  searchInput.addEventListener("input", () => {
    debouncedFilter();
    debouncedAC();
  });

  searchInput.addEventListener("focus", () => {
    if (searchInput.value.trim() && acItems.length) {
      suggestions.classList.remove("hidden");
    }
  });

  regionSelect.addEventListener("change", filter);

  document.addEventListener("click", (e) => {
    if (!suggestions.contains(e.target) && e.target !== searchInput) {
      hideSuggestions();
    }
  });

  // ── Boot ─────────────────────────────────────────────
  render(allCountries);
})();