n, m , cross = 0, 0, {}
count = 0
$stdin.read.lines.each do |line|
  count += 1
  if count == 1
    n, m = line.split(" ").map(&:to_i)
  end

  if count > 2
    cross[line] = 1
  end
end


result = Array.new(m+2) { Array.new(n+2, 0.0) }
(1..n+1).each { |x| result[1][x] = x-1.0 }
(1..m+1).each { |y| result[y][1] = y-1.0 }
diagonal = 2 ** 0.5

(2..m+1).each do |y|
  (2..n+1).each do |x|
    if cross["#{x-1} #{y-1}\n"]
      result[y][x] = result[y-1][x-1] + diagonal
    else
      if result[y-1][x] < result[y][x-1]
        result[y][x] = 1.0 + result[y-1][x]
      else
        result[y][x] = 1.0 + result[y][x-1]
      end
    end
  end
end

puts (result[m+1][n+1] * 100).round
