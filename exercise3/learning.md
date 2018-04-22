# Learning


* Need to serve CSS/JS files within a web application. Use something akin to `m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))` for this

