<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
<link rel="stylesheet" href="/static/css/countries.css">

<div class="ts-page">

  <!-- Header -->
  <div class="mb-8">
    <h1 class="ts-title">Country Explorer</h1>
    <p class="ts-subtitle mt-2">
      Browse every destination on first load. Search and filter update only the results
      below — no full page reload.
    </p>
  </div>

  <!-- Filter Bar -->
  <div class="ts-filter-bar mb-7">
    <div class="flex flex-wrap gap-5 items-end">
      <div class="flex-1 min-w-48 relative">
        <label class="ts-label" for="search-input">Search</label>
        <input
          id="search-input"
          class="ts-input"
          type="text"
          placeholder="Country or capital..."
          autocomplete="off"
        />
        <div id="search-suggestions" class="absolute left-0 right-0 top-full z-50 bg-white border border-gray-200 rounded-lg shadow-lg mt-1 max-h-60 overflow-y-auto hidden"></div>
      </div>
      <div class="min-w-44">
        <label class="ts-label" for="region-select">Region</label>
        <select id="region-select" class="ts-select">
          <option value="">All regions</option>
        </select>
      </div>
    </div>
  </div>

  <!-- Meta row -->
  <div class="flex items-center justify-between flex-wrap gap-2 mb-5">
    <p id="result-count" class="ts-result-count"></p>
  </div>

  <!-- Grid -->
  <div id="country-results" class="ts-grid">
    <!-- Cards injected by JS -->
  </div>

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
