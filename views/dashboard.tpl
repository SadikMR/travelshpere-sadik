<h1>Travel Dashboard</h1>

<div class="row">

    <div class="card">
        <h3>Total Wishlist</h3>
        <p id="wishlist-count">
            {{ .Summary.TotalWishlist }}
        </p>
    </div>

    <div class="card">
        <h3>Planned Trips</h3>
        <p id="planned-count">
            {{ .Summary.PlannedTrips }}
        </p>
    </div>

    <div class="card">
        <h3>Visited Destinations</h3>
        <p id="visited-count">
            {{ .Summary.VisitedTrips }}
        </p>
    </div>

</div>

<hr>

<h2>Saved Destinations</h2>

<table border="1">

    <tr>
        <th>Country</th>
        <th>Note</th>
        <th>Status</th>
        <th>Added</th>
    </tr>

    {{ range .Destinations }}
    <tr>
        <td>{{ .CountryName }}</td>
        <td>{{ .Note }}</td>
        <td>{{ .Status }}</td>
        <td>{{ .CreatedAt }}</td>
    </tr>
    {{ end }}

</table>

<script>
function refreshDashboardSummary() {
    fetch('/api/dashboard/summary')
        .then(res => res.json())
        .then(data => {
            document.getElementById('wishlist-count').innerText =
                data.totalWishlist;

            document.getElementById('planned-count').innerText =
                data.plannedTrips;

            document.getElementById('visited-count').innerText =
                data.visitedTrips;
        });
}
</script>