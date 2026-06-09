/**
 * TravelSphere — search.js
 *
 * Autocomplete dropdown (unchanged) + live grid filtering as extra.
 * - Dropdown: fetches /api/countries/search?q= with keyboard nav, selecting navigates to country page.
 * - Grid: as you type, /api/countries?search= filters .country-grid in place.
 * - Empty input restores original SSR grid.
 * - "From the left" match: results are already ordered by the API; grid preserves that order.
 */
(function () {
  "use strict";

  /* ── Elements ─────────────────────────────────────────── */
  var input    = document.getElementById("country-search");
  var dropdown = document.getElementById("search-results");
  if (!input || !dropdown) return;

  var grid      = document.querySelector(".country-grid");
  var gridTitle = grid ? grid.closest(".section").querySelector(".section-title") : null;
  var origHTML  = grid ? grid.innerHTML : "";
  var origTitle = gridTitle ? gridTitle.textContent : "";

  /* ── Autocomplete state ───────────────────────────────── */
  var acTimer     = null;
  var gridTimer   = null;
  var activeIdx   = -1;
  var suggestions = [];

  /* ── Debounce ─────────────────────────────────────────── */
  function debounce(fn, ms) {
    return function () { clearTimeout(arguments.callee._t); arguments.callee._t = setTimeout(fn, ms); };
  }

  /* ── AUTOCOMPLETE: fetch suggestions ─────────────────── */
  function fetchSuggestions() {
    var q = input.value.trim();
    if (!q) { suggestions = []; hideDropdown(); return; }

    fetch("/api/countries/search?q=" + encodeURIComponent(q), {
      headers: { "X-Requested-With": "XMLHttpRequest" }
    })
      .then(function (res) { if (!res.ok) throw new Error("HTTP " + res.status); return res.json(); })
      .then(function (data) {
        var all = Array.isArray(data) ? data : [];
        suggestions = prefixFilter(all, q, "name");
        renderDropdown();
      })
      .catch(function () { suggestions = []; hideDropdown(); });
  }

  /* ── AUTOCOMPLETE: render dropdown ───────────────────── */
  function renderDropdown() {
    dropdown.innerHTML = "";
    activeIdx = -1;

    if (!suggestions.length) {
      var empty = document.createElement("li");
      empty.className   = "dd-empty";
      empty.textContent = "No results found";
      dropdown.appendChild(empty);
      showDropdown();
      return;
    }

    suggestions.forEach(function (item, idx) {
      var li  = document.createElement("li");
      li.setAttribute("role", "option");
      li.setAttribute("data-idx", String(idx));

      var name = document.createElement("span");
      name.className   = "dd-name";
      name.textContent = item.name;

      var cap = document.createElement("span");
      cap.className   = "dd-capital";
      cap.textContent = item.capital ? "— " + item.capital : "";

      li.appendChild(name);
      li.appendChild(cap);

      li.addEventListener("mousedown", function (e) {
        e.preventDefault();
        selectItem(item);         // navigate — unchanged behaviour
      });

      dropdown.appendChild(li);
    });

    showDropdown();
  }

  /* ── AUTOCOMPLETE: select → navigate to country page ─── */
  function selectItem(item) {
    input.value = item.name;
    hideDropdown();
    window.location.href = "/countries/" + (item.slug || encodeURIComponent(item.name));
  }

  /* ── AUTOCOMPLETE: keyboard nav ──────────────────────── */
  input.addEventListener("keydown", function (e) {
    var items = dropdown.querySelectorAll("li:not(.dd-empty)");
    if (!items.length || dropdown.classList.contains("hidden")) return;

    if (e.key === "ArrowDown") {
      e.preventDefault();
      activeIdx = Math.min(activeIdx + 1, items.length - 1);
      updateHighlight(items);
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      activeIdx = Math.max(activeIdx - 1, 0);
      updateHighlight(items);
    } else if (e.key === "Enter") {
      e.preventDefault();
      if (activeIdx >= 0 && suggestions[activeIdx]) selectItem(suggestions[activeIdx]);
    } else if (e.key === "Escape") {
      hideDropdown();
      input.blur();
    }
  });

  function updateHighlight(items) {
    items.forEach(function (li, i) {
      li.classList.toggle("active", i === activeIdx);
      if (i === activeIdx) li.scrollIntoView({ block: "nearest" });
    });
  }

  /* ── GRID: build a card matching SSR markup ───────────── */
  function buildCard(c) {
    var a = document.createElement("a");
    a.className = "country-card";
    a.href = "/countries/" + encodeURIComponent(c.name || "");

    var flagDiv = document.createElement("div");
    flagDiv.className = "country-flag";

    var img = document.createElement("img");
    img.src     = c.flag || "";
    img.alt     = "Flag of " + (c.name || "");
    img.loading = "lazy";
    img.onerror = function () { this.style.background = "#e5e2da"; this.removeAttribute("src"); };
    flagDiv.appendChild(img);

    var info = document.createElement("div");
    info.className = "country-info";

    var nameSpan = document.createElement("span");
    nameSpan.className   = "country-name";
    nameSpan.textContent = c.name || "";

    var meta = document.createElement("span");
    meta.className   = "country-meta";
    meta.textContent = (c.capital || "") + (c.region ? " · " + c.region : "");

    info.appendChild(nameSpan);
    info.appendChild(meta);
    a.appendChild(flagDiv);
    a.appendChild(info);
    return a;
  }

  /* ── GRID: render results ─────────────────────────────── */
  function renderGrid(countries) {
    if (!grid) return;
    grid.innerHTML = "";

    if (!countries || !countries.length) {
      var p = document.createElement("p");
      p.className   = "empty-state";
      p.textContent = "No countries match your search.";
      grid.appendChild(p);
    } else {
      countries.forEach(function (c) { grid.appendChild(buildCard(c)); });
    }

    if (gridTitle) gridTitle.textContent = "Search results";
  }

  /* ── GRID: restore SSR ────────────────────────────────── */
  function restoreGrid() {
    if (!grid) return;
    grid.innerHTML = origHTML;
    if (gridTitle) gridTitle.textContent = origTitle;
  }

  /* ── Prefix filter: only names starting with query ──── */
  function prefixFilter(list, q, key) {
    var lower = q.toLowerCase();
    return list.filter(function (item) {
      return (item[key] || "").toLowerCase().indexOf(lower) === 0;
    });
  }

  /* ── GRID: fetch and render ───────────────────────────── */
  function updateGrid() {
    var q = input.value.trim();
    if (!q) { restoreGrid(); return; }

    fetch("/api/countries?search=" + encodeURIComponent(q))
      .then(function (r) { return r.ok ? r.json() : []; })
      .then(function (data) {
        var filtered = prefixFilter(Array.isArray(data) ? data : [], q, "name");
        renderGrid(filtered);
      })
      .catch(function () { renderGrid([]); });
  }

  /* ── Input events ─────────────────────────────────────── */
  input.addEventListener("input", function () {
    // Autocomplete dropdown — 300ms debounce
    clearTimeout(acTimer);
    acTimer = setTimeout(fetchSuggestions, 300);

    // Grid update — 300ms debounce
    clearTimeout(gridTimer);
    gridTimer = setTimeout(updateGrid, 300);
  });

  input.addEventListener("focus", function () {
    if (suggestions.length > 0) showDropdown();
  });

  document.addEventListener("click", function (e) {
    if (!input.contains(e.target) && !dropdown.contains(e.target)) hideDropdown();
  });

  /* ── Show / hide ──────────────────────────────────────── */
  function showDropdown() { dropdown.classList.remove("hidden"); }
  function hideDropdown()  { dropdown.classList.add("hidden"); activeIdx = -1; }

}());