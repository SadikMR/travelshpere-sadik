<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
<link rel="stylesheet" href="/static/css/countries.css">

<div class="ts-page">

  <div style="margin-bottom:1.5rem;">
    <h1 class="ts-title">Country Explorer</h1>
    <p class="ts-subtitle">Browse every destination. Search and filter update only the results below.</p>
  </div>

  <!-- Filter bar: ALWAYS side-by-side via inline flex -->
  <div class="ts-filter-bar" style="margin-bottom:1.25rem; display:flex; flex-wrap:wrap; gap:1rem; align-items:flex-end;">
    <div style="flex:1; min-width:160px; position:relative;">
      <label class="ts-label" for="search-input">Search</label>
      <input
        id="search-input"
        class="ts-input"
        type="text"
        placeholder="Country or capital..."
        autocomplete="off"
        aria-autocomplete="list"
        aria-haspopup="listbox"
        aria-controls="search-suggestions"
      />
      <ul id="search-suggestions" class="ts-suggestions hidden" role="listbox"></ul>
    </div>
    <div style="min-width:160px; width:200px;">
      <label class="ts-label" for="region-select">Region</label>
      <select id="region-select" class="ts-select">
        <option value="">All regions</option>
      </select>
    </div>
  </div>

  <div style="margin-bottom:1rem;">
    <p id="result-count" class="ts-result-count"></p>
  </div>

  <div id="country-results" class="ts-grid"></div>

  <div id="pagination-controls" class="ts-pagination"></div>

</div>

<script>
  window.__COUNTRIES__ = [
    {{range $i, $c := .Countries}}{{if $i}},{{end}}{
      "name":       "{{$c.Name}}",
      "capital":    "{{$c.Capital}}",
      "currency":   "{{$c.Currency}}",
      "languages":  [{{range $j, $l := $c.Languages}}{{if $j}},"{{$l}}"{{else}}"{{$l}}"{{end}}{{end}}],
      "population": {{$c.Population}},
      "flag":       "{{$c.Flag}}",
      "region":     "{{$c.Region}}"
    }{{end}}
  ];
</script>
<script src="/static/js/countries.js"></script>