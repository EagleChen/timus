n = gets.to_i

puts gets.split(" ").map { |s| s.to_i / 2 + 1 }.sort[0, n / 2 + 1].reduce(:+)
