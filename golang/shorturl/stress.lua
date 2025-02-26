wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

math.randomseed(os.clock()^5)
local random = math.random

local function uuid()
    local template ='xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function (c)
        local v = (c == 'x') and random(0, 0xf) or random(8, 0xb)
        return string.format('%x', v)
    end)
end

request = function()
    local randomUUID = uuid()
    -- apply latency with toxiproxy later
    local body = string.format('{"original": "http://localhost:8080/metrics?query=%s"}', randomUUID)
    return wrk.format(nil, nil, nil, body)
end
