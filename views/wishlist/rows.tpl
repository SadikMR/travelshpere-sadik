{{if .Wishlists}}
<table border="1" cellpadding="8">
    <thead>
        <tr>
            <th>Country</th>
            <th>Status</th>
            <th>Note</th>
            <th>Actions</th>
        </tr>
    </thead>

    <tbody>
        {{range .Wishlists}}
        <tr>
            <td>{{.CountryName}}</td>

            <td>{{.Status}}</td>

            <td>{{.Note}}</td>

            <td>
                <button
                    hx-put="/api/wishlist/{{.ID}}"
                    hx-vals='{
                        "note":"{{.Note}}",
                        "status":"planned"
                    }'
                    hx-get="/wishlist/rows"
                    hx-target="#wishlist-rows">
                    Mark Planned
                </button>

                <button
                    hx-delete="/api/wishlist/{{.ID}}"
                    hx-get="/wishlist/rows"
                    hx-target="#wishlist-rows">
                    Delete
                </button>
            </td>
        </tr>
        {{end}}
    </tbody>
</table>
{{else}}
<p>No wishlist entries found.</p>
{{end}}