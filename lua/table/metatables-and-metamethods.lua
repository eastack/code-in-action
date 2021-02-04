f1 = {a = 1, b = 2}
f2 = {a = 2, b = 3}

-- s = f1 + f2

metafraction = {}
function metafraction.__add(f1, f2)
	print("f1.a", f1.a)
	print("f2.a", f2.a)
	sum = {}
	sum.b = f1.b * f2.b
	sum.a = f1.a * f2.b + f2.a * f1.b
	return sum
end

metafraction2 = {}
function metafraction2.__add(f1, f2)
	print("2f1.a", f1.a)
	print("2f2.a", f2.a)
	sum = {}
	sum.b = f1.b * f2.b
	sum.a = f1.a * f2.b + f2.a * f1.b
	return sum
end

setmetatable(f1, metafraction)
setmetatable(f2, metafraction2)

s = f1 + f2
for key, val in pairs(getmetatable(f1)['__add'](f1, f2)) do
	print(key, val)
end

