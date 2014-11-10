gets.to_i

arr = gets.chomp.split(" ").map(&:to_i)

weights = [0]
total = arr.reduce(:+)
half =  total / 2
arr.each do |stone|
  tmp = []
  weights.each do |weight|
    sum = weight + stone
    tmp << sum if sum <= half
  end
  weights += tmp
end

max = weights.max
puts (total - max * 2).abs
