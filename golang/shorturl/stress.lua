wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

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
    local body = string.format('{"original": "http://example.com/%s"}', randomUUID)
    return wrk.format(nil, nil, nil, body)
end
