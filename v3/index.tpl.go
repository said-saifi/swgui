package v3

var indexTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }} - Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="./swagger-ui.css">
    <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32"/>
    <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16"/>
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }

        *,
        *:before,
        *:after {
            box-sizing: inherit;
        }

        body {
            margin: 0;
            background: #fafafa;
        }
    </style>
</head>

<body>
<div id="swagger-ui"></div>

<script src="./swagger-ui-bundle.js"></script>
<script src="./swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {

        var cfg = {
            "title": "API Document",
            "swaggerJsonUrl": "/swagger.json",
            "basePath": "/",
            "showTopBar": false,
            "jsonEditor": false,
            "preAuthorizeApiKey": null
        };

        var cfg = {{ .ConfigJson }};

        var url = "./"+cfg.swaggerJsonUrl;

        // Build a system
        var settings = {
            url: url,
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout",
            showExtensions: true,
            showCommonExtensions: true,
            validatorUrl: null,
            onComplete: function() {
                if (cfg.preAuthorizeApiKey) {
                    for (var name in cfg.preAuthorizeApiKey) {
                        ui.preauthorizeApiKey(name, cfg.preAuthorizeApiKey[name]);
                    }
                }

                var dom = document.querySelector('.scheme-container select');
                for (var key in dom) {
                    if (key.startsWith("__reactInternalInstance$")) {
                        var compInternals = dom[key]._currentElement;
                        var compWrapper = compInternals._owner;
                        compWrapper._instance.setScheme(window.location.protocol.slice(0,-1));
                    }
                }
            }
        };

        if (cfg.showTopBar == false) {
            settings.plugins.push(function () {
                return {
                    components: {
                        Topbar: function () {
                            return null;
                        }
                    }
                }
            });

        }

        window.ui = SwaggerUIBundle(settings);
    }
</script>
</body>
</html>
`
