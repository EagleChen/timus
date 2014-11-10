num = gets.to_i

arr = Array.new(num) { [] }
counts = [0] * num
num.times do |index|
  children = gets.split(" ").map { |i| i.to_i - 1 }
  children.pop
  children.each { |child| counts[child] += 1 }
  arr[index] = children
end

order = []
loop do
  break if order.size == num
  index = counts.index(0)
  order << index
  counts[index] = -1 # mark as ordered
  arr[index].each { |child| counts[child] -= 1}
end

puts order.map { |i| i + 1 }.join(" ")
