
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;700;800&display=swap" rel="stylesheet">
<link rel="stylesheet" href="/static/css/dashboard.css">

<div class="db-page">

  <h1 class="db-title">Travel Dashboard</h1>
  <p class="db-subtitle">Your saved trips at a glance. Stats refresh automatically when your wishlist changes.</p>

  <!-- Stat cards -->
  <div class="db-stats" id="dashboard-stats">
    <div class="db-stat-card">
      <span class="db-stat-label">Total Saved</span>
      <span class="db-stat-value" id="wishlist-count">{{ .Summary.TotalWishlist }}</span>
    </div>
    <div class="db-stat-card">
      <span class="db-stat-label">Planned</span>
      <span class="db-stat-value" id="planned-count">{{ .Summary.PlannedTrips }}</span>
    </div>
    <div class="db-stat-card">
      <span class="db-stat-label">Visited</span>
      <span class="db-stat-value" id="visited-count">{{ .Summary.VisitedTrips }}</span>
    </div>
  </div>

  <!-- Destinations list -->
  <h2 class="db-section-title">Saved destinations</h2>

  {{ if .Destinations }}
  <div class="db-dest-list">
    {{ range .Destinations }}
    <div class="db-dest-row">
      <a href="/countries/{{ .CountryName }}" class="db-dest-country">{{ .CountryName }}</a>
      <span class="db-dest-sep">—</span>
      {{if eq (printf "%s" .Status) "Visited"}}
        <span class="db-dest-status db-dest-status--visited">{{.Status}}</span>
        {{else}}
        <span class="db-dest-status db-dest-status--planned">{{.Status}}</span>
      {{end}}
      <span class="db-dest-note">{{ .Note }}</span>
    </div>
    {{ end }}
  </div>
  {{ else }}
  <p class="db-empty">No saved destinations yet.</p>
  {{ end }}

</div>

<script>
(function() {
  function refreshDashboardSummary() {
    fetch('/api/dashboard/summary')
      .then(function(res) { return res.json(); })
      .then(function(data) {
        document.getElementById('wishlist-count').innerText = data.totalWishlist;
        document.getElementById('planned-count').innerText  = data.plannedTrips;
        document.getElementById('visited-count').innerText  = data.visitedTrips;
      })
      .catch(function(err) { console.error('Dashboard refresh failed:', err); });
  }

  setInterval(refreshDashboardSummary, 30000);
  window.refreshDashboardSummary = refreshDashboardSummary;
})();
</script>