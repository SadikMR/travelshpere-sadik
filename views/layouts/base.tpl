<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>

    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="min-h-screen flex flex-col bg-slate-50">

    <nav class="bg-white shadow-sm border-b">
        <div class="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">

            <a href="/" class="text-xl font-bold text-blue-600">
                TravelSphere
            </a>

            <div class="flex items-center gap-6">

                <a href="/" class="hover:text-blue-600">
                    Home
                </a>

                <a href="/countries" class="hover:text-blue-600">
                    Countries
                </a>

                <a href="/wishlist" class="hover:text-blue-600">
                    Wishlist
                </a>

                <a href="/dashboard" class="hover:text-blue-600">
                    Dashboard
                </a>

                {{if .IsLoggedIn}}
                <form action="/logout" method="post" class="inline">
                    <button
                        type="submit"
                        class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">
                        Logout
                    </button>
                </form>
                {{else}}
                <a href="/login" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                    Login
                </a>
                {{end}}

            </div>
        </div>
    </nav>

    <main class="flex-1">
        {{.LayoutContent}}
    </main>

    <footer class="bg-white border-t mt-auto">
        <div class="max-w-7xl mx-auto px-6 py-4 text-center text-gray-500">
            TravelSphere • Learning Project
        </div>
    </footer>

</body>
</html>