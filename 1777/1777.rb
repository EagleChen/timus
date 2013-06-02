arr = gets.split(" ").map { |s| s.to_i }.sort_by { |i| -i }

count = 1
c1, c2 = arr[0] - arr[1], arr[1] - arr[2]
min = [c1, c2].min
arr[0] = arr[1]
arr[1] = arr[2]
arr[2] = min

loop do
  count += 1
  break if arr[2] == arr[1]
  
  if arr[2] > arr[1]
    min = [arr[0] - arr[2], arr[2] - arr[1]].min
    arr[0] = arr[2]
    arr[2] = min
  else
    min = [arr[2], arr[1] - arr[2]].min
    arr[0] = arr[1]
    arr[1] = arr[2]
    arr[2] = min
  end
end

puts count
