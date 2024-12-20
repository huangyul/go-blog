local key = KEYS[1]

local cntKey = key .. ":cnt"

local inputCode = ARGV[1]

local code = redis.call("GET", key)

local cnt = tonumber(redis.call("GET", cntKey))

if cnt == nil or cnt <= 0 then
    -- code not exist or too many validations
    return -1
end

if inputCode == code then
    -- code correct
    redis.call("SET", cntKey, 0)
    return 0
else
    -- code error
    redis.call("decr", cntKey)
    return -2
end
