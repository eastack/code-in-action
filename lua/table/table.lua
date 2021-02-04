t = {key1 = 'value1', key2 = false}

print(t.key1)
t.newKey = {}
t.key2 = nil

u = {['@!#'] = 'qbert', [{}] = 1729, [6.28] = 'tau'}
print(u[6.28])

a = u['@!#']
b = u[{}]

function h(x) print(x.key1) end
h{key1 = 'Sonmi!451'}

print(_G['_G'] == _G)

v = {'value1', 'value2', 1.21, 'gigawatts'}
for i = 1, #v do
	print(v[i])
end

