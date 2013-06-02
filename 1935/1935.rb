n = gets.to_i

arr = gets.split(" ").map { |s| s.to_i }.sort
puts arr.reduce(:+) + arr.last
