<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="stylesheet" href="/static/css/destination.css">

<div class="dest-page">

  <!-- Hero card -->
  <div class="dest-hero">
    <span class="dest-region-badge">{{.Country.Region}}</span>

    <div class="dest-hero__top">
      <img
        class="dest-flag"
        src="{{.Country.Flag}}"
        alt="Flag of {{.Country.Name}}"
      />
      <div>
        <h1 class="dest-name">{{.Country.Name}}</h1>
        <p class="dest-subregion">{{.Country.SubRegion}}</p>
      </div>
    </div>

    <div class="dest-stats">
      <div>
        <span class="dest-stat__label">Capital</span>
        <span class="dest-stat__value">{{.Country.Capital}}</span>
      </div>
      <div>
        <span class="dest-stat__label">Population</span>
        <span class="dest-stat__value" id="js-population" data-pop="{{.Country.Population}}">{{.Country.Population}}</span>
      </div>
      <div>
        <span class="dest-stat__label">Region</span>
        <span class="dest-stat__value">{{.Country.Region}}</span>
      </div>
      <div>
        <span class="dest-stat__label">Currency</span>
        <span class="dest-stat__value">{{.Country.Currency}}</span>
      </div>
      <div>
        <span class="dest-stat__label">Languages</span>
        <span class="dest-stat__value">{{range $i, $l := .Country.Languages}}{{if $i}}, {{end}}{{$l}}{{end}}</span>
      </div>
    </div>
  </div>

  <!-- Wishlist button -->
  <div class="dest-wishlist-wrap">
    <button
      id="wishlist-btn"
      class="dest-wishlist-btn{{if .IsWishlisted}} added{{end}}"
      data-country="{{.Country.Name}}"
      data-wishlisted="{{.IsWishlisted}}"
    >
      {{if .IsWishlisted}}&#10003; Added to Wishlist{{else}}+ Add to Wishlist{{end}}
    </button>
    <span id="wishlist-feedback" class="dest-wishlist-feedback"></span>
  </div>

  <!-- Bottom panels -->
  <div class="dest-bottom">

    <div class="dest-panel">
  <h2 class="dest-panel__title">Travel Weather</h2>

  <div id="weather-content" class="dest-weather-note">
      <div><strong>Current Conditions</strong></div>
      <div>🌤️ Partly Cloudy</div>
      <div>Temperature: 28°C</div>
      <div>Humidity: 68%</div>
      <div>Wind: 12 km/h</div>
    </div>
  </div>
    <div class="dest-panel">
      <h2 class="dest-panel__title">Attractions &amp; landmarks</h2>
      <div id="attractions-list" class="dest-attractions">
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