num = gets.to_i

first, last = num < 1 ? [num, 1] : [1, num]

puts (first..last).to_a.reduce(0, &:+)
