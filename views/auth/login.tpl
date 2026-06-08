<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TravelSphere Login</title>

    <script src="https://cdn.tailwindcss.com"></script>

    <link rel="stylesheet" href="/static/css/login.css">
</head>

<body class="bg-slate-100 min-h-screen flex items-center justify-center">

    <div class="w-full max-w-md bg-white rounded-xl shadow-lg p-8">

        <h1 class="text-3xl font-bold text-center mb-6">
            TravelSphere
        </h1>

        {{if .Error}}
        <div class="mb-4 p-3 rounded bg-red-100 text-red-700">
            {{.Error}}
        </div>
        {{end}}

        <form
            id="loginForm"
            action="/login"
            method="POST"
            class="space-y-4">

            <div>
                <label class="block mb-2 text-sm font-medium">
                    Username
                </label>

                <input
                    id="username"
                    name="username"
                    type="text"
                    placeholder="Enter username"
                    class="w-full border rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
            </div>

            <button
                type="submit"
                class="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition">

                Login
            </button>

        </form>

    </div>

    <script src="/static/js/login.js"></script>

</body>
</html>