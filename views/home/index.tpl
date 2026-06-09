<div class="hero">
  <div class="hero-card">
    <h1 class="hero-title">Discover your next destination</h1>
    <p class="hero-sub">Search countries, explore attractions, and curate your personal travel wishlist.</p>

    <div class="search-wrapper">
      <label class="search-label" for="country-search">WHERE TO NEXT?</label>
      <div class="search-box">
        <input
          id="country-search"
          class="search-input"
          type="text"
          placeholder="Search countries..."
          autocomplete="off"
          aria-autocomplete="list"
          aria-haspopup="listbox"
          aria-controls="search-results"
        />
        <ul id="search-results" class="search-dropdown hidden" role="listbox"></ul>
      </div>
    </div>
  </div>
</div>

<main class="page-content">

  <section class="section">
    <h2 class="section-title">Featured destinations</h2>
    {{if .FeaturedCountries}}
    <div class="country-grid">
      {{range .FeaturedCountries}}
      <a class="country-card" href="/countries/{{.Name}}">
        <div class="country-flag">
          <img src="{{.Flag}}" alt="Flag of {{.Name}}" loading="lazy" />
        </div>
        <div class="country-info">
          <span class="country-name">{{.Name}}</span>
          <span class="country-meta">{{.Capital}}{{if .Region}} · {{.Region}}{{end}}</span>
        </div>
      </a>
      {{end}}
    </div>
    {{else}}
    <p class="empty-state">No featured destinations available right now.</p>
    {{end}}
  </section>

  <section class="section">
    <h2 class="section-title">Popular attractions</h2>
    {{if .PopularAttractions}}
    <ul class="attraction-list">
      {{range .PopularAttractions}}
      <li class="attraction-item">
        <a href="#" class="attraction-link">
          <span class="attraction-name">{{.Name}}</span>
          {{if .Category}}<span class="attraction-tag">{{.Category}}</span>{{end}}
          {{if gt .Rate 0}}<span class="attraction-tag attraction-tag--rate">★ {{.Rate}}</span>{{end}}
        </a>
      </li>
      {{end}}
    </ul>
    {{else}}
    <p class="empty-state">No attractions available right now.</p>
    {{end}}
  </section>

</main>

<script src="/static/js/search.js"></script>