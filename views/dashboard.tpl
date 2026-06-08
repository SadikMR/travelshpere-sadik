<div class="max-w-4xl mx-auto px-4 py-8">

<h1 class="text-2xl font-bold text-gray-900 mb-6">Travel Dashboard</h1>

<div id="dashboard-stats">
  <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-8">

    <div class="bg-white border border-gray-200 rounded-xl p-5 text-center">
      <h3 class="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-1">Total Wishlist</h3>
      <p id="wishlist-count" class="text-3xl font-bold text-gray-900">
        {{ .Summary.TotalWishlist }}
      </p>
    </div>

    <div class="bg-white border border-gray-200 rounded-xl p-5 text-center">
      <h3 class="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-1">Planned Trips</h3>
      <p id="planned-count" class="text-3xl font-bold text-blue-600">
        {{ .Summary.PlannedTrips }}
      </p>
    </div>

    <div class="bg-white border border-gray-200 rounded-xl p-5 text-center">
      <h3 class="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-1">Visited Destinations</h3>
      <p id="visited-count" class="text-3xl font-bold text-green-600">
        {{ .Summary.VisitedTrips }}
      </p>
    </div>

  </div>
</div>

<h2 class="text-xl font-bold text-gray-900 mb-4">Saved Destinations</h2>

{{ if .Destinations }}
<table class="w-full border-collapse bg-white rounded-lg overflow-hidden shadow-sm">
  <thead>
    <tr class="bg-gray-50 text-left text-xs font-semibold uppercase tracking-wider text-gray-500">
      <th class="px-4 py-3">Country</th>
      <th class="px-4 py-3">Note</th>
      <th class="px-4 py-3">Status</th>
      <th class="px-4 py-3">Added</th>
    </tr>
  </thead>
  <tbody>
    {{ range .Destinations }}
    <tr class="border-t border-gray-100">
      <td class="px-4 py-3 font-medium">
        <a href="/countries/{{ .CountryName }}" class="text-blue-600 hover:underline">{{ .CountryName }}</a>
      </td>
      <td class="px-4 py-3 text-sm text-gray-600">{{ .Note }}</td>
      <td class="px-4 py-3 text-sm">{{ .Status }}</td>
      <td class="px-4 py-3 text-sm text-gray-400">{{ .CreatedAt.Format "Jan 02, 2006" }}</td>
    </tr>
    {{ end }}
  </tbody>
</table>
{{ else }}
<p class="text-gray-400">No saved destinations yet.</p>
{{ end }}

</div>

<script>
(function() {
  // Refresh dashboard stats from API (AJAX, updates only #dashboard-stats counters)
  function refreshDashboardSummary() {
    fetch('/api/dashboard/summary')
      .then(function(res) { return res.json(); })
      .then(function(data) {
        document.getElementById('wishlist-count').innerText = data.totalWishlist;
        document.getElementById('planned-count').innerText = data.plannedTrips;
        document.getElementById('visited-count').innerText = data.visitedTrips;
      })
      .catch(function(err) { console.error('Dashboard refresh failed:', err); });
  }

  // Auto-refresh stats every 30 seconds
  setInterval(refreshDashboardSummary, 30000);

  // Expose for manual calls
  window.refreshDashboardSummary = refreshDashboardSummary;
})();
</script>