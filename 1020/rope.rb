num, radius = gets.split(" ")
num = num.to_i
radius = radius.to_f

nails = []
num.times { nails << gets.split(" ").map(&:to_f) }

def len(nails, i1, i2)
  x = nails[i1][0] - nails[i2][0]
  y = nails[i1][1] - nails[i2][1]
  (x ** 2 + y ** 2) ** (1.0 / 2)
end

length = len(nails, num-1, 0)
(1..num-1).each do |index|
  length += len(nails, index, index-1)
end

length += 3.1415926 * radius * 2

puts length.round(2)
