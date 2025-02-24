wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

request = function()
    local randomDomain = "example" .. math.random(1, 1000) .. ".com"
    local randomPath = "path" .. math.random(1, 1000)
    local body = string.format('{"original": "http://%s/%s"}', randomDomain, randomPath)
    return wrk.format(nil, nil, nil, body)
end
