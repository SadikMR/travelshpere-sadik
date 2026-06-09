
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
<link rel="stylesheet" href="/static/css/wishlist.css">

<div class="wl-page">
  <h1 class="wl-title">Travel Wishlist</h1>
  <p class="wl-subtitle">Edit notes, update trip status, or remove destinations. Changes save without reloading the page.</p>

  <div id="wishlist-rows">
    {{template "wishlist/rows.tpl" .}}
  </div>
</div>

<script src="/static/js/wishlist.js"></script>