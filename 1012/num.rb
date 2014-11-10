n_digit, k_base = gets.to_i, gets.to_i

arr = Array.new(k_base) { Array.new(n_digit) }

arr[0][0] = 1
arr[1][0] = k_base - 1

(1..n_digit-1).each do |col|
  arr[0][col] = arr[1][col-1]
  arr[1][col] = (arr[0][col-1] + arr[1][col-1]) * (k_base - 1)
end

puts arr[1][n_digit-1]
