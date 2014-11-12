n, m = gets.split(" ").map { |s| s.to_i - 1 }

puts n+m-n.gcd(m)
