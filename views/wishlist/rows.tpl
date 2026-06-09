{{if .Wishlists}}
<div class="wl-card">
  <table class="wl-table">
    <colgroup>
      <col class="col-country">
      <col class="col-note">
      <col class="col-status">
      <col class="col-actions">
    </colgroup>
    <thead>
      <tr>
        <th>Country</th>
        <th>Note</th>
        <th>Status</th>
        <th class="col-actions">Actions</th>
      </tr>
    </thead>
    <tbody>
      {{range .Wishlists}}
      <tr data-wishlist-id="{{.ID}}">
        <td>
          <a href="/countries/{{.CountryName}}" class="wl-country-link">{{.CountryName}}</a>
        </td>
        <td>
          <input
            type="text"
            class="js-note-input wl-note-input"
            data-id="{{.ID}}"
            value="{{.Note}}"
            placeholder="Add a note..."
          />
        </td>
        <td>
          <select class="js-status-select wl-status-select" data-id="{{.ID}}">
            <option value="Planned" {{if eq (printf "%s" .Status) "Planned"}}selected{{end}}>Planned</option>
            <option value="Visited" {{if eq (printf "%s" .Status) "Visited"}}selected{{end}}>Visited</option>
          </select>
        </td>
        <td>
          <div class="wl-actions">
            <button class="js-save-note wl-btn-save" data-id="{{.ID}}">Save</button>
            <button class="js-delete-btn wl-btn-delete" data-id="{{.ID}}">Delete</button>
          </div>
        </td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>
{{else}}
<div class="wl-card">
  <div class="wl-empty">
    <p>No wishlist entries yet.</p>
    <p style="margin-top:6px;font-size:0.8rem;">Visit a <a href="/countries">country page</a> and click "Add to Wishlist".</p>
  </div>
</div>
{{end}}