n, m, y = gets.split(" ").map(&:to_i)

result = []
(0..m-1).each do |i|
  result << i if (i ** n) % m == y
end

if result.empty?
  puts "-1"
else
  puts result.join(" ")
end
