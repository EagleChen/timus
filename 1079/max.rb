arr = [0, 1]
max = [0, 1]

(2..100000).each do |i|
  arr[i] = i%2 == 0 ? arr[i/2] : arr[i/2] + arr[i/2 + 1]
  max[i] = max[i-1] > arr[i] ? max[i-1] : arr[i]
end

loop do
  num = gets.to_i
  break if num == 0

  puts max[num]
end
