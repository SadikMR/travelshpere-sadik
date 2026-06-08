<div class="max-w-4xl mx-auto px-5 py-8">

  <!-- Hero Card -->
  <div class="bg-white border border-gray-200 rounded-2xl p-6 mb-6">

    <!-- Region badge -->
    <span class="inline-block text-xs font-bold uppercase tracking-widest text-blue-500 bg-blue-50 px-3 py-1 rounded-full mb-4">
      {{.Country.Region}}
    </span>

    <!-- Flag + Name -->
    <div class="flex items-center gap-6 mb-6">
      <img
        src="{{.Country.Flag}}"
        alt="Flag of {{.Country.Name}}"
        class="w-32 h-24 object-cover rounded-xl shadow-sm flex-shrink-0"
      />
      <div>
        <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight">{{.Country.Name}}</h1>
        <p class="text-sm text-gray-400 mt-1">{{.Country.SubRegion}}</p>
      </div>
    </div>

    <!-- Stats row -->
    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4 pt-4 border-t border-gray-100">

      <div class="flex flex-col gap-1">
        <span class="text-xs font-semibold uppercase tracking-widest text-gray-400">Capital</span>
        <span class="text-sm font-medium text-gray-800">{{.Country.Capital}}</span>
      </div>

      <div class="flex flex-col gap-1">
        <span class="text-xs font-semibold uppercase tracking-widest text-gray-400">Population</span>
        <span class="text-sm font-medium text-gray-800">
          {{.FormattedPopulation}}
        </span>
      </div>

      <div class="flex flex-col gap-1">
        <span class="text-xs font-semibold uppercase tracking-widest text-gray-400">Region</span>
        <span class="text-sm font-medium text-gray-800">{{.Country.Region}}</span>
      </div>

      <div class="flex flex-col gap-1">
        <span class="text-xs font-semibold uppercase tracking-widest text-gray-400">Currency</span>
        <span class="text-sm font-medium text-gray-800">{{.FormattedCurrency}}</span>
      </div>

      <div class="flex flex-col gap-1">
        <span class="text-xs font-semibold uppercase tracking-widest text-gray-400">Languages</span>
        <span class="text-sm font-medium text-gray-800">
          {{.FormattedLanguages}}
        </span>
      </div>

    </div>
  </div>

  <!-- Wishlist Button -->
  <div class="mb-8">
    <button
      id="wishlist-btn"
      data-country="{{.Country.Name}}"
      data-wishlisted="{{.IsWishlisted}}"
      data-wishlist-id="{{.WishlistID}}"
      class="px-5 py-2 rounded-lg border-2 text-sm font-semibold transition-all duration-150
             {{if .IsWishlisted}}bg-gray-800 text-white border-gray-800{{else}}border-gray-800 text-gray-800 hover:bg-gray-800 hover:text-white{{end}}"
    >
      {{if .IsWishlisted}}✓ Added to Wishlist{{else}}+ Add to Wishlist{{end}}
    </button>
    <div id="wishlist-feedback" class="mt-2 text-sm"></div>
  </div>

  <!-- Bottom Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

    <!-- Travel Weather -->
    <div class="bg-white border border-gray-200 rounded-2xl p-5">
      <h2 class="text-base font-bold text-gray-900 mb-4">Travel weather</h2>
      <div id="weather-content"
           class="text-sm text-gray-400 bg-gray-50 rounded-lg p-4 leading-relaxed">
        Weather data is optional. Add
        <code class="text-xs bg-gray-200 text-gray-700 px-1.5 py-0.5 rounded font-mono">WEATHER_API_KEY</code>
        to your
        <code class="text-xs bg-gray-200 text-gray-700 px-1.5 py-0.5 rounded font-mono">.env</code>
        file to enable live conditions.
      </div>
    </div>

    <!-- Attractions -->
    <div class="bg-white border border-gray-200 rounded-2xl p-5">
      <h2 class="text-base font-bold text-gray-900 mb-4">Attractions &amp; landmarks</h2>
      <div id="attractions-list" class="flex flex-col gap-2">
        <!-- Injected by JS -->
      </div>
    </div>

  </div>

</div>

<script>
  window.__COUNTRY_DETAIL__ = {
    name: "{{.Country.Name}}",
    isWishlisted: {{.IsWishlisted}},
    attractions: [
      {{range $i, $a := .Attractions}}{{if $i}},{{end}}{
        "name":     "{{$a.Name}}",
        "category": "{{$a.Category}}",
        "rate":     {{$a.Rate}},
        "distance": {{$a.Distance}},
        "lat":      {{$a.Lat}},
        "lon":      {{$a.Lon}}
      }{{end}}
    ]
  };
</script>
<script src="/static/js/destination.js"></script>