total = gets.to_i

squares = Hash[(1..245).map { |i| [i, i * i] }]
result = Array.new(total + 1, 4)
top = (total ** 0.5).to_f.floor

squares.each { |k, v| result[v] = 1 }
if result[total] < 4
  puts result[total]
  exit
end

(1..top).each do |i|
  to = ((total - squares[i]) ** 0.5).to_f.floor
  (i..to).each do |j|
    sum = squares[i] + squares[j]
    result[sum] = 2
  end
end

if result[total] < 4
  puts result[total]
  exit
end

(1..top).each do |i|
  to = ((total - squares[i]) ** 0.5).to_f.floor
  (i..to).each do |j|
    two_sums = squares[i] + squares[j]
    extra = ((total - two_sums) ** 0.5).to_f.floor
    (j..extra).each do |k|
      sum = two_sums + squares[k]
      if sum == total
        puts 3
        exit
      end
    end
  end
end

puts 4

#total = gets.to_i
#
#squares = Hash[(1..245).map { |i| [i, i * i] }]
#two_sums = Array.new(246) { Array.new(246) }
#top = (total ** 0.5).ceil
#(1..top).each do |i|
#  if squares[i] == total
#    puts 1
#    exit
#  end
#end
#
#
#(1..top).each do |i|
#  (1..top).each do |j|
#    sum = squares[i] + squares[j]
#    two_sums[i][j] = sum
#    if sum == total
#      puts 2
#      exit
#    elsif sum > total
#      break
#    end
#  end
#end
#
#(1..top).each do |i|
#  (1..top).each do |j|
#    break if two_sums[i][j] > total
#    (1..top).each do |k|
#      sum = two_sums[i][j] + squares[k]
#      if sum == total
#        puts 3
#        exit
#      elsif sum > total
#        break
#      end
#    end
#  end
#end
#
#puts 4

# version 3
#total = gets.to_i
#
#top = (total ** 0.5).ceil
#(1..top).each do |i|
#  if i * i == total
#    puts 1
#    exit
#  end
#end
#
#
#(1..top).each do |i|
#  first = i * i
#  (1..top).each do |j|
#    sum = first + j * j
#    if sum == total
#      puts 2
#      exit
#    elsif sum > total
#      break
#    end
#  end
#end
#
#(1..top).each do |i|
#  first = i * i
#  (1..top).each do |j|
#    second = j * j
#    two_sum = first + second
#    break if two_sum > total
#    (1..top).each do |k|
#      sum = two_sum + k * k
#      if sum == total
#        puts 3
#        exit
#      elsif sum > total
#        break
#      end
#    end
#  end
#end
#
#puts 4
# version 3

# version 2
#choices = Hash[(1..245).map { |i| [i * i, 1] }]
#
#total = gets.to_i
#
#arr = Array.new(total+1, 4)
#arr[1] = 1
#
#(2..total).each do |target|
#  arr[target] = 1 if choices.key?(target)
#end
#
#(2..total).each do |target|
#  if arr[target] == 4
#    top = target / 2
#    (1..top).each do |i|
#      if arr[i] == 1 && arr[target-i] == 1
#        arr[target] = 2
#        break
#      end
#    end
#  end
#end
#
#(2..total).each do |target|
#  if arr[target] == 4
#    top = target / 2
#    (1..top).each do |i|
#      if (arr[i] == 1 && arr[target-i] == 2) ||
#        (arr[i] == 2 && arr[target-i] == 1)
#        arr[target] = 3
#        break
#      end
#    end
#  end
#end
# puts arr[total]
# version 2 end



# version 1
#  choices = Hash[(1..245).map { |i| [i * i, 1] }]
#(2..total).each do |target|
#  if choices.key?(target)
#    arr[target] = 1
#  else
#    arr[target] = 60000
#    top = target / 2
#    (1..top).each do |i|
#      sum = arr[i] + arr[target-i]
#      arr[target] = sum if sum < arr[target]
#    end
#  end
#end
# puts arr[total]
# version 1 end
