n, m = gets.split(" ").map(&:to_i)

cross = Array.new(n+1) { Array.new(m+1, 0) }
k = gets.to_i

k.times do
  x, y = gets.split(" ").map(&:to_i)
  cross[x][y] = 1
end

max_y = 0
count = 0
(1..n).each do |x|
  (max_y+1..m).each do |y|
    if cross[x][y] == 1
      count += 1
      max_y = y
      break
    end
  end
end

puts ((n + m - 2 * count + 2 ** 0.5 * count) * 100).round
