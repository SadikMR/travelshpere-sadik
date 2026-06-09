/**
 * TravelSphere — search.js
 *
 * AJAX autocomplete for the home page country search.
 * - Initial page load is SSR (no JS needed for content).
 * - Typing in the search box calls /api/countries/search?q=
 *   and updates #search-results ONLY — zero full-page reloads.
 * - 300 ms debounce on every keystroke.
 * - Full keyboard navigation (↑ ↓ Enter Escape).
 */
(function () {
  "use strict";

  /* ── Elements ─────────────────────────────────────────── */
  var input    = document.getElementById("country-search");
  var dropdown = document.getElementById("search-results");
  if (!input || !dropdown) return;          // not on home page

  /* ── State ────────────────────────────────────────────── */
  var timer       = null;
  var activeIdx   = -1;
  var suggestions = [];                     // last successful response

  /* ── Debounce ─────────────────────────────────────────── */
  function debounce(fn, ms) {
    return function () {
      clearTimeout(timer);
      timer = setTimeout(fn, ms);
    };
  }

  /* ── Fetch suggestions from server ───────────────────── */
  function fetchSuggestions() {
    var q = input.value.trim();

    if (q.length === 0) {
      suggestions = [];
      hideDropdown();
      return;
    }

    fetch("/api/countries/search?q=" + encodeURIComponent(q), {
      headers: { "X-Requested-With": "XMLHttpRequest" }
    })
      .then(function (res) {
        if (!res.ok) throw new Error("HTTP " + res.status);
        return res.json();
      })
      .then(function (data) {
        suggestions = Array.isArray(data) ? data : [];
        renderDropdown();
      })
      .catch(function (err) {
        console.warn("Search error:", err);
        suggestions = [];
        hideDropdown();
      });
  }

  var debouncedFetch = debounce(fetchSuggestions, 300);

  /* ── Render dropdown items ────────────────────────────── */
  function renderDropdown() {
    dropdown.innerHTML = "";
    activeIdx = -1;

    if (suggestions.length === 0) {
      var empty = document.createElement("li");
      empty.className = "dd-empty";
      empty.textContent = "No results found";
      dropdown.appendChild(empty);
      showDropdown();
      return;
    }

    suggestions.forEach(function (item, idx) {
      var li = document.createElement("li");
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

      /* mousedown fires before blur so the input keeps focus */
      li.addEventListener("mousedown", function (e) {
        e.preventDefault();
        selectItem(item);
      });

      dropdown.appendChild(li);
    });

    showDropdown();
  }

  /* ── Navigate to country page ─────────────────────────── */
  function selectItem(item) {
    input.value = item.name;
    hideDropdown();
    window.location.href = "/countries/" + item.slug;
  }

  /* ── Keyboard navigation ──────────────────────────────── */
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
      if (activeIdx >= 0 && suggestions[activeIdx]) {
        selectItem(suggestions[activeIdx]);
      }
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

  /* ── Input events ─────────────────────────────────────── */
  input.addEventListener("input", debouncedFetch);

  input.addEventListener("focus", function () {
    if (suggestions.length > 0) showDropdown();
  });

  /* close when clicking anywhere outside the search box */
  document.addEventListener("click", function (e) {
    if (!input.contains(e.target) && !dropdown.contains(e.target)) {
      hideDropdown();
    }
  });

  /* ── Show / hide helpers ──────────────────────────────── */
  function showDropdown() { dropdown.classList.remove("hidden"); }
  function hideDropdown()  { dropdown.classList.add("hidden"); activeIdx = -1; }

}());