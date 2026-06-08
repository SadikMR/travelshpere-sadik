{{if .Wishlists}}
<table class="w-full border-collapse bg-white rounded-lg overflow-hidden shadow-sm">
  <thead>
    <tr class="bg-gray-50 text-left text-xs font-semibold uppercase tracking-wider text-gray-500">
      <th class="px-4 py-3">Country</th>
      <th class="px-4 py-3">Status</th>
      <th class="px-4 py-3">Note</th>
      <th class="px-4 py-3 text-right">Actions</th>
    </tr>
  </thead>
  <tbody>
    {{range .Wishlists}}
    <tr class="border-t border-gray-100" data-wishlist-id="{{.ID}}">
      <td class="px-4 py-3 font-medium text-gray-900">
        <a href="/countries/{{.CountryName}}" class="text-blue-600 hover:underline">{{.CountryName}}</a>
      </td>

      <td class="px-4 py-3">
        <select
          class="js-status-select text-sm border border-gray-300 rounded px-2 py-1 bg-white focus:outline-none focus:ring-1 focus:ring-blue-400"
          data-id="{{.ID}}"
        >
          <option value="Planned" {{if eq (printf "%s" .Status) "Planned"}}selected{{end}}>Planned</option>
          <option value="Visited" {{if eq (printf "%s" .Status) "Visited"}}selected{{end}}>Visited</option>
        </select>
      </td>

      <td class="px-4 py-3">
        <input
          type="text"
          class="js-note-input text-sm border border-gray-300 rounded px-2 py-1 w-full focus:outline-none focus:ring-1 focus:ring-blue-400"
          data-id="{{.ID}}"
          value="{{.Note}}"
          placeholder="Add a note..."
        />
      </td>

      <td class="px-4 py-3 text-right">
        <div class="flex items-center justify-end gap-2">
          <button
            class="js-save-note text-xs px-3 py-1 bg-blue-600 text-white rounded hover:bg-blue-700"
            data-id="{{.ID}}"
          >Save</button>
          <button
            class="js-delete-btn text-xs px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
            data-id="{{.ID}}"
          >Delete</button>
        </div>
      </td>
    </tr>
    {{end}}
  </tbody>
</table>
{{else}}
<div class="text-center py-12 text-gray-400">
  <p class="text-lg">No wishlist entries yet.</p>
  <p class="text-sm mt-1">Visit a <a href="/countries" class="text-blue-600 hover:underline">country page</a> and click "Add to Wishlist".</p>
</div>
{{end}}