<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} — TravelSphere</title>

    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>

    {{template "partials/navbar.tpl" .}}

    {{.LayoutContent}}

    {{template "partials/footer.tpl" .}}

    <script src="/static/js/app.js"></script>
</body>
</html>