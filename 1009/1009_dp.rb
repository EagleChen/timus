n = gets.split()[0].to_i
k = gets.split()[0].to_i

arr = Array.new(n, Array.new(k, 1))

(n-2).downto(0) do |i|
  sum = arr[i+1].reduce(:+)
  arr[i][0] = sum - arr[i+1][0]
  (1..k-1).each do |j|
    arr[i][j] = sum
  end
end

arr[0][0] = 0
puts arr[0].reduce(:+)
