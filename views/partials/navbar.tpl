{{define "partials/navbar.tpl"}}
<nav class="navbar">
  <div class="nav-container">

    <a class="nav-brand" href="/">TravelSphere</a>

    <button class="nav-toggle" id="nav-toggle" aria-label="Toggle navigation" aria-expanded="false">
      <span></span><span></span><span></span>
    </button>

    <ul class="nav-links" id="nav-links">
      <li><a href="/"          class="nav-link{{if eq .ActivePage "home"}}      active{{end}}">Home</a></li>
      <li><a href="/countries" class="nav-link{{if eq .ActivePage "countries"}} active{{end}}">Countries</a></li>
      <li><a href="/wishlist"  class="nav-link{{if eq .ActivePage "wishlist"}}  active{{end}}">Wishlist</a></li>
      <li><a href="/dashboard" class="nav-link{{if eq .ActivePage "dashboard"}} active{{end}}">Dashboard</a></li>
    </ul>

    <div class="nav-right">
      {{if .IsLoggedIn}}
        <a href="/logout" class="nav-login">Logout</a>
      {{else}}
        <a href="/login" class="nav-login">Login</a>
      {{end}}
    </div>

  </div>
</nav>
{{end}}