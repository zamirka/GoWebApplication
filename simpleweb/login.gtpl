<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            <ol>
                <li><input type="checkbox" name="interest" value="football">Football</li>
                <li><input type="checkbox" name="interest" value="basketball">Basketball</li>
                <li><input type="checkbox" name="interest" value="tennis">Tennis</li>
            </ol>
            <br />
            Username:<input type="text" name="username">
            <br />
            Password:<input type="password" name="password">
            <input type="hidden" name="token" value="{{.}}">
            <br />
            <input type="submit" value="Login">
        </form>
    </body>
</html>